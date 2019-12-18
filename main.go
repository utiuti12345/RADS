package main

import "RosterAutomaticDeliverySystem/infra/web"

func main(){
	web.LoadRouter().Start(":1323")
}