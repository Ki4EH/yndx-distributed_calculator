package utils

import (
	"fmt"
	"github.com/Ki4EH/yndx-distributed_calculator/internal/logger"
	"regexp"
	"strings"
)

// var loggerInst, _ = logger.NewFileLogger()
var loggerInst logger.Logger
var err error

func SetNewLogger() {
	loggerInst, err = logger.NewFileLogger()
	if err != nil {
		fmt.Println("Logger create file error ", err)
	}
}

// IsValidExpression проверяем валидность отправленного выражения
func IsValidExpression(expression string) bool {
	SetNewLogger()
	regex := regexp.MustCompile(`^[0-9+\-*/().\s]+$`)
	if !regex.MatchString(expression) {
		loggerInst.Info(fmt.Sprintf("[VALID/EXP] expression {%s} is false", expression))
		return false
	}
	expression = strings.ReplaceAll(expression, " ", "")
	for i := 0; i < len(expression)-1; i++ {
		if (expression[i] == '+' || expression[i] == '-' || expression[i] == '/' || expression[i] == '*') &&
			(expression[i+1] == '+' || expression[i+1] == '-' || expression[i+1] == '/' || expression[i+1] == '*') {
			loggerInst.Info(fmt.Sprintf("[VALID/EXP] expression {%s} is false", expression))
			return false
		}
	}
	loggerInst.Info(fmt.Sprintf("[VALID/EXP] expression {%s} is true", expression))
	return true
}
