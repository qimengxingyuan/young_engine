package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func HandleAddExpression(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, utils.H{
		"message": "please implement your code logic",
	})
}

func HandleDeleteExpression(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, utils.H{
		"message": "please implement your code logic",
	})
}

func HandleGetAllExpression(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, utils.H{
		"message": "please implement your code logic",
	})
}

func HandleRunExpression(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, utils.H{
		"message": "please implement your code logic",
	})
}
