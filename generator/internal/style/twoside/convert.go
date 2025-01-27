package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
	"log"
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

		// about me
		AboutMeLabel: customDefaultString(app.Profile.CustomAboutMeTextLabel, lang, lang.Profile()),
		AboutMe:      util.DefaultValue(app.JobPosition.ProfileText, lang.String(app.Profile.GeneralAboutMeText)),

		// motivation
		MyMotivationLabel: customDefaultString(app.Profile.CustomMyMotivationTextLabel, lang, lang.Motivation()),
		MyMotivation:      util.DefaultValue(app.JobPosition.MyMotivationText, lang.String(app.Profile.GeneralMyMotivationText)),

		// time amount
		TimeAmountLabel: customDefaultString(app.Profile.CustomMyMotivationTextLabel, lang, lang.Motivation()),
		TimeAmounts:     util.DefaultLeftSideList(app.JobPosition.TimeAmount, app.Profile.TimeAmount),

		// main question
		MainQuestionLabel: customDefaultString(app.Profile.CustomMainQuestionTextLabel, lang, lang.MainQuestion()),
		MainQuestion:      util.DefaultValue(app.JobPosition.MainQuestionText, lang.String(app.Profile.GeneralMainQuestionText)),

		Attachment: util.DefaultStringArray(app.JobPosition.Attachment, app.Profile.Attachment),
	}

	addLsa(&data, app, lang)
	addExperiences(&data, app, lang)
	addEducation(&data, app, lang)
	log.Printf("addPublication before")
	addPublication(&data, app, lang)
	addAward(&data, app, lang)

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

func addHasDocumentLinks(data *texmodel.Index) {
	// TODO improve
	if checkTwoExperience(data, func(e texmodel.Experience) string {
		if len(e.DocumentLinks) > 0 {
			return e.DocumentLinks[0]
		}
		return ""
	}) || checkTwoEducation(data, func(e texmodel.Education) string {
		if len(e.DocumentLinks) > 0 {
			return e.DocumentLinks[0]
		}
		return ""
	}) || checkTwoPublication(data, func(e texmodel.Publication) string {
		if len(e.DocumentLinks) > 0 {
			return e.DocumentLinks[0]
		}
		return ""
	}) || checkTwoAward(data, func(e texmodel.Award) string {
		if len(e.DocumentLinks) > 0 {
			return e.DocumentLinks[0]
		}
		return ""
	}) {
		data.HasDocumentLinks = true
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

func checkTwoEducation(data *texmodel.Index, getter func(e texmodel.Education) string) bool {
	if checkEducation(data.SideOneEducation, getter) {
		return true
	}
	if checkEducation(data.SideTwoEducation, getter) {
		return true
	}
	return false
}

func checkEducation(exps []texmodel.Education, getter func(e texmodel.Education) string) bool {
	for _, e := range exps {
		if len(getter(e)) > 0 {
			return true
		}
	}
	return false
}

func checkTwoPublication(data *texmodel.Index, getter func(e texmodel.Publication) string) bool {
	if checkPublication(data.SideOnePublication, getter) {
		return true
	}
	if checkPublication(data.SideTwoPublication, getter) {
		return true
	}
	return false
}

func checkPublication(exps []texmodel.Publication, getter func(e texmodel.Publication) string) bool {
	for _, e := range exps {
		if len(getter(e)) > 0 {
			return true
		}
	}
	return false
}

func checkTwoAward(data *texmodel.Index, getter func(e texmodel.Award) string) bool {
	if checkAward(data.SideOneAward, getter) {
		return true
	}
	if checkAward(data.SideTwoAward, getter) {
		return true
	}
	return false
}

func checkAward(exps []texmodel.Award, getter func(e texmodel.Award) string) bool {
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
