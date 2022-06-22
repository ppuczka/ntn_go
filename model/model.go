package model



type Pages struct {

	Pages []Page `json:"results"` 
}

type Page struct {
	Url             string 		   `json:"url"`
	Id              string 		   `json:"id"`
    Parent          Parent 		   `json:"parent"`
	Properties    	PageProperties `json:"properties"`
	Children		[]PageChildren   `json:"children"`    
}
  
type Parent struct {
	Type     string `json:"type"`
	PageId   string `json:"page_id"`
}

type PageProperties struct {
	Title    Title `json:"title"`
}

type Title struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	InnerTitles  []InnerTitle `json:"title"`
}

type InnerTitle struct {
	Type        string `json:"type"`
	Text        Text   `json:"text"`
	PlainText   string `json:"plain_text"` 
}

type Text struct {
	Content  string `json:"content"`
	Link     string `json:"link"`
}

type PageChildren struct {
	Object   string `json:"object"`
	Type     string `json:"paragraph"`
}

type Paragraph struct {
	RichText []RichText `json:"rich_text"`
}

type RichText struct {
	Type     string `json:"type"`
	Text     Text   `json:"content"`
}