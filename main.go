package main

import (
	"fmt"

	"github.com/yasseresi/Gator/internal/config"
)

func main() {

	conf := config.ReadConfig()
	fmt.Println(("DbUrl: " + conf.DbUrl))
	fmt.Println(("UserName before: " + conf.UserName))

	conf.UserName = "Yasser"
	if err := config.Write(conf); err != nil {
		fmt.Println("Error in writing file")
	}
	conf = config.ReadConfig()
	fmt.Println(("DbUrl: " + conf.DbUrl))
	fmt.Println(("UserName after: " + conf.UserName))

}
