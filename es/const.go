package es

type CodeIndexField string

type IndexName string

const (
	CodeIndexFieldPlain  CodeIndexField = "code-plain-text"
	CodeIndexFieldGolang CodeIndexField = "code-plain-text.golang"
	CodeIndexFieldJava   CodeIndexField = "code-plain-text.java"
	CodeIndexFieldPython   CodeIndexField = "code-plain-text.python"

	CodeIndex IndexName = "code-index"
)
