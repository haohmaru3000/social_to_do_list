package ginuserlikeitem

import (
	"gorm.io/gorm"
	"net/http"
	"time"

	"social_todo_list/common"
	"social_todo_list/module/userlikeitem/biz"
	"social_todo_list/module/userlikeitem/model"
	"social_todo_list/module/userlikeitem/storage"
	"social_todo_list/pubsub"

	"github.com/gin-gonic/gin"
	goservice "github.com/haohmaru3000/go_sdk"
)

func LikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		ps := serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)

		store := storage.NewSQLStore(db)
		//itemStore := itemStorage.NewSQLStore(db)
		business := biz.NewUserLikeItemBiz(store, ps)
		now := time.Now().UTC()

		if err := business.LikeItem(c.Request.Context(), &model.Like{
			UserId:    requester.GetUserId(),
			ItemId:    int(id.GetLocalID()),
			CreatedAt: &now,
		}); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
