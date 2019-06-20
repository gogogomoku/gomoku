package player

const (
	NotYetPlaying = iota
	Playing
	Lost
	Won
)

type Player struct {
	id             int
	piecesLeft     int
	capturedPieces int
	status         int
}

func CreatePlayer(id int, piecesLeft int) *Player {
	player := Player{id: id, piecesLeft: 45}
	return &player
}
