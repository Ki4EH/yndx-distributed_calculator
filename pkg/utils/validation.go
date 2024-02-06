package utils

import (
	"fmt"
	"github.com/Ki4EH/yndx-distributed_calculator/internal/logger"
	"regexp"
)

var loggerInst logger.Logger

// IsValidExpression проверяем валидность отправленного выражения
func IsValidExpression(expression string) bool {
	regex := regexp.MustCompile(`^[0-9+\-*/().\s]+$`)
	loggerInst.Info(fmt.Sprintf("[VALID/EXP] expression {%s} is %t", expression, regex.MatchString(expression)))
	return regex.MatchString(expression)
}
