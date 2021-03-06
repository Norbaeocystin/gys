package pkg

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
	Extractor struct {
		Filewithurls string `yaml:"filewithurls"`
		Urls string `yaml:"Urls"`
		Selector string `yaml:"selector"`
		Type string  `yaml:"type"`
		Subselectors []struct{
			Attribute string `yaml:"attribute"`
			Selector string`yaml:"selector"`
			Name string `yaml:"name"`
			Split string `yaml:"split"`
			Default string `yaml:"default"`
		}
	}
	Output struct {
		Filename string `yaml:"filename"`
	}
}

type GysServer struct {
	Url string `json:"Url"`
	Selector string `json:"selector"`
	Type string  `json:"type"`
	Subselectors []struct{
		Attribute string `json:"attribute"`
		Selector string`json:"selector"`
		Name string `json:"name"`
		Split string `json:"split"`
		Default string `json:"default"`
	}
}