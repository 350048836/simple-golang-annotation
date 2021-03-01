package main

import (
	"./amyxframework"
	"./controller"
)

func main() {
	amyxframework.Start(amyxframework.StartParams{
		Application: main,
		APIs:        []interface{}{controller.GetTest, controller.PostTest},
	})
}
