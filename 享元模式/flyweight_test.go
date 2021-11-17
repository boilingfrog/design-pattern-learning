package 享元模式

import "testing"

func TestPlayChess(t *testing.T) {

	ChessBoard := ChessBoard{}
	ChessBoard.InitChessBoard()
	t.Log(ChessBoard.chessPieces[1])
	t.Log(ChessBoard.chessPieces[2])

}
