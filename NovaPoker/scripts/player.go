package NovaPoker

func draw_player() (drawed_card card) {
	drawed_card = draw_from_deck()
	return
}

func see() {
	return
}

func check() {
	return
}

func bet(game Game, amount int, player Player) bool {
	if player.fiches >= amount {
		player.fiches -= amount
		game.lastBet = amount
		game.plate += amount
		return true
	} else {
		return false
	}
}

func leave() {
	return
}
