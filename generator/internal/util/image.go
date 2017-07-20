package util

import (
	"github.com/maprost/application/generator/internal/image"
	"strings"
)

func WebsiteIcon(website string) (icon string) {
	imageRootPath := image.ImagePath()
	icon = imageRootPath + "website"

	if strings.HasPrefix(website, "https://github.com") {
		icon = imageRootPath + "github"

	} else if strings.HasPrefix(website, "https://www.linkedin.com") {
		icon = imageRootPath + "linkedIn"

	} else if strings.HasPrefix(website, "https://www.xing.com") {
		icon = imageRootPath + "xing"
	}

	return
}
