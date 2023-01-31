package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/techschool/simplebank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currencry, ok := fl.Field().Interface().(string); ok {
		// Check currency is supported
		return util.IsSupportedCurrency(currencry)
	}
	return false
}