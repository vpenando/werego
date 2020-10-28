package werego

import (
	"errors"
)

// Language is used for define available ones.
type Language int

// Available languages.
const (
	LanguageFrench Language = iota
	LanguageEnglish

	CurrentLanguage = LanguageFrench
)

// RoleToString returns a role as string, depending
// on the given language.
func RoleToString(role Role, language Language) (string, error) {
	switch language {
	case LanguageFrench:
		return roleToFrench(role)
	case LanguageEnglish:
		return roleToEnglish(role)
	default:
		return "", errors.New("Unknown language")
	}
}

func roleToEnglish(role Role) (r string, e error) {
	switch role {
	case RoleVillager:
		r = "Villager"
	case RoleWerewolf:
		r = "Werewolf"
	case RoleSeer:
		r = "Seer"
	case RoleWitch:
		r = "Witch"
	case RoleHunter:
		r = "Hunter"
	case RoleCupid:
		r = "Cupid"
	case RoleThief:
		r = "Thief"
	case RoleIdiot:
		r = "Idiot"
	case RoleGuard:
		r = "Guard"
	case RoleRaven:
		r = "Raven"
	default:
		e = errors.New("Unknown role")
	}
	return
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
