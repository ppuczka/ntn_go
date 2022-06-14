package model



type Pages struct {

	Pages []Page `json:"results"` 
}

type Page struct {
	Url       string `json:"url"`
	Id        string `json:"id"`
    Parent    Parent `json:"parent"`

}
  
type Parent struct {
	Type     string `json:"type"`
	PageId   string `json:"page_id"`
}
