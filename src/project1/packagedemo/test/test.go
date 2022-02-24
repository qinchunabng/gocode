package test

import (
	"project1/packagedemo/utils" //循环引用会报错
)

var Test = utils.Test
