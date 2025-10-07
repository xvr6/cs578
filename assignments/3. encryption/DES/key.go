package main

// amount to left shift by, respective of SegKey count var.
var shiftCount [16]uint8 = [...]uint8{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

type SegKey struct {
	C     []uint8
	D     []uint8
	count uint8
}

func NewSegKey(inputKey []uint8) *SegKey {
	println("creating segkey")
	c0 := diffusion(inputKey, &pc1[0])
	d0 := diffusion(inputKey, &pc1[1])

	return &SegKey{c0, d0, 0}
}

func (k SegKey) getFullKey() []uint8 {
	return append(k.C, k.D...)
}

func (k SegKey) getNextKey() []uint8 {
	cyclicLeftShift(&k.C, shiftCount[k.count])
	cyclicLeftShift(&k.D, shiftCount[k.count])

	fk := k.getFullKey()
	k.count++

	return diffusion(fk, &pc2)

}

// cyclic left shift operation
// ex: (1,2,3,4) -> (2,3,4,1)
func cyclicLeftShift(in *[]uint8, count uint8) {
	x, b := (*in)[:count], (*in)[count:]
	*in = append(b, x...)
}
