package werego

import (
	"flag"
	"fmt"
	"log"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

var (
	Bot *discord.Session
)

func init() {
	token := flag.String("t", "", "Your discord token")
	flag.Parse()
	var err error
	Bot, err = discord.New(fmt.Sprintf("Bot %s", *token))
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
	Bot.AddHandler(listen)
}

func listen(s *discord.Session, m *discord.MessageCreate) {
	if m.Content[0] != '!' || len(m.Content) <= 2 {
		return
	}

	tokens := strings.Split(m.Content[1:], " ")
	command := tokens[0]

	switch command {
	case "reset":
		reset()
	case "test":
		test(tokens[1:])
	}
}

func test(args []string) {
	user, err := Bot.User(args[0])
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(user)
}
