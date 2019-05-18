package controllers

import "api/controllers/requests"

type CategoryController struct {
	ApiController
}

func (c *CategoryController) URLMapping()  {
	c.Mapping("Index", c.Index)
	c.Mapping("Store", c.Store)
}

// @router / [get]
func (c *CategoryController) Index()  {

}

// @router / [post]
func (c *CategoryController) Store()  {
	r := requests.CategoryStoreRequest{}
	c.ValidateRequest(r)
	c.JsonReturn("创建类别接口", "", 200)
}


