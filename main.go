package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"url/db"
	"url/services"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func IndexHandler(c echo.Context) error {
	urls := services.GetAllUrls()

	return c.Render(http.StatusOK, "index", urls)
}

func Shorten(c echo.Context) error {
	type Body struct {
		Url string `form:"url"`
	}

	var body Body

	err := c.Bind(&body)
	if err != nil {
		fmt.Println(err)
	}

	id := uuid.New()
	db.Db.Exec("INSERT INTO urls(url, public_id) VALUES($1, $2)", body.Url, id)

	return c.JSON(http.StatusOK, body.Url)
}

func UrlsHandler(c echo.Context) error {
	urls := services.GetAllUrls()

	return c.JSON(http.StatusOK, urls)
}

func GetUrlHandler(c echo.Context) error {
	id := c.Param("id")

	url := services.GetUrlById(id)

	c.Redirect(http.StatusFound, url.Url)

	return nil

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.Init()

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Static("/public", "public")

	e.GET("/:id", GetUrlHandler)
	e.GET("/", IndexHandler)
	e.POST("/shorten", Shorten)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
