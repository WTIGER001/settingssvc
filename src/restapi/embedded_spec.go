// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Definition for the Preferences Server",
    "title": "Preferences Store",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/configuration": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Get all configuration definitions",
        "operationId": "getConfig",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Config"
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/definition": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Get all preference definitions",
        "operationId": "getDefinitions",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/PreferenceDefinitionArray"
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Add a new Preference Definition to the set of available definitions",
        "operationId": "addDefinition",
        "parameters": [
          {
            "description": "Preference Definition that needs to be added",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PreferenceDefinition"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/PreferenceDefinition"
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/definition/{id}": {
      "get": {
        "description": "Returns a single definition",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Find definition by ID",
        "operationId": "getDefinition",
        "parameters": [
          {
            "type": "string",
            "description": "ID of definition to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/PreferenceDefinition"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Update an existing Definition",
        "operationId": "updateDefinition",
        "parameters": [
          {
            "type": "string",
            "description": "ID of definition to return",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Preference Definition needs to be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PreferenceDefinition"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/PreferenceDefinition"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Definition not found"
          },
          "405": {
            "description": "Validation exception"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Deletes a definition",
        "operationId": "deleteDefinition",
        "parameters": [
          {
            "type": "string",
            "description": "Definition id to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Definition not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/my": {
      "get": {
        "description": "Returns a single owner",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Preferences"
        ],
        "summary": "Find Preference for the JWT supplied",
        "operationId": "getMyActiveProfile",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/owner/{id}": {
      "get": {
        "description": "Returns a single owner",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Preferences"
        ],
        "summary": "Find Preference Owner by Id",
        "operationId": "getOwner",
        "parameters": [
          {
            "type": "string",
            "description": "ID of type to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/PreferenceOwner"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Preferences"
        ],
        "summary": "Update an existing owner",
        "operationId": "updateOwner",
        "parameters": [
          {
            "type": "string",
            "description": "ID of type to return",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Owner needs to be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PreferenceOwner"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Definition not found"
          },
          "405": {
            "description": "Validation exception"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Preferences"
        ],
        "summary": "Deletes an Owner",
        "operationId": "deleteOwner",
        "parameters": [
          {
            "type": "string",
            "description": "Type id to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Type not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/profile/{id}": {
      "get": {
        "description": "Returns a single owner",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Preferences"
        ],
        "summary": "Find profile  by Id",
        "operationId": "getProfile",
        "parameters": [
          {
            "type": "string",
            "description": "ID of type to return",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int",
            "description": "ID of type to return",
            "name": "version",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Preferences"
        ],
        "summary": "Update an existing profile",
        "operationId": "updateProfile",
        "parameters": [
          {
            "type": "string",
            "description": "ID of type to return",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Owner needs to be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Definition not found"
          },
          "405": {
            "description": "Validation exception"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Preferences"
        ],
        "summary": "Deletes a profile",
        "operationId": "deleteProfile",
        "parameters": [
          {
            "type": "string",
            "description": "Type id to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Type not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/profile/{id}/version": {
      "get": {
        "description": "Returns a list of all profile versions",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Preferences"
        ],
        "summary": "Versions of a profile",
        "operationId": "getProfileVersions",
        "parameters": [
          {
            "type": "string",
            "description": "ID of type to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ProfileVersions"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/profiles": {
      "get": {
        "description": "Returns a single owner",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Preferences"
        ],
        "summary": "Find profile  by Id",
        "operationId": "getProfiles",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "ID of type to return",
            "name": "id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ProfileArray"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/type": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Get all owner types",
        "operationId": "getOwnerTypes",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/OwnerTypeArray"
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Add a new Owner type to the set of available types",
        "operationId": "addOwnerType",
        "parameters": [
          {
            "description": "Type that needs to be added",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/OwnerType"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/type/{id}": {
      "get": {
        "description": "Returns a single owner type",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Find type by ID",
        "operationId": "getType",
        "parameters": [
          {
            "type": "string",
            "description": "ID of type to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/OwnerType"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Update an existing type",
        "operationId": "updateOwnerType",
        "parameters": [
          {
            "type": "string",
            "description": "ID of type to return",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Owner type needs to be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/OwnerType"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/OwnerType"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Definition not found"
          },
          "405": {
            "description": "Validation exception"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Configuration"
        ],
        "summary": "Deletes a type",
        "operationId": "deleteType",
        "parameters": [
          {
            "type": "string",
            "description": "Type id to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Type not found"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Config": {
      "type": "object",
      "properties": {
        "definitions": {
          "$ref": "#/definitions/PreferenceDefinitionArray"
        },
        "ownerTypes": {
          "$ref": "#/definitions/OwnerTypeArray"
        }
      }
    },
    "OwnerType": {
      "type": "object",
      "properties": {
        "definitions": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "OwnerTypeArray": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/OwnerType"
      }
    },
    "Preference": {
      "type": "object",
      "properties": {
        "definition-id": {
          "type": "string"
        },
        "value": {
          "type": "object"
        }
      }
    },
    "PreferenceDefinition": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "layout": {
          "description": "Layout Instructions",
          "type": "object"
        },
        "name": {
          "type": "string"
        },
        "order": {
          "type": "integer",
          "format": "int"
        },
        "schema": {
          "description": "JSON Schema",
          "type": "object"
        }
      }
    },
    "PreferenceDefinitionArray": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/PreferenceDefinition"
      }
    },
    "PreferenceOwner": {
      "type": "object",
      "properties": {
        "active": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "owner-type": {
          "type": "string"
        },
        "profile-ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Profile": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "preferences": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Preference"
          }
        },
        "version": {
          "type": "integer",
          "format": "int"
        }
      }
    },
    "ProfileArray": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Profile"
      }
    },
    "ProfileVersion": {
      "type": "object",
      "properties": {
        "version": {
          "type": "integer",
          "format": "int"
        },
        "version-date": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "ProfileVersions": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "versions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ProfileVersion"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "Preference Management including instnaces of types and profiles",
      "name": "Preferences"
    },
    {
      "description": "Preferences Definition Management",
      "name": "Configuration"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}