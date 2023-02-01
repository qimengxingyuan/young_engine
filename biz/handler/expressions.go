package handler

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qimengxingyuan/young_engine/biz/dal"
)

// Add expression
type AddExpressionMessage struct {
	ID int `json:"id"`
}

func HandleAddExpression(ctx context.Context, c *app.RequestContext) {
	var req RuleRunRequest
	if err := c.Bind(&req); err != nil {
		BindResp(c, ParamErrCode, err.Error(), nil)
		return
	}

	_, err := Compiler(req.Exp)
	if err != nil {
		BindResp(c, CompileErrCode, err.Error(), &AddExpressionMessage{ID: 0})
		return
	}

	dal.AddExpression(req.Exp)
	exp, _ := dal.GetGetExpressionByExp(req.Exp)

	BindResp(c, SuccessCode, SuccessMsg, &AddExpressionMessage{ID: int(exp.ID)})
}

type ExpressionMessage struct {
	ID  int    `json:"id"`
	Exp string `json:"exp"`
}

// Delete expressions by ID
func HandleDeleteExpression(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		BindResp(c, RuleNotExistCode, "id is not an integer", nil)
		return
	}

	exp, err := dal.GetExpressionByID(uint(id))
	if err != nil {
		BindResp(c, RuleNotExistCode, err.Error(), nil)
		return
	}

	dal.DeleteExpressionByID(uint(id))
	BindResp(c, SuccessCode, SuccessMsg, &ExpressionMessage{ID: id, Exp: exp.Exp})
}

// Get all expressions
func HandleGetAllExpression(ctx context.Context, c *app.RequestContext) {
	allExpression, err := dal.GetAllExpression()
	if err != nil {
		BindResp(c, -1, err.Error(), &AddExpressionMessage{ID: 0})
		return
	}

	var res []ExpressionMessage
	for _, exp := range allExpression {
		tmp := ExpressionMessage{
			ID:  int(exp.ID),
			Exp: exp.Exp,
		}

		res = append(res, tmp)
	}

	BindResp(c, SuccessCode, SuccessMsg, res)
}

// Run expression by id
type RunRequest struct {
	ID     uint                   `json:"exp_id"`
	Params map[string]interface{} `json:"params"`
}

func HandleRunExpression(ctx context.Context, c *app.RequestContext) {
	var req RunRequest
	if err := c.Bind(&req); err != nil {
		BindResp(c, ParamErrCode, err.Error(), nil)
		return
	}

	exp, err := dal.GetExpressionByID(req.ID)
	if err != nil {
		BindResp(c, RuleNotExistCode, err.Error(), nil)
		return
	}

	evaluatedExp, _ := Compiler(exp.Exp)
	if err != nil {
		BindResp(c, CompileErrCode, err.Error(), nil)
		return
	}

	params, _ := getParams(req.Params)
	err = evaluatedExp.Eval(params)
	if err != nil {
		BindResp(c, RuleExecErrCode, err.Error(), nil)
		return
	}

	resp, _ := evaluatedExp.GetVal()

	BindResp(c, SuccessCode, SuccessMsg, resp)
}
