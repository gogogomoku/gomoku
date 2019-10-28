package player

const (
	NotYetPlaying = iota
	Playing
	Lost
	Won
)

type Player struct {
	Id               int16
	PiecesLeft       int16
	CapturedPieces   int16
	Status           int16
	AiStatus         int16
	WinningSequences [][]int16
}

func CreatePlayer(id int16, piecesLeft int16, AiStatus int16) *Player {
	player := Player{Id: id, PiecesLeft: piecesLeft, AiStatus: AiStatus}
	return &player
}

func ResetPlayers(p1 *Player, p2 *Player, piecesLeft int16, AiStatus1 int16, AiStatus2 int16) {
	p1.PiecesLeft = piecesLeft
	p1.CapturedPieces = 0
	p1.Status = NotYetPlaying
	p1.AiStatus = AiStatus1
	p1.WinningSequences = [][]int16{}
	p2.PiecesLeft = piecesLeft
	p2.CapturedPieces = 0
	p2.Status = NotYetPlaying
	p2.AiStatus = AiStatus2
	p2.WinningSequences = [][]int16{}
}
