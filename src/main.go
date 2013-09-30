package main

import "fmt"
const ( // iota is reset to 0
	SPACE_EMPTY = iota // c0 == 0
	SPACE_WHITE = iota // c1 == 1
	SPACE_BLACK = iota // c2 == 2
)
func main() {
	fmt.Println("Starting")
	theboard := Board2D{10, 10, make([][]uint8, 10)}
	theboard.Init()

	definitions = loadDefinitionsFromFile("definitions.txt")

	for theboard.IsValid() {
		fmt.Println("Iterating")
	}
	fmt.Println("Complete")
}
func loadDefinitionsFromFile()
{
	
}


// Piece
type Piece struct {
	group      int
	definition []int
}

// End Piece

// Board
type Board2D struct {
	numRows    int
	numColumns int
	matrix     [][]uint8
}

func (b Board2D) Init() {
	matrix := b.matrix
	for i := 0; i < b.numRows; i++ {
		matrix[i] = make([]uint8, b.numColumns)
		for j := 0; j < b.numColumns; j++ {
			matrix[i][j] = SPACE_EMPTY
		}
	}
}
func (b Board2D) Area() int {
	return b.numRows * b.numColumns
}
func (b Board2D) IsValid() bool {
	return false
}

// End Board
