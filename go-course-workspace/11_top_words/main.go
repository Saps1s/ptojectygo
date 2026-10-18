package main

import (
	"fmt"
	"sort"
	"strings"
)

type WordStat struct {
	Word  string
	Count int
}

func getTopWords(freq map[string]int, topN int) []WordStat {
	stats := make([]WordStat, 0, len(freq))
	for w, c := range freq {
		stats = append(stats, WordStat{Word: w, Count: c})
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})
	if topN > len(stats) {
		topN = len(stats)
	}
	return stats[:topN]
}

func main() {
	text := "apple banana apple orange banana apple cherry cherry cherry"
	freq := make(map[string]int)
	for _, w := range strings.Fields(text) {
		freq[w]++
	}
	for i, ws := range getTopWords(freq, 5) {
		fmt.Printf("%d. %s: %d\n", i+1, ws.Word, ws.Count)
	}
}