package service

import (
	imageser "go-blog/service/image_ser"
	userser "go-blog/service/user_ser"
)

type ServiceGroup struct {
	ImageService imageser.ImageService
	UserService  userser.UserService
}

var ServiceApp = new(ServiceGroup)
