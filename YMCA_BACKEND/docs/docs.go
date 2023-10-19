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
        "/addPrivateData": {
            "post": {
                "description": "This endpoint is used to add Private Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "add Private Data",
                "parameters": [
                    {
                        "description": "enter the details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PrivateDataInputReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/addPublicData": {
            "post": {
                "description": "This endpoint is used to add Public Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "add Public Data",
                "parameters": [
                    {
                        "description": "enter details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PublicDataInputReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AddPublicDataResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/deploy-smart-contract": {
            "post": {
                "description": "This endpoint is used to deploy the smart contract token and token chain to the network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "This function deploys the smart contract token",
                "parameters": [
                    {
                        "description": "Give the input",
                        "name": "smart_contract_input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.DeploySmartContractInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.RubixResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/execute-smart-contract": {
            "post": {
                "description": "This endpoint is used to execute the smart contract. When a smart contract is executed the tokenchain is updated, this updation happens here.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "This function update the token chain",
                "parameters": [
                    {
                        "description": "Give the input",
                        "name": "smart_contract_input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ExecuteSmartContractInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.RubixResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/generate-smart-contract": {
            "post": {
                "description": "This endpoint is used to generate the smart contract token and the genesis block of the tokenchain",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "This function generates the smart contract token",
                "parameters": [
                    {
                        "description": "Give the input",
                        "name": "smart_contract_input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SmartContractInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.RubixResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/subscribe-smart-contract": {
            "post": {
                "description": "This endpoint is used to subscribe the smart contract.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "This function subscribes the smart contract",
                "parameters": [
                    {
                        "description": "Give the input",
                        "name": "subscribe_smart_contract_input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.SubscribeContractRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.RubixResponse"
                        }
                    }
                }
            }
        },
        "/getAllAccessDatabyID": {
            "get": {
                "description": "Get rivate data that has been given access to a  ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return user private data that has been given access to a  ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User's ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.PrivateDataResponse"
                            }
                        }
                    }
                }
            }
        },
        "/getAllPrivateDataByID": {
            "get": {
                "description": "Get private data for a user by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return user private data by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User's ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.PrivateDataResponse"
                            }
                        }
                    }
                }
            }
        },
        "/getAllPublicDataByID": {
            "get": {
                "description": "Get public data for a user by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return user public data by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User's ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.PublicDataResponse"
                            }
                        }
                    }
                }
            }
        },
        "/getDIDbyUserID": {
            "get": {
                "description": "Get user DID when ID is given",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return user DID when ID is given",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User's ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/getPvtDatabyID": {
            "get": {
                "description": "Get user DID when ID is given",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return user DID when ID is given",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User's ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PvtDataResponse"
                        }
                    }
                }
            }
        },
        "/getUserIDbyDID": {
            "get": {
                "description": "Get user id when DID is given",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return user id when DID is given",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "This endpoint is used to authenticate existing user log in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return user data",
                "parameters": [
                    {
                        "description": "enter email and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LogInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LogInResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "This endpoint is used to when new user signs up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return user data",
                "parameters": [
                    {
                        "description": "enter email and phone number",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SignUpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AddPublicDataResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "pub_data_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "model.BasicResponse": {
            "type": "object",
            "properties": {
                "did": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.LogInRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.LogInResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.PrivateData": {
            "type": "object",
            "properties": {
                "capsule": {
                    "type": "string"
                },
                "cipher_text": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "pvt_data_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.PrivateDataInputReq": {
            "type": "object",
            "properties": {
                "communities": {
                    "type": "string"
                },
                "decrypt_user_id": {
                    "type": "integer"
                },
                "focus_area": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.PrivateDataResponse": {
            "type": "object",
            "properties": {
                "capsule": {
                    "type": "string"
                },
                "cipher_text": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.PublicDataInputReq": {
            "type": "object",
            "properties": {
                "communities": {
                    "type": "string"
                },
                "focus_area": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.PublicDataResponse": {
            "type": "object",
            "properties": {
                "communities": {
                    "type": "string"
                },
                "focus_area": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.PvtDataResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "privateData": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PrivateData"
                    }
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "model.SignUpRequest": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "model.SignUpResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "service.DeploySmartContractInput": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "deployerAddress": {
                    "type": "string"
                },
                "port": {
                    "type": "string"
                },
                "quorumType": {
                    "type": "integer"
                },
                "rbtAmount": {
                    "type": "integer"
                },
                "smartContractToken": {
                    "type": "string"
                }
            }
        },
        "service.ExecuteSmartContractInput": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "executorAddress": {
                    "type": "string"
                },
                "port": {
                    "type": "string"
                },
                "quorumType": {
                    "type": "integer"
                },
                "smartContractData": {
                    "type": "string"
                },
                "smartContractToken": {
                    "type": "string"
                }
            }
        },
        "service.RubixResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "result": {},
                "status": {
                    "type": "boolean"
                }
            }
        },
        "service.SmartContractInput": {
            "type": "object",
            "properties": {
                "binaryCodePath": {
                    "type": "string"
                },
                "did": {
                    "type": "string"
                },
                "port": {
                    "type": "string"
                },
                "rawCodePath": {
                    "type": "string"
                },
                "schemaFilePath": {
                    "type": "string"
                }
            }
        },
        "service.SubscribeContractRequest": {
            "type": "object",
            "properties": {
                "contract": {
                    "type": "string"
                },
                "port": {
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
