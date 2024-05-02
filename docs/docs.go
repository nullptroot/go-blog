// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/adverts": {
            "get": {
                "description": "广告列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "广告列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "模糊查询使用的关键字",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "一页限制多少记录",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "从第几页开始，后续会生成offset",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-models_AdvertModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "创建广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "创建广告",
                "parameters": [
                    {
                        "description": "表示多个参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/advertapi.AdvertRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "批量删除广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "批量删除广告",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "广告id列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/adverts/{id}": {
            "put": {
                "description": "更新广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "更新广告",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "广告的一些参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/advertapi.AdvertRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/images": {
            "get": {
                "description": "图片列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "模糊查询使用的关键字",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "一页限制多少记录",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "从第几页开始，后续会生成offset",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-models_BannerModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "更新图片",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "更新图片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "图片的一些参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/imagesapi.ImageUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "上传图片",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "上传图片",
                "parameters": [
                    {
                        "type": "file",
                        "description": "上传的图片文件",
                        "name": "data",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/imageser.FileUploadResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "删除图片",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "删除图片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "图片id列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/images_names": {
            "get": {
                "description": "图片名称列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片名称列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/imagesapi.ImagesResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/menus": {
            "get": {
                "description": "菜单列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/menuapi.MenuResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "更新菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "更新菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "图片的一些参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/menuapi.MenuRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "添加菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "添加菜单",
                "parameters": [
                    {
                        "description": "表示多个参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/menuapi.MenuRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "删除菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "图片id列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/menus/{id}": {
            "get": {
                "description": "菜单细节",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单细节",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/menuapi.MenuResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/menus_names": {
            "get": {
                "description": "菜单名称列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单名称列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/menuapi.MenuNameResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "advertapi.AdvertRequest": {
            "type": "object",
            "required": [
                "href",
                "images",
                "title"
            ],
            "properties": {
                "href": {
                    "description": "跳转链接",
                    "type": "string"
                },
                "images": {
                    "description": "图片",
                    "type": "string"
                },
                "is_show": {
                    "description": "是否展示",
                    "type": "boolean"
                },
                "title": {
                    "description": "显示的标题",
                    "type": "string"
                }
            }
        },
        "ctype.ImageType": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-comments": {
                "Local": "ImageType",
                "QiNiu": "ImageType"
            },
            "x-enum-varnames": [
                "Local",
                "QiNiu"
            ]
        },
        "imagesapi.ImageUpdateRequest": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "imagesapi.ImagesResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "图片名称",
                    "type": "string"
                },
                "path": {
                    "description": "图片路径",
                    "type": "string"
                }
            }
        },
        "imageser.FileUploadResponse": {
            "type": "object",
            "properties": {
                "file_name": {
                    "description": "文件名",
                    "type": "string"
                },
                "is_success": {
                    "description": "是否上传成功",
                    "type": "boolean"
                },
                "msg": {
                    "description": "消息",
                    "type": "string"
                }
            }
        },
        "menuapi.Banner": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "menuapi.ImageSort": {
            "type": "object",
            "properties": {
                "image_id": {
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                }
            }
        },
        "menuapi.MenuNameResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "menuapi.MenuRequest": {
            "type": "object",
            "required": [
                "path",
                "sort",
                "title"
            ],
            "properties": {
                "abstract": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "abstract_time": {
                    "description": "切换的时间，单位秒",
                    "type": "integer"
                },
                "banner_time": {
                    "description": "切换的时间，单位秒",
                    "type": "integer"
                },
                "image_sort_list": {
                    "description": "具体图片的顺序",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/menuapi.ImageSort"
                    }
                },
                "path": {
                    "type": "string"
                },
                "slogan": {
                    "type": "string"
                },
                "sort": {
                    "description": "菜单的序号",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "menuapi.MenuResponse": {
            "type": "object",
            "properties": {
                "abstract": {
                    "description": "简介",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "abstract_time": {
                    "description": "简介的切换时间",
                    "type": "integer"
                },
                "banner_time": {
                    "description": "菜单图片的切换时间 为 0 表示不切换",
                    "type": "integer"
                },
                "banners": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/menuapi.Banner"
                    }
                },
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "slogan": {
                    "description": "slogan",
                    "type": "string"
                },
                "sort": {
                    "description": "菜单的顺序",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "models.AdvertModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "href": {
                    "description": "跳转链接",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "images": {
                    "description": "图片",
                    "type": "string"
                },
                "is_show": {
                    "description": "是否展示",
                    "type": "boolean"
                },
                "title": {
                    "description": "显示的标题",
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "models.BannerModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "hash": {
                    "description": "图片的hash值，用于判断重复图片",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "image_type": {
                    "description": "图片的类型， 本地还是七牛",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ctype.ImageType"
                        }
                    ]
                },
                "name": {
                    "description": "图片名称",
                    "type": "string"
                },
                "path": {
                    "description": "图片路径",
                    "type": "string"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "models.RemoveRequest": {
            "type": "object",
            "properties": {
                "id_list": {
                    "description": "要删除的记录id",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "res.ListResponse-models_AdvertModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "$ref": "#/definitions/models.AdvertModel"
                }
            }
        },
        "res.ListResponse-models_BannerModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "$ref": "#/definitions/models.BannerModel"
                }
            }
        },
        "res.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "10.0.0.3:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "gvb-blog文档",
	Description:      "API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
