package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
type (
	AppError struct {
		Code     string `json:"code"`
		Message  string `json:"message"`
	}
)



func (appError *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", appError.Code, appError.Message)
}

//
// Middleware Error Handler in server package
//
func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *AppError
			switch err.(type) {
			case *AppError:
				parsedError = err.(*AppError)
				break
			default:
				parsedError = &AppError{
					Code:  "ERR???",
					Message: err.Error(),
				}
			}
			// Put the error into response
			if parsedError.Code == "" {
				parsedError.Code = "ERR???"
			}

			c.JSON(400, parsedError)
			c.Abort()
			// or c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}

	}

}