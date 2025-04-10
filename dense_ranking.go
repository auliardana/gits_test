package main

import (
	"fmt"
	"sort"
)

func denseRanking(scores []int, gits []int) []int {
	unique := []int{}
	seen := map[int]bool{}
	for _, score := range scores {
		if !seen[score] {
			unique = append(unique, score)
			seen[score] = true
		}
	}
	sort.Slice(unique, func(i, j int) bool { return unique[i] > unique[j] })

	result := []int{}
	for _, score := range gits {
		rank := sort.Search(len(unique), func(i int) bool { return unique[i] <= score })
		result = append(result, rank+1)
	}
	return result
}

func main() {
	scores := []int{100, 100, 50, 40, 40, 20, 10}
	gits1 := []int{5, 25, 50, 100}
	gits2 := []int{1, 20, 60}
	gits3 := []int{100, 80, 100, 5, 4, 3}

	fmt.Println(denseRanking(scores, gits1))
	fmt.Println(denseRanking(scores, gits2))
	fmt.Println(denseRanking(scores, gits3))
}
