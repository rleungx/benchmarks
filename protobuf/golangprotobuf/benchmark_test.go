package pb

import (
	"testing"

	"github.com/golang/protobuf/proto"
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
		_, err := proto.Marshal(message)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark Proto3 Unmarshal
func BenchmarkProto3Unmarshal(b *testing.B) {
	data, err := proto.Marshal(message)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var message Message
		err := proto.Unmarshal(data, &message)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark Proto3 Clone
func BenchmarkProto3Clone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = proto.Clone(message)
	}
}
