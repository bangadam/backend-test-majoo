package main

import (
	"github.com/bangadam/backend-test-majoo/pkg/server"
)

// @title           API documentaion of Backend Test Majoo
// @version         1.0
// @description     API documentaion of Backend Test Majoo
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	server.Run()
}