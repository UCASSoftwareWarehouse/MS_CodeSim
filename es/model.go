package es

type Document interface {
	getID() string
}

type CodePlainText struct {
	CodePlainText  string `json:"code-plain-text"`
	CodeUniquePath string `json:"code-unique-path"`
	Tag            string `json:"tag"`
	ID             string `json:"id"`
}

func NewCodePlainText(plainText, uniquePath, tag string) *CodePlainText {
	return &CodePlainText{
		CodePlainText:  plainText,
		CodeUniquePath: uniquePath,
		Tag:            tag,
		ID:             uniquePath + ":" + tag,
	}
}

func (c *CodePlainText) getID() string {
	return c.ID
}

type CodeTransformedText struct {
	CodeTransformedText  string `json:"code-transformed-text"`
	CodeUniquePath string `json:"code-unique-path"`
	Tag            string `json:"tag"`
	ID             string `json:"id"`
}

func NewCodeTransformedText(transformed, uniquePath, tag string) *CodeTransformedText {
	return &CodeTransformedText{
		CodeTransformedText:  transformed,
		CodeUniquePath: uniquePath,
		Tag:            tag,
		ID:             uniquePath + ":" + tag,
	}
}