package simple_redis

import (
	"github.com/AkashKanteti/simple-redis/serializer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleString(t *testing.T) {
	t.Run("when proper string is given, avg case", func(t *testing.T) {
		result, err := serializer.Serialize(`+OK\r\n`)
		assert.Equal(t, "OK", result)
		assert.NoError(t, err)
	})

	t.Run("when empty string is given", func(t *testing.T) {
		result, err := serializer.Serialize(`+\r\n`)
		assert.Equal(t, "", result)
		assert.NoError(t, err)
	})

	t.Run("when CRLF delimiter is not provided", func(t *testing.T) {
		result, err := serializer.Serialize(`+ok\r`)
		assert.Equal(t, "", result)
		assert.Error(t, err)
	})
}
