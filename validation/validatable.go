package validation

type Validatable interface {
	Validate() Errors
}
