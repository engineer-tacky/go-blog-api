package services

import (
	"github.com/engineer-tacky/go-blog-api/apperrors"
	"github.com/engineer-tacky/go-blog-api/models"
	"github.com/engineer-tacky/go-blog-api/repositories"
)

// PostCommentHandlerで使用することを想定したサービス
// 引数の情報をもとに新しいコメントを作り、結果を返却
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
