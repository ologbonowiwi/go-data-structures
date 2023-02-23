package datastructures_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	datastructures "github.com/ologbonowiwi/go-data-structures/pkg/data-structures"
)

func TestStack(t *testing.T) {
	// test BuildStack function
	stack := datastructures.BuildStack[int](5)
	assert.Equal(t, 5, stack.Size(), "stack size should be 5")
	assert.Equal(t, -1, stack.TopIndex(), "stack top should be -1")
	assert.Empty(t, stack.Vector(), "stack vector should be empty")

	// test Push function
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 1, stack.Vector()[0], "stack should contain 1 as the first element")
	assert.Equal(t, 2, stack.Vector()[1], "stack should contain 2 as the second element")
	assert.Equal(t, 1, stack.TopIndex(), "stack top should be 1")

	// test Push when stack is full
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6) // should not be added to the stack
	assert.Equal(t, 5, stack.Size(), "stack size should still be 5")
	assert.Equal(t, 4, stack.TopIndex(), "stack top should still be 4")
	assert.Equal(t, 1, stack.Vector()[0], "stack should still contain 1 as the first element")
	assert.Equal(t, 2, stack.Vector()[1], "stack should still contain 2 as the second element")

	// test Pop function
	popped := stack.Pop()
	assert.Equal(t, 5, popped, "popped value should be 5")
	assert.Equal(t, 3, stack.TopIndex(), "stack top should be 3")
	assert.Equal(t, 1, stack.Vector()[0], "stack should still contain 1 as the first element")
	assert.Equal(t, 2, stack.Vector()[1], "stack should still contain 2 as the second element")
	assert.Equal(t, 3, stack.Vector()[2], "stack should contain 3 as the third element")

	// test Peek function
	peeked := stack.Peek()
	assert.Equal(t, 4, peeked, "peeked value should be 4")
	assert.Equal(t, 3, stack.TopIndex(), "stack top should still be 3")

	// test IsEmpty function
	assert.False(t, stack.IsEmpty(), "stack should not be empty")
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	assert.True(t, stack.IsEmpty(), "stack should be empty")
}
