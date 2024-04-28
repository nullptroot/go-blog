package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 // QQ
	SignGitee SignStatus = 2 //Gitee
	SigEmail  SignStatus = 3 //Email
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "Gitee"
	case SigEmail:
		str = "Email"
	default:
		str = "其他"
	}
	return str
}
