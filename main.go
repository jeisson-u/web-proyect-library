package main

import (
	"bufio"
	"fmt"
	"github.com/jeisson-u/web-proyect-library.git/controllers"
	"github.com/jeisson-u/web-proyect-library.git/models"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}


func main(){

	fmt.Println("DIGIT PORT:")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	port:=scanner.Text()

	localLibrary:=[]models.NewBook{}

	e := echo.New()


	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/library/*.html")),
	}
	e.Renderer = renderer

	libraryController:=controllers.LibraryController{
		&localLibrary,
	}
	e.GET("/", libraryController.Index)
	e.GET("/book", libraryController.Book)
	e.GET("/book/category/:categoryName", libraryController.BookCategory)
	e.GET("/book/author/:author", libraryController.BookAuthor)
	e.GET("/book/:id", libraryController.BookId)
	e.GET("/library/add", libraryController.AddBookPage)
	e.POST("/book", libraryController.AddBook)

	e.Static("/assets", "public/assets")

	webBrowserStart("http://localhost:"+port)

	e.Logger.Fatal(e.Start(":"+port))



}


func webBrowserStart(url string){
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)

	 cmd.Start()

}