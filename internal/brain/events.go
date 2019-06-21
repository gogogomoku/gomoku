package brain

import "gomoku/internal/player"

func StartRound() {
	GameRound.Status = Running
	player.ResetPlayers(GameRound.P1, GameRound.P2, MAXPIECES)
	GameRound.CurrentPlayer = GameRound.P1
}

func checkValidMove() bool {
	return true
}

func updateWhoseTurn() {
	if GameRound.CurrentPlayer == GameRound.P1 {
		GameRound.CurrentPlayer = GameRound.P2
	} else {
		GameRound.CurrentPlayer = GameRound.P1
	}
}

func HandleMove(id int, position int) {
	if GameRound.CurrentPlayer.Id != id {
		return
	}
	if !checkValidMove() {
		return
	}
	if GameRound.CurrentPlayer.PiecesLeft == 0 {
		return
	}
	GameRound.Goban.Tab[position] = byte(id)
	GameRound.CurrentPlayer.PiecesLeft--
	updateWhoseTurn()
}
