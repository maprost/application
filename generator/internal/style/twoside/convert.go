package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
)

func style(app *genmodel.Application) genmodel.TwoSideStyle {
	return app.JobPosition.TwoSideStyle
}

func initData(app *genmodel.Application) (data texmodel.Index, err error) {
	local := app.JobPosition.Lang

	data = texmodel.Index{
		// config
		MainColor: util.DefaultColor(app.JobPosition.MainColor),
		SideColor: util.DefaultColor(app.JobPosition.SideColor),
		Label:     local,
		Icon: texmodel.Icon{
			Project:   util.ProjectIconPath,
			Role:      util.RoleIconPath,
			TechStack: util.TechStackIconPath,
			Document:  util.DocumentIconPath,
		},

		// main infos
		Name:        util.JoinStrings(app.Profile.FirstName, " ", app.Profile.LastName),
		Title:       util.DefaultValue(app.JobPosition.Title, app.Profile.Title),
		Image:       util.DefaultImage(app.Profile.Image),
		Email:       app.Profile.Email,
		Phone:       app.Profile.Phone,
		Location:    util.JoinStrings(local.String(app.Profile.Address.City), ", ", local.String(app.Profile.Address.Country)),
		Nationality: local.String(app.Profile.Nationality),
		Websites:    convertWebsites(app),

		Attachment: util.DefaultStringArray(app.JobPosition.Attachment, app.Profile.Attachment),
	}

	addLsa(&data, app, local)
	addRSA(&data, app, local)

	addHasDocumentLinks(&data)
	addHasProjects(&data)
	addHasRole(&data)
	addHasTechStack(&data)

	return
}

func convertWebsites(application *genmodel.Application) (websites []texmodel.Website) {
	for _, website := range application.Profile.Websites {
		websites = append(websites, texmodel.Website{
			Icon: util.WebsiteIcon(website),
			Url:  website,
		})
	}
	return
}

func addHasProjects(data *texmodel.Index) {
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

func findId(checkId genmodel.ID, ids []genmodel.ID) bool {
	for _, id := range ids {
		if checkId == id {
			return true
		}
	}
	return false
}
