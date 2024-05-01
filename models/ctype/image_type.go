package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1 // ImageType
	QiNiu ImageType = 2 //ImageType
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "ImageType"
	case QiNiu:
		str = "ImageType"
	}
	return str
}
