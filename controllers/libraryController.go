package controllers

import (
	"github.com/jeisson-u/web-proyect-library.git/models"
	"github.com/labstack/echo"
	"net/http"
)

type LibraryController struct {
	LocalLibrary *[]models.NewBook
}

func (libraryController LibraryController) Index(c echo.Context) error {

	library:=models.Library{}

	books:=library.GetRandomBooks()

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"books": books,
	})
}

func (LibraryController) Book(c echo.Context) error {
	library:=models.Library{}
	var books []models.Book
	query := c.QueryParam("q")

	if query!=""{
		books=library.GetSearchBook(query)
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"books": books,
		"query":query,
	})
}

func (LibraryController) BookId(c echo.Context) error {
	bookId := c.Param("id")
	library:=models.Library{}

	books:=library.GetSearchBook(bookId)

	return c.Render(http.StatusOK, "book.html", map[string]interface{}{
		"book": books[0],
	})
}

func (LibraryController) BookCategory(c echo.Context) error {
	categoryName := c.Param("categoryName")
	library:=models.Library{}

	books:=library.GetSearchBook(categoryName)

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"books": books,
	})
}

func (LibraryController) BookAuthor(c echo.Context) error {
	author := c.Param("author")
	library:=models.Library{}

	books:=library.GetSearchBook(author)

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"books": books,
	})
}

func (libraryController LibraryController) AddBook(c echo.Context) error {

	modelIn:= models.NewBook{
		Title: c.FormValue("Title"),
		Category: c.FormValue("Category"),
		Description: c.FormValue("Description"),
		Date: c.FormValue("Date"),
		Author: c.FormValue("Author"),
		Image: c.FormValue("Image"),
		Url: c.FormValue("Url"),
	}


	newList:=append(*libraryController.LocalLibrary, modelIn)

	libraryController.LocalLibrary=&newList

	return c.Render(http.StatusOK, "newBook.html", map[string]interface{}{
	})
}

func (LibraryController) AddBookPage(c echo.Context) error {


	return c.Render(http.StatusOK, "newBook.html", map[string]interface{}{
	})
}