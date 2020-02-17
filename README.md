# builder

`builder` предоставляет сервис для кодирования текста для демонстрации паттерна Builder.

Реализовано два кодировщика:
   1. Побайтово в бинарную форму.
   2. Поруново в бинарную форму.
   
```(go)
type Encoder interface {
	SymbolType() Encoder
	MaxLength() Encoder
	EncodingType () Encoder
	Message() (message, error)
}
```

## Структура проекта
1. `/pkg/product` - интерфейс выходного сообщения.

2.  `/pkg/director ` - director билдера.

3.  `/pkg/builder` - интерфейс builder.

4.  `pkg/concrete_builder/binary_bytes` - кодировщик побайтово в бинарную форму

5.  `pkg/concrete_builder/binary_runes` - кодировщик поруново в бинарную форму

## Запуск тестов
`go test ./...`
   
