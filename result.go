package jsonapivalidator

// Result is the container for wrapping validation results
type Result struct {
	Errors map[error]error
}

// NewResult instantiates a Result struct
func NewResult() (r *Result) {
	r = &Result{
		Errors: make(map[error]error),
	}
	return
}

// AddError will append the error to the Result
func (r *Result) AddError(err error) {
	r.Errors[err] = err
}

// HasError checks whether the Result has the particular Error
func (r *Result) HasError(err error) bool {
	_, exists := r.Errors[err]
	return exists
}

// HasErrors is whether or not the Result has Errors in general
func (r *Result) HasErrors() bool {
	return len(r.Errors) > 0
}
