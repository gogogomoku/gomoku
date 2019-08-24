package player

const (
	NotYetPlaying = iota
	Playing
	Lost
	Won
)

type Player struct {
	Id             int16
	PiecesLeft     int16
	CapturedPieces int16
	Status         int16
}

func CreatePlayer(id int16, piecesLeft int16) *Player {
	player := Player{Id: id, PiecesLeft: piecesLeft}
	return &player
}

func ResetPlayers(p1 *Player, p2 *Player, piecesLeft int16) {
	p1.PiecesLeft = piecesLeft
	p1.CapturedPieces = 0
	p1.Status = NotYetPlaying
	p2.PiecesLeft = piecesLeft
	p2.CapturedPieces = 0
	p2.Status = NotYetPlaying
}
