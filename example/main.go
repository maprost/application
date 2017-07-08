package main

import (
	"flag"
	"log"
	"os"

	"github.com/maprost/application/example/july"
	"github.com/maprost/application/example/max"
	"github.com/maprost/application/generator"
	"github.com/maprost/application/generator/genmodel"
)

func main() {
	flag.Parse()

	var name string
	var company string

	if len(flag.Args()) == 2 {
		name = flag.Arg(0)
		company = flag.Arg(1)
	} else {
		name = "max"
		company = "santa"
	}

	log.Println("Name:", name, " Company:", company)

	var profile genmodel.Profile
	switch name {
	case "max":
		profile = max.Profile(company)
	case "july":
		profile = july.Profile(company)
	}

	err := generator.Build(profile)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
