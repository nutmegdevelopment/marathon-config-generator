package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetFlagStringSlice(t *testing.T) {
	var ss stringslice
	ss.Set("I am item one")
	ss.Set("I am item two")

	assert.Equal(t, "I am item one", ss[0])
	assert.Equal(t, "I am item two", ss[1])
}

func TestIterateFlagStringSlice(t *testing.T) {
	var ss stringslice
	ss.Set("I am item one")
	ss.Set("I am item two")
	ss.Set("I am item three")

	for i, f := range ss {
		if i == 0 {
			assert.Equal(t, "I am item one", f)
		} else if i == 1 {
			assert.Equal(t, "I am item two", f)
		} else if i == 2 {
			assert.Equal(t, "I am item three", f)
		}

	}
}
