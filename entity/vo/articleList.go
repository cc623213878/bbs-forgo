package vo

// ArticleList 文章列表
type ArticleList struct {
	Title     string
	Keywords  string
	CoverPath string
	Content   string
	CreatedAt int64
	UpdatedAt int64
	Username  string
}
