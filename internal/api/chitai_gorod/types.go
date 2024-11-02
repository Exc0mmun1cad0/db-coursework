package chitaigorod

const (
	ProductType = "product"
)

type Response struct {
	Data []Data `json:"data"`
}

type Data struct {
	Attributes struct {
		Authors []struct {
			FirstName  string `json:"firstName"`
			ID         int    `json:"id"`
			LastName   string `json:"lastName"`
			MiddleName string `json:"middleName"`
		} `json:"authors"`
		Category struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"category"`
		Description string `json:"description"`
		ID          int    `json:"id"`
		IsBook      bool   `json:"isBook"`
		Pages       string `json:"pages"`
		Publisher   struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"publisher"`
		Title          string `json:"title"`
		URL            string `json:"url"`
		YearPublishing int    `json:"yearPublishing"`
	} `json:"attributes"`
	ID   string `json:"id"`
	Type string `json:"type"`
}
