package main

import (
	"fmt"

	"github.com/leekchan/accounting"
	"github.com/matkinhig/echo-fw/config"
)

func main() {
	fmt.Println("start golang")

	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))

	fmt.Println("The Config value:", config.Config)

}
