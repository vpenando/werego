package werego

import (
	"fmt"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

// WereBot handles the game commands.
// It also contains the logic and game data.
type WereBot struct {
	Bot     *discord.Session
	players []*Player
	users   []*discord.User
}

// NewWereBot returns a new bot that is authenticated
// with a given token.
func NewWereBot(token string) (*WereBot, error) {
	bot, err := discord.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}
	wb := WereBot{Bot: bot}
	bot.AddHandler(func(s *discord.Session, m *discord.MessageCreate) {
		listen(&wb, s, m)
	})
	wb.Reset()
	return &wb, nil
}

func listen(wb *WereBot, s *discord.Session, m *discord.MessageCreate) {
	if m.Content[0] != '!' || len(m.Content) <= 2 {
		return
	}

	tokens := strings.Split(m.Content[1:], " ")
	command := tokens[0]

	fmt.Println("Mention:", m.Author.Mention())
	fmt.Println("Tokens:", tokens[1])
	s.ChannelMessageSend(m.ChannelID, tokens[1])

	switch command {
	case "reset":

	case "test":
		//test(tokens[1:])
	}
}

func (wb *WereBot) Reset() {
	wb.players = make([]*Player, 0)
	wb.users = make([]*discord.User, 0)
}
