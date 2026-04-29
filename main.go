package main

import (
	"webscaper/db"
	"webscaper/templetes"
)

func main() {

	db.Init(".env")

	templetes.InitHTML()

}
