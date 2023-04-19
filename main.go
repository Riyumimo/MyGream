package main

import (
	"MyGream/database"
	_ "MyGream/docs"
	"MyGream/router"
)

// @title           MyGram Final Project
// @version         1.0
// @description     Ini Adalah Tugas final project dari DTS kominfo "Project MyGram".
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	database.StartDb()
	r := router.StartApp()
	r.Run(":8080")
}
