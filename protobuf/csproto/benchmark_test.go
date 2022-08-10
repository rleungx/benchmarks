package pb

import (
	"testing"
)

var message = &Message{
	Query:         "abcdefg",
	PageNumber:    100,
	ResultPerPage: 1,
	Comment:       []byte("xyz"),
	Corpus:        Message_UNIVERSAL,
}

// Benchmark Proto3 Marshal
func BenchmarkProto3Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := message.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark Proto3 Unmarshal
func BenchmarkProto3Unmarshal(b *testing.B) {
	data, err := message.Marshal()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var message Message
		err := message.Unmarshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
