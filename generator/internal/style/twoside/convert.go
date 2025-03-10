package twoside

import (
	"log"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
)

func style(app *genmodel.Application) genmodel.TwoSideStyle {
	return app.JobPosition.TwoSideStyle
}

func initData(app *genmodel.Application, outputPath string) (data texmodel.Index, err error) {
	local := app.JobPosition.Lang
	style := style(app)

	nationality := local.String(app.Profile.Nationality)
	if local.String(app.Profile.NationalityMig) != "" && app.JobPosition.UseMig {
		nationality = local.String(app.Profile.NationalityMig)
	}

	var attachments []string
	if !style.NoAttachments {
		attachments = util.DefaultStringArray(app.JobPosition.Attachment, app.Profile.Attachment)
	}

	data = texmodel.Index{
		// config
		MainColor:   util.GetSpecificDefaultColor(app.JobPosition.MainColor, util.DefaultMainColorValue),
		SideColor:   util.GetSpecificDefaultColor(app.JobPosition.SideColor, util.DefaultSideColorValue),
		ShadowColor: util.GetSpecificDefaultColor(app.JobPosition.ShadowColor, util.DefaultShadowColorValue),
		ScaleLineBG: util.GetSpecificDefaultColor(app.JobPosition.ScaleLineBG, util.DefaultScaleBgColorValue),
		Color1:      app.JobPosition.Color1,
		Color2:      app.JobPosition.Color2,
		Color3:      app.JobPosition.Color3,
		Color4:      app.JobPosition.Color4,
		Color5:      app.JobPosition.Color5,
		Label:       local,
		Icon: texmodel.Icon{
			Project:   util.ProjectIconPath,
			Role:      util.RoleIconPath,
			TechStack: util.TechStackIconPath,
			Document:  util.DocumentIconPath,
			ExtLink:   util.ExtLinkIconPath,
		},

		// main infos
		Name:        util.JoinStrings(app.Profile.FirstName, " ", app.Profile.LastName),
		Title:       util.DefaultValue(app.JobPosition.Title, local.String(app.Profile.Title)),
		Image:       util.DefaultImage(app.Profile.Image),
		Email:       app.Profile.Email,
		Phone:       app.Profile.Phone,
		Location:    util.JoinStrings(local.String(app.Profile.Address.City), ", ", local.String(app.Profile.Address.Country)),
		Nationality: nationality,
		Websites:    convertWebsites(app),

		AttachmentAndHintsPage: !style.NoAttachmentAndHintsPage,
		Attachment:             attachments,
	}

	checkColor := func(color *string, lsaLine string) {
		if *color == "" {
			*color = "ffffff"
		} else {
			data.LSALineLatex = lsaLine
		}
	}
	checkColor(&data.Color1, "")
	checkColor(&data.Color2, `\twoColorLine`)
	checkColor(&data.Color3, `\threeColorLine`)
	checkColor(&data.Color4, `\fourColorLine`)
	checkColor(&data.Color5, `\fiveColorLine`)

	addCoverLetter(&data, app, outputPath, local)
	addLsa(&data, app, local)
	addRSA(&data, app, local)

	addHasDocumentLinks(&data)
	addHasExternalLinks(&data)
	addHasProjects(&data)
	addHasRole(&data)
	addHasTechStack(&data)

	hasCount := 0
	if data.HasDocumentLinks {
		hasCount += 1
	}
	if data.HasExternalLinks {
		hasCount += 1
	}
	if data.HasProject {
		hasCount += 1
	}
	if data.HasRole {
		hasCount += 1
	}
	if data.HasTechStack {
		hasCount += 1
	}
	data.HasLegendCount = hasCount

	log.Printf("data after legend checks %+v\n", data)

	return
}

func convertWebsites(application *genmodel.Application) (websites []texmodel.Website) {
	for _, website := range application.Profile.Websites {
		websites = append(websites, texmodel.Website{
			Icon:        util.WebsiteIcon(website),
			FontAwesome: util.WebsiteFontAwesome(website),
			Url:         website,
		})
	}
	return
}

func addHasProjects(data *texmodel.Index) {
	log.Printf("addHasProjects")
	if checkTwoRSA(data, func(e texmodel.RSA) bool {
		return e.HasProjects
	}) {
		data.HasProject = true
	}
}

func addHasRole(data *texmodel.Index) {
	if checkTwoRSA(data, func(e texmodel.RSA) bool {
		return e.HasRole
	}) {
		data.HasRole = true
	}
}

func addHasTechStack(data *texmodel.Index) {
	if checkTwoRSA(data, func(e texmodel.RSA) bool {
		return e.HasTechStack
	}) {
		data.HasTechStack = true
	}
}

func addHasExternalLinks(data *texmodel.Index) {
	if checkTwoRSA(data, func(e texmodel.RSA) bool {
		return e.HasExternalLinks
	}) {
		data.HasExternalLinks = true
	}
}

func addHasDocumentLinks(data *texmodel.Index) {
	if checkTwoRSA(data, func(e texmodel.RSA) bool {
		return e.HasDocumentLinks
	}) {
		data.HasDocumentLinks = true
	}
}

func checkTwoRSA(data *texmodel.Index, getter func(e texmodel.RSA) bool) bool {
	if checkRSA(data.SideOneRSA, getter) {
		return true
	}
	if checkRSA(data.SideTwoRSA, getter) {
		return true
	}
	return false
}

func checkRSA(exps []texmodel.RSA, getter func(e texmodel.RSA) bool) bool {
	for _, e := range exps {
		if getter(e) {
			return true
		}
	}
	return false
}

func customDefaultString(customMap lang.TranslationMap, lang lang.Language, def string) string {
	custom := lang.String(customMap)
	if custom == "" {
		custom = def
	}
	return custom
}

func customCustomDefaultString(custom string, customMap2 lang.TranslationMap, lang lang.Language, def string) string {
	if custom == "" {
		custom = lang.String(customMap2)
	}
	if custom == "" {
		custom = def
	}
	return custom
}

func findId(checkId genmodel.ID, ids []genmodel.ID) bool {
	for _, id := range ids {
		if checkId == id {
			return true
		}
	}
	return false
}
