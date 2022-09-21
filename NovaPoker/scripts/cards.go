package NovaPoker

type card struct {
	value  int
	symbol string
	color  string
}

func draw_card() (drawed_card card) {
	drawed_card = draw_from_deck()
	return
}

func get_cards() (card_list []card) {
	card_list = get_deck_cards()
	return
}
