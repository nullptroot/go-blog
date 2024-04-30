package res

type ErrorCode int

const (
	SettingsError = 1001
	ArgumentError = 1002
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误·",
	}
)
