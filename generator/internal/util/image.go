package util

import (
	"github.com/maprost/application/generator/internal/image"
	"strings"
)

var (
	WebsiteIconPath   = image.ImagePath() + "website"
	XingIconPath      = image.ImagePath() + "xing"
	GithubIconPath    = image.ImagePath() + "github"
	LinkedinIconPath  = image.ImagePath() + "linkedIn"
	TechStackIconPath = image.ImagePath() + "techstack"
	ProjectIconPath   = image.ImagePath() + "project"
	RoleIconPath      = image.ImagePath() + "role"
	DocumentIconPath  = image.ImagePath() + "document"
	ExtLinkIconPath   = image.ImagePath() + "extlink"
)

func WebsiteIcon(website string) (icon string) {
	icon = `\faWebsite`
	website = strings.ToLower(website)

	// repair urls
	containsHttpPrefix := strings.HasPrefix(website, "http")
	containsHttpsPrefix := strings.HasPrefix(website, "https")
	if containsHttpPrefix && !containsHttpsPrefix {
		website = "https" + website[4:]
	} else if !containsHttpPrefix {
		website = "https://" + website
	}

	if strings.HasPrefix(website, "https://github.com") {
		icon = GithubIconPath

	} else if strings.HasPrefix(website, "https://www.linkedin.com") || strings.HasPrefix(website, "https://linkedin.com") {
		icon = LinkedinIconPath

	} else if strings.HasPrefix(website, "https://www.xing.com") || strings.HasPrefix(website, "https://xing.com") {
		icon = XingIconPath
	}

	return
}

func WebsiteFontAwesome(website string) (faIcon string) {
	faIcon = `\faGlobe`
	website = strings.ToLower(website)

	// repair urls
	containsHttpPrefix := strings.HasPrefix(website, "http")
	containsHttpsPrefix := strings.HasPrefix(website, "https")
	if containsHttpPrefix && !containsHttpsPrefix {
		website = "https" + website[4:]
	} else if !containsHttpPrefix {
		website = "https://" + website
	}

	if strings.HasPrefix(website, "https://github.com") {
		faIcon = `\faGithub`
		//faIcon = `\faGithubSquare`

	} else if strings.HasPrefix(website, "https://www.linkedin.com") || strings.HasPrefix(website, "https://linkedin.com") {
		faIcon = `\faLinkedin`
		//faIcon = `\faLinkedinSquare`

	} else if strings.HasPrefix(website, "https://www.xing.com") || strings.HasPrefix(website, "https://xing.com") {
		faIcon = `\faXing`
		//faIcon = `\faXingSquare`
	}

	return
}
