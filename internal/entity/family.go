package entity

type Family struct {
	ID       int64
	Name     string
	Lft      int64
	Rgt      int64
	ParentId int64
}
