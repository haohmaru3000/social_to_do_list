package biz

import (
	"context"

	"social_todo_list/common"
	"social_todo_list/module/userlikeitem/model"
)

type ListUserLikeItemStore interface {
	ListUsers(
		ctx context.Context,
		itemId int,
		paging *common.Paging,
	) ([]common.SimpleUser, error)
}

type listUserLikeItemBiz struct {
	store ListUserLikeItemStore
}

func NewListUserLikeItemBiz(store ListUserLikeItemStore) *listUserLikeItemBiz {
	return &listUserLikeItemBiz{store: store}
}

func (biz *listUserLikeItemBiz) ListUserLikedItem(
	ctx context.Context,
	itemId int,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	result, err := biz.store.ListUsers(ctx, itemId, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return result, nil
}
