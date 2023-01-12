package point

import (
	"context"
	"net/http"
	"wpm-project/cmd/logic/point"
	"wpm-project/cmd/svc"
	"wpm-project/cmd/types"
	"wpm-project/core/http_request"
	"wpm-project/core/http_response"
	"wpm-project/core/logger"

	"github.com/gin-gonic/gin"
)

func CreatePointHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add trace_id to context
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("get-point-api"))
		// Init log helper with context (have trace_id)
		logHelper := logger.NewContextLog(ctx)
		// New object logic (all logic code will implement in this object)
		createpointLogic := point.NewCreatePointLogic(ctx, svcCtx, logHelper)

		// Call functions in logic to process
		request := &types.CreatePointRequest{}
		err := http_request.BindBodyJson(c, request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		response, err := createpointLogic.CreatePoint(request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, response)
	}
}