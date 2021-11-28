package main

import "mini-node-exporter/services/routes"

func main() {
	r := routes.NewRouter()
	r.Run(":23333")
}
