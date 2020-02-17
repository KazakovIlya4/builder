package binarybytes

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	textRunesShortTest       = "世"
	textBinaryBytesShortTest = "111001001011100010010110"
	maxBytesLength           = 256
)

func TestBinaryRunes_Encode(t *testing.T) {
	builder := NewBuilder(textRunesShortTest)
	builder.EncodingType()
	builder.MaxLength()
	builder.SymbolType()
	message, err := builder.Message()
	assert.Equal(t, nil, err)
	assert.Equal(t, message.Encoding(), "binary")
	assert.Equal(t, message.Symbol(), "byte")
	assert.Equal(t, textBinaryBytesShortTest, message.Get())
}

func TestBinaryRunes_MaxFail(t *testing.T) {
	stringBuilder := strings.Builder{}

	for i := 0; i < maxBytesLength+1; i++ {
		stringBuilder.WriteByte('a')
	}

	builder := NewBuilder(stringBuilder.String())
	builder.EncodingType()
	builder.MaxLength()
	builder.SymbolType()
	_, err := builder.Message()
	assert.Equal(t, errors.As(err, &errTextTooLong), true)
}
