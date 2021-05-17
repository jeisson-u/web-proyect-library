package models

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)
const(
	googleVolumesApiUrl="https://www.googleapis.com/books/v1/volumes"
)



type Library struct{
	LocalLibrary *[]NewBook
}
func (library Library) GetRandomBooks() []Book{


	resp, err := http.Get(googleVolumesApiUrl+"?q=harry+potter&key=AIzaSyB0wgQppx-gZkFQiHVGpeppgcMLe2gC3QA&maxResults=6&printType=books")

	if err!=nil{
		print(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var modelResponse GoogleBookResponse
	if err = json.Unmarshal(body,&modelResponse); err != nil {
		print(err.Error())
	}

	if modelResponse.TotalItems==0{
		return []Book{}
	}

	return modelResponse.Items

}

func (Library) GetSearchBook(title string) []Book{

	q:=url.QueryEscape(title)
	resp, err := http.Get(googleVolumesApiUrl+"?q="+q+"&key=AIzaSyB0wgQppx-gZkFQiHVGpeppgcMLe2gC3QA&maxResults=6&printType=books")

	if err!=nil{
		print(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var modelResponse GoogleBookResponse
	if err = json.Unmarshal(body,&modelResponse); err != nil {
		print(err.Error())
	}

	if modelResponse.TotalItems==0{
		return []Book{}
	}

	return modelResponse.Items

}

func (Library) GetById(bookId string) []Book{


	resp, err := http.Get(googleVolumesApiUrl+"/"+bookId+"&key=AIzaSyB0wgQppx-gZkFQiHVGpeppgcMLe2gC3QA")

	if err!=nil{
		print(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var modelResponse GoogleBookResponse
	if err = json.Unmarshal(body,&modelResponse); err != nil {
		print(err.Error())
	}

	if modelResponse.TotalItems==0{
		return []Book{}
	}

	return modelResponse.Items

}


func (Library) NewBook(bookId string) []Book{


	resp, err := http.Get(googleVolumesApiUrl+"/"+bookId+"&key=AIzaSyB0wgQppx-gZkFQiHVGpeppgcMLe2gC3QA")

	if err!=nil{
		print(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var modelResponse GoogleBookResponse
	if err = json.Unmarshal(body,&modelResponse); err != nil {
		print(err.Error())
	}

	if modelResponse.TotalItems==0{
		return []Book{}
	}

	return modelResponse.Items

}