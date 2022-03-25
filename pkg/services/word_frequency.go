package services

import "fmt"

type WordFreq struct {
	Word string
	Freq int
}

func (p WordFreq) String() string {
	return fmt.Sprintf("%s %d", p.Word, p.Freq)
}

type ByFreq []WordFreq

func (a ByFreq) Len() int           { return len(a) }
func (a ByFreq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFreq) Less(i, j int) bool { return a[i].Freq < a[j].Freq }
