package common

// 更新模型
// 1, 全量更新：对象的替换
// 2, 部分更新：(old obj)Patch机制 --> new obj --> save

// 更新模式。这是枚举
type UpdateMode int 

const (
	// 全量更新
	UPDATE_MODE_PUT UpdateMode = iota
	// 部分更新
	UPDATE_MODE_PATCH
)


// 分页参数
type Pager struct {
	// 分页大小，一页多少个
	PageSize int
	// 当前页，查询哪一页的数据
	PageNumber int
}
