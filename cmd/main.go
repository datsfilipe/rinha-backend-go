package main

import "github.com/datsfilipe/rinha-backend-go/pkg/routes"

func main() {
	r := routes.SetupRouter()

	r.Run(":80")
}
