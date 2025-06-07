package serializer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializer(t *testing.T) {
	t.Run("simple string", func(t *testing.T) {
		assert.Equal(t, "+hello world\r\n", deserializeString("hello world"))
	})

	t.Run("simple error", func(t *testing.T) {
		assert.Equal(t, "-Error message\r\n", deserializeError("Error message"))
	})

	t.Run("simple integer", func(t *testing.T) {
		assert.Equal(t, ":123\r\n", deserializeInteger(123))
	})
}
