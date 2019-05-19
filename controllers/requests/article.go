package requests

type ArticleStoreRequest struct {
	Title       string `form:"title" valid:"Required"`
	Description string `form:"description" valid:"Required"`
	ImageUrl    string `form:"image_url"`
	CategoryId  int    `form:"category_id"`
}

type ArticleIndexRequest struct {
	IndexRequest
}
