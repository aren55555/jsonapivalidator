package jsonapivalidator

// Result is the container for wrapping validation results
type Result struct {
	errors map[error]error
}

// NewResult instantiates a Result struct
func NewResult() (r *Result) {
	r = &Result{
		errors: make(map[error]error),
	}
	return
}

// AddError will append the error to the Result
func (r *Result) AddError(err error) {
	r.errors[err] = err
}

// HasError checks whether the Result has the particular Error
func (r *Result) HasError(err error) bool {
	_, exists := r.errors[err]
	return exists
}

// HasErrors is whether or not the Result has Errors in general
func (r *Result) HasErrors() bool {
	return len(r.errors) > 0
}

// Errors will return all the errors in the result as a slice
func (r *Result) Errors() (errors []error) {
	for err := range r.errors {
		errors = append(errors, err)
	}
	return
}
