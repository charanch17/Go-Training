package data

import (
	models "comments.com/charan/comments/Models"
)

var CommentsTable = []models.Comment{
	{PostId: 1, CommentorId: 2, CommentString: "super dp", CommentId: 21, TagList: []string{"dp", "selflove"}, Likes: 10},
	{PostId: 2, CommentorId: 3, CommentString: "super profile", CommentId: 22, TagList: []string{"dp", "selflove", "respect"}, Likes: 5},
	{PostId: 2, CommentorId: 2, CommentString: "super car", CommentId: 23, TagList: []string{"dp", "selflove", "motivation"}, Likes: 7},
	{PostId: 3, CommentorId: 5, CommentString: "super hero", CommentId: 31, TagList: []string{"dp", "selflove", "hatred"}, Likes: 8},
	{PostId: 4, CommentorId: 2, CommentString: "super dp", CommentId: 51, TagList: []string{"dp", "selflove", "jai janasena"}, Likes: 4},
	{PostId: 1, CommentorId: 6, CommentString: "helll no", CommentId: 28, TagList: []string{"dp", "selflove", "pawankalyan"}, Likes: 2},
	{PostId: 5, CommentorId: 2, CommentString: "super ghost", CommentId: 29, TagList: []string{"dp", "selflove", "damn"}, Likes: 1},
}
