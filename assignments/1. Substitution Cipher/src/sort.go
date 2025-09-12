package main

import "sort"

// sortRunesByCount takes a map of runes to their counts and returns a slice of runes
// sorted in descending order by their count values (highest first).
func sortRunesByCount(in map[rune]int) []rune {

	// Define a key-value struct to hold rune and its count.
	type kv struct {
		Key   rune
		Count int
	}
	var toSort []kv

	// Convert the map to a slice of kv structs.
	for k, v := range in {
		toSort = append(toSort, kv{k, v})
	}

	// Sort the slice by count in descending order, and by rune value ascending for ties.
	sort.Slice(toSort, func(i, j int) bool {
		if toSort[i].Count == toSort[j].Count {
			return toSort[i].Key < toSort[j].Key
		}
		return toSort[i].Count > toSort[j].Count
	})

	// Extract the sorted runes into a result slice.
	result := make([]rune, len(toSort))
	for i, kv := range toSort {
		result[i] = kv.Key
	}
	return result
}
