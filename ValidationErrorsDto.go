package response

type ValidationErrorsDto struct {
	Message string
	Errors  map[string]string
	Code    int
}

func (dto ValidationErrorsDto) GetMessage() string {
	return dto.Message
}

func (dto ValidationErrorsDto) GetErrors() map[string]string {
	return dto.Errors
}

func (dto ValidationErrorsDto) GetCode() int {
	return dto.Code
}
