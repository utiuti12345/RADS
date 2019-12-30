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
	sc,_ := config.LoadFile(*f)
	gc := config.NewGoogleConfig(*f, *f)
	c := config.NewConfig(gc,sc)
	web.LoadRouter(c).Start(":1323")
}