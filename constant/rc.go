package constant

const (
	RCSuccess    = "00"
	RCSuccessMsg = "Success"

	GeneralErrorCode = "99"
	GeneralErrorDesc = "General Error"

	NotFoundErrorCode     = "DN"
	NotFoundErrorCodeDesc = "Data Not Found"

	RCValidationError        = "99"
	RCValidationErrorDesc    = "Validation Error"
	RCDatabaseError          = "DE"
	RCDatabaseErrorDesc      = "something went wrong with Database (check query or connection)"
	RCPasswordNotMatch       = "NM"
	RCPasswordNotMatchDesc   = "Sorry, password not match"
	RCUserBlockError         = "UB"
	RCUserBlockErrorDesc     = "Sorry, user blocked"
	RCDataAlreadyExist       = "DX"
	RCUserCannotLogin        = "CL"
	RCDataCannotLoginDesc    = "sorry user cannot login"
	RCOtpMaxRequestError     = "MR"
	RCOtpMaxRequestErrorDesc = "sorry otp max request"
	RCUserNotBlockError      = "NB"
	RCUserNotBlockErrorDesc  = "Sorry, user blocked"
	RCNoCookie               = "NC"
	RCTokenNotValid          = "JNF"
)
