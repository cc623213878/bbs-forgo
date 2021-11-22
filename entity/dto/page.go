package dto

type Page struct {
	PageNum  int `binding:"min=0,max=10240"`
	PageSize int `binding:"min=1,max=1024"`
}
