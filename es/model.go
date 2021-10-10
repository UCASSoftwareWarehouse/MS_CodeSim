package es

type Document interface {
	getID() string
}

type CodePlainText struct {
	CodePlainText string `json:"code-plain-text"`
	*ProjectFileIdentifier
}

type ProjectFileIdentifier struct {
	CodeUniquePath string `json:"code-unique-path"`
	Tag            string `json:"tag"`
	ID             string `json:"id"`
}

func NewCodePlainText(plainText string, esInfo *ProjectFileIdentifier) *CodePlainText {
	return &CodePlainText{
		CodePlainText: plainText,
		ProjectFileIdentifier: &ProjectFileIdentifier{
			CodeUniquePath: esInfo.CodeUniquePath,
			Tag:            esInfo.Tag,
			ID:             esInfo.ID,
		},
	}
}

func (c *CodePlainText) getID() string {
	return c.ID
}

type CodeTransformedText struct {
	CodeTransformedText string `json:"code-transformed-text"`
	*ProjectFileIdentifier
}

func (c *CodeTransformedText) getID() string {
	return c.ID
}

func NewCodeTransformedText(transformed string, esInfo *ProjectFileIdentifier) *CodeTransformedText {
	return &CodeTransformedText{
		CodeTransformedText: transformed,
		ProjectFileIdentifier: &ProjectFileIdentifier{
			CodeUniquePath: esInfo.CodeUniquePath,
			Tag:            esInfo.Tag,
			ID:             esInfo.ID,
		},
	}
}
