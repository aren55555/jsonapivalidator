package jsonapivalidator

// Result is the container for wrapping validation results
type Result struct {
	// errors will be populated with violations of MUST, MUST NOT, REQUIRED,
	// SHALL and SHALL NOT
	errors map[error]*interface{}
	// warnings will be populated with violations of SHOULD, SHOULD NOT,
	// RECOMMENDED, MAY and OPTIONAL
	warnings map[error]*interface{}
}

// NewResult instantiates a Result struct
func NewResult() (r *Result) {
	r = &Result{
		errors:   make(map[error]*interface{}),
		warnings: make(map[error]*interface{}),
	}
	return
}

// IsValid summarizes if the validation result was valid or not
func (r *Result) IsValid() bool {
	if !r.HasErrors() && !r.HasWarnings() {
		return true
	}
	return false
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
func (r *Result) AddWarning(warning error) {
	r.warnings[warning] = nil
}

// HasWarning checks whether the Result has the particular Error
func (r *Result) HasWarning(warning error) bool {
	_, exists := r.warnings[warning]
	return exists
}

// HasWarnings is whether or not the Result has Errors in general
func (r *Result) HasWarnings() bool {
	return len(r.warnings) > 0
}

// Warnings will return all the warnings in teh result as a slice
func (r *Result) Warnings() (warnings []error) {
	for warning := range r.warnings {
		warnings = append(warnings, warning)
	}
	return
}
