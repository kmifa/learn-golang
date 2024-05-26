package gocloc

// ClocOptions is gocloc processor options.
type ClocOptions struct {
	Debug bool
}

// NewClocOptions create new ClocOptions with default values.
func NewClocOptions() *ClocOptions {
	return &ClocOptions{
		Debug: false,
	}
}
