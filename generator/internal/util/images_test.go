package util_test

import (
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/assertion"
	"testing"
)

func TestWebsiteIcons(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(util.WebsiteIcon("https://www.linkedin.com/myname"), util.LinkedinIconPath)
	assert.Equal(util.WebsiteIcon("https://github.com/myaccount"), util.GithubIconPath)
	assert.Equal(util.WebsiteIcon("https://www.xing.com/myself"), util.XingIconPath)
	assert.Equal(util.WebsiteIcon("https://mywebsite.de"), util.WebsiteIconPath)
}

func TestWebsiteIcons_onlyHttp(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(util.WebsiteIcon("http://www.linkedin.com/myname"), util.LinkedinIconPath)
	assert.Equal(util.WebsiteIcon("http://github.com/myaccount"), util.GithubIconPath)
	assert.Equal(util.WebsiteIcon("http://www.xing.com/myself"), util.XingIconPath)
}

func TestWebsiteIcons_noHttp(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(util.WebsiteIcon("www.linkedin.com/myname"), util.LinkedinIconPath)
	assert.Equal(util.WebsiteIcon("github.com/myaccount"), util.GithubIconPath)
	assert.Equal(util.WebsiteIcon("www.xing.com/myself"), util.XingIconPath)
}

func TestWebsiteIcons_noWWW(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(util.WebsiteIcon("https://linkedin.com/myname"), util.LinkedinIconPath)
	assert.Equal(util.WebsiteIcon("https://xing.com/myself"), util.XingIconPath)
}

func TestWebsiteIcons_noHttp_noWWW(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(util.WebsiteIcon("linkedin.com/myname"), util.LinkedinIconPath)
	assert.Equal(util.WebsiteIcon("github.com/myaccount"), util.GithubIconPath)
	assert.Equal(util.WebsiteIcon("xing.com/myself"), util.XingIconPath)
}

func TestWebsiteIcons_withUpperCase(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(util.WebsiteIcon("http://wWw.LinkedIn.Com/myname"), util.LinkedinIconPath)
	assert.Equal(util.WebsiteIcon("http://GitHub.com/myaccount"), util.GithubIconPath)
	assert.Equal(util.WebsiteIcon("HTTP://www.Xing.com/myself"), util.XingIconPath)
}
