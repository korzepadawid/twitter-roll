package roll

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	capacity = 10
)

func TestNew(t *testing.T) {
	t.Run("should create a new empty roll with given capacity", func(t *testing.T) {
		// given
		capacity := capacity
		// when
		r := New[int](capacity)
		// then
		assert.Equal(t, 0, len(r.items))
		assert.Equal(t, capacity, cap(r.items))
	})
}

func TestAdd(t *testing.T) {
	t.Run("should add a new item to roll when length is smaller than capacity", func(t *testing.T) {
		// given
		r := New[int](capacity)
		// when
		for i := 0; i < capacity; i++ {
			r.Add(i)
		}
		// then
		assert.Equal(t, 10 ,len(r.items))
		assert.Equal(t, 10 ,cap(r.items))
		assert.Equal(t, 9, r.items[0])
		assert.Equal(t, 0, r.items[9])
	})

	t.Run("should add a new item and remove the oldest when length is the same as capacity", func(t *testing.T) {
		// given
		r := New[int](capacity)
		// when
		for i := 0; i < capacity; i++ {
			r.Add(i)
		}
		r.Add(100)
		// then
		assert.Equal(t, 10 ,len(r.items))
		assert.Equal(t, 10 ,cap(r.items))
		assert.Equal(t, 100, r.items[0])
		assert.Equal(t, 1, r.items[9])
	})

	t.Run("should roll to last ten numbers when 100 numbers added", func(t *testing.T) {
		// give
		r := New[int](capacity)
		// when
		for i := 0 ; i < 100; i++ {
			r.Add(i)
		}
		// then
		assert.Equal(t, 10, len(r.items))
		assert.Equal(t, 10, cap(r.items))
		assert.Equal(t, 99, r.items[0])
		assert.Equal(t, 90, r.items[9])
	})

	t.Run("should roll items when added concurrently", func(t *testing.T) {
		// give
		r := New[int](capacity)
		var wg sync.WaitGroup
		// when
		for i := 0 ; i < 10000; i++ {
			wg.Add(1)
			i := i
			go func ()  {
				defer wg.Done()
				r.Add(i)
			}()
		}
		wg.Wait()
		// then
		assert.Equal(t, 10, len(r.items))
		assert.Equal(t, 10, cap(r.items))
		resultMap := make(map[int]int, 10)
		for i, v := range r.ReadAll() {
			resultMap[v] = i	
		}
		assert.Equal(t, 10, len(resultMap))
	})
}
