package initialize

import (
	"github.com/go-playground/validator/v10"
	"github.com/hainguyen27798/open-search-engine/global"
)

func InitValidator() {
	global.Validate = validator.New(validator.WithRequiredStructEnabled())
}
