package model

// 角色
type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// 权限
type Permission struct {
	ID       int64  `json:"id"`
	Action   string `json:"action"`   // "read", "write", "delete" 等
	Resource string `json:"resource"` // "article", "user", "order" 等
}
