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
	Icon        string // path
	FontAwesome string
	Url         string
}

type Index struct {
	// config
	MainColor        string
	SideColor        string
	ShadowColor      string
	LSALineLatex     string
	Color1           string
	Color2           string
	Color3           string
	Color4           string
	Color5           string
	ScaleLineBG      string
	HasDocumentLinks bool
	HasExternalLinks bool
	HasProject       bool
	HasRole          bool
	HasTechStack     bool
	HasLegendCount   int
	Icon             Icon
	Label            lang.Language

	// cover letter
	CoverLetter string

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

	AttachmentAndHintsPage bool
	Attachment             []string
}

func (x Index) HasSideOne() bool {
	return len(x.SideOneLSA) > 0 ||
		len(x.SideOneRSA) > 0
}

func (x Index) HasSideTwo() bool {
	return len(x.SideTwoLSA) > 0 ||
		len(x.SideTwoRSA) > 0
}
