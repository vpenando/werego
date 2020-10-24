package werego

import "testing"

func TestAllocateRoles(t *testing.T) {

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
