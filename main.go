package main

import (
	"championship-football/championship"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	championship.Run()
}
