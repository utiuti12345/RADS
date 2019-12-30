package main

import (
	"RosterAutomaticDeliverySystem/config"
	"RosterAutomaticDeliverySystem/infra/web"
	"flag"
)

var (
	f = flag.String("f", "", "credentials path")
)

func main(){
	flag.Parse()
	gc := config.NewGoogleConfig(*f, *f)
	web.LoadRouter(gc).Start(":1323")
}