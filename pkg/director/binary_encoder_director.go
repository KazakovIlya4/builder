package director

type message = interface {
	GetOriginal() string
	Get() string
	Symbol() string
	Encoding() string
}

type encoder interface {
	SymbolType()
	MaxLength()
	EncodingType ()
	Message() (message, error)
}

// EncoderDirectorService represents director interface
type EncoderDirectorService interface {
	BuildMessage(builder encoder) (message, error)
}

type director struct {}

func (d *director) BuildMessage(builder encoder) (encodedMessage message, err error) {
	builder.EncodingType()
	builder.MaxLength()
	builder.SymbolType()
	encodedMessage, err = builder.Message()
	return
}

// NewEncoderService returns new instance of encoder director
func NewEncoderService() EncoderDirectorService {
	return &director{}
}
