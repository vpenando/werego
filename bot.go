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
	Bot       *discord.Session
	players   []*Player
	users     []*discord.User
	started   bool
	startDate time.Time
	votes     map[string]int
}

// NewWereBot returns a new bot that is authenticated
// with a given token.
func NewWereBot(token string) (*WereBot, error) {
	bot, err := discord.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}
	wb := &WereBot{Bot: bot}
	bot.AddHandler(func(s *discord.Session, m *discord.MessageCreate) {
		listen(wb, s, m)
	})
	wb.reset()
	return wb, nil
}

func (wb WereBot) humansWon() bool {
	for _, p := range wb.players {
		if p.alive && p.Role&RoleWerewolf != 0 {
			return false
		}
	}
	return true
}

func (wb WereBot) werewolvesWon() bool {
	wwCount := 0
	aliveCount := 0
	for _, p := range wb.players {
		if p.alive {
			aliveCount++
		}
	}
	for _, p := range wb.players {
		if p.alive && p.Role&RoleWerewolf != 0 {
			wwCount++
		}
	}
	return wwCount >= (aliveCount / 2)
}

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

	if err := checkCommand(command, len(args)); err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	switch command {
	case CommandJoin:
		if err := wb.Join(m.Author); err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s registered!", m.Author.Mention()))
		}
	case CommandStart:
		if err := wb.Start(); err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err.Error()))
		} else {
			s.ChannelMessageSend(m.ChannelID, "Started!")
		}
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
		if err := wb.Kill(args[0]); err != nil {
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
	case CommandRole:
		role, err := wb.Role(args[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err.Error()))
		} else {
			s.ChannelMessageSend(m.ChannelID, role)
		}
	case CommandCleanVotes:
		wb.votes = make(map[string]int, 0)
	case CommandHelp:
		s.ChannelMessageSend(m.ChannelID, help())
	case CommandAlive:
		alivePlayers, err := wb.AlivePlayers()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %s", err.Error()))
		} else {
			s.ChannelMessageSend(m.ChannelID, alivePlayers)
		}

	}

}

// Start launches the game. Thus, nobody can join
// the game untill it is stopped.
//
// It requires at least 'MinPlayers' to start.
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
	for _, u := range wb.users {
		fmt.Println("User: ", u.Mention())
		fmt.Println("Mention: ", mention)
		if u.Mention() == mention {
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
	for _, p := range wb.players {
		if mention == p.User.Mention() {
			return RoleToString(p.Role, CurrentLanguage)
		}
	}
	return "", errors.New("player not found")
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

func (wb *WereBot) sendPlayersRoles() {
	for _, player := range wb.players {
		userChannel, _ := wb.Bot.UserChannelCreate(player.User.ID)
		roleName, _ := RoleToString(player.Role, CurrentLanguage)
		wb.Bot.ChannelMessageSend(userChannel.ID, roleName)
	}
}

// AlivePlayers returns the names of all
// players that has not been killed.
func (wb WereBot) AlivePlayers() (string, error) {
	if wb.started {
		return "", errors.New("game is not started")
	}
	alivePlayers := make([]string, 0)
	for _, player := range wb.players {
		if player.alive {
			alivePlayers = append(alivePlayers, player.User.Username)
		}
	}
	return strings.Join(alivePlayers, ", "), nil
}

func (wb *WereBot) reset() {
	wb.players = make([]*Player, 0)
	wb.users = make([]*discord.User, 0)
	wb.votes = make(map[string]int, 0)
	wb.started = false
}
