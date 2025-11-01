package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializer(t *testing.T) {
	t.Run("simple string", func(t *testing.T) {
		assert.Equal(t, "+hello world\r\n", SerializeString("hello world"))
	})

	t.Run("simple error", func(t *testing.T) {
		assert.Equal(t, "-Error message\r\n", serializeError("Error message"))
	})

	t.Run("simple integer", func(t *testing.T) {
		assert.Equal(t, ":123\r\n", serializeInteger(123))
	})

	t.Run("simple bulk string", func(t *testing.T) {
		assert.Equal(t, "$3\r\nget\r\n", serializeBulkStrings("get"))
	})

	t.Run("simple integer", func(t *testing.T) {
		assert.Equal(t, "*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n", SerializeArray([]string{"echo", "hello world"}))
	})
}
