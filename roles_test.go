package werego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllocateRoles(t *testing.T) {
	roles, err := AllocateRoles(6)
	if err != nil {
		assert.Fail(t, "Error: %s", err.Error())
	}
	assert.True(t, containsExactly(roles, 1, RoleWerewolf))

	roles, err = AllocateRoles(11)
	if err != nil {
		assert.Fail(t, "Error: %s", err.Error())
	}
	assert.True(t, containsExactly(roles, 2, RoleWerewolf))

	roles, err = AllocateRoles(15)
	if err != nil {
		assert.Fail(t, "Error: %s", err.Error())
	}
	assert.True(t, containsExactly(roles, 3, RoleWerewolf))

	roles, err = AllocateRoles(16)
	if err != nil {
		assert.Fail(t, "Error: %s", err.Error())
	}
	assert.True(t, containsExactly(roles, 4, RoleWerewolf))
}

func containsExactly(roles []Role, count int, role Role) bool {
	total := 0
	for _, r := range roles {
		if r == role {
			total++
		}
	}
	return total == count
}
