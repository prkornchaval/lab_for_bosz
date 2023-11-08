// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/v1": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "create customer",
                "parameters": [
                    {
                        "description": "Create customer",
                        "name": "create-customer-request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.CreateCustomerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CreateCustomerResponse"
                        }
                    }
                }
            }
        },
        "/v1/customer-address": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "create customer and address",
                "parameters": [
                    {
                        "description": "Create customer",
                        "name": "create-customer-request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.CreateCustomerAddressRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CreateCustomerResponse"
                        }
                    }
                }
            }
        },
        "/v1/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "get customer by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CustomerRow"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Address": {
            "type": "object",
            "properties": {
                "district": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "subdistrict": {
                    "type": "string"
                }
            }
        },
        "model.CreateCustomerAddressRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/model.Address"
                },
                "created_address_error": {
                    "type": "boolean"
                },
                "created_by": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "mobile_no": {
                    "type": "string"
                },
                "the1_member_id": {
                    "type": "string"
                },
                "the1_mobile_no": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "model.CreateCustomerRequest": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "mobile_no": {
                    "type": "string"
                },
                "the1_member_id": {
                    "type": "string"
                },
                "the1_mobile_no": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "model.CreateCustomerResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.CustomerRow": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "mobile_no": {
                    "type": "string"
                },
                "the1_member_id": {
                    "type": "string"
                },
                "the1_mobile_no": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "API key",
            "type": "apiKey",
            "name": "X-API-KEY",
            "in": "header"
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
