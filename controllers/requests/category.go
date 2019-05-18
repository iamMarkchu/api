package requests

type CategoryStoreRequest struct {
	Name string	`form:"name" valid:"Required"`
	ParentCate int  `form:"parent_cate_id" valid:"Required"`
}
