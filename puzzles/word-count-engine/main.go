package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stripNonAlpha(s string) string {
	var bs []byte
	for i := 0; i < len(s); i++ {
		var b = s[i]
		if b >= 'a' && b <= 'z' {
			bs = append(bs, b)
		}
	}
	return string(bs)
}

func WordCountEngine(document string) [][]string {
	// lower-case-ify document
	document = strings.ToLower(document)
	// split into words by whitespace
	var words = strings.Split(document, " ")
	var counts = make(map[string]int)
	var tokens []string
	var maxCount = 0
	// iterate through document and build a map[string]payload
	for _, w := range words {
		// remove all non-alpha numerica chars from word
		w = stripNonAlpha(w)
		if w == "" {
			continue
		}
		if _, ok := counts[w]; !ok {
			counts[w] = 0
			tokens = append(tokens, w)
		}
		counts[w]++
		if maxCount < counts[w] {
			maxCount = counts[w]
		}
	}

	// Create buckets of words
	var buckets = make([][]string, maxCount+1)

	var cnt int
	for _, t := range tokens {
		cnt = counts[t]
		buckets[cnt] = append(buckets[cnt], t)
	}

	var out [][]string
	for i := len(buckets) - 1; i > 0; i-- {
		for _, j := range buckets[i] {
			out = append(out, []string{j, strconv.Itoa(i)})
		}
	}

	return out
}

func main() {
	fmt.Println("vim-go")
}
