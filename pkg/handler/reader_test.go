package handler

import (
	"mobydick/pkg/services"

	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordFreqSortDesc(t *testing.T) {
	var wordFreqs services.ByFreq = []services.WordFreq{
		{Word: "the", Freq: 30},
		{Word: "is", Freq: 56},
		{Word: "at", Freq: 34},
		{Word: "or", Freq: 25},
	}

	sort.Sort(sort.Reverse(services.ByFreq(wordFreqs)))
	assert.NotEmpty(t, wordFreqs)
	assert.Equal(t, "is", wordFreqs[0].Word)
	assert.Equal(t, 56, wordFreqs[0].Freq)
}


func TestWordFreqSortAsc(t *testing.T) {
	var wordFreqs services.ByFreq = []services.WordFreq{
		{Word: "the", Freq: 30},
		{Word: "is", Freq: 56},
		{Word: "at", Freq: 34},
		{Word: "or", Freq: 25},
	}

	sort.Sort(services.ByFreq(wordFreqs))
	assert.NotEmpty(t, wordFreqs)
	assert.Equal(t, "or", wordFreqs[0].Word)
	assert.Equal(t, 25, wordFreqs[0].Freq)
}
