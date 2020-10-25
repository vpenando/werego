package werego

import (
	"errors"
	"math/rand"
	"time"
)

// Role defines each player's goal and abilities.
type Role = int

// Here are the different players' roles.
const (
	RoleVillager = (1 << iota)
	RoleWerewolf
	RoleSeer
	//LittleGirl
	RoleWitch
	RoleHunter
	RoleCupid
	RoleThief
	//Shaman

	// Additional roles
	RoleIdiot
	RoleGuard
	RoleRaven
)

// AllocateRoles picks a given number of roles,
// depending on the number of players.
func AllocateRoles(playersCount int) ([]Role, error) {
	roles, err := computeRoles(playersCount)
	if err != nil {
		return nil, err
	}
	suffleRoles(&roles)
	return roles, nil
}

func computeRoles(playersCount int) ([]Role, error) {
	if playersCount < 6 {
		return nil, errors.New("not enough players")
	}
	roles := make([]Role, 0, playersCount)
	fillWerewolves(&roles, playersCount)
	fillSpecialSlots(&roles, playersCount)
	fillRemainingSlots(&roles, playersCount)
	return roles, nil
}

func fillWerewolves(roles *[]Role, playersCount int) {
	werewolvesCount := computeWerewolvesCount(playersCount)
	for i := 0; i < werewolvesCount; i++ {
		*roles = append(*roles, RoleWerewolf)
	}
}

func fillSpecialSlots(roles *[]Role, playersCount int) {
	// Special role: all but villager and ww
	availableRoles := map[Role]bool{
		RoleSeer: true,

		RoleRaven: playersCount >= 6,
		RoleGuard: playersCount >= 6,
		RoleThief: playersCount >= 6,
		RoleIdiot: playersCount >= 6,

		RoleWitch:  playersCount >= 7,
		RoleHunter: playersCount >= 7,

		RoleCupid: playersCount >= 8,
	}
	for role, isRoleAvailable := range availableRoles {
		availableSlots := int(playersCount) - len(*roles)
		if availableSlots == 0 {
			break
		}
		if isRoleAvailable {
			*roles = append(*roles, role)
		}
	}
}

func fillRemainingSlots(roles *[]Role, playersCount int) {
	availableSlots := playersCount - len(*roles)
	for i := 0; i < availableSlots; i++ {
		*roles = append(*roles, RoleVillager)
	}
}

func computeWerewolvesCount(playersCount int) int {
	switch {
	case playersCount <= 6:
		return 1
	case playersCount <= 11:
		return 2
	case playersCount <= 15:
		return 3
	}
	return 4
}

func suffleRoles(roles *[]Role) {
	r := *roles
	indexes := make([]int, len(r))
	for i := 0; i < len(indexes); i++ {
		indexes[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(indexes), func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})
	newRoles := make([]Role, len(r))
	for i := 0; i < len(indexes); i++ {
		newRoles[i] = r[indexes[i]]
	}
	*roles = newRoles
}
