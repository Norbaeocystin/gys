package gysyaml

type Gys struct {
	Name string `yaml:"name"`
	Description       string `yaml:"description"`
	Version       string `yaml:"version"`
	Iterator   struct {
		Url     string `yaml:"url"`
		Replace string `yaml:"replace"`
		Min     int    `yaml:"min"`
		Max     int `yaml:"max"`
	}
	Identificator struct {
		Attribute string `yaml:"attribute"`
		Selector string`yaml:"selector"`
		Name string `yaml:"url"`
		Type string `yaml:"type"`
		Default string `yaml:"default"`
		Base string `yaml:"base"`
		}
	Output struct {
		Format string `yaml:"format"`
		Delimiter string`yaml:"delimiter"`
	}
}
