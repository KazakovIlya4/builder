package director

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	buildermock "builder/pkg/builder"
	"builder/pkg/models"
)

var errTextTooLong = fmt.Errorf("input symbolType is too long")

func Test_BuildMessageSucces(t *testing.T) {

	director := NewEncoderService()

	builder := &buildermock.Mock{
		Mock: mock.Mock{},
	}

	builder.On("Message").Return(&models.Message{
		Original:   "test1",
		Text:       "test2",
		SymbolType: "test3",
		Type:       "test4",
	}, nil)
	builder.On("EncodingType").Return()
	builder.On("MaxLength").Return()
	builder.On("SymbolType").Return()

	message, err := director.BuildMessage(builder)
	assert.Equal(t, nil, err)
	assert.Equal(t, message.GetOriginal(), "test1")
	assert.Equal(t, message.Get(), "test2")
	assert.Equal(t, message.Symbol(), "test3")
	assert.Equal(t, message.Encoding(), "test4" )
}

func Test_BuildMessageFail(t *testing.T) {

	director := NewEncoderService()

	builder := &buildermock.Mock{
		Mock: mock.Mock{},
	}

	builder.On("Message").Return(&models.Message{}, errTextTooLong)
	builder.On("EncodingType").Return()
	builder.On("MaxLength").Return()
	builder.On("SymbolType").Return()

	_, err := director.BuildMessage(builder)
	assert.Equal(t, errors.As(err, &errTextTooLong), true)
}
