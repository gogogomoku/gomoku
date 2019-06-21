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

func HandleMove(id int, position int) (code int, msg string) {
	if GameRound.CurrentPlayer.Id != id {
		return 1, "It is not your turn"
	}
	if !checkValidMove() {
		return 1, "Move isn't valid"
	}
	if GameRound.CurrentPlayer.PiecesLeft == 0 {
		return 1, "You have no pieces left"
	}
	GameRound.Goban.Tab[position] = int8(id)
	GameRound.CurrentPlayer.PiecesLeft--
	updateWhoseTurn()
	return 0, "Move done"
}
