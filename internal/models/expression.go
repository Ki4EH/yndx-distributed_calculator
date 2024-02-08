package models

type Expression struct {
	ID     string  `json:"id"`
	Value  string  `json:"value"`
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

func NewExpression() *Expression {
	return &Expression{
		ID:     "",
		Value:  "",
		Status: "",
		Result: 0,
	}
}
