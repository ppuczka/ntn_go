package model



type Pages struct {
	Pages []Page `json:"results"` 
}

type Page struct {
	Url             string 		   `json:"url,omitempty"`
	Id              string 		   `json:"id,omitempty"`
	// Type            string         `json:"type"`
    Parent          Parent 		   `json:"parent"`
	Properties    	PageProperties `json:"properties"`
	Children		[]PageChildren `json:"children"`   
	// ChildPage       ChildPage      `json:"child_page,omitempty"`
	// Code            Code           `json:"code,omitempty"` 
}

func CreateSnippetPageModel(parentPage Page, newSnippetTitle, newSnippetText, newSnippetCaption string) (snippetPage Page) {
	snippetTitle := Text{newSnippetTitle, ""}
	innerTitle := InnerTitle{"text", snippetTitle, "innerText"}
	pageTitle := Title{"id", "title", []InnerTitle{innerTitle}}
	properties := PageProperties {pageTitle}

	snippetContent := Text{newSnippetText, ""}
	snippetCaptionText := Text{newSnippetCaption, ""}
	snippetRichText := RichText{"text", snippetContent}
	snippetCaption := RichText{"text", snippetCaptionText}
	
	// paragraph := Paragraph{[]RichText {snippetRichText}}
	codeType := Code{[]RichText {snippetRichText}, []RichText {snippetCaption}, "bash"}
	child := PageChildren{"block", "code", codeType}

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
	// Paragraph Paragraph `json:"paragraph,omitempty"`
	Code      Code      `json:"code"`
}
type Paragraph struct {
	RichText []RichText `json:"rich_text"`
}
type RichText struct {
	Type     string `json:"type"`
	Text     Text   `json:"text"`
}

type Code struct {
	RichText  []RichText `json:"rich_text"`
	Caption   []RichText `json:"caption"`
	Language  string     `json:"language"`
}

type ChildPage struct {
	Title     string `json:"title"`
}