package post

import (
	"context"
	"net/http"
	"wpm-project/cmd/logic/post"
	"wpm-project/cmd/svc"
	"wpm-project/core/http_response"
	"wpm-project/core/logger"

	"github.com/gin-gonic/gin"
)

func GetPostHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add trace_id to context
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("get-post-api"))
		// Init log helper with context (have trace_id)
		logHelper := logger.NewContextLog(ctx)
		// New object logic (all logic code will implement in this object)
		getPostLogic := post.NewGetPostLogic(ctx, svcCtx, logHelper)

		// Call functions in logic to process
		response, err := getPostLogic.GetPost()
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, response)
	}
}