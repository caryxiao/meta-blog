package request

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=255"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=255"`
	Content string `json:"content" binding:"required"`
}

type ListPostRequest struct {
	Page     int   `form:"page" binding:"min=1"`
	PageSize int   `form:"page_size" binding:"min=1,max=100"`
	UserID   *uint `form:"user_id"`
}
