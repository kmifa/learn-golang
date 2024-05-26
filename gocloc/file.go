package gocloc

// ClocFile is collecting to line count result.
type ClocFile struct {
	Code     int32
	Comments int32
	Blanks   int32
	Name     string
	Lang     string
}
