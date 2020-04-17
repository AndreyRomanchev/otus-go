package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
)

func Top10(text string) []string {
	const top = 10

	if len(text) == 0 {
		return nil
	}

	count := make(map[string]int)
	r := regexp.MustCompile(`[^\s]+`)
	for _, word := range r.FindAllString(text, -1) {
		count[word]++
	}

	pairs := make([]string, 0, top)
	for key := range count {
		pairs = append(pairs, key)
	}

	sort.Slice(pairs, func(i, j int) bool {
		return count[pairs[i]] > count[pairs[j]]
	})

	return pairs[:top]
}
