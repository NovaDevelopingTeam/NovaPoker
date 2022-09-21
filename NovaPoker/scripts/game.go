package NovaPoker

import (
	"math/rand"
	"time"
)

var cards []card

func initialize_cards() {
	for i := 1; i < 11; i++ {
		cards = append(cards, card{i, "diamond", "red"})
	}
	for i := 1; i < 11; i++ {
		cards = append(cards, card{i, "heart", "red"})
	}
	for i := 1; i < 11; i++ {
		cards = append(cards, card{i, "club", "black"})
	}
	for i := 1; i < 11; i++ {
		cards = append(cards, card{i, "spade", "black"})
	}
	// Jack
	cards = append(cards, card{11, "diamond", "red"})
	cards = append(cards, card{11, "heart", "red"})
	cards = append(cards, card{11, "club", "black"})
	cards = append(cards, card{11, "spade", "black"})
	// Queen
	cards = append(cards, card{12, "diamond", "red"})
	cards = append(cards, card{12, "heart", "red"})
	cards = append(cards, card{12, "club", "black"})
	cards = append(cards, card{12, "spade", "black"})
	// King
	cards = append(cards, card{13, "diamond", "red"})
	cards = append(cards, card{13, "heart", "red"})
	cards = append(cards, card{13, "club", "black"})
	cards = append(cards, card{13, "spade", "black"})
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func draw_from_deck() (drawed_card card) {
	drawed_card = cards[0]
	cards = cards[:0]
	return
}

func get_deck_cards() (card_list []card) {
	card_list = cards
	return
}
