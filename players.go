package werego

// Players is a pool of players
// with additionnal filters.
type Players []*Player

// Werewolves returns ONLY the players
// who are considered werewolves.
func (p Players) Werewolves() Players {
	capacity := computeWerewolvesCount(len(p))
	werewolves := make(Players, 0, capacity)
	for _, player := range p {
		if player.IsWerewolf() {
			werewolves = append(werewolves, player)
		}
	}
	return werewolves
}

// Alive returns ONLY the players
// who are still alive.
func (p Players) Alive() Players {
	alivePlayers := make(Players, 0, len(p))
	for _, player := range p {
		if player.IsAlive() {
			alivePlayers = append(alivePlayers, player)
		}
	}
	return alivePlayers
}
