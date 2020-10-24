package werego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWerewolves(t *testing.T) {
	players := Players{
		NewPlayer(RoleWerewolf),
		NewPlayer(RoleWerewolf),
		// Others
		NewPlayer(RoleVillager),
		NewPlayer(RoleSeer),
		NewPlayer(RoleRaven),
		NewPlayer(RoleGuard),
		NewPlayer(RoleThief),
		NewPlayer(RoleIdiot),
		NewPlayer(RoleWitch),
		NewPlayer(RoleHunter),
		NewPlayer(RoleCupid),
	}
	expected := 2
	assert.Equal(t, expected, len(players.Werewolves()))
	diff := len(players) - len(players.Werewolves())
	assert.Equal(t, len(players)-expected, diff)
}

func TestAlive(t *testing.T) {
	players := Players{
		NewPlayer(RoleWerewolf),
		NewPlayer(RoleWerewolf),
		// Others
		NewPlayer(RoleVillager),
		NewPlayer(RoleSeer),
		NewPlayer(RoleRaven),
		NewPlayer(RoleGuard),
		NewPlayer(RoleThief),
		NewPlayer(RoleIdiot),
		NewPlayer(RoleWitch),
		NewPlayer(RoleHunter),
		NewPlayer(RoleCupid),
	}
	players[0].Kill()
	players[1].Kill()
	expected := len(players) - 2
	assert.Equal(t, expected, len(players.Alive()))
	diff := len(players) - len(players.Alive())
	assert.Equal(t, len(players)-expected, diff)
}
