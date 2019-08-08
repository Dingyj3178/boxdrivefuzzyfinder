package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// type File struct {
// 	Filename string `json:"filename"`
// }
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initRead := os.Getenv("INIT_READ")
	println(initRead)
	if initRead == "0" {
		dirwalk(os.Getenv("DIR"))
	}

}

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./dist", false)))
	router.GET("/search/:inputText", handlefiles)
	router.POST("/open", openfiles)
	router.Run()

}

func handlefiles(c *gin.Context) {
	f := c.Param("inputText")
	println(f)
	test := finder(`./files.in`, f)
	// test := []string{"UI5", "Web", "Components"}
	c.JSON(200, test)
	// c.HTML(http.StatusOK, "index.html", gin.H{"a": "a"})
	// if err != nil {
	// 	panic(err)
	// }
}
func openfiles(c *gin.Context) {
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	f := string(buf[0:n])
	trimIndex := strings.LastIndex(f, "/")
	dir := f[:trimIndex]
	println(dir)
	println(f)
	err := exec.Command("open", dir).Run()
	// files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	c.JSON(200, "")
}
