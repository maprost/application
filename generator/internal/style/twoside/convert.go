package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
)

func initData(app *genmodel.Application) (data texmodel.Index, err error) {
	lang := app.JobPosition.Lang

	data = texmodel.Index{
		// config
		MainColor: util.DefaultColor(app.JobPosition.MainColor),
		SideColor: util.DefaultColor(app.JobPosition.SideColor),
		Label:     lang,
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
		Location:    util.JoinStrings(lang.String(app.Profile.Address.City), ", ", lang.String(app.Profile.Address.Country)),
		Nationality: lang.String(app.Profile.Nationality),
		Websites:    convertWebsites(app),

		// motivation
		AboutMeLabel: customDefaultString(app.Profile.CustomMotivationTextLabel, lang, lang.Profile()),
		AboutMe:      util.DefaultValue(app.JobPosition.MotivationText, lang.String(app.Profile.GeneralMotivationText)),

		Attachment: util.DefaultStringArray(app.JobPosition.Attachment, app.Profile.Attachment),
	}

	addLsa(&data, app, lang)
	addExperiences(&data, app, lang)
	addEducation(&data, app, lang)

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
	if checkTwoExperience(data, func(e texmodel.Experience) string {
		return e.Project
	}) {
		data.HasProject = true
	}
}

func addHasRole(data *texmodel.Index) {
	if checkTwoExperience(data, func(e texmodel.Experience) string {
		return e.Role
	}) {
		data.HasRole = true
	}
}

func addHasTechStack(data *texmodel.Index) {
	if checkTwoExperience(data, func(e texmodel.Experience) string {
		return e.Tech
	}) {
		data.HasTechStack = true
	}
}

func checkTwoExperience(data *texmodel.Index, getter func(e texmodel.Experience) string) bool {
	if checkExperience(data.SideOneExperience, getter) {
		return true
	}
	if checkExperience(data.SideTwoExperience, getter) {
		return true
	}
	return false
}

func checkExperience(exps []texmodel.Experience, getter func(e texmodel.Experience) string) bool {
	for _, e := range exps {
		if len(getter(e)) > 0 {
			return true
		}
	}
	return false
}

func convertLeftSideActions(application *genmodel.Application) []int {
	list := application.Profile.LeftSideActionType
	if len(list) == 0 {
		return []int{
			int(genmodel.TechSkill),
			int(genmodel.Interests),
			int(genmodel.SoftSkills),
			int(genmodel.Languages),
			int(genmodel.Hobbies),
		}
	}

	res := make([]int, 0, len(list))
	for _, lsa := range list {
		res = append(res, int(lsa))
	}

	return res
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
