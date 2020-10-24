package werego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	p := NewPlayer(RoleVillager)
	// human
	assert.Equal(t, RoleVillager, p.Role)
	assert.True(t, p.IsAlive())
	assert.True(t, p.IsHuman())
	assert.False(t, p.IsSheriff())
	assert.False(t, p.IsWerewolf())
	// ww
	p = NewPlayer(RoleWerewolf)
	assert.Equal(t, RoleWerewolf, p.Role)
	assert.True(t, p.IsWerewolf())
	assert.False(t, p.IsHuman())
}
