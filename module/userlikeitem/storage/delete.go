package storage

import (
	"context"

	"social_todo_list/common"
	"social_todo_list/module/userlikeitem/model"
)

func (s *sqlStore) Delete(ctx context.Context, userId, itemId int) error {
	var data model.Like

	if err := s.db.Table(data.TableName()).
		Where("user_id = ? and item_id = ?", userId, itemId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
