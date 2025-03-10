package twoside_test

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/image"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	style2 "github.com/maprost/application/generator/style"
	"github.com/maprost/testbox/must"
	"github.com/maprost/testbox/should"
	"strings"
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

	containsRSA := func(t *testing.T, rsa texmodel.RSA, exp string) {
		should.Contain(t, strings.Join(rsa.TexList, "\n"), exp)
	}
	notContainsRSA := func(t *testing.T, rsa texmodel.RSA, exp string) {
		should.NotContain(t, strings.Join(rsa.TexList, "\n"), exp)
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
			should.BeTrue(t, data.AttachmentAndHintsPage)

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
		t.Run("style.NoAttachmentAndHintsPage = true", func(t *testing.T) {
			profile := genmodel.Profile{
				Attachment: []string{"file1", "file2"},
			}
			jobProfile := genmodel.JobPosition{
				TwoSideStyle: genmodel.TwoSideStyle{
					NoAttachmentAndHintsPage: true,
				},
			}
			data := run(t, profile, jobProfile)
			must.HaveLength(t, data.Attachment, 2)
			should.BeEqual(t, data.Attachment[0], "file1")
			should.BeEqual(t, data.Attachment[1], "file2")
			should.BeFalse(t, data.AttachmentAndHintsPage)

		})
	})

	t.Run("check document links", func(t *testing.T) {
		t.Run("experience", func(t *testing.T) {
			t.Run("normal", func(t *testing.T) {
				profile := genmodel.Profile{
					Experience: []genmodel.Experience{
						{
							Company:       "Company",
							DocumentLinks: []string{"document_link1", "document_link2"},
						},
					},
				}
				jobProfile := genmodel.JobPosition{}
				data := run(t, profile, jobProfile)
				must.HaveLength(t, data.SideOneRSA, 1)
				containsRSA(t, data.SideOneRSA[0], "\\documentlink{document_link1}")
				containsRSA(t, data.SideOneRSA[0], "\\documentlink{document_link2}")
			})

			t.Run("style.NoDocumentLinks = true", func(t *testing.T) {
				profile := genmodel.Profile{
					Experience: []genmodel.Experience{
						{
							Company:       "Company",
							DocumentLinks: []string{"document_link1", "document_link2"},
						},
					},
				}
				jobProfile := genmodel.JobPosition{
					TwoSideStyle: genmodel.TwoSideStyle{
						NoDocumentLinks: true,
					},
				}
				data := run(t, profile, jobProfile)
				must.HaveLength(t, data.SideOneRSA, 1)
				notContainsRSA(t, data.SideOneRSA[0], "\\documentlink{document_link1}")
				notContainsRSA(t, data.SideOneRSA[0], "\\documentlink{document_link2}")
			})
		})

		t.Run("education", func(t *testing.T) {
			t.Run("normal", func(t *testing.T) {
				profile := genmodel.Profile{
					Education: []genmodel.Education{
						{
							Institute:     "University",
							DocumentLinks: []string{"document_link1", "document_link2"},
						},
					},
				}
				jobProfile := genmodel.JobPosition{}
				data := run(t, profile, jobProfile)
				must.HaveLength(t, data.SideTwoRSA, 1)
				containsRSA(t, data.SideTwoRSA[0], "\\documentlink{document_link1}")
				containsRSA(t, data.SideTwoRSA[0], "\\documentlink{document_link2}")
			})
			t.Run("style.NoDocumentLinks = true", func(t *testing.T) {
				profile := genmodel.Profile{
					Education: []genmodel.Education{
						{
							Institute:     "University",
							DocumentLinks: []string{"document_link1", "document_link2"},
						},
					},
				}
				jobProfile := genmodel.JobPosition{
					TwoSideStyle: genmodel.TwoSideStyle{
						NoDocumentLinks: true,
					},
				}
				data := run(t, profile, jobProfile)
				must.HaveLength(t, data.SideTwoRSA, 1)
				notContainsRSA(t, data.SideTwoRSA[0], "\\documentlink{document_link1}")
				notContainsRSA(t, data.SideTwoRSA[0], "\\documentlink{document_link2}")
			})

		})

	})

	t.Run("image", func(t *testing.T) {
		t.Run("profile", func(t *testing.T) {
			profile := genmodel.Profile{
				Image: "profilePath",
			}
			jobProfile := genmodel.JobPosition{}
			data := run(t, profile, jobProfile)
			should.BeEqual(t, data.Image, "profilePath")
		})

		t.Run("jobPosition", func(t *testing.T) {
			profile := genmodel.Profile{}
			jobProfile := genmodel.JobPosition{
				Image: "jpPath",
			}
			data := run(t, profile, jobProfile)
			should.BeEqual(t, data.Image, "jpPath")
		})

		t.Run("profile + jobPosition", func(t *testing.T) {
			profile := genmodel.Profile{
				Image: "profilePath",
			}
			jobProfile := genmodel.JobPosition{
				Image: "jpPath",
			}
			data := run(t, profile, jobProfile)
			should.BeEqual(t, data.Image, "jpPath")
		})

		t.Run("none", func(t *testing.T) {
			profile := genmodel.Profile{}
			jobProfile := genmodel.JobPosition{}
			data := run(t, profile, jobProfile)
			should.BeEqual(t, data.Image, image.ImagePath()+"noimage")
		})
	})
}
