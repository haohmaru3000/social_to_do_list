package storage

import (
	"context"

	"to_do_list/common"
	"to_do_list/module/item/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	deletedStatus := "Deleted"

	if err := s.db.Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deletedStatus,
		}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
