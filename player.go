package werego

// Player is a Discord player.
type Player struct {
	*character
}

// NewPlayer create a new player with a given role.
func NewPlayer(role Role) *Player {
	p := Player{character: NewCharacter(role)}
	return &p
}
