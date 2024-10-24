package upload

import (
	"fmt"
	"net/http"
	"time"

	"social_todo_list/common"

	"github.com/gin-gonic/gin"
	goservice "github.com/haohmaru3000/go_sdk"
)

// NOTE: this handler is very simple and please do not use it in practice
// Instead, I recommend you should check "Upload Image to AWS S3 and CDN with CloudFront"

func Upload(serviceCtx goservice.ServiceContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		dst := fmt.Sprintf("static/%d.%s", time.Now().UTC().UnixNano(), fileHeader.Filename)

		if err := c.SaveUploadedFile(fileHeader, dst); err != nil {

		}

		img := common.Image{
			Id:        0,
			Url:       dst,
			Width:     100,
			Height:    100,
			CloudName: "local",
			Extension: "",
		}

		img.Fulfill("http://localhost:3000")

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
