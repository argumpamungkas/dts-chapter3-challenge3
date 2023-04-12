package main

import (
	"DTS/Chapter-3/chapter3-challenge3/repo"
	"DTS/Chapter-3/chapter3-challenge3/router"
)

func main() {

	repo.StartDB()

	router.StartServer().Run(":3000")

}
