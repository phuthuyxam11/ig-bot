package request

type RecoveryPasswordVerifyBody struct {
	Email string `uri:"email" form:"email" binding:"required" message_key:" validate.verifyAcc.code.require, validate.verifyAcc.code.exist"`
}
