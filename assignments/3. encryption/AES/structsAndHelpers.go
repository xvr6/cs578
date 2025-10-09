package main

import (
	"fmt"
	"strconv"
	"strings"
)

type StateMatrix struct {
	list [4][4]*SingleHex
}

type SingleHex struct {
	value  uint8
	str    string
	binary [8]uint8
}

func NewStateMatrix(inputByte [128]uint8) *StateMatrix {
	hexArr := [4][4]*SingleHex{}

	for i := range 16 {
		row := i / 4
		newhex := NewSingleHex([8]uint8(inputByte[8*i : 8*(i+1)]))
		hexArr[row][i%4] = newhex
	}
	fmt.Printf("%v", hexArr)

	return &StateMatrix{hexArr}
}

func NewSingleHex(in [8]uint8) *SingleHex {
	var sb strings.Builder

	for i := range 8 {
		sb.WriteString(strconv.Itoa(int(in[i])))
	}

	out, _ := strconv.ParseInt(sb.String(), 2, 16)
	printable := fmt.Sprintf("%x", out)

	return &SingleHex{uint8(out), printable, in}

}

// prints out current state matrix in a table
func (SM StateMatrix) printable() string {
	spacer := ("-----------------\n")
	var mainb strings.Builder
	mainb.WriteString(spacer)

	for r := range 4 {
		var innerb strings.Builder
		innerb.WriteString(fmt.Sprintf("%s|", spacer))
		for c := range 4 {
			innerb.WriteString(fmt.Sprintf("% v |", SM.list[r][c].str))
		}
		innerb.WriteString("\n")

		mainb.WriteString(innerb.String())
	}
	mainb.WriteString(spacer)

	return mainb.String()
}

// func onvToInt() {

// }

// func convToHex() {

// }
