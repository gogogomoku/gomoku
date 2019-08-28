package player

import "testing"

func TestCreatePlayer(t *testing.T) {
	tables := []struct {
		player                 *Player
		expectedId             int16
		expectedPiecesLeft     int16
		expectedCapturedPieces int16
		expectedStatus         int16
	}{
		{CreatePlayer(1, 50), 1, 50, 0, 0},
		{CreatePlayer(2, 1000), 2, 1000, 0, 0},
	}
	for _, table := range tables {
		if table.player.Id != table.expectedId {
			t.Errorf("Player initialized with wrong id: %d, want: %d.", table.player.Id, table.expectedId)
		} else if table.player.PiecesLeft != table.expectedPiecesLeft {
			t.Errorf("Player initialized with wrong piecesLeft: %d, want: %d.", table.player.PiecesLeft, table.expectedPiecesLeft)
		} else if table.player.CapturedPieces != table.expectedCapturedPieces {
			t.Errorf("Player initialized with wrong capturedPieces: %d, want: %d.", table.player.CapturedPieces, table.expectedCapturedPieces)
		} else if table.player.Status != table.expectedStatus {
			t.Errorf("Player initialized with wrong status: %d, want: %d.", table.player.Status, table.expectedStatus)
		}
	}
}

func TestResetPlayers(t *testing.T) {
	t.Skip("Skipping, function prototype likely to change")
}
