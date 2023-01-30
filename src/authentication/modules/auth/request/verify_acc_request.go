package request

type VerifyAccBody struct {
	Code string `uri:"code" form:"code" binding:"required" message_key:"validate.verifyAcc.code.require, validate.verifyAcc.code.exist"`
	Id   string `uri:"id" form:"id" binding:"required,userexist" message_key:"validate.verifyAcc.id.require, validate.verifyAcc.id.userexist"`
}
