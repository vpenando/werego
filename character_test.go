package werego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCharacter(t *testing.T) {
	humanRoles := []Role{
		RoleVillager,
		RoleSeer,
		RoleWitch,
		RoleHunter,
		RoleCupid,
		RoleThief,
		RoleIdiot,
		RoleGuard,
		RoleRaven,
	}
	for _, role := range humanRoles {
		c := NewCharacter(role)
		t.Logf("Creating a new character with role '%d'...", c.Role)
		// human
		assert.Equal(t, role, c.Role)
		assert.True(t, c.IsAlive())
		assert.True(t, c.IsHuman())
		assert.False(t, c.IsSheriff())
		assert.False(t, c.IsWerewolf())
	}
	// ww
	c := NewCharacter(RoleWerewolf)
	assert.Equal(t, RoleWerewolf, c.Role)
	assert.True(t, c.IsWerewolf())
	assert.False(t, c.IsHuman())
}

func TestKill(t *testing.T) {
	c := NewCharacter(RoleVillager)
	assert.True(t, c.IsAlive())
	c.Kill()
	assert.False(t, c.IsAlive())
}
