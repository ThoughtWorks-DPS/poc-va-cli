package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T){
	s := get()

	assert.Equal(t, "2001 OK", s, "Status code should equal 200 OK" )
}