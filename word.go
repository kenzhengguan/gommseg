package gommseg

type Word struct {
	Text string
	Freq int
}

func NewWord(text string, freq int) *Word {
	return &Word{text, freq}
}
