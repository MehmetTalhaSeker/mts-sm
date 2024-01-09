package dto

type PostLikeCreateRequest struct {
	PostID uint `json:"postId'" validate:"required"`
	UserID uint `validate:"required"`
}

type CommentLikeCreateRequest struct {
	CommentID uint `json:"commentId'" validate:"required"`
	UserID    uint `validate:"required"`
}
