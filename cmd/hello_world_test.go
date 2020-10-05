package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T){
	s := "Hello world"
	assert.Equal(t,s, "Hello world", "The two phrases should be the same")
}