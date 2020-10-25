package werego

func (wb *WereBot) Start() error {
	count := len(wb.users)
	roles, err := AllocateRoles(count)
	if err != nil {
		return err
	}
	for i := 0; i < count; i++ {
		player := NewPlayer(roles[i])
		player.User = wb.users[i]
		wb.players = append(wb.players, player)
	}
	return nil
}

func stop() {

}

func vote() {

}
