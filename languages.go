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

// RoleDescription returns a description for a given role
// in a given language.
func RoleDescription(role Role, language Language) (string, error) {
	switch language {
	case LanguageFrench:
		return roleDetailsToFrench(role)
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

func roleDetailsToFrench(role Role) (r string, e error) {
	switch role {
	case RoleVillager:
		r = "Ton rôle est de débusquer tous les loups-garous !"
	case RoleWerewolf:
		r = "Ton rôle est de tuer tous les villageois !"
	case RoleSeer:
		r = "Ton rôle est de débusquer tous les loups-garous ! Tu peux regarder le rôle de quelqu'un chaque nuit."
	case RoleWitch:
		r = "Ton rôle est de débusquer tous les loups-garous ! Tu peux utiliser une potion de vie et une potion de mort."
	case RoleHunter:
		r = "Ton rôle est de débusquer tous les loups-garous ! À ta mort, tu peux emporter quelqu'un avec toi."
	case RoleCupid:
		r = "Ton rôle est de débusquer tous les loups-garous ! Tu dois désigner deux amoureux. Si l'un meurt, l'autre se suicidera tant le chagrin sera immense :("
	case RoleThief:
		r = "Voleur"
	case RoleIdiot:
		r = "Ton rôle est de débusquer tous les loups-garous ! Tu ne peux pas être tué par les villageois, mais tu ne pourras plus voter s'ils votent contre toi."
	case RoleGuard:
		r = "Ton rôle est de débusquer tous les loups-garous ! Tu peux protéger un joueur différent chaque nuit contre les loups-garous."
	case RoleRaven:
		r = "Ton rôle est de débusquer tous les loups-garous ! Tu peux attribuer deux votes contre un joueur chaque nuit."
	default:
		e = errors.New("Unknown role")
	}
	return
}
