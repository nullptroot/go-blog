package service

import imageser "go-blog/service/image_ser"

type ServiceGroup struct {
	ImageService imageser.ImageService
}

var ServiceApp = new(ServiceGroup)
