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

func TestBulkString(t *testing.T) {
	t.Run("when empty string", func(t *testing.T) {
		result, err := deserialize(`$0\r\n\r\n`)
		assert.Equal(t, "", result)
		assert.NoError(t, err)
	})

	t.Run("when null string", func(t *testing.T) {
		result, err := deserialize(`$-1\r\n`)
		assert.Equal(t, nil, result)
		assert.NoError(t, err)
	})

	t.Run("when real string", func(t *testing.T) {
		result, err := deserialize(`$11\r\nhello world\r\n`)
		assert.Equal(t, "hello world", result)
		assert.NoError(t, err)
	})
}

func TestArrays(t *testing.T) {
	t.Run("when empty array", func(t *testing.T) {
		result, err := deserialize(`*0\r\n\r\n`)
		assert.Equal(t, "", result)
		assert.NoError(t, err)
	})

	t.Run("when null array", func(t *testing.T) {
		result, err := deserialize(`*-1\r\n`)
		assert.Equal(t, nil, result)
		assert.NoError(t, err)
	})

	t.Run("when real array", func(t *testing.T) {
		result, err := deserialize(`*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n`)
		assert.Equal(t, []string{"echo", "hello world"}, result)
		assert.NoError(t, err)
	})
}
