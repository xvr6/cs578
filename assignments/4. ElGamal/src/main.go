package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// encryptInt performs ElGamal encryption on a single integer message, with explicit ephemeral key k
func encryptInt(m, q, h, g, k *big.Int) (*big.Int, *big.Int) {
	s := modExp(h, k, q)
	p := modExp(g, k, q)
	fmt.Println("g^k used :", p)
	fmt.Println("g^ak used :", s)
	enMsg := new(big.Int).Mul(s, m)
	return enMsg, p
}

// decryptInt performs ElGamal decryption for a single integer message
func decryptInt(enMsg, p, key, q *big.Int) *big.Int {
	h := modExp(p, key, q)
	m := new(big.Int).Div(enMsg, h)
	return m
}

// encrypt performs ElGamal encryption on a string message
func encrypt(msg string, q, h, g *big.Int) ([]*big.Int, *big.Int) {
	enMsg := make([]*big.Int, len(msg))
	k := genKey(q) // Private key for sender
	s := modExp(h, k, q)
	p := modExp(g, k, q)

	fmt.Println("g^k used :", p)
	fmt.Println("g^ak used :", s)

	for i := 0; i < len(msg); i++ {
		m := big.NewInt(int64(msg[i]))
		enMsg[i] = new(big.Int).Mul(s, m)
	}
	return enMsg, p
}

// decrypt performs ElGamal decryption on the encrypted message
func decrypt(enMsg []*big.Int, p, key, q *big.Int) string {
	drMsg := make([]rune, len(enMsg))
	h := modExp(p, key, q)
	for i := 0; i < len(enMsg); i++ {
		m := new(big.Int).Div(enMsg[i], h)
		drMsg[i] = rune(m.Int64())
	}
	return string(drMsg)
}

// modExp computes (base^exp) % mod.
// Replaced with my own implementation of square-and-multiply.
func modExp(base, exp, mod *big.Int) *big.Int {
	// result := new(big.Int).Exp(base, exp, mod)

	e := new(big.Int).Set(exp)
	b := new(big.Int).Mod(new(big.Int).Set(base), mod)
	result := big.NewInt(1)

	zero := big.NewInt(0)

	for e.Cmp(zero) > 0 {
		// if least significant bit is 1, multiply result by current base
		if e.Bit(0) == 1 {
			result.Mul(result, b)
			result.Mod(result, mod)
		}
		// square base: b = b*b mod mod
		b.Mul(b, b)
		b.Mod(b, mod)
		// shift exponent right by 1
		e.Rsh(e, 1)
	}
	return result
}

// greatest common denominator using math/big
func gcd(a, b *big.Int) *big.Int {
	return new(big.Int).GCD(nil, nil, a, b)
}

// genKey generates a random key coprime with q using math/big
func genKey(q *big.Int) *big.Int {
	one := big.NewInt(1)
	min := new(big.Int).Set(q)
	for {
		// Generate a random 80-bit number and add q to ensure it's large
		key, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil))
		if err != nil {
			panic(err)
		}
		key.Add(key, min)
		if gcd(q, key).Cmp(one) == 0 {
			return key
		}
	}
}

func main() {
	// Check for command-line arguments: go run main.go x k m
	if len(os.Args) == 4 {
		xInt, _ := strconv.ParseInt(os.Args[1], 10, 64)
		kInt, _ := strconv.ParseInt(os.Args[2], 10, 64)
		mInt, _ := strconv.ParseInt(os.Args[3], 10, 64)
		x := big.NewInt(xInt)
		k := big.NewInt(kInt)
		m := big.NewInt(mInt)
		q := big.NewInt(409)
		g := big.NewInt(19)
		h := modExp(g, x, q)
		fmt.Printf("Test values: q=%v, g=%v, x=%v, k=%v, m=%v\n", q, g, x, k, m)
		fmt.Printf("h (g^x mod q): %v\n", h)
		enc, p := encryptInt(m, q, h, g, k)
		fmt.Printf("Encrypted message: %v\n", enc)
		fmt.Printf("Ephemeral public key p: %v\n", p)
		dec := decryptInt(enc, p, x, q)
		fmt.Printf("Decrypted message: %v\n", dec)
		return
	}

	// Fallback to interactive
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter private key x (or press Enter to use random):")
	xStr, _ := reader.ReadString('\n')
	xStr = strings.TrimSpace(xStr)
	fmt.Println("Enter ephemeral key k (or press Enter to use random):")
	kStr, _ := reader.ReadString('\n')
	kStr = strings.TrimSpace(kStr)
	fmt.Println("Enter integer message m (or press Enter to use string mode):")
	mStr, _ := reader.ReadString('\n')
	mStr = strings.TrimSpace(mStr)

	if xStr != "" && kStr != "" && mStr != "" {
		// Integer mode: use provided x, k, m
		xInt, _ := strconv.ParseInt(xStr, 10, 64)
		kInt, _ := strconv.ParseInt(kStr, 10, 64)
		mInt, _ := strconv.ParseInt(mStr, 10, 64)
		x := big.NewInt(xInt)
		k := big.NewInt(kInt)
		m := big.NewInt(mInt)
		q := big.NewInt(409)
		g := big.NewInt(19)
		h := modExp(g, x, q)
		fmt.Printf("Test values: q=%v, g=%v, x=%v, k=%v, m=%v\n", q, g, x, k, m)
		fmt.Printf("h (g^x mod q): %v\n", h)
		enc, p := encryptInt(m, q, h, g, k)
		fmt.Printf("Encrypted message: %v\n", enc)
		fmt.Printf("Ephemeral public key p: %v\n", p)
		dec := decryptInt(enc, p, x, q)
		fmt.Printf("Decrypted message: %v\n", dec)
	} else {
		// Default string mode
		msg := "Hello, ElGamal!"
		fmt.Printf("Message: %v\n", msg)

		// q in range [10^20, 10^50)
		minQ := new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil)
		maxQ := new(big.Int).Exp(big.NewInt(10), big.NewInt(50), nil)
		diffQ := new(big.Int).Sub(maxQ, minQ)
		q, err := rand.Int(rand.Reader, diffQ)
		if err != nil {
			panic(err)
		}
		q.Add(q, minQ)
		fmt.Printf("q: %v\n", q)

		// g in range [2, q)
		g, err := rand.Int(rand.Reader, new(big.Int).Sub(q, big.NewInt(2)))
		if err != nil {
			panic(err)
		}
		g.Add(g, big.NewInt(2))
		fmt.Printf("g: %v\n", g)

		key := genKey(q) // private key for sender
		fmt.Printf("key: %v\n", key)

		h := modExp(g, key, q)
		fmt.Printf("h: %v\n", h)

		msgEnc, p := encrypt(msg, q, h, g)
		fmt.Printf("Encrypted message: %v\n", msgEnc)
		msgDec := decrypt(msgEnc, p, key, q)
		fmt.Printf("Decrypted message: %v\n", msgDec)
	}
}
