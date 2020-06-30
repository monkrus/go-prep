package main

import (
	"testing"
)

//Attempt 1
func TestWrongType(t *testing.T) {
	result := WrongType(error)
	expected := (nil)
	if result != expected {
		t.Errorf("TestWrongType() test returned an unexpected result: got %v want %v", result, expected)
	}
}

/* Attempt 2
func TestExampleTestSuite(t *testing.T) {
suite.Run(t, new(ExampleTestSuite))
}

func TestWrongType(t *testing.T) {

	assert.Equal(t)

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, object)

	// assert for not nil (good when you expect something)
	if assert.NotNil(t, object) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", object.Value)

	}

}
*/
