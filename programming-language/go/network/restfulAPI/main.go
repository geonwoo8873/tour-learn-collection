package main

import (
	"github.com/gin-gonic/gin"
)

/*
200 OK
201 Created
400 Bad Request
401 Unauthorized
403 Forbidden
404 Not Found
500 Internal Server Error
*/

const (
	HTTPStatusOK                  = 200
	HTTPStatusCreated             = 201
	HTTPStatusBadRequest          = 400
	HTTPStatusUnauthorized        = 401
	HTTPStatusForbidden           = 403
	HTTPStatusNotFound            = 404
	HTTPStatusInternalServerError = 500
)

type AppError struct {
	Status int
	Code   string
	Msg    string
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		// `http.Status<Code>`대신 직접적으로 상태 코드를 사용하여 응답을 반환할 수 있다.
		ctx.JSON(HTTPStatusOK, gin.H{
			"status":  HTTPStatusOK,
			"message": "OK",
			"error":   nil,
		})
	})

	router.Run(":8080")
}

//// Result
/*
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2026/10/14 - 14:18:47 | 200 |       0s |             ::1 | GET      "/"
[GIN] 2026/10/14 - 14:18:47 | 404 |       0s |             ::1 | GET      "/favicon.ico"
*/
