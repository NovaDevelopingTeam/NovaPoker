package NovaPoker

type Player struct {
	id        int
	isPlaying bool
	cards     []card
	fiches    int
}

type Game struct {
	id      int
	players []Player
	board   []card
	lastBet int
	plate   int
}

func players_draw(game Game) {
	for _, player := range game.players {
		var new_card card = draw_player()
		player.cards = append(player.cards, new_card)
		new_card = draw_player()
		player.cards = append(player.cards, new_card)
	}
}

func board_draw(game Game) {
	var new_card card = draw_board()
	game.board = append(game.board, new_card)
}

func start_game(id int) {
	var game Game
	game = Game{id, []Player{}, []card{}, 2000, 0}
	players_draw(game)
	// force first bet [revisited gameplay]
	for i, player := range game.players {
		if bet(game, game.lastBet, player) == true {
			continue
		} else {
			game.players = game.players[:i]
		}
	}
	// first 3 board cards
	board_draw(game)
	board_draw(game)
	board_draw(game)
	// TODO: players actions and board draws
	for _, player := range game.players {
		check_best_win(player, game)
	}
}
