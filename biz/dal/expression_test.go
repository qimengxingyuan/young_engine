package dal

import (
	"testing"
)

func init() {
	InitDB()
}

func TestAddExpression(t *testing.T) {
	exp, err := AddExpression("uid2 != 12345 && did > 0")
	t.Log(exp)
	t.Log(err)
}

func TestGetAllExpression(t *testing.T) {
	exps, err := GetAllExpression()
	t.Log(exps)
	t.Log(err)
}

func TestGetAllExpressionByID(t *testing.T) {
	exps, err := GetExpressionByID(1)
	t.Log(exps)
	t.Log(err)
}

func TestDeleteExpression(t *testing.T) {
	err := DeleteExpressionByID(1)
	t.Log(err)
}
