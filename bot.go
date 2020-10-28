package werego

import (
	"errors"
	"fmt"
	"strings"
	"time"

	discord "github.com/bwmarrin/discordgo"
)

// Constants
const (
	MinPlayers = 6
)

// WereBot handles the game commands.
// It also contains the logic and game data.
type WereBot struct {
	session   *discord.Session
	voice     *discord.VoiceConnection
	players   Players
	users     []*discord.User
	started   bool
	startDate time.Time
	votes     map[string]int
}

// NewWereBot returns a new bot that is authenticated
// with a given token.
//
// If the authentication fails, an error is returned.
func NewWereBot(token string) (*WereBot, error) {
	session, err := discord.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}
	wb := &WereBot{session: session}
	wb.session.AddHandler(func(s *discord.Session, m *discord.MessageCreate) {
		go listen(wb, s, m)
	})
	wb.reset()
	err = wb.session.Open()
	return wb, err
}

// Close closes the connection.
//
// session.Close() can return an error
// but it is ignored here.
func (wb *WereBot) Close() {
	wb.session.Close()
}

// Returns true if only humans (aka non-ww)
// players remain.
func (wb WereBot) humansWon() bool {
	for _, player := range wb.players {
		if player.alive && player.Role&RoleWerewolf != 0 {
			return false
		}
	}
	return true
}

// Returns true if there is at least 50%
// of werewolves among players.
func (wb WereBot) werewolvesWon() bool {
	wwCount := len(wb.players.Werewolves().Alive())
	aliveCount := len(wb.players.Alive())
	return wwCount >= (aliveCount / 2)
}

// Handles text commands.
func listen(wb *WereBot, s *discord.Session, m *discord.MessageCreate) {
	if len(m.Content) == 0 || m.Content[0] != '!' || len(m.Content) <= 2 {
		return
	}

	tokens := strings.Split(m.Content, " ")
	command := tokens[0]

	fmt.Println("user:", m.Author.Username)
	fmt.Println("command:", command)
	args := make([]string, 0)
	if len(tokens) > 1 {
		args = tokens[1:]
		fmt.Println("args:", args)
	}

	// If a command is sent without a correct number of args
	if err := checkCommand(command, len(args)); err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	switch command {
	case CommandJoin:
		wb.handleJoin(s, m)
	case CommandStart:
		wb.handleStart(s, m)
	case CommandStop:
		wb.Stop()
	case CommandJoined:
		joined := wb.JoinedUsers()
		s.ChannelMessageSend(m.ChannelID, joined)
	case CommandVote:
		if err := wb.Vote(args[0]); err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err.Error()))
		}
	case CommandVotes:
		s.ChannelMessageSend(m.ChannelID, wb.Votes())
	case CommandKill:
		wb.handleKill(args[0], s, m)
	case CommandRole:
		wb.handleRole(args[0], s, m)
	case CommandRoles:
		wb.sendDM(m.Author, wb.Roles())
	case CommandCleanVotes:
		wb.votes = make(map[string]int, 0)
	case CommandHelp:
		s.ChannelMessageSend(m.ChannelID, help())
	case CommandAlive:
		wb.handleAlive(s, m)
	case CommandConnect:
		wb.handleConnectAudio(args[0], s, m)
	case CommandDisconnect:
		wb.disconnect()
	}
}

func (wb *WereBot) handleJoin(s *discord.Session, m *discord.MessageCreate) {
	if err := wb.Join(m.Author); err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s registered!", m.Author.Mention()))
	}
}

func (wb *WereBot) handleStart(s *discord.Session, m *discord.MessageCreate) {
	if err := wb.Start(); err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err.Error()))
	} else {
		s.ChannelMessageSend(m.ChannelID, "Started!")
	}
}

func (wb *WereBot) handleKill(mention string, s *discord.Session, m *discord.MessageCreate) {
	if err := wb.Kill(mention); err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err.Error()))
	} else {
		if wb.humansWon() {
			s.ChannelMessageSend(m.ChannelID, "Humans won!!")
		} else if wb.werewolvesWon() {
			s.ChannelMessageSend(m.ChannelID, "Werewolve won!!")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Beware, some wolves are still among you...")
		}
	}
}

func (wb *WereBot) handleRole(roleName string, s *discord.Session, m *discord.MessageCreate) {
	role, err := wb.Role(roleName)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err.Error()))
	} else {
		s.ChannelMessageSend(m.ChannelID, role)
	}
}

func (wb *WereBot) handleAlive(s *discord.Session, m *discord.MessageCreate) {
	alivePlayers, err := wb.AlivePlayers()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err.Error()))
	} else {
		s.ChannelMessageSend(m.ChannelID, alivePlayers)
	}
}

func (wb *WereBot) handleConnectAudio(channelName string, s *discord.Session, m *discord.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Cound not find channel.")
		fmt.Println(err.Error())
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		// Could not find guild.
		fmt.Println("Could not find guild.")
		fmt.Println(err.Error())
		return
	}

	for _, c := range g.Channels {
		if c.Type == discord.ChannelTypeGuildVoice && c.Name == channelName {
			voice, err := s.ChannelVoiceJoin(c.GuildID, c.ID, false, false)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if wb.voice != nil {
				wb.voice.Disconnect()
			}
			wb.voice = voice
			break
		}
	}
}

func (wb *WereBot) disconnect() {
	if wb.voice != nil {
		wb.voice.Disconnect()
		wb.voice = nil
	}
}

// Start launches the game. Thus, nobody can join
// the game untill it is stopped.
//
// It needs at least 'MinPlayers' to start.
func (wb *WereBot) Start() error {
	if wb.started {
		return fmt.Errorf("already started (%s)", wb.startDate.Format("2006-01-02 15:04:05"))
	}
	count := len(wb.users)
	roles, err := AllocateRoles(count)
	if err != nil {
		return err
	}
	for i := 0; i < count; i++ {
		player := NewPlayer(roles[i])
		player.User = wb.users[i]
		wb.players = append(wb.players, player)
	}
	wb.sendPlayersRoles()
	wb.started = true
	wb.startDate = time.Now()
	return nil
}

// Join adds the user to the next game.
//
// It requires the game to NOT be started.
func (wb *WereBot) Join(user *discord.User) error {
	if wb.started {
		return errors.New("sorry, the game has already started :(")
	}
	alreadyJoined := false
	for _, u := range wb.users {
		if u.ID == user.ID {
			alreadyJoined = true
			break
		}
	}
	if alreadyJoined {
		return fmt.Errorf("%s, you are already registered ;)", user.Mention())
	}
	wb.users = append(wb.users, user)
	return nil
}

// JoinedUsers returns all the players
// who have joined the game.
func (wb WereBot) JoinedUsers() string {
	if len(wb.users) == 0 {
		return "nobody joined :("
	}
	joined := make([]string, 0)
	for _, user := range wb.users {
		joined = append(joined, user.Username)
	}
	return strings.Join(joined, ",")
}

func (wb WereBot) isMentionInUsers(mention string) bool {
	mention = strings.ReplaceAll(mention, "!", "")
	for _, user := range wb.users {
		if user.Mention() == mention {
			return true
		}
	}
	return false
}

// Vote adds a vote to the given user. It can
// be used to elect a sheriff or kill a player.
func (wb *WereBot) Vote(mention string) error {
	if !wb.isMentionInUsers(mention) {
		return errors.New("player not found")
	}
	if _, found := wb.votes[mention]; !found {
		wb.votes[mention] = 1
	} else {
		wb.votes[mention]++
	}
	return nil
}

// Votes returns the number of votes
// for each player.
func (wb *WereBot) Votes() string {
	results := make([]string, 0)
	if len(wb.votes) == 0 {
		return "no vote"
	}
	for mention, votes := range wb.votes {
		results = append(results, fmt.Sprintf("%s: %d", mention, votes))
	}
	return strings.Join(results, ", ")
}

// Role returns the role of a given user.
//
// The role is visible to EACH USER that can see
// the channel.
func (wb WereBot) Role(mention string) (string, error) {
	mention = strings.ReplaceAll(mention, "!", "")
	for _, player := range wb.players {
		if mention == player.User.Mention() {
			return RoleToString(player.Role, CurrentLanguage)
		}
	}
	for _, user := range wb.users {
		if mention == user.Mention() {
			return "", errors.New("role not set")
		}
	}
	return "", errors.New("player not found")
}

// Roles prints the roles of EACH player.
// Everyone can read it.
func (wb WereBot) Roles() string {
	roles := ""
	for _, player := range wb.players {
		role, _ := RoleToString(player.Role, CurrentLanguage)
		roles += fmt.Sprintf("%s: %s\n", player.User.Username, role)
	}
	return roles
}

// Kill removes a player from the game.
//
// It requires the game to be started.
func (wb *WereBot) Kill(mention string) error {
	if !wb.started {
		return errors.New("game is not started")
	}
	if !wb.isMentionInUsers(mention) {
		return errors.New("player not found")
	}
	mention = strings.ReplaceAll(mention, "!", "")
	for _, player := range wb.players {
		if player.User.Mention() == mention {
			if !player.alive {
				return fmt.Errorf("%s is already dead", player.User.Mention())
			}
			player.alive = false
		}
	}
	wb.votes = make(map[string]int, 0)
	return nil
}

// Stop resets the game.
//
// Nothing happens if the game is not started.
func (wb *WereBot) Stop() {
	wb.reset()
}

func (wb *WereBot) sendDM(user *discord.User, message string) {
	userChannel, _ := wb.session.UserChannelCreate(user.ID)
	wb.session.ChannelMessageSend(userChannel.ID, message)
}

func (wb *WereBot) sendPlayersRoles() {
	for _, player := range wb.players {
		roleName, _ := RoleToString(player.Role, CurrentLanguage)
		wb.sendDM(player.User, roleName)
	}
}

// AlivePlayers returns the names of all
// players that has not been killed.
func (wb WereBot) AlivePlayers() (string, error) {
	if !wb.started {
		return "", errors.New("game is not started")
	}
	alivePlayers := make([]string, 0)
	for _, player := range wb.players.Alive() {
		alivePlayers = append(alivePlayers, player.User.Username)
	}
	return strings.Join(alivePlayers, ", "), nil
}

func (wb *WereBot) reset() {
	wb.players = make([]*Player, 0)
	wb.users = make([]*discord.User, 0)
	wb.votes = make(map[string]int, 0)
	wb.started = false
}
