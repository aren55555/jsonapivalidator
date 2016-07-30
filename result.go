package jsonapivalidator

// Result is the container for wrapping validation results
type Result struct {
	// errors will be populated with violations of MUST, MUST NOT, REQUIRED,
	// SHALL and SHALL NOT
	errors map[error]*empty
	// warnings will be populated with violations of SHOULD, SHOULD NOT,
	// RECOMMENDED, MAY and OPTIONAL
	warnings map[string]*empty
}
type empty struct{}

// NewResult instantiates a Result struct
func NewResult() (r *Result) {
	r = &Result{
		errors: make(map[error]*empty),
	}
	return
}

// AddError will append the error to the Result
func (r *Result) AddError(err error) {
	r.errors[err] = nil
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

// AddWarning will append the error to the Result
func (r *Result) AddWarning(warning string) {
	r.warnings[warning] = nil
}

// HasWarning checks whether the Result has the particular Error
func (r *Result) HasWarning(warning string) bool {
	_, exists := r.warnings[warning]
	return exists
}

// HasWarnings is whether or not the Result has Errors in general
func (r *Result) HasWarnings() bool {
	return len(r.warnings) > 0
}

// Warnings will return all the warnings in teh result as a slice
func (r *Result) Warnings() (warnings []string) {
	for warning := range r.warnings {
		warnings = append(warnings, warning)
	}
	return
}
