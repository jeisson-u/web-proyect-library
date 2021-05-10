package models

type Book struct {
	Id string `json:"id"`
	SelfLink string `json:"selfLink"`
	VolumeInfo struct{
		Title string `json:"title"`
		Subtitle string `json:"subtitle"`
		Authors []string `json:"authors"`
		Publisher string `json:"publisher"`
		PublishedDate string `json:"publishedDate"`
		Description string `json:"description"`
		IndustryIdentifiers []struct{
			Type string `json:"type"`
			Identifier string `json:"identifier"`
		}`json:"industryIdentifiers"`
		PageCount int `json:"pageCount"`
		PrintType string `json:"printType"`
		Categories []string `json:"categories"`
		AverageRating float64 `json:"averageRating"`
		ImageLinks struct{
			SmallThumbnail string `json:"smallThumbnail"`
			Thumbnail string `json:"thumbnail"`
		} `json:"imageLinks"`
		Language string `json:"language"`
		PreviewLink string `json:"previewLink"`
	} `json:"volumeInfo"`

}
