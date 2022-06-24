package model


type Pages struct {
	Pages []Page `json:"results"` 
}

type Page struct {
	Url             string 		   `json:"url,omitempty"`
	Id              string 		   `json:"id,omitempty"`
    Parent          Parent 		   `json:"parent"`
	Properties    	PageProperties `json:"properties"`
	Children		[]PageChildren `json:"children"`    
}

func CreateSnippetPageModel(parentPage Page, newSnippetTitle, newSnippetText string) (snippetPage Page) {
	snippetTitle := Text{newSnippetTitle, ""}
	innerTitle := InnerTitle{"text", snippetTitle, "innerText"}
	pageTitle := Title{"id", "title", []InnerTitle{innerTitle}}
	properties := PageProperties {pageTitle}

	richTextContent := Text{newSnippetText, ""}
	richText := RichText{"text", richTextContent}
	paragraph := Paragraph{[]RichText {richText}}
	child := PageChildren{"block", "paragraph", paragraph}

	children := []PageChildren{child}
	parent := Parent{"page_id", parentPage.Id}
	
	page  := Page {
		Parent : parent,
		Properties: properties,
		Children : children,	
	}
	return page
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
	Link     string `json:"link,omitempty"`
}
type PageChildren struct {
	Object    string    `json:"object"`
	Type      string    `json:"type"`
	Paragraph Paragraph `json:"paragraph"`
}
type Paragraph struct {
	RichText []RichText `json:"rich_text"`
}
type RichText struct {
	Type     string `json:"type"`
	Text     Text   `json:"text"`
}