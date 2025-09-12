package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.geeksforgeeks.org/dsa/ascii-table/

// sorted in descending order
// https://en.wikipedia.org/wiki/Letter_frequency
// 'e': 12.7, 't': 9.1, 'a': 8.2, 'o': 7.5, 'i': 7.0, 'n': 6.7, 's': 6.3,
// 'h': 6.1, 'r': 6.0, 'd': 4.3, 'l': 4.0, 'c': 2.8, 'u': 2.8, 'm': 2.4,
// 'w': 2.4, 'f': 2.2, 'g': 2.0, 'y': 2.0, 'p': 1.9, 'b': 1.5, 'v': 0.98,
// 'k': 0.77, 'j': 0.15, 'x': 0.15, 'q': 0.095, 'z': 0.074

// I got to a certain point during my first attempt to decipher this text and realized:
// 	Sometimes why is 'Wcy' and sometimes its 'Wry' -> therefore
// 	some letters must have the same frequency.
// 	This was fixed by changing function used in sort.go to account for ties
//  sort the ties value-ascending.
// After noticing this, I restarted from the base string:
// 'e','t','a','o','i','n','s','h','r','d','l','c','u','m','w','f','g','y','p','b','v','k','j','x','q','z'
// perhaps it could be a fun exercise to expand this app into something more functional for a CLI tool as a personal project

var standardFreq = []rune{'e', 't', 'a', 's', 'i', 'r', 'n', 'o', 'c', 'h', 'l', 'd', 'u', 'f', 'p', 'y', 'm', 'w', 'b', 'g', 'k', 'v', 'j', 'x', 'q', 'z'}

/* conversions made for this ciphertext
* etr.     		 -> etc.       	   | swap r/c
* Mus Mact 		 -> Fun Fact   	   | swap m/f, s/n
* Pecauos  		 -> Because   	   | swap p/b, o/s
* belieke  		 -> believe	 	   | swap k/v
* instear  		 -> instead  	   | swap r/d
* tre 	 		 -> the			   | swap r/h
* Yhm		 	 -> why			   | swap y/w, m/y
* caooy	 		 -> carry		   | swap o/r
* crymtopramhers -> cryptomraphers | swap m/p
* cryptomraphers -> cryptographers | swap m/g
 */

// completed deciphering - now i must fix " ’ " being converted into " â "
// used: https://www.babelstone.co.uk/Unicode/whatisit.html to determine that it is U+2019
// used: https://www.compart.com/en/unicode/U+2019			to determine that this is '0xE2' in UTF-8 or 226.
// This was fixed by changing input method from os.ReadFile to bufio.NewReader(file).ReadRune()

func main() {
	charMap := map[rune]int{}
	file, err := os.Open("./ciphertext.txt")
	if err != nil {
		panic(err) //crash
	}
	stat, _ := file.Stat()

	reader := bufio.NewReader(file)
	var dat []rune = make([]rune, 0, stat.Size())
	for { //loop until no more runes can be found
		data, _, err := reader.ReadRune()
		if err != nil {
			break // completed
		}
		dat = append(dat, data)

	}

	file.Close()
	fmt.Println(dat)

	for _, c := range dat {
		if 'a' <= c && c <= 'z' {
			charMap[rune(c)]++
		} else if 'A' <= c && c <= 'Z' {
			lower := c + ('a' - 'A') //convert to ascii value of lower case letter
			charMap[rune(lower)]++
		}

	}

	// Debug: print frequency analysis
	fmt.Println("\n--- Frequency Analysis ---")
	for k, v := range charMap {
		fmt.Printf("%c: %d\n", k, v)
	}

	inOrder := sortRunesByCount(charMap)
	fmt.Println("\n--- Sorted by Frequency ---")
	for i, r := range inOrder {
		fmt.Printf("%d. %c (%d)\n", i+1, r, charMap[r])
	}

	convMap := map[rune]rune{}
	minLen := len(inOrder)
	if len(standardFreq) < minLen {
		minLen = len(standardFreq)
	}
	for i := 0; i < minLen; i++ {
		convMap[inOrder[i]] = standardFreq[i]
	}

	// Debug: print mapping
	fmt.Println("\n--- Mapping (Ciphertext -> Plaintext) ---")
	for k, v := range convMap {
		fmt.Printf("%c -> %c\n", k, v)
	}

	var converted []rune
	var conv rune
	var toFind rune
	var convToUpper bool

	for _, b := range dat {
		fmt.Print(string(b)) //print out raw string for reference.
		convToUpper = false
		br := rune(b)

		// determine what kind of rune we are dealing with
		if 'a' <= br && br <= 'z' { //lower case letter conversion
			toFind = br
		} else if 'A' <= br && br <= 'Z' { // upper case letter conversion.
			convToUpper = true
			toFind = br + ('a' - 'A') //convert to ascii value of lower case letter
		} else {
			converted = append(converted, br)
			continue
		}

		convTo, found := convMap[toFind]
		if !found {
			continue
		}

		conv = convTo
		if convToUpper {
			conv = convTo - ('a' - 'A')
		}
		converted = append(converted, conv)

	}

	fmt.Println("\n\n--- \n\nConversion Attempt: ---")
	fmt.Println(string(converted))

}
