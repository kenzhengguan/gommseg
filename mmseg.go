package mmseg

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	ChineseCharLength = 3
)

type Word struct {
	Text string
	Freq int
}

func NewWord(text string, freq int) *Word {
	return &Word{text, freq}
}

type Chunk struct {
	Words []*Word
}

func NewChunk(words []*Word) *Chunk {
	return &Chunk{words}
}

func (c *Chunk) Length() int {
	var length int
	for _, word := range c.Words {
		length += len(word.Text)
	}

	return length
}

func (c *Chunk) AverageLength() float64 {
	return float64(c.Length()) / float64(len(c.Words))
}

func (c *Chunk) Variance() float64 {
	var averageLength float64 = c.AverageLength()
	var sumDistance float64
	for _, word := range c.Words {
		sumDistance += math.Pow(float64(len(word.Text))-averageLength, 2.0)
	}

	return math.Sqrt(sumDistance / float64(len(c.Words)))
}

func (c *Chunk) Freq() int {
	var freq int
	for _, word := range c.Words {
		freq += word.Freq
	}

	return freq
}

type Analysis struct {
	WordMap map[string]*Word
}

// "/Users/raquelken/Desktop/mini.txt"
func NewAyalysis(fileName string) *Analysis {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var (
		er   error = nil
		line []byte
	)

	var wordMap map[string]*Word
	wordMap = make(map[string]*Word)

	for er == nil {
		line, _, er = reader.ReadLine()
		a := bytes.Split(line, []byte("\t"))

		if len(a) >= 2 {
			text := string(a[0])
			freq, e := strconv.Atoi(string(a[1]))
			if e == nil {
				wordMap[text] = &Word{text, freq}
			}
		}
	}

	return &Analysis{wordMap}
}

func (ana *Analysis) Get(text string) (*Word, bool) {
	word, ok := ana.WordMap[text]
	return word, ok
}

func (ana *Analysis) MatchWords(text string) []*Word {
	var (
		pos        int
		matchWords []*Word
	)

	for pos = ChineseCharLength; pos <= len(text); pos += ChineseCharLength {
		t := string(text[0:pos])
		if word, ok := ana.Get(t); ok {
			matchWords = append(matchWords, word)
		}
	}

	return matchWords
}

func (ana *Analysis) Chunks(text string) []*Chunk {
	var chunks []*Chunk
	for _, word1 := range ana.MatchWords(text) {
		textLength := len(text)
		wordLength1 := len(word1.Text)
		if wordLength1 < textLength {
			text1 := string([]byte(text)[wordLength1:textLength])
			for _, word2 := range ana.MatchWords(text1) {
				wordLength2 := len(word2.Text)
				if wordLength1+wordLength2 < textLength {
					text2 := string([]byte(text)[wordLength1+wordLength2 : textLength])
					for _, word3 := range ana.MatchWords(text2) {
						chunks = append(chunks, NewChunk([]*Word{word1, word2, word3}))
					}
				} else {
					chunks = append(chunks, NewChunk([]*Word{word1, word2}))
				}
			}
		} else {
			chunks = append(chunks, NewChunk([]*Word{word1}))
		}
	}
	for idx, chunk := range chunks {
		fmt.Printf("chunk %d", idx)
		for _, word := range chunk.Words {
			fmt.Printf("%s ", word.Text)
		}
		fmt.Println("")
	}
	return chunks
}
