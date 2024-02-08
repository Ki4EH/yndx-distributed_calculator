package database

import (
	"fmt"
	"github.com/Ki4EH/yndx-distributed_calculator/internal/logger"
	"github.com/Ki4EH/yndx-distributed_calculator/internal/models"
)

type ExpressionDB struct {
	expList map[string]models.Expression
}

var log, _ = logger.NewFileLogger()

func Find(ID string) bool {
	for id := range ExpList.expList {
		if ID == id {
			return true
		}
	}
	return false
}

func Add(ID string, expressionInfo *models.Expression) {
	log.Info(fmt.Sprintf("[DATABASE] Add new expression info, id: %s", ID))
	ExpList.expList[ID] = *expressionInfo
}

var ExpList = NewExpressionDB()

func NewExpressionDB() *ExpressionDB {
	return &ExpressionDB{expList: map[string]models.Expression{}}
}
