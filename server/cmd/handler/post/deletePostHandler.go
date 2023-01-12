package post

import (
	"context"
	"net/http"
	"wpm-project/cmd/logic/post"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"
	"wpm-project/core/http_request"
	"wpm-project/core/http_response"
	"wpm-project/core/logger"

	"github.com/gin-gonic/gin"
)

func DeletePostHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add trace_id to context
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("Delete-post-api"))
		// Init log helper with context (have trace_id)
		logHelper := logger.NewContextLog(ctx)
		// New object logic (all logic code will implement in this object)
		deletePostLogic := post.NewDeletePostLogic(ctx, svcCtx, logHelper)

		request := &types.DeletePostRequest{}
		err := http_request.BindUri(c, request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		// Call functions in logic to process
		response, err := deletePostLogic.DeletePost(request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, response)
	}
}