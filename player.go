package werego

import (
	discord "github.com/bwmarrin/discordgo"
)

// Player is a Discord player.
type Player struct {
	*Character

	User *discord.User
}

// NewPlayer create a new player with a given role.
func NewPlayer(role Role) *Player {
	p := Player{Character: NewCharacter(role)}
	return &p
}
