package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_vector_Empty(t *testing.T) {
	vec := New()
	assert.True(t, vec.Empty())

	vecNotEmpty := NewWithSize(1)
	assert.False(t, vecNotEmpty.Empty())
}

func Test_vector_Size(t *testing.T) {
	vec := New()
	assert.Equal(t, 0, vec.Size())

	vecNotEmpty := NewWithSize(1)
	assert.Equal(t, 1, vecNotEmpty.Size())
}

func Test_vector(t *testing.T) {
	vec := New()
	vec.PushBack(1)
	vec.PushBack(2)
	vec.PushBack(3)

	assert.Equal(t, 1, vec.At(0))
	assert.Equal(t, 2, vec.At(1))
	assert.Equal(t, 3, vec.At(2))
	assert.Equal(t, 1, vec.Front())
	assert.Equal(t, 3, vec.Back())

	capacity := vec.Capacity()
	vec.Clear()
	assert.Equal(t, 0, vec.Size())
	assert.Equal(t, capacity, vec.Capacity())
}

func Test_vector_Insert(t *testing.T) {
	vec := New()
	vec.PushBack(1)
	vec.PushBack(2)
	vec.PushBack(3)

	vec.Insert(0, 12)
	assert.True(t, validateVector(vec, []int{12, 1, 2, 3}))
	vec.Insert(2, 23)
	assert.True(t, validateVector(vec, []int{12, 1, 23, 2, 3}))
	vec.Insert(5, 56)
	assert.True(t, validateVector(vec, []int{12, 1, 23, 2, 3, 56}))
}

func validateVector(vec *vector, values []int) bool {
	if vec.Size() != len(values) {
		return false
	}
	for i, v := range values {
		if vec.At(i).(int) != v {
			return false
		}
	}
	return true
}

func Test_vector_Erase(t *testing.T) {
	vec := New()
	vec.PushBack(1)
	vec.PushBack(2)
	vec.PushBack(3)
	vec.PushBack(4)

	vec.Erase(3)
	assert.True(t, validateVector(vec, []int{1, 2, 3}))
	vec.Erase(0)
	assert.True(t, validateVector(vec, []int{2, 3}))
}

func Test_vector_PopBack(t *testing.T) {
	vec := New()
	vec.PushBack(1)
	vec.PushBack(2)
	vec.PushBack(3)
	vec.PushBack(4)

	initValues := []int{1, 2, 3, 4}
	assert.True(t, validateVector(vec, initValues))

	for i := 0; i < 4; i++ {
		vec.PopBack()
		assert.True(t, validateVector(vec, initValues[:3-i]))
	}
}
