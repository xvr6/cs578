package main

import (
	"fmt"
)

// consts needed for the logic
var (
	rounds    uint8 = 16
	blockSize uint8 = 64
	// -1 are the parity bits which were removed for this assignment
	key        = [64]int8{0, 0, 1, 0, 0, 0, 0, -1, 0, 0, 0, 0, 1, 1, 1, -1, 1, 0, 1, 1, 1, 0, 0, -1, 0, 1, 0, 0, 1, 0, 1, -1, 1, 1, 1, 1, 1, 1, 0, -1, 1, 0, 1, 1, 0, 1, 1, -1, 1, 1, 1, 1, 1, 0, 1, -1, 1, 1, 1, 1, 1, 1, 1, -1}
	plaintext1 = [64]uint8{1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0}
	plaintext2 = [64]uint8{}
	plaintext3 = [64]uint8{}
)

// turn array of len 32 into array of len 48
func expand(in [32]uint8) [48]uint8 {

	var out [48]uint8

	for i, b := range eBlock {
		for j, c := range b {
			curr := i + j
			out[c] = in[curr]
		}
	}

	return out
}

// function which is recursively called to print out a result.
func round(l []uint8, r []uint8, ctr uint8) {
	fmt.Printf("\nlen(l): %v\nlen(r): %v", len(l), len(r))
	//step 1 - turn each side into 48 bits

}

func main() {
	fmt.Printf("\n1) Plaintext: %v", plaintext1)

	var initPermutation [64]uint8 = [64]uint8(diffusion(plaintext1, ip))

	fmt.Printf("\n2) ip: %v", initPermutation)

	round(initPermutation[0:32], initPermutation[32:64], 0)

	// var finalPermutation [64]uint8 = permutation(plaintext1, fp)

	// fmt.Printf("\n5) ip: %v", finalPermutation)
}
