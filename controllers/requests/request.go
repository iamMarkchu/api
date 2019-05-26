package requests

type Request interface {}

type IndexRequest struct {
	Status int8 `form:"status"`
	Page   int  `form:"page"`
	Limit  int  `form:"limit"`
}
