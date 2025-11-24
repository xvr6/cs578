package main

import (
	"fmt"
)

type Point struct {
	x, y     int
	infinity bool
}

// Modular inverse using extended Euclidean algorithm
func modInv(a, m int) int {
	a = a % m
	if a < 0 {
		a += m
	}
	t, newt := 0, 1
	r, newr := m, a
	for newr != 0 {
		quotient := r / newr
		t, newt = newt, t-quotient*newt
		r, newr = newr, r-quotient*newr
	}
	if r > 1 {
		return -1 // not invertible
	}
	if t < 0 {
		t += m
	}
	return t
}

// Elliptic curve: y^2 = x^3 + ax + b mod p
type Curve struct {
	a, b, p int
}

// Check if point is on curve
func (c Curve) isOnCurve(P Point) bool {
	if P.infinity {
		return true
	}
	left := (P.y * P.y) % c.p
	right := (P.x*P.x*P.x + c.a*P.x + c.b) % c.p
	return left == ((right + c.p) % c.p)
}

// Point addition
func (c Curve) add(P, Q Point) Point {
	if P.infinity {
		return Q
	}
	if Q.infinity {
		return P
	}
	if P.x == Q.x && (P.y != Q.y || P.y == 0) {
		return Point{0, 0, true}
	}
	var m int
	if P.x == Q.x && P.y == Q.y {
		// Point doubling
		denom := modInv(2*P.y, c.p)
		if denom == -1 {
			return Point{0, 0, true}
		}
		m = ((3*P.x*P.x + c.a) * denom) % c.p
	} else {
		denom := modInv(Q.x-P.x, c.p)
		if denom == -1 {
			return Point{0, 0, true}
		}
		m = ((Q.y - P.y) * denom) % c.p
	}
	if m < 0 {
		m += c.p
	}
	x3 := (m*m - P.x - Q.x) % c.p
	y3 := (m*(P.x-x3) - P.y) % c.p
	if x3 < 0 {
		x3 += c.p
	}
	if y3 < 0 {
		y3 += c.p
	}
	return Point{x3, y3, false}
}

// Scalar multiplication
func (c Curve) mul(P Point, k int) Point {
	R := Point{0, 0, true}
	Q := P
	for k > 0 {
		if k%2 == 1 {
			R = c.add(R, Q)
		}
		Q = c.add(Q, Q)
		k /= 2
	}
	return R
}

// Find any point on curve
func findPoint(c Curve) Point {
	for x := 0; x < c.p; x++ {
		rhs := (x*x*x + c.a*x + c.b) % c.p
		for y := 0; y < c.p; y++ {
			if (y*y)%c.p == ((rhs + c.p) % c.p) {
				return Point{x, y, false}
			}
		}
	}
	return Point{0, 0, true}
}

// Count points on curve
func countPoints(c Curve) int {
	count := 1 // point at infinity
	for x := 0; x < c.p; x++ {
		rhs := (x*x*x + c.a*x + c.b) % c.p
		for y := 0; y < c.p; y++ {
			if (y*y)%c.p == ((rhs + c.p) % c.p) {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println("==============================")
	fmt.Println("Assignment 4: Elliptic Curve Cryptography")
	fmt.Println("==============================")
	fmt.Println()

	// Part 1: ECC over mod 29
	fmt.Println("Part 1: ECC Arithmetic (mod 29)")
	curve1 := Curve{a: 1, b: 7, p: 29}
	a := findPoint(curve1)
	fmt.Printf("(a) Example point α on curve: (x=%d, y=%d)\n", a.x, a.y)

	fmt.Println("(b) Scalar multiples of α:")
	for i := 2; i <= 5; i++ {
		pt := curve1.mul(a, i)
		fmt.Printf("    %dα = (x=%d, y=%d)\n", i, pt.x, pt.y)
	}

	pt3 := curve1.mul(a, 3)
	pt5 := curve1.mul(a, 5)
	pt4 := curve1.mul(a, 4)
	pt8a := curve1.add(pt3, pt5)
	pt8b := curve1.add(pt4, pt4)
	fmt.Println("(c) 8α computed two ways:")
	fmt.Printf("    3α + 5α = (x=%d, y=%d)\n", pt8a.x, pt8a.y)
	fmt.Printf("    4α + 4α = (x=%d, y=%d)\n", pt8b.x, pt8b.y)
	if pt8a == pt8b {
		fmt.Println("    Both methods yield the same point.")
	} else {
		fmt.Println("    Methods yield different points!")
	}

	npts := countPoints(curve1)
	fmt.Printf("(d) Total number of points on curve: %d\n", npts)

	fmt.Println("\n------------------------------")
	fmt.Println("Part 2: Elliptic Curve Diffie-Hellman (mod 751)")
	curve2 := Curve{a: -1, b: 188, p: 751}
	a2 := Point{0, 376, false}
	aA, aB := 3, 5
	pubA := curve2.mul(a2, aA)
	pubB := curve2.mul(a2, aB)
	fmt.Printf("(a) Alice's public key: (x=%d, y=%d)\n", pubA.x, pubA.y)
	fmt.Printf("(a) Bob's public key:   (x=%d, y=%d)\n", pubB.x, pubB.y)

	sharedA := curve2.mul(pubB, aA)
	sharedB := curve2.mul(pubA, aB)
	fmt.Printf("(b) Shared key computed by Alice: (x=%d, y=%d)\n", sharedA.x, sharedA.y)
	fmt.Printf("(b) Shared key computed by Bob:   (x=%d, y=%d)\n", sharedB.x, sharedB.y)
	if sharedA == sharedB {
		fmt.Println("    Shared keys match!")
	} else {
		fmt.Println("    Shared keys do NOT match!")
	}
	fmt.Println("==============================")
}
