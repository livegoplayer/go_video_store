package config

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/livegoplayer/go_helper/validate"
)

func InitValidate() {
	//更换校验器
	binding.Validator = validate.ValidatorV10
}
