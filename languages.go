package werego

import (
	"errors"
)

// Language is used for define available ones.
type Language int

// Available languages.
const (
	French = iota
)

// RoleToString returns a role as string, depending
// on the given language.
func RoleToString(role Role, language Language) (string, error) {
	switch language {
	case French:
		return roleToFrench(role)
	default:
		return "", errors.New("Unknown language")
	}
}

func roleToFrench(role Role) (r string, e error) {
	switch role {
	case RoleVillager:
		r = "Villageois"
	case RoleWerewolf:
		r = "Loup-garou"
	case RoleSeer:
		r = "Voyante"
	case RoleWitch:
		r = "Sorcière"
	case RoleHunter:
		r = "Chasseur"
	case RoleCupid:
		r = "Cupidon"
	case RoleThief:
		r = "Voleur"
	case RoleIdiot:
		r = "Idiot du village"
	case RoleGuard:
		r = "Salvateur"
	case RoleRaven:
		r = "Corbeau"
	default:
		e = errors.New("Unknown role")
	}
	return
}
