package utils

import (
	"fmt"
	"strconv"

	data "comments.com/charan/comments/Db"
	models "comments.com/charan/comments/Models"
)

func getCommentsByPostId(postid int) []models.Comment {
	CommentsList := []models.Comment{}
	for _, comment := range data.CommentsTable {
		if comment.PostId == postid {
			CommentsList = append(CommentsList, comment)
		}

	}
	return CommentsList

}

func getCommentsByCommentorId(commentorid int) []models.Comment {
	CommentsList := []models.Comment{}
	for _, comment := range data.CommentsTable {
		if comment.CommentorId == commentorid {
			CommentsList = append(CommentsList, comment)
		}

	}
	fmt.Println(CommentsList)
	return CommentsList

}

func GetAllComments() []models.Comment {
	return data.CommentsTable
}

func GetCommentsByField(FieldName string, Value interface{}) []models.Comment {

	switch FieldName {
	case "PostId":
		postId, _ := strconv.Atoi(Value.(string))
		return getCommentsByPostId(postId)
	case "CommentorId":
		commentorId, _ := strconv.Atoi(Value.(string))
		return getCommentsByCommentorId(commentorId)

	}
	return []models.Comment{}

}
func AppendComment(comment models.Comment) {
	data.CommentsTable = append(data.CommentsTable, comment)
}

func SendCommentRef(commentid int) *models.Comment {
	for _, comment := range data.CommentsTable {
		if comment.CommentId == commentid {
			return &comment
		}
	}
	return &models.Comment{}
}

func DeleteCommentRecord(commentid int) map[string]string {
	for index, comment := range data.CommentsTable {
		if comment.CommentId == commentid {
			data.CommentsTable = append(data.CommentsTable[:index], data.CommentsTable[index+1:]...)
			return map[string]string{"msg": "deleted successfully", "status": "200"}
		}
	}
	return map[string]string{"msg": "no comment found unable to delete", "status": "400"}
}
