package errno


var(

	//common errors
	OK	= &Errno{Code:0,Message:"OK"}
	InternalServerError = &Errno{Code:10001,Message:"Internal server error."}
	ErrBind = &Errno{Code:10001,Message:"Error occurred while binding the request body to the struct."}

	ErrValidation =&Errno{20001,"Validation failed."}
	ErrDatabase = &Errno{20002,"Database error."}
	ErrToken = &Errno{20003,"Error occurred while signing the JSON web token."}


	//user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
)