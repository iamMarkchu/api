package requests

type CategoryStoreRequest struct {
	Name       string `form:"name" valid:"Required"`
	ParentCate int    `form:"parent_id"`
}

//type CategoryIndexRequest struct {
//	IndexRequest
//}
