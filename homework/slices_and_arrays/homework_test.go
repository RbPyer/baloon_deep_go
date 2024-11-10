package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

const EmptyValue = -1

type Int interface {
	int | int8 | int16 | int32 | int64
}

type CircularQueue[T Int] struct {
	values   []T
	head     int
	cursor   int // index of the next insertion
	fullness int
}

func NewCircularQueue[T Int](size T) CircularQueue[T] {
	return CircularQueue[T]{
		values: make([]T, size),
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}
	q.values[q.cursor] = value
	q.cursor = (q.cursor + 1) % len(q.values)
	q.fullness++
	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}
	q.head = (q.head + 1) % len(q.values)
	q.fullness--
	return true
}

func (q *CircularQueue[T]) Front() T {
	if q.Empty() {
		return EmptyValue
	}
	return q.values[q.head]
}

func (q *CircularQueue[T]) Back() T {
	if q.Empty() {
		return EmptyValue
	}
	if q.cursor == 0 {
		return q.values[len(q.values)-1]
	}
	return q.values[q.cursor-1]
}

func (q *CircularQueue[T]) Empty() bool {
	if q.fullness == 0 {
		return true
	}
	return false
}

func (q *CircularQueue[T]) Full() bool {
	return q.fullness == len(q.values)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())
	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))
	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
