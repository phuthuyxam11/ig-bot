package request

type VerifyAccBody struct {
	Code string `uri:"code" form:"code" binding:"required"`
	Id   string `uri:"id" form:"id" binding:"required"`
}
