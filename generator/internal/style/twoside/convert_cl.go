package twoside

import (
	"fmt"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/compiler"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
)

func addCoverLetter(data *texmodel.Index, app *genmodel.Application, outputPath string, local lang.Language) {
	style := app.JobPosition.TwoSideStyle
	if !style.ActivateCoverLetter {
		return
	}

	var res texmodel.CoverLetter
	res.JobAddress = texmodel.Address{
		Name:   app.JobPosition.Company,
		Street: app.JobPosition.Address.Street,
		Zip:    local.CityZip(app.JobPosition.Address.City, app.JobPosition.Address.Zip),
	}
	res.MyAddress = texmodel.Address{
		Name:   fmt.Sprintf("%s %s", app.Profile.FirstName, app.Profile.LastName),
		Phone:  app.Profile.Phone,
		Email:  app.Profile.Email,
		Street: app.Profile.Address.Street,
		Zip:    local.CityZip(local.String(app.Profile.Address.City), app.Profile.Address.Zip),
	}
	res.Subject = util.DefaultValue(app.JobPosition.CoverLetterSubject, local.String(app.Profile.CoverLetterSubject))
	res.Text = util.DefaultValue(app.JobPosition.CoverLetterTxt, local.String(app.Profile.CoverLetterTxt))
	res.Date = local.Date()

	if len(app.JobPosition.CoverLetterAttachment) > 0 {
		for _, att := range app.JobPosition.CoverLetterAttachment {
			res.AttachmentList = append(res.AttachmentList, att)
		}
	} else {
		for _, att := range app.Profile.CoverLetterAttachment {
			res.AttachmentList = append(res.AttachmentList, local.String(att))
		}
	}

	clText, err := compiler.CompileSubTex(templatePath(), "cover_letter.tex", res)
	if err != nil {
		clText = "error: " + err.Error()
	}

	if style.CoverLetterOnSeparatePdf {
		err := buildCoverLetter(*data, app, clText, outputPath) // data texmodel will copy here, no interaction!
		if err != nil {
			panic(err)
		}
	} else {
		data.CoverLetter = clText
	}
}

func buildCoverLetter(data texmodel.Index, app *genmodel.Application, clText string, outputPath string) error {
	// prep data (no first + second page)
	data.SideOneLSA = ""
	data.SideTwoLSA = ""
	data.SideOneRSA = nil
	data.SideTwoRSA = nil
	data.Attachment = nil
	data.CoverLetter = clText

	outputPath, file := compiler.GenerateOutput(outputPath, app)
	file += "_cl"

	path, mainFile, subFiles := Files()
	err := compiler.CreateTexFile(outputPath, file, data, path, mainFile, subFiles...)
	if err != nil {
		return err
	}

	err = compiler.Compile(outputPath, file)
	if err != nil {
		return err
	}

	err = compiler.CleanUp(outputPath, file)
	if err != nil {
		return err
	}

	return nil
}
