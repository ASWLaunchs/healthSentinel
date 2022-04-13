package main

import (
	"fmt"
	"healthSentinel/models"
	"healthSentinel/routers/routers_admin"
	"healthSentinel/routers/routers_office"
	"healthSentinel/routers/routers_verify"
	"os"

	"gopkg.in/ini.v1"

	gin "github.com/gin-gonic/gin"
)

func main() {
	//initialize database.
	models.CreateDB()

	//get configuration.
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	address := cfg.Section("website").Key("address").String()
	port := cfg.Section("website").Key("port").String()
	server := address + ":" + port

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	routers_admin.Admin(r)
	routers_office.Office(r)
	routers_verify.Verify(r)

	r.Run(server) // listen and serve on 0.0.0.0:22018 (for windows "localhost:22018")
}
