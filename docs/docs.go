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
        "/v1/appData/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "App data"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.AppDataListResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "User register",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.UserRegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorMessagePrototype"
                        }
                    }
                }
            }
        },
        "/v2/appData/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "App data"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.AppDataListResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.AppDataListResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "items": {},
                "pagination": {
                    "$ref": "#/definitions/utils.PaginationObject"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "domain.UserRegister": {
            "type": "object",
            "properties": {
                "firstName": {
                    "description": "ชื่อ",
                    "type": "string"
                },
                "lastName": {
                    "description": "นามสกุล",
                    "type": "string"
                },
                "midName": {
                    "description": "ชื่อกลาง",
                    "type": "string"
                }
            }
        },
        "domain.UserRegisterRequest": {
            "type": "object",
            "required": [
                "firstName",
                "lastName"
            ],
            "properties": {
                "firstName": {
                    "description": "ชื่อ",
                    "type": "string"
                },
                "lastName": {
                    "description": "นามสกุล",
                    "type": "string"
                },
                "midName": {
                    "description": "ชื่อกลาง",
                    "type": "string"
                }
            }
        },
        "domain.UserRegisterResponse": {
            "type": "object",
            "properties": {
                "apiVersion": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "properties": {
                        "description": {
                            "type": "string"
                        },
                        "item": {
                            "$ref": "#/definitions/domain.UserRegister"
                        },
                        "title": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "utils.ErrorMessagePrototype": {
            "type": "object",
            "properties": {
                "apiVersion": {
                    "type": "string"
                },
                "error": {
                    "$ref": "#/definitions/utils.errorResponse"
                }
            }
        },
        "utils.PaginationObject": {
            "type": "object",
            "properties": {
                "currentPage": {
                    "description": "หน้าปัจจุบัน",
                    "type": "integer"
                },
                "lastPage": {
                    "description": "หน้าสุดท้าย",
                    "type": "integer"
                },
                "nextPage": {
                    "description": "หน้าถัดไป",
                    "type": "integer"
                },
                "perPage": {
                    "description": "จำนวนต่อหน้า",
                    "type": "integer"
                },
                "previousPage": {
                    "description": "หน้าก่อนหน้า",
                    "type": "integer"
                },
                "total": {
                    "description": "จำนวนรายการทั้งหมด",
                    "type": "integer"
                }
            }
        },
        "utils.errorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "security": [
        {
            "bearer": []
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "base code",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
