package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
)

type RuleRunRequest struct {
	Exp    string                 `json:"exp"`
	Params map[string]interface{} `json:"params"`
}

func getParams(param map[string]interface{}) (map[string]interface{}, error) {
	newParams := make(map[string]interface{})

	paramsByte, _ := json.Marshal(param)
	decoder := json.NewDecoder(strings.NewReader(string(paramsByte)))
	decoder.UseNumber()
	if err := decoder.Decode(&newParams); err != nil {
		return nil, fmt.Errorf("invalid request params: %v", err)
	}
	for k, v := range newParams {
		// 用最优的方式断定类型。
		if num, ok := v.(json.Number); ok {
			var err error
			if strings.Contains(num.String(), ".") || strings.Contains(num.String(), "e") {
				newParams[k], err = num.Float64()
			} else {
				newParams[k], err = num.Int64()
				if err != nil {
					newParams[k], err = num.Float64()
				}
			}
			if err != nil {
				return nil, fmt.Errorf("invalid request params: %v", err)
			}
		}
	}

	return newParams, nil
}

func HandleRunRule(ctx context.Context, c *app.RequestContext) {
	var req RuleRunRequest
	if err := c.Bind(&req); err != nil {
		BindResp(c, ParamErrCode, err.Error(), nil)
		return
	}

	evaluatedExp, err := Compiler(req.Exp)
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
