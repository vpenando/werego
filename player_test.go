package werego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	// humans
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
		p := NewPlayer(role)
		assert.Equal(t, role, p.Role)
		assert.True(t, p.IsAlive())
		assert.True(t, p.IsHuman())
		assert.False(t, p.IsSheriff())
		assert.False(t, p.IsWerewolf())
	}
	// ww
	p := NewPlayer(RoleWerewolf)
	assert.Equal(t, RoleWerewolf, p.Role)
	assert.True(t, p.IsWerewolf())
	assert.False(t, p.IsHuman())
}
