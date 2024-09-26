package forms

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/noeeekr/sch-server/pkg/models"
)

type Form struct {
	Fields     *models.Users
	Errors     map[string]string
	emailRegex *regexp.Regexp
}

type FormField struct {
	form *Form    // The Parent Form.
	CWFS []string // Current Working Fields.
}

// Returns a new form for request data validation.
func New(data *models.Users) *Form {
	return &Form{
		Fields:     data,
		Errors:     make(map[string]string),
		emailRegex: regexp.MustCompile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`),
	}
}

// Set a field to work on.
func (f *Form) SetField(_field string) *FormField {
	field := FormField{
		form: f,
		CWFS: []string{_field},
	}

	return &field
}

// Define multiple fields to work on.
func (f *Form) SetFields(fields ...string) *FormField {
	field := FormField{
		f,
		fields,
	}

	return &field
}

// Define required fields.
func (f *Form) Required(fields ...string) {
	// Check if f.Fields is a pointer and dereference it
	var val reflect.Value = reflect.ValueOf(f.Fields).Elem()

	for _, field := range fields {

		field := val.FieldByName(field)

		if !field.IsValid() {
			f.Errors["Required"] += fmt.Sprintf("%s não pode estar vazio. ", field)
		}

		if strings.TrimSpace(field.String()) == "" {
			f.Errors["Required"] += fmt.Sprintf("%s não pode estar vazio. ", field)
		}
	}
}

func (f *Form) IsValid() error {

	for errs := range f.Errors {
		return errors.New(errs)
	}
	return nil

}

// Checks if the string is in the required length. Set max to 0 if no max length is necessary.
func (f *FormField) Length(min int, max int) *FormField {

	// Get form values
	form_vals := reflect.ValueOf(f.form.Fields).Elem()

	for i, _field := range f.CWFS {

		// Get field
		field := form_vals.FieldByName(_field)

		length := utf8.RuneCount([]byte(field.String()))

		if length > max && max > 0 {
			f.form.Errors[_field] += fmt.Sprintf("%s é maior que o máximo permitido (%d). ", f.CWFS[i], max)
		} else if length < min {
			f.form.Errors[_field] += fmt.Sprintf("%s é menor que o mínimo permitido (%d). ", f.CWFS[i], min)
		}
	}

	return f
}

// Check if a field in form is a valid email.
func (f *FormField) IsValidEmail() *FormField {
	fieldName := f.CWFS[0]

	val := reflect.ValueOf(f.form.Fields).Elem()
	field := val.FieldByName(fieldName)

	if !field.IsValid() {
		f.form.Errors[fieldName] = "O email recebido não é válido."
		return f
	}

	var isValid bool = f.form.emailRegex.MatchString(field.String())

	if !isValid {
		f.form.Errors[fieldName] = "O email recebido não é válido."
	}

	return f
}

// Define the field as required.
func (f *FormField) Required() *FormField {
	fieldName := f.CWFS[0]

	if strings.TrimSpace(fieldName) == "" {
		f.form.Errors[fieldName] += fmt.Sprintf("%s não pode estar vazio. ", fieldName)
	}

	return f
}
