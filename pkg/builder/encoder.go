package builder

type message = interface {
	Get() string
	Symbol() string
	Encoding() string
	GetOriginal() string
}

// Encoder represents construction interface for builder pattern
type Encoder interface {
	SymbolType() Encoder
	MaxLength() Encoder
	EncodingType () Encoder
	Message() (message, error)
}
