package gocloc

// Processor is gocloc analyzing processor.
type Processor struct {
	langs *DefinedLanguages
	opts  *ClocOptions
}

// Result defined processing result.
type Result struct {
	Total         *Language
	Files         map[string]*ClocFile
	Language      map[string]*Language
	MaxPathLength int
}

// NewProcessor returns Processor.
func NewProcessor(langs *DefinedLanguages, options *ClocOptions) *Processor {
	return &Processor{
		langs: langs,
		opts:  options,
	}

}
