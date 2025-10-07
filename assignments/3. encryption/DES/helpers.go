package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func convToBinary(in uint8) []uint8 {
	var out []uint8
	conv := strconv.FormatInt(int64(in), 2)
	strs := strings.Split(conv, "")

	fmt.Printf("%v", strs)
	for _, v := range strs {
		i, _ := strconv.Atoi(v)
		out = append(out, uint8(i))
	}
	return out
}

func convToI(in []uint8) uint8 {
	var out uint8 = 0
	for i, v := range in {
		out += v * uint8(math.Pow(2, float64(i)))

	}
	return out
}

func XOR(a1 []uint8, a2 []uint8) []uint8 {
	out := make([]uint8, len(a1))

	for i, v := range a1 {
		out[i] = v ^ a2[i]
	}

	return out
}
