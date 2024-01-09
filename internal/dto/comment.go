package dto

type CommentCreateRequest struct {
	Text     string `json:"text" validate:"required,min=5,max=30"`
	PostID   uint   `json:"postId'" validate:"required"`
	ParentID *uint  `json:"parentId'"`
	UserID   uint   `validate:"required"`
}
