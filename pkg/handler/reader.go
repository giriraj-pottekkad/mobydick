package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	word "mobydick/pkg/services"
	"regexp"
	"sort"
	"strings"
)

func ReadFile(filepath string, count int) {
	bs, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}
	
	text := strings.ToLower(string(bs))

	re := regexp.MustCompile("[a-zA-Z]+")
	matches := re.FindAllString(text, -1)

	words := make(map[string]int)

	for _, match := range matches {
		words[match]++
	}

	var wordFreqs []word.WordFreq
	for k, v := range words {
		wordFreqs = append(wordFreqs, word.WordFreq{Word: k, Freq: v})
	}

	sort.Sort(sort.Reverse(word.ByFreq(wordFreqs)))

	for i := 0; i < count; i++ {
		fmt.Printf("%v\n", wordFreqs[i])
	}
}
