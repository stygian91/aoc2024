package inputs

import "embed"

//go:embed files/**
var files embed.FS

func GetInputFile(path string) (string, error) {
	contents, err := files.ReadFile("files/" + path)
	return string(contents), err
}
