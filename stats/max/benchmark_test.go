package max

import (
	"math/rand"
	"testing"
	"time"

	"github.com/elliotchance/pie/v2"
	"github.com/montanaflynn/stats"
)

func BenchmarkMax(b *testing.B) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	f := make([]float64, 100000)
	for i := 0; i < 100000; i++ {
		f[i] = r.Float64()
	}
	b.Run("elliotchance/pie/v2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pie.Max(f)
		}
	})
	b.Run("montanaflynn/stats", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stats.Max(f)
		}
	})
}
