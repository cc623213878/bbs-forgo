package response

type Page struct {
	PageNum  int   `json:"page_num"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

type resopnsePage struct {
	Page    `json:"page"`
	Content interface{} `json:"content"`
}

func ResponsePage(pageNum int, pageSize int, total int64, content interface{}) *resopnsePage {
	return &resopnsePage{
		Page: Page{
			PageNum:  pageNum,
			PageSize: pageSize,
			Total:    total,
		},
		Content: content,
	}
}
