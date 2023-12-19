package main

import "amartha/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
