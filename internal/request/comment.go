package request

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

type GetCommentsRequest struct {
	Page     int `form:"page,default=1" binding:"min=1"`
	PageSize int `form:"page_size,default=10" binding:"min=1,max=100"`
}
