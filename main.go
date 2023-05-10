package main

import "demo1/conf"
import "demo1/router"
func main() {

	conf.Init()
	r:= router.NewRouter()
	r.Run(":3000")
}
