package main

import (
	"fmt"
)

// consts needed for the logic
var (
	DEBUG           = false
	rounds    uint8 = 16
	blockSize uint8 = 64
	// bits 8,16,24,32,40,48,56,64 are the parity bits which were removed for this assignment and replaced as 0s
	key        = []uint8{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0}
	sKey       = NewSegKey(key) // convert into SegKey struct
	plaintext1 = []uint8{1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0}
	plaintext2 = [64]uint8{}
	plaintext3 = [64]uint8{}
)

// function which is recursively called to print out a result.
func round(inL []uint8, inR []uint8, ctr uint8) []uint8 {

	fout := f(inL, inR)

	// xor function f output with L0

	outR := XOR(inL, fout)
	outL := inR

	if ctr == rounds-1 {
		if DEBUG {
			fmt.Println("done!")

		}
		return append(outL, outR...)
	}
	ctr++
	return round(outL, outR, ctr)

}

// function f as described by the paper
func f(inL []uint8, inR []uint8) []uint8 {
	// ----- BEGIN function f -----
	if DEBUG {
		fmt.Printf("\nlen(l): %v\nlen(r): %v", len(inL), len(inR))
	}

	//step 1 - right side into 48 bits
	Ri := diffusion(inL, &eBlock)
	//step 2 - get current key, then xor with ri
	Ki := sKey.getNextKey()
	o := XOR(Ri, Ki)
	//step 3 - S-box substitution; each 6bit sub-array -> 4bit
	var step3 []uint8

	for i := range 8 {
		si := o[i*6 : (i+1)*6]
		//check notes 2.3: 3: S-box substitution for more detailed explanation about what is happening here.
		row := append(si[:1], si[5:]...)
		col := si[1:5]
		x, y := convToI(row), convToI(col)
		if DEBUG {
			fmt.Printf("\n%v:\nsi: %v\nRow: %v; Col: %v\n[%v,%2v]\n", i, si, row, col, x, y)
		}

		siFinal := convToBinary(lookUp(x, y, &s[i]))
		step3 = append(step3, siFinal...)
	}
	if DEBUG {
		fmt.Printf("\n%v", step3)

	}
	//step 4 - bitwise permutation p

	step4 := diffusion(step3, &p)
	if DEBUG {
		fmt.Printf("step4: %v\n", step4)
	}

	// ----- END function f -----

	return step4
}

func main() {
	fmt.Printf("\nPlaintext: %v\n", plaintext1)

	var initPermutation []uint8 = diffusion(plaintext1, &ip)

	// fmt.Printf("\n2) ip: %v\n", initPermutation)

	preout := round(initPermutation[0:32], initPermutation[32:64], 0)
	// fmt.Printf("\nPreout: %v\n", preout)

	var finalPermutation []uint8 = diffusion(preout, &fp)

	fmt.Printf("\nfinal: %v", finalPermutation)
}
