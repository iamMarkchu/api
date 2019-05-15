package requests

type ArticleStoreRequest struct {
	Title       string `form:"title" valid:"Required;Match(/^Bee.*/)"`
	Description string `form:"description" valid:"Required"`
	ImageUrl    string `form:"imageurl"`
}
