// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://cihantaylan.com",
            "email": "cihantaylan@cihantaylan.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ldap/find": {
            "post": {
                "description": "EG;\u003cbr\u003eByRG81IDDPQFY9+9dSaWFKIA3Xp1vZhrpCjCg4XXR7gnNxLM9SvgTK1PFKMrsdE5s4mNRSIo8qJhzeZAdMi5zQfAhJOV8FDdmEs=\u003cbr\u003eSearchBase: \"dc=example,dc=com\"\u003cbr\u003eSearchFilter: \"(objectClass=person)\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ldap"
                ],
                "summary": "Find",
                "operationId": "Find",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Your Auth Token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.FindRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.FindResponse"
                        }
                    }
                }
            }
        },
        "/ldap/login": {
            "post": {
                "description": "EG;\u003cbr\u003eLdapURL: ldap://ldap.forumsys.com:389\u003cbr\u003eBindDN: cn=read-only-admin,dc=example,dc=com\u003cbr\u003eBindPassword: password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ldap"
                ],
                "summary": "Login",
                "operationId": "Login",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.FindRequest": {
            "type": "object",
            "required": [
                "search_base",
                "search_filter"
            ],
            "properties": {
                "attributes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "search_base": {
                    "type": "string"
                },
                "search_filter": {
                    "type": "string"
                }
            }
        },
        "controllers.FindResponse": {
            "type": "object",
            "required": [
                "data",
                "message",
                "status"
            ],
            "properties": {
                "data": {
                    "type": "object",
                    "required": [
                        "session_token"
                    ],
                    "properties": {
                        "session_token": {
                            "type": "string"
                        }
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "controllers.LoginRequest": {
            "type": "object",
            "required": [
                "bind_dn",
                "bind_password",
                "ldap_url"
            ],
            "properties": {
                "bind_dn": {
                    "type": "string"
                },
                "bind_password": {
                    "type": "string"
                },
                "ldap_url": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginResponse": {
            "type": "object",
            "required": [
                "data",
                "message",
                "status"
            ],
            "properties": {
                "data": {
                    "type": "object",
                    "required": [
                        "session_token"
                    ],
                    "properties": {
                        "session_token": {
                            "type": "string"
                        }
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8088",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Go Ldap Rest API",
	Description:      "This is a go ldap rest API Documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
