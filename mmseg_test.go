package mmseg

import "testing"

var ana *Analysis = NewAyalysis("./d/data.txt")

func TestNewWord(t *testing.T) {
	text := "test word"
	freq := 5
	word := NewWord(text, freq)

	if word.Text != text {
		t.Errorf("expected %s, got %s", text, word.Text)
	}

	if word.Freq != freq {
		t.Errorf("expected %d, got %d", freq, word.Freq)
	}
}

func TestNewChunk(t *testing.T) {
	words := []*Word{
		{"test", 4},
		{"word", 5},
	}
	chunk := NewChunk(words)

	if chunk.Length() != 8 {
		t.Errorf("expected 8, got %d", chunk.Length())
	}
	if chunk.AverageLength() != 4 {
		t.Errorf("expected 4, got %d", chunk.AverageLength())
	}
	if chunk.Variance() != 0 {
		t.Errorf("expected 0, got %d", chunk.Variance())
	}
	if chunk.Freq() != 9 {
		t.Errorf("expected 9, got %d", chunk.Freq())
	}
}

func TestAnalysisGetWord(t *testing.T) {
	text := "你好"
	_, ok := ana.Get(text)

	if !ok {
		t.Errorf("shit happen\n")
	}
}

func TestAnalysisWord(t *testing.T) {
	text := "你好"
	word, _ := ana.Get(text)

	if word.Text != text {
		t.Errorf("expected %s, got %s", text, word.Text)
	}

	if word.Freq != 974 {
		t.Errorf("word freq should be 974, got %d", word.Freq)
	}
}

func TestAnalysisMatchWordsLength(t *testing.T) {
	text := "学术危机"
	// text := []byte("你好吗")

	words := ana.MatchWords(text)
	if len(words) != 2 {
		t.Errorf("words length should be 2, got %d, got %v", len(words), words)
	}
}

func TestAnalysisMatchWords(t *testing.T) {
	text := "南京市长江大桥欢迎你"

	words := ana.MatchWords(text)
	if words[0].Text != "南" {
		t.Errorf("match word error")
	}
}

func TestAnalysisChunks(t *testing.T) {
	text := "南京市长江大桥欢迎你"
	chunks := ana.Chunks(text)

	if len(chunks) != 16 {
		t.Errorf("expected 2, got %d", len(chunks))
	}
}
