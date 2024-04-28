package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 //用户
	PermissionVisitor     Role = 3 //参观者
	PermissionDisableUser Role = 4 //被禁言用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisitor:
		str = "参观者"
	case PermissionDisableUser:
		str = "被禁言用户"
	default:
		str = "其他"
	}
	return str
}
