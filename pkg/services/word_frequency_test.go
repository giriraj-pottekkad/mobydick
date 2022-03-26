package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordFreqToString(t *testing.T) {
	wordFreq := WordFreq{
		Word: "the",
		Freq: 30,
	}
	assert.Equal(t, "30 the", wordFreq.String())
}

func TestWordFreqLength(t *testing.T) {
	wordFreq := WordFreq{
		Word: "the",
		Freq: 30,
	}
	var frequencices ByFreq = []WordFreq{wordFreq}
	assert.NotEmpty(t, frequencices)
	assert.Equal(t, 1, frequencices.Len())
}

func TestWordFreqSwap(t *testing.T) {
	wordFreq1 := WordFreq{
		Word: "the",
		Freq: 30,
	}
	wordFreq2 := WordFreq{
		Word: "at",
		Freq: 20,
	}
	var frequencices ByFreq = []WordFreq{wordFreq1, wordFreq2}
	assert.NotEmpty(t, frequencices)
	assert.Equal(t, frequencices[0].Word, "the")
	assert.Equal(t, frequencices[1].Word, "at")

	frequencices.Swap(0, 1)
	assert.Equal(t, frequencices[0].Word, "at")
	assert.Equal(t, frequencices[1].Word, "the")
}

func TestWordFreqLess(t *testing.T) {
	wordFreq1 := WordFreq{
		Word: "the",
		Freq: 30,
	}
	wordFreq2 := WordFreq{
		Word: "at",
		Freq: 20,
	}
	var frequencices ByFreq = []WordFreq{wordFreq1, wordFreq2}
	assert.NotEmpty(t, frequencices)
	assert.True(t, frequencices.Less(1, 0))
}

func TestWordFreqGreater(t *testing.T) {
	wordFreq1 := WordFreq{
		Word: "the",
		Freq: 30,
	}
	wordFreq2 := WordFreq{
		Word: "at",
		Freq: 20,
	}
	var frequencices ByFreq = []WordFreq{wordFreq1, wordFreq2}
	assert.NotEmpty(t, frequencices)
	assert.False(t, frequencices.Less(0, 1))
}
