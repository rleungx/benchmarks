package pb

import (
	"testing"

	"github.com/cloudwego/fastpb"
)

var message = &Message{
	Query:         "abcdefg",
	PageNumber:    100,
	ResultPerPage: 1,
	Comment:       []byte("xyz"),
	Corpus:        Message_UNIVERSAL,
}

// Marshal .
func Marshal(msg fastpb.Writer) []byte {
	// TODO: buffer can be reused.
	buf := make([]byte, msg.Size())

	msg.FastWrite(buf)
	return buf
}

// Unmarshal .
func Unmarshal(buf []byte, msg fastpb.Reader) error {
	_, err := fastpb.ReadMessage(buf, int8(fastpb.SkipTypeCheck), msg)
	return err
}

// Benchmark Proto3 Marshal
func BenchmarkProto3Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Marshal(message)
	}
}

// Benchmark Proto3 Unmarshal
func BenchmarkProto3Unmarshal(b *testing.B) {
	data := Marshal(message)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var message Message
		err := Unmarshal(data, &message)
		if err != nil {
			b.Fatal(err)
		}
	}
}
