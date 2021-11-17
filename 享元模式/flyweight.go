package 享元模式

// ChessPieceUnit ...
type ChessPieceUnit struct {
	id    int
	text  string
	color string
}

var pieces map[int]*ChessPieceUnit

func init() {
	pieces = make(map[int]*ChessPieceUnit, 32)
	pieces[1] = &ChessPieceUnit{
		id:    1,
		text:  "马",
		color: "BLACK",
	}
	pieces[2] = &ChessPieceUnit{
		id:    2,
		text:  "炮",
		color: "BLACK",
	}
	// ...
}

func getChessPiece(chessPieceId int) *ChessPieceUnit {
	return pieces[chessPieceId]
}

type ChessPiece struct {
	chessPieceUnit *ChessPieceUnit
	positionX      int
	positionY      int
}

func newChessPiece(chessPieceId int, positionX, positionY int) *ChessPiece {
	return &ChessPiece{
		chessPieceUnit: getChessPiece(chessPieceId),
		positionX:      positionX,
		positionY:      positionY,
	}
}

// 棋盘
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func (cb *ChessBoard) InitChessBoard() {
	cb.chessPieces = make(map[int]*ChessPiece, 32)
	cb.chessPieces[1] = newChessPiece(1, 0, 1)
	cb.chessPieces[2] = newChessPiece(2, 0, 2)
	// ...
}

// Move 下棋
func (cb *ChessBoard) Move(chessPieceId int, positionX, positionY int) {
	// TODO
}
