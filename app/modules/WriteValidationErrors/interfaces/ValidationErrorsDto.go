package interfaces

type ValidationErrorsDto interface {
	GetMessage() string
	GetErrors() map[string]string
	GetCode() int
}
