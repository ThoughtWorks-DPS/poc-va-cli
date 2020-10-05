package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T){
	s := helloWorld()
	assert.Equal(t,"Hello world", s, "The two phrases should be the same")
}