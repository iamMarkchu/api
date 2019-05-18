package models

type Category struct {
	Id         int       `orm:"auto"`
	ParentCate *Category `orm:"rel(fk);null;on_delete(set_null)" json:"parent_cate"`
	Author     *User     `orm:"rel(fk);null;on_delete(set_null)" json:"author"`
	Name       string    `orm:"size(30);description(类别名称)" json:"name"`
	BaseModel
}
