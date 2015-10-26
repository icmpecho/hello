package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetReturnHelloName(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Hello foo\n", Greet("foo"))
}
