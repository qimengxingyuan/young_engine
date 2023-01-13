package dal

type Expression struct {
	ID  uint   `json:"id"`
	Exp string `json:"exp"`
}

func AddExpression(exp string) (*Expression, error) {
	var res *Expression
	err := DB.Where("exp = ?", exp).Find(&res).Error
	if err == nil {
		if res.ID == 0 {
			res = &Expression{
				Exp: exp,
			}
			return res, DB.Save(res).Error
		}
		return res, nil
	} else {
		return nil, err
	}
}

func GetAllExpression() ([]*Expression, error) {
	var res []*Expression
	return res, DB.Model(&Expression{}).Find(&res).Error
}

func GetExpressionByID(id uint) (*Expression, error) {
	var res *Expression
	return res, DB.Model(&Expression{}).Where("id = ?", id).Find(&res).Error
}

func DeleteExpressionByID(id uint) error {
	return DB.Where("id = ?", id).Delete(&Expression{}).Error
}
