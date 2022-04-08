package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkAlphaChar(charVariable rune) bool {
	if (charVariable >= 'a' && charVariable <= 'h') || (charVariable >= 'A' && charVariable <= 'H') {
		return true
		// fmt.Printf("%s is an alphabet in chess\n", string(charVariable))
	} else {
		//fmt.Printf("**** %s is not an alphabet position in chess *****\n", string(charVariable))
		return false
	}
}

func checkNumber(numVariable int8) bool {
	if numVariable >= 1 && numVariable <= 8 {
		//fmt.Print("\n Number Position of the given peice : ", currentPositionNumber)
		return true
	} else {
		return false
		//fmt.Println("****", currentPositionNumber, "is not a valid Number Position. Please enter again******")
	}
}

func main() {

	var cPiece string
	var currentPositionNumber int8
	var charCPA rune
	var nextMoves string
	fmt.Print("\n-----------------START---------------")
	for 1 > 0 {
		fmt.Print("\n Please select a chess piece using short code(Q-Queen, R-Rook, N-Knight): ")
		fmt.Scanln(&cPiece)
		if cPiece == "Q" || cPiece == "N" || cPiece == "R" || cPiece == "q" || cPiece == "n" || cPiece == "r" {
			fmt.Println("\n Selected chess peice is: ", cPiece)
			break
		} else {
			fmt.Println("****", cPiece, "is not a valid chess piece. Please select again******")
		}
	}

	fmt.Print("\n-----------------Alphabet Position---------------")
	for 1 > 0 {
		fmt.Print("\n Please select alphabet position of the choosen piece (select in between A-H) : ")
		currentPositionAlphabet := bufio.NewReader(os.Stdin)
		charCPA, _, _ = currentPositionAlphabet.ReadRune()
		if checkAlphaChar(charCPA) {
			fmt.Print("\n Alphabet Position of the given peice : ", string(charCPA))
			break
		} else {
			fmt.Println("****", string(charCPA), "is not a valid Alphabet Position. Please enter again******")
		}

	}

	fmt.Print("\n \n-----------------Number Position---------------")
	for 1 > 0 {
		fmt.Print("\n Please select number position of the choosen piece is(select in between 1-8) : ")
		fmt.Scanln(&currentPositionNumber)

		if checkNumber(currentPositionNumber) {
			fmt.Print("\n Number Position of the given piece : ", currentPositionNumber)
			break
		} else {
			fmt.Println("****", currentPositionNumber, "is not a valid Number Position. Please enter again******")
		}
	}

	if cPiece == "N" || cPiece == "n" {
		fmt.Print("\n \n-----------------Potential Board Postions---------------")
		// All possible moves of a knight
		X := [8]int8{2, 1, -1, -2, -2, -1, 1, 2}
		Y := [8]int8{1, 2, 2, 1, -1, -2, -2, -1}

		count := 0

		//Check if each possible move is valid or not
		for i := 0; i < 8; i++ {
			x := charCPA + rune(X[i])
			y := currentPositionNumber + Y[i]
			// count valid moves
			if checkNumber(y) && checkAlphaChar(x) {
				numberStr := strconv.Itoa(int(y))
				nextMoves = nextMoves + " " + string(x) + numberStr
				count++
			}
		}
		fmt.Print("\n \n Total number of potential board postions of the Knight: ", count)
		fmt.Print("\n \n List of all the potential board postions of the Knight : ", nextMoves)
	}

	if cPiece == "R" || cPiece == "r" {
		fmt.Print("\n \n-----------------Potential Board Postions---------------")
		// All possible moves of a Rook
		X := [14]int8{-7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7}
		Y := [14]int8{-7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7}

		count := 0

		//Check if each possible move is valid or not
		for i := 0; i < 14; i++ {
			x := charCPA
			y := currentPositionNumber + Y[i]
			// count valid moves
			if checkNumber(y) && checkAlphaChar(x) {
				numberStr := strconv.Itoa(int(y))
				nextMoves = nextMoves + " " + string(x) + numberStr
				count++
			}
		}

		for i := 0; i < 14; i++ {
			x := charCPA + rune(X[i])
			y := currentPositionNumber
			// count valid moves
			if checkNumber(y) && checkAlphaChar(x) {
				numberStr := strconv.Itoa(int(y))
				nextMoves = nextMoves + " " + string(x) + numberStr
				count++
			}
		}

		fmt.Print("\n \n Total number of potential board postions of the Rook: ", count)
		fmt.Print("\n \n List of all the potential board postions of the Rook : ", nextMoves)
	}

	if cPiece == "Q" || cPiece == "q" {
		fmt.Print("\n \n-----------------Potential Board Postions---------------")
		// All possible moves of a Queen
		X := [14]int8{-7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7}
		Y := [14]int8{-7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7}
		M := [14]int8{-7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7}
		N := [14]int8{7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7}

		count := 0

		//Check if each possible move is valid or not
		for i := 0; i < 14; i++ {
			x := charCPA
			y := currentPositionNumber + Y[i]
			// count valid moves
			if checkNumber(y) && checkAlphaChar(x) {
				numberStr := strconv.Itoa(int(y))
				nextMoves = nextMoves + " " + string(x) + numberStr
				count++
			}
		}

		for i := 0; i < 14; i++ {
			x := charCPA + rune(X[i])
			y := currentPositionNumber
			// count valid moves
			if checkNumber(y) && checkAlphaChar(x) {
				numberStr := strconv.Itoa(int(y))
				nextMoves = nextMoves + " " + string(x) + numberStr
				count++
			}
		}

		for i := 0; i < 14; i++ {
			x := charCPA + rune(X[i])
			y := currentPositionNumber + Y[i]
			// count valid moves
			if checkNumber(y) && checkAlphaChar(x) {
				numberStr := strconv.Itoa(int(y))
				nextMoves = nextMoves + " " + string(x) + numberStr
				count++
			}
		}

		for i := 0; i < 14; i++ {
			x := charCPA + rune(M[i])
			y := currentPositionNumber + N[i]
			// count valid moves
			if checkNumber(y) && checkAlphaChar(x) {
				numberStr := strconv.Itoa(int(y))
				nextMoves = nextMoves + " " + string(x) + numberStr
				count++
			}
		}

		fmt.Print("\n \n Total number of potential board postions of the Queen: ", count)
		fmt.Print("\n \n List of all the potential board postions of the Queen : ", nextMoves)
	}
	fmt.Print("\n-----------------END---------------")
}
