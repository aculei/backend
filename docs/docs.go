// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Michele Dinelli",
            "email": "dinellimichele00@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/dataset": {
            "get": {
                "description": "Return info about dataset like number of records, number of columns, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dataset"
                ],
                "summary": "return dataset info",
                "responses": {
                    "200": {
                        "description": "The dataset info",
                        "schema": {
                            "$ref": "#/definitions/models.Dataset"
                        }
                    },
                    "500": {
                        "description": "An error occurred",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorInternalServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Dataset": {
            "type": "object",
            "properties": {
                "badger": {
                    "type": "string"
                },
                "buzzard": {
                    "type": "string"
                },
                "camera_1_coordinates": {
                    "type": "string"
                },
                "camera_2_coordinates": {
                    "type": "string"
                },
                "camera_3_coordinates": {
                    "type": "string"
                },
                "camera_4_coordinates": {
                    "type": "string"
                },
                "camera_5_coordinates": {
                    "type": "string"
                },
                "camera_6_coordinates": {
                    "type": "string"
                },
                "camera_7_coordinates": {
                    "type": "string"
                },
                "cat": {
                    "type": "string"
                },
                "deer": {
                    "type": "string"
                },
                "fox": {
                    "type": "string"
                },
                "hare": {
                    "type": "string"
                },
                "heron": {
                    "type": "string"
                },
                "horse": {
                    "type": "string"
                },
                "mallard": {
                    "type": "string"
                },
                "marten": {
                    "type": "string"
                },
                "photos_autumn": {
                    "type": "string"
                },
                "photos_camera_1": {
                    "type": "string"
                },
                "photos_camera_2": {
                    "type": "string"
                },
                "photos_camera_3": {
                    "type": "string"
                },
                "photos_camera_4": {
                    "type": "string"
                },
                "photos_camera_5": {
                    "type": "string"
                },
                "photos_camera_6": {
                    "type": "string"
                },
                "photos_camera_7": {
                    "type": "string"
                },
                "photos_spring": {
                    "type": "string"
                },
                "photos_summer": {
                    "type": "string"
                },
                "photos_winter": {
                    "type": "string"
                },
                "porcupine": {
                    "type": "string"
                },
                "squirrel": {
                    "type": "string"
                },
                "total_records": {
                    "type": "string"
                },
                "wild_boar": {
                    "type": "string"
                },
                "wolf": {
                    "type": "string"
                }
            }
        },
        "models.ErrorInternalServerError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "internal server error"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "aculei-be",
	Description:      "The purpose of this microservice is to serve aculei.xyz",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
