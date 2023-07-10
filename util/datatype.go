package util

type FilesConfig struct {
	Include []string `yaml:"includePath"`
	Exclude []string `yaml:"excludepath"`
}

type Files struct {
	List []string
}
