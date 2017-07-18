package main

import (
	"flag"
	"log"
	"os"

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
		company = "bunny"
	}

	log.Println("Name:", name, " JobPosition:", company)

	var application genmodel.Application
	switch name {
	case "max":
		application = max.Application(company)
	}

	err := generator.Build(application, generator.TwoSide)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
