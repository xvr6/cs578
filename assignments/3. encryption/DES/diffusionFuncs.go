package main

import "fmt"

// copilot was used here to turn these from strings into arrays.
var eBlock = [][]uint8{
	{32, 1, 2, 3, 4, 5},
	{4, 5, 6, 7, 8, 9},
	{8, 9, 10, 11, 12, 13},
	{12, 13, 14, 15, 16, 17},
	{16, 17, 18, 19, 20, 21},
	{20, 21, 22, 23, 24, 25},
	{24, 25, 26, 27, 28, 29},
	{28, 29, 30, 31, 32, 1},
}

// [0][*][*] is for C0 - og left key;
// [1][*][*] is for D0 - og right key;
var pc1 = [2][][]uint8{
	{
		{57, 49, 41, 33, 25, 17, 9},
		{1, 58, 50, 42, 34, 26, 18},
		{10, 2, 59, 51, 43, 35, 27},
		{19, 11, 3, 60, 52, 44, 36},
	},
	{
		{63, 55, 47, 39, 31, 23, 15},
		{7, 62, 54, 46, 38, 30, 22},
		{14, 6, 61, 53, 45, 37, 29},
		{21, 13, 5, 28, 20, 12, 4},
	},
}

var pc2 = [][]uint8{
	{14, 17, 11, 24, 1, 5},
	{3, 28, 15, 6, 21, 10},
	{23, 19, 12, 4, 26, 8},
	{16, 7, 27, 20, 13, 2},
	{41, 52, 31, 37, 47, 55},
	{30, 40, 51, 45, 33, 48},
	{44, 49, 39, 56, 34, 53},
	{46, 42, 50, 36, 29, 32},
}

var ip = [][]uint8{
	{58, 50, 42, 34, 26, 18, 10, 2},
	{60, 52, 44, 36, 28, 20, 12, 4},
	{62, 54, 46, 38, 30, 22, 14, 6},
	{64, 56, 48, 40, 32, 24, 16, 8},
	{57, 49, 41, 33, 25, 17, 9, 1},
	{59, 51, 43, 35, 27, 19, 11, 3},
	{61, 53, 45, 37, 29, 21, 13, 5},
	{63, 55, 47, 39, 31, 23, 15, 7},
}

var fp = [][]uint8{
	{40, 8, 48, 16, 56, 24, 64, 32},
	{39, 7, 47, 15, 55, 23, 63, 31},
	{38, 6, 46, 14, 54, 22, 62, 30},
	{37, 5, 45, 13, 53, 21, 61, 29},
	{36, 4, 44, 12, 52, 20, 60, 28},
	{35, 3, 43, 11, 51, 19, 59, 27},
	{34, 2, 42, 10, 50, 18, 58, 26},
	{33, 1, 41, 9, 49, 17, 57, 25},
}

var p = [][]uint8{
	{16, 7, 20, 21},
	{29, 12, 28, 17},
	{1, 15, 23, 26},
	{5, 18, 31, 10},
	{2, 8, 24, 14},
	{32, 27, 3, 9},
	{19, 13, 30, 6},
	{22, 11, 4, 25},
}

/*
TODO: write documentation
*/
func diffusion(in []uint8, sMap *[][]uint8) []uint8 {

	//calculated size needed for the output.
	var size int = (len(*sMap) * len((*sMap)[0]))
	var out []uint8 = make([]uint8, size)

	if size >= len(in) { //map in values to smap position
		println("\nin -> smap")
		//this shuffles the bits and either increases or keeps size same.

		for i, b := range *sMap {
			//j is ctr, c is value in second dimension array
			for j, c := range b {
				//current index in array
				curr := i + j
				pos := c - 1
				toInsert := in[curr]
				if toInsert == 0 {
					continue
				}
				fmt.Printf("%v <- pos: %2v; insert: %v\n", out, pos, toInsert)
				out[pos] = in[curr]
			}
		}

	} else { // take values from 'in' and turn them into a string based upon the positions stated .
		// ex if sMap is {2,1,3,4} then we take values from position 2 of 'in' and assign it to the new position 0 in output.
		// this can also shrink input
		println("\nin <- smap")

		pos := 0
		for _, r := range *sMap {
			for _, c := range r {
				toInsert := in[c-1]
				if toInsert == 0 {
					pos++
					continue
				}
				fmt.Printf("%v <- pos: %2v; insert %v from in[%2v]\n", out, pos, toInsert, c)

				out[pos] = toInsert
				pos++
			}
		}
	}
	return out

}
