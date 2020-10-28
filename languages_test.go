package werego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var languages = []Language{
	LanguageEnglish,
	LanguageFrench,
}

var roles = []Role{
	RoleVillager,
	RoleWerewolf,
	RoleSeer,
	RoleWitch,
	RoleHunter,
	RoleCupid,
	RoleThief,
	RoleIdiot,
	RoleGuard,
	RoleRaven,
}

func TestRoleToString(t *testing.T) {
	for _, language := range languages {
		for _, role := range roles {
			_, err := RoleToString(role, language)
			assert.Nil(t, err)
		}
	}
}
