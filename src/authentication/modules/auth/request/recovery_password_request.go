package request

type RecoveryPasswordVerifyBody struct {
	Email string `json:"email" uri:"email" form:"email" binding:"required,email,db_exist=user_login_data.email" message_key:"validate.recoveryAcc.email.require, validate.recoveryAcc.email.type, validate.recoveryAcc.email.exist"`
}
