package dll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing length property of the dll
func TestLength1(t *testing.T) {
	/*
		Objective:
			Test if the `Length` field is working properly
			while performing operations
	*/

	list := CreateDoublyLinkedListEmpty()
	list.Append(CreateNode("name", "Rajab", -1, nil, nil))
	assert.Equal(t, uint32(1), list.LengthC)
	list.Prepend(CreateNode("age", "16", -1, nil, nil))
	assert.Equal(t, uint32(2), list.LengthC)
	list.Pop()
	assert.Equal(t, uint32(1), list.LengthC)
	list.Pop()
	assert.Equal(t, uint32(0), list.LengthC)
	list.Pop()
	assert.Equal(t, uint32(0), list.LengthC)
}

func TestLength2(t *testing.T) {
	/*
		Objective:
			Test if the `Length` field is working properly
			while performing operations
	*/

	list := CreateDoublyLinkedListEmpty()
	list.Prepend(CreateNode("age", "16", -1, nil, nil))
	assert.Equal(t, uint32(1), list.LengthC)
}
