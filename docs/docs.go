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
        "/user/login": {
            "post": {
                "tags": [
                    "admin - 用户相关"
                ],
                "summary": "admin - 登陆接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "路径ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "请求体",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserResp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResp"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "tags": [
                    "admin - 用户相关"
                ],
                "summary": "admin - 用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "路径ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UserResp": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "年龄",
                    "type": "integer"
                },
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "gender": {
                    "description": "用户性别",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "nickname": {
                    "description": "用户名",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
