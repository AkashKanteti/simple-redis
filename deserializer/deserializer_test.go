package deserializer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleString(t *testing.T) {
	t.Run("when proper string is given, avg case", func(t *testing.T) {
		result, err := deserialize(`+OK\r\n`)
		assert.Equal(t, "OK", result)
		assert.NoError(t, err)
	})

	t.Run("when empty string is given", func(t *testing.T) {
		result, err := deserialize(`+\r\n`)
		assert.Equal(t, "", result)
		assert.NoError(t, err)
	})

	t.Run("when CRLF delimiter is not provided", func(t *testing.T) {
		result, err := deserialize(`+ok\r`)
		assert.Equal(t, "", result)
		assert.Error(t, err)
	})
}

func TestInteger(t *testing.T) {
	t.Run("when positive integer is given, avg case", func(t *testing.T) {
		result, err := deserialize(`:123\r\n`)
		assert.Equal(t, 123, result)
		assert.NoError(t, err)
	})

	t.Run("when positive integer is given, avg case", func(t *testing.T) {
		result, err := deserialize(`:-123\r\n`)
		assert.Equal(t, -123, result)
		assert.NoError(t, err)
	})

	t.Run("when positive integer is given, avg case", func(t *testing.T) {
		_, err := deserialize(`:\r\n`)
		assert.Error(t, err)
	})
}
