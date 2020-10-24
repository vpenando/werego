package werego

// Character represent an in game character.
// It has a role and some additional informations.
type Character struct {
	Role Role

	// Private fields
	sheriff bool
	alive   bool
}

// NewCharacter creates a new character.
// This character has theses properties:
//
//  - The character has a given role;
//  - The character is considered alive;
//  - The character is NOT sheriff.
func NewCharacter(role Role) Character {
	c := Character{
		Role:    role,
		alive:   true,
		sheriff: false,
	}
	return c
}

// IsAlive returns true if the current character
// has not been killed.
func (c Character) IsAlive() bool {
	return c.alive
}

// ElectAsSheriff sets the current character as
// the sheriff.
func (c *Character) ElectAsSheriff() {
	c.sheriff = true
}

// IsSheriff returns true if the current character
// was elected sheriff.
func (c Character) IsSheriff() bool {
	return c.sheriff
}

// IsHuman returns true if the current character
// is NOT a werewolf.
func (c Character) IsHuman() bool {
	return !c.IsWerewolf()
}

// IsWerewolf returns true if the current character
// is a werewolf.
func (c Character) IsWerewolf() bool {
	return c.Role&RoleWerewolf != 0
}

// Kill sets 'alive' to false, meaning the current
// character has been killed.
func (c *Character) Kill() {
	c.alive = false
}
