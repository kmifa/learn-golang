package gocloc

// Language is a type used to definitions and store statistics for one programming language.
type Language struct {
	Name         string
	lineComments []string
	multiLines   [][]string
	Files        []string
	Code         int32
	Comments     int32
	Blanks       int32
	Total        int32
}

// NewLanguage create Language data store.
func NewLanguage(name string, lineComments []string, multiLines [][]string) *Language {
	return &Language{
		Name:         name,
		lineComments: lineComments,
		multiLines:   multiLines,
		Files:        []string{},
	}
}

// DefinedLanguages is the type information for mapping language name(key) and NewLanguage.
type DefinedLanguages struct {
	Langs map[string]*Language
}

// NewDefinedLanguages create DefinedLanguages.
func NewDefinedLanguages() *DefinedLanguages {
	return &DefinedLanguages{
		Langs: map[string]*Language{
			"Go": NewLanguage("Go", []string{"//"}, [][]string{{"/*", "*/"}}),
		},
	}
}
