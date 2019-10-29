// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package render

import (
	"github.com/micro/go-micro/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CODE_OK                     = 0
	CODE_ERR_SYSTEM             = 1000
	CODE_ERR_APP                = 1001
	CODE_ERR_PARAM              = 1002
	CODE_ERR_DATA_REPEAT        = 1003
	CODE_ERR_LOGIN_FAILED       = 1004
	CODE_ERR_NO_LOGIN           = 1005
	CODE_ERR_NO_PRIV            = 1006
	CODE_ERR_TASK_ERROR         = 1007
	CODE_ERR_USER_OR_PASS_WRONG = 1008
	CODE_ERR_NO_DATA            = 1009
)

func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   CODE_OK,
		"detail": "success",
		"data":   data,
	})
}

func CustomerError(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"detail": message,
	})
}

func RepeatError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":   CODE_ERR_DATA_REPEAT,
		"detail": message,
	})
}

func NoDataError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":   CODE_ERR_NO_DATA,
		"detail": message,
	})
}

func ParamError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":   CODE_ERR_PARAM,
		"detail": message,
	})
}

func AppError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":   CODE_ERR_APP,
		"detail": message,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":   CODE_OK,
		"detail": "success",
	})
}

func MicroParse(c *gin.Context, message string) {
	err := errors.Parse(message)
	if err.Code == 0 {
		c.JSON(http.StatusOK, gin.H{
			"id":     err.Id,
			"code":   500,
			"detail": err.Detail,
			"status": err.Status,
		})
		return
	}
	c.JSON(http.StatusOK, err)
}
