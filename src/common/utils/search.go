package utils

// 搜索条件主体
type Search struct {
	Page    int
	Size    int
	Limit   int
	Sort    interface{}
	GroupBy interface{}
	Where   interface{}
}