package dequeue

import (
	"container/list"
	"testing"

	ed "github.com/edwingeng/deque/v2"
	efd "github.com/ef-ds/deque"
	gd "github.com/gammazero/deque"
	juju "github.com/juju/collections/deque"
	gostl "github.com/liyue201/gostl/ds/deque"
	"github.com/phf/go-queue/queue"
)

func BenchmarkPushBack(b *testing.B) {
	b.Run("container/list", func(b *testing.B) {
		q := list.New()
		for i := 0; i < b.N; i++ {
			q.PushBack(i)
		}
	})
	b.Run("edwingeng/deque", func(b *testing.B) {
		q := ed.NewDeque[any]()
		for i := 0; i < b.N; i++ {
			q.PushBack(i)
		}
	})
	b.Run("gammazero/deque", func(b *testing.B) {
		q := gd.New[any]()
		for i := 0; i < b.N; i++ {
			q.PushBack(i)
		}
	})
	b.Run("phf/go-queue/queue", func(b *testing.B) {
		q := queue.New()
		q.Init()
		for i := 0; i < b.N; i++ {
			q.PushBack(i)
		}
	})
	b.Run("juju/collections/deque", func(b *testing.B) {
		q := juju.New()
		for i := 0; i < b.N; i++ {
			q.PushBack(i)
		}
	})
	b.Run("ef-ds/deque", func(b *testing.B) {
		q := efd.New()
		for i := 0; i < b.N; i++ {
			q.PushBack(i)
		}
	})
	b.Run("liyue201/gostl/ds/deque", func(b *testing.B) {
		q := gostl.New[int]()
		for i := 0; i < b.N; i++ {
			q.PushBack(i)
		}
	})
}

func BenchmarkPushFront(b *testing.B) {
	b.Run("container/list", func(b *testing.B) {
		q := list.New()
		for i := 0; i < b.N; i++ {
			q.PushFront(i)
		}
	})
	b.Run("edwingeng/deque", func(b *testing.B) {
		q := ed.NewDeque[any]()
		for i := 0; i < b.N; i++ {
			q.PushFront(i)
		}
	})
	b.Run("gammazero/deque", func(b *testing.B) {
		q := gd.New[any]()
		for i := 0; i < b.N; i++ {
			q.PushFront(i)
		}
	})
	b.Run("phf/go-queue/queue", func(b *testing.B) {
		q := queue.New()
		q.Init()
		for i := 0; i < b.N; i++ {
			q.PushFront(i)
		}
	})
	b.Run("juju/collections/deque", func(b *testing.B) {
		q := juju.New()
		for i := 0; i < b.N; i++ {
			q.PushFront(i)
		}
	})
	b.Run("ef-ds/deque", func(b *testing.B) {
		q := efd.New()
		for i := 0; i < b.N; i++ {
			q.PushFront(i)
		}
	})
	b.Run("liyue201/gostl/ds/deque", func(b *testing.B) {
		q := gostl.New[int]()
		for i := 0; i < b.N; i++ {
			q.PushFront(i)
		}
	})
}

func BenchmarkClone(b *testing.B) {
	b.Run("container/list", func(b *testing.B) {
		q := list.New()
		for i := 0; i < 1000000; i++ {
			q.PushBack(i)
		}
		clone := func() {
			q1 := list.New()
			v := q.Remove(q.Front())
			q.PushBack(v)
			q1.PushBack(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			clone()
		}
	})
	b.Run("edwingeng/deque", func(b *testing.B) {
		q := ed.NewDeque[any]()
		for i := 0; i < 1000000; i++ {
			q.PushBack(i)
		}
		clone := func() {
			q1 := ed.NewDeque[any]()
			v := q.PopFront()
			q.PushBack(v)
			q1.PushBack(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			clone()
		}
	})
	b.Run("gammazero/deque", func(b *testing.B) {
		q := gd.New[any]()
		for i := 0; i < 1000000; i++ {
			q.PushBack(i)
		}
		clone := func() {
			q1 := gd.New[any]()
			v := q.PopFront()
			q.PushBack(v)
			q1.PushBack(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			clone()
		}
	})
	b.Run("phf/go-queue/queue", func(b *testing.B) {
		q := queue.New().Init()
		for i := 0; i < 1000000; i++ {
			q.PushBack(i)
		}
		clone := func() {
			q1 := queue.New().Init()
			v := q.PopFront()
			q.PushBack(v)
			q1.PushBack(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			clone()
		}
	})
	b.Run("juju/collections/deque", func(b *testing.B) {
		q := juju.New()
		for i := 0; i < 1000000; i++ {
			q.PushBack(i)
		}
		clone := func() {
			q1 := juju.New()
			v, _ := q.PopFront()
			q.PushBack(v)
			q1.PushBack(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			clone()
		}
	})
	b.Run("ef-ds/deque", func(b *testing.B) {
		q := efd.New()
		for i := 0; i < 1000000; i++ {
			q.PushBack(i)
		}
		clone := func() {
			q1 := efd.New()
			v, _ := q.PopFront()
			q.PushBack(v)
			q1.PushBack(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			clone()
		}
	})
	b.Run("liyue201/gostl/ds/deque", func(b *testing.B) {
		q := gostl.New[int]()
		for i := 0; i < 1000000; i++ {
			q.PushBack(i)
		}
		clone := func() {
			q1 := gostl.New[int]()
			v := q.PopFront()
			q.PushBack(v)
			q1.PushBack(v)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			clone()
		}
	})
}
