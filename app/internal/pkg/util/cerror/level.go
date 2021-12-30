package cerror

type customErrorLevel string

const (
	ErrorLevel customErrorLevel = customErrorLevel("Error")
	WarnLevel                   = customErrorLevel("Warn")
	PanicLevel                  = customErrorLevel("Panic")
	FatalLevel                  = customErrorLevel("Fatal")
)

func (cel customErrorLevel) ToString() string {
	return string(cel)
}
