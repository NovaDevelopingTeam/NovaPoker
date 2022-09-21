package NovaPoker

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove_by_value(s []int, e int) []int {
	for _, i := range s {
		if i == e {
			s = s[:i]
			return s
		}
	}
	return s
}

func check_pairs_tris(player Player, game Game, hand []card) (pairs int, tris int) {
	var values []int
	var paired []int
	for _, card := range hand {
		if contains(values, card.value) == true {
			remove_by_value(values, card.value)
			pairs += 1
			paired = append(paired, card.value)
		} else {
			if contains(paired, card.value) {
				tris += 1
				pairs -= 1
			} else {
				values = append(values, card.value)
			}
		}
	}
	return
}

func check_best_win(player Player, game Game) {
	var hand []card
	for _, card := range game.board {
		hand = append(hand, card)
	}
	for _, card := range player.cards {
		hand = append(hand, card)
	}
	check_pairs_tris(player, game, hand)
}
