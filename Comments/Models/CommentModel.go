package models

type Comment struct {
	PostId        int      `json:"postid,omitempty"`
	CommentorId   int      `json:"commentorid,omitempty"`
	CommentString string   `json:"comment,omitempty"`
	CommentId     int      `json:"commentid,omitempty"`
	TagList       []string `json:"tags,omitempty"`
	Likes         int      `json:"likes,omitempty"`
}
