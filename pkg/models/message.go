package models

// Message represents product of text encoder
type Message struct {
	Original   string
	Text       string
	SymbolType string
	Type       string
}

// Get returns encoded text
func (m *Message) Get() (text string) {
	text = m.Text
	return
}

// Symbol returns type of symbol
func (m *Message) Symbol() (symbolType string) {
	symbolType = m.SymbolType
	return
}

// Encoding returns encoded text
func (m *Message) Encoding() (encoding string) {
	encoding = m.Type
	return
}

// GetOriginal returns unencoded text
func (m *Message) GetOriginal() (original string) {
	original = m.Original
	return
}
