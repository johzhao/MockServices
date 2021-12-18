package encoder

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewJSONResponseEncoder() ResponseEncoder {
	return &JSONResponseEncoder{}
}

type JSONResponseEncoder struct {
}

func (e *JSONResponseEncoder) ResponseWithData(ctx *gin.Context, data interface{}) error {
	responseBody := make(map[string]interface{}, 3)
	responseBody["code"] = 0
	responseBody["message"] = "success"
	if data != nil {
		responseBody["data"] = data
	}

	ctx.JSON(http.StatusOK, responseBody)

	return nil
}

func (e *JSONResponseEncoder) ResponseWithError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": err.Error(),
	})
}
