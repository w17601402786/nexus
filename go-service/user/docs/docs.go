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
        "/get_user": {
            "get": {
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "description": "登录并获取 token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户服务"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "账号（邮箱、手机、用户名）",
                        "name": "UsernameOrEmail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "Password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/register": {
            "post": {
                "responses": {}
            }
        },
        "/third_party_login": {
            "post": {
                "responses": {}
            }
        },
        "/update_user_profile": {
            "post": {
                "responses": {}
            }
        }
    },
    "tags": [
        {
            "name": "用户服务"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "用户微服务",
	Description:      "用户微服务",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}