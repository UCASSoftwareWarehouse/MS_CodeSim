package es

type Document interface {
	getID() string
}

type CodePlainText struct {
	CodePlainText string `json:"code-plain-text"`
	*ProjectFileIdentifier
}

type ProjectIdentifier struct {
	ProjectName string `json:"project_name"`
	Tag         string `json:"tag"`
}

type ProjectFileIdentifier struct {
	CodeUniquePath string `json:"code-unique-path"`
	ProjectIdentifier
	ID string `json:"id"`
}

func NewCodePlainText(plainText string, esInfo *ProjectFileIdentifier) *CodePlainText {
	return &CodePlainText{
		CodePlainText: plainText,
		ProjectFileIdentifier: &ProjectFileIdentifier{
			CodeUniquePath: esInfo.CodeUniquePath,
			ProjectIdentifier: ProjectIdentifier{
				ProjectName: esInfo.ProjectName,
				Tag:         esInfo.Tag,
			},
			ID: esInfo.ID,
		},
	}
}

func (c *CodePlainText) getID() string {
	return c.ID
}
