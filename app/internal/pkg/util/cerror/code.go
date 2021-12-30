package cerror

type customErrorCode string

const (
	InterServerErrorCode          customErrorCode = customErrorCode("Internal Server Error")
	InvalidArgumentErrorCode                      = customErrorCode("Invalid Argument Error")
	NotFoundErrorCode                             = customErrorCode("Not Found Error")
	AlreadyExistsErrorCode                        = customErrorCode("Already Exists Error")
	NotImplementedErrorCode                       = customErrorCode("Not Implemented Error")
	AuthenticationErrorCode                       = customErrorCode("Authentication Error")
	UnauthorizedErrorCode                         = customErrorCode("Unauthozied Error")
	InvalidTokenErrorCode                         = customErrorCode("Invalid Token Error")
	UnregisterdUserErrorCode                      = customErrorCode("Unregistered User Error")
	AlreadyRegisterdUserErrorCode                 = customErrorCode("Already Registered User Error")
)

func (cec customErrorCode) ToString() string {
	return string(cec)
}
