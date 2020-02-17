package binaryrunes

import "github.com/stretchr/testify/mock"

// Mock of binary runes builder
type Mock struct {
	mock.Mock
}

// SymbolType ...
func (m *Mock) SymbolType() {
	m.Called()
}

// MaxLength ...
func (m *Mock) MaxLength() {
	m.Called()
}

// EncodingType ...
func (m *Mock) EncodingType() {
	m.Called()
}

// Message ...
func (m *Mock) Message() (encodedMessage message, err error) {
	args := m.Called()
	err = args.Error(1)
	if err == nil {
		encodedMessage = args.Get(0).(message)
	}
	return
}
