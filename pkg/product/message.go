package product

// Message represents interface of product of text encoder
type Message interface {
	Get() string
	Symbol() string
	Encoding() string
	GetOriginal() string
}
