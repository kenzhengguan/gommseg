package gommseg

import (
	"math"
)

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
