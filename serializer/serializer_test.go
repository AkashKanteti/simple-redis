package serializer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeserializer(t *testing.T) {
	t.Run("simple string", func(t *testing.T) {
		assert.Equal(t, "+hello world\r\n", deserializeString("hello world"))
	})
}
