package twoside_test

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	style2 "github.com/maprost/application/generator/style"
	"github.com/maprost/testbox/must"
	"github.com/maprost/testbox/should"
	"testing"
)

func TestConvert(t *testing.T) {
	const outputPath = "myFolder"
	run := func(t *testing.T, profile genmodel.Profile, jobPosition genmodel.JobPosition) texmodel.Index {
		style := style2.TwoSide
		iData, err := style.Data(&genmodel.Application{
			Profile:     profile,
			JobPosition: jobPosition,
		}, outputPath)
		must.BeNoError(t, err)
		data, ok := iData.(texmodel.Index)
		must.BeTrue(t, ok)
		return data
	}

	t.Run("check attachments", func(t *testing.T) {
		t.Run("normal", func(t *testing.T) {
			profile := genmodel.Profile{
				Attachment: []string{"file1", "file2"},
			}
			jobProfile := genmodel.JobPosition{}
			data := run(t, profile, jobProfile)
			must.HaveLength(t, data.Attachment, 2)
			should.BeEqual(t, data.Attachment[0], "file1")
			should.BeEqual(t, data.Attachment[1], "file2")

		})
		t.Run("empty", func(t *testing.T) {
			profile := genmodel.Profile{
				Attachment: nil,
			}
			jobProfile := genmodel.JobPosition{}
			data := run(t, profile, jobProfile)
			should.HaveLength(t, data.Attachment, 0)

		})
		t.Run("style.NoAttachments = true", func(t *testing.T) {
			profile := genmodel.Profile{
				Attachment: []string{"file1", "file2"},
			}
			jobProfile := genmodel.JobPosition{
				TwoSideStyle: genmodel.TwoSideStyle{
					NoAttachments: true,
				},
			}
			data := run(t, profile, jobProfile)
			should.HaveLength(t, data.Attachment, 0)
		})
	})
}
