package werego

type character struct {
	//Name string
	Role Role

	sheriff bool
	alive   bool
}

func NewCharacter(role Role) *character {
	c := character{
		Role:    role,
		alive:   true,
		sheriff: false,
	}
	return &c
}

// Character is... a character.
type Character interface {
	IsAlive() bool
	ElectAsSheriff()
	IsSheriff() bool
	Kill()
}

// IsAlive returns true if the current character
// has not been killed.
func (p character) IsAlive() bool {
	return p.alive
}

// ElectAsSheriff sets the current character as
// the sheriff.
func (p *character) ElectAsSheriff() {
	p.sheriff = true
}

// IsSheriff returns true if the current character
// was elected sheriff.
func (p character) IsSheriff() bool {
	return p.sheriff
}

// IsWerewolf returns true if the current character
// is a werewolf.
func (p character) IsWerewolf() bool {
	return p.Role&RoleWerewolf != 0
}

func (p *character) Kill() {
	p.alive = false
}
