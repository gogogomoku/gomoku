package player

const (
	NotYetPlaying = iota
	Playing
	Lost
	Won
)

type Player struct {
	Id             int
	PiecesLeft     int
	CapturedPieces int
	Status         int
}

func CreatePlayer(id int, piecesLeft int) *Player {
	player := Player{Id: id, PiecesLeft: piecesLeft}
	return &player
}

func ResetPlayers(p1 *Player, p2 *Player, piecesLeft int) {
	p1.PiecesLeft = piecesLeft
	p1.CapturedPieces = 0
	p1.Status = NotYetPlaying
	p2.PiecesLeft = piecesLeft
	p2.CapturedPieces = 0
	p2.Status = NotYetPlaying
}
