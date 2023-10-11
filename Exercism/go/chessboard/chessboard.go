package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (numSquares int) {
	fileSquares, exists := cb[file]

	if !exists {
		return 0
	}

	for _, occupied := range fileSquares {
		if occupied {
			numSquares++
		}
	}

	return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (numRanks int) {
	if rank < 1 || rank > 8 {
		return 0
	}

	for _, v := range cb {
		if v[rank-1] {
			numRanks++
		}
	}

	return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (totalSquares int) {
	totalSquares = len(cb) * len(cb["A"])
	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (occupiedSquares int) {
	for _, ranks := range cb {
		for _, occupied := range ranks {
			if occupied {
				occupiedSquares++
			}
		}
	}

	return
}
