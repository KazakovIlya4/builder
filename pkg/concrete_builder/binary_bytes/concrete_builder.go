package binarybytes

import (
	"bytes"
	"fmt"

	"builder/pkg/models"
)

var errTextTooLong = fmt.Errorf("input symbolType is too long")

type message = interface {
	Get() string
	Symbol() string
	Encoding() string
	GetOriginal() string
}

type builder interface {
	SymbolType()
	MaxLength()
	EncodingType()
	Message() (message, error)
}

type binaryBytes struct {
	text        string
	symbolType  string
	maxBytes    uint32
	messageType string
}

// SymbolType sets info about character type
func (b *binaryBytes) SymbolType() {
	b.symbolType = "byte"
}

// MaxLength sets maximum allowed input string length in bytes
func (b *binaryBytes) MaxLength() {
	b.maxBytes = 256
}

// EncodingType sets info about encoding type
func (b *binaryBytes) EncodingType() {
	b.messageType = "binary"
}

// Message encodes input text and returns it with additional info
func (b *binaryBytes) Message() (encodedMessage message, err error) {
	if len(b.text) > int(b.maxBytes) {
		err = fmt.Errorf("%w: max length is %d", errTextTooLong, b.maxBytes)
		return
	}
	var buffer bytes.Buffer
	for i := 0; i < len(b.text); i++ {
		fmt.Fprintf(&buffer, "%b", b.text[i])
	}
	encodedMessage = &models.Message{
		Original:   b.text,
		Text:       buffer.String(),
		SymbolType: b.symbolType,
		Type:       b.messageType,
	}
	return
}

// NewBuilder returns new instance of binary bytes encoder concrete builder implementation
func NewBuilder(text string) builder {
	return &binaryBytes{
		text: text,
	}
}
