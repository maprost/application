package texmodel

import (
	"github.com/maprost/application/generator/lang"
)

type Icon struct {
	Project   string
	Role      string
	TechStack string
	Document  string
	ExtLink   string
}

type Website struct {
	Icon string // path
	Url  string
}

type Index struct {
	// config
	MainColor        string
	SideColor        string
	HasDocumentLinks bool
	HasProject       bool
	HasRole          bool
	HasTechStack     bool
	Icon             Icon
	Label            lang.Language

	// main infos
	Image       string // path
	Name        string
	Title       string
	Nationality string
	Location    string
	Email       string
	Phone       string
	Websites    []Website

	SideOneLSA string
	SideTwoLSA string

	SideOneRSA []RSA
	SideTwoRSA []RSA

	Attachment []string
}

func (x Index) HasSideTwo() bool {
	return len(x.SideTwoLSA) > 0 ||
		len(x.SideTwoRSA) > 0
}
