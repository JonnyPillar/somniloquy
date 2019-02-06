package fake

// Converter ...
type Converter struct {
	Error error
}

// Execute ...
func (c Converter) Execute(...string) error {
	if c.Error != nil {
		return c.Error
	}

	return nil
}
