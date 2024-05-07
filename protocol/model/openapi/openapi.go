package openapi

type (
	Info struct {
		Version     string
		Description string
		Title       string
	}

	Contact struct {
		Name  string
		Email string
		URL   string
	}

	Server struct {
		URL         string
		Description string
	}

	Tag struct {
		Name         string
		Description  string
		ExternalDocs *ExternalDocs
	}

	// The name is plural by open api definition
	ExternalDocs struct {
		Description string
		URL         string
	}

	NameGetter interface {
		GetName() string
	}

	DescriptionGetter interface {
		GetDescription() string
	}
)
