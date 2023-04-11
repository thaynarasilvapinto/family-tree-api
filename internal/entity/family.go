package entity

type Family struct {
	Id         int64
	Name       string
	ParentId1  *int64
	ParentId2  *int64
	Generation *int
}
