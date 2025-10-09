package main

import (
	"fmt"
	"strconv"
	"strings"
)

type StateMatrix struct {
	list [4][4]*SingleHex
}

type KeyMatrix struct {
	StateMatrix
}

type SingleHex struct {
	value uint8
	str   string
}

// type byte = [8]uint8

// key scheduling algorithm
func (KM *KeyMatrix) nextKey() {
	// AES key scheduling algorithm for 128-bit key (generates next round key)
	rcon := [10]uint8{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1B, 0x36}

	// Rotate last column up by one
	temp := [4]*SingleHex{
		KM.list[1][3],
		KM.list[2][3],
		KM.list[3][3],
		KM.list[0][3],
	}

	// Substitute bytes using S-box
	for i := 0; i < 4; i++ {
		val := temp[i].str

		// row := val >> 4
		// col := val & 0x0F
		temp[i] = NewSingleHex(lookup(val, &sbox))
	}

	// XOR first column with temp and rcon
	KM.list[0][0] = NewSingleHex(KM.list[0][0].value ^ temp[0].value ^ rcon[0])
	for i := 1; i < 4; i++ {
		KM.list[i][0] = NewSingleHex(KM.list[i][0].value ^ temp[i].value)
	}

	// XOR remaining columns
	for col := 1; col < 4; col++ {
		for row := 0; row < 4; row++ {
			KM.list[row][col] = NewSingleHex(KM.list[row][col].value ^ KM.list[row][col-1].value)
		}
	}

}

func NewKeyMatrix(inputByte [128]uint8) *KeyMatrix {
	SM := NewStateMatrix(inputByte)
	KM := KeyMatrix{}
	KM.list = SM.list

	return &KM
}

func (SM *StateMatrix) leftShift(row int, count int) {
	in := SM.list[row]
	x, b := (in)[:count], (in)[count:]

	SM.list[row] = [4]*SingleHex(append(b, x...))

}

func (SM *StateMatrix) updateState(row int, col int, val uint8) {
	new := NewSingleHex(val)
	SM.list[row][col] = new
}

func NewStateMatrix(inputByte [128]uint8) *StateMatrix {
	hexArr := [4][4]*SingleHex{}

	for i := range 16 {
		row := i / 4
		toAdd := [8]uint8(inputByte[8*i : 8*(i+1)])
		newHex := NewSingleHex(Ltoi(toAdd))
		hexArr[row][i%4] = newHex
	}

	return &StateMatrix{hexArr}
}

func NewSingleHex(in uint8) *SingleHex {
	printable := fmt.Sprintf("%x", in)

	return &SingleHex{in, printable}

}

func Ltoi(in [8]uint8) uint8 {
	var sb strings.Builder

	for i := range 8 {
		sb.WriteString(strconv.Itoa(int(in[i])))
	}

	out, _ := strconv.ParseInt(sb.String(), 2, 16)
	return uint8(out)

}

// prints out current state matrix in a table
func (SM StateMatrix) printable() string {
	spacer := ("-------------\n")
	var mainB strings.Builder

	for r := range 4 {
		var innerB strings.Builder
		innerB.WriteString(fmt.Sprintf("%s|", spacer))
		for c := range 4 {
			innerB.WriteString(fmt.Sprintf("%02v|", SM.list[r][c].str))
		}
		innerB.WriteString("\n")

		mainB.WriteString(innerB.String())
	}
	mainB.WriteString(spacer)

	return mainB.String()
}

func lookup(val string, sbox *[][]uint8) uint8 {
	x, _ := strconv.ParseInt(val[1:], 16, 64)
	y, _ := strconv.ParseInt(val[:1], 16, 64)

	return (*sbox)[x][y]

}

// func onvToInt() {

// }

// func convToHex() {

// }
