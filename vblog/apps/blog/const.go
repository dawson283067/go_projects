package blog

type Status int

const (
	// 草稿
	STATUS_DRAFT Status = iota
	// 已发布
	STATUS_PUBLISHED
)