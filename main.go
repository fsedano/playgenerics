package main

import (
	"log"

	"fsedano.net/gen/rest"
	"fsedano.net/gen/resttypes"
)

func main() {
	var d resttypes.Device
	var r resttypes.Resource

	err := rest.GetData("https://reqres.in/api/users?page=2", &d)
	if err != nil {
		log.Printf("error!!: %s", err)
	}
	for _, v := range d.Data {
		log.Printf("Email is: %s", v.Email)
	}

	err = rest.GetData("https://reqres.in/api/{resources}", &r)
	if err != nil {
		log.Printf("error!!: %s", err)
	}

	for _, w := range r.Data {
		log.Printf("Color: %s", w.Color)
	}

	err = rest.GetData("https://reqres.in/a1pi/{resources}", &r)
	if err != nil {
		log.Printf("error!!: %s", err)
	} else {

		for _, w := range r.Data {
			log.Printf("Color: %s", w.Color)
		}
	}
}
