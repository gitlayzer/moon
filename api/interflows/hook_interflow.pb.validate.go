// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: interflows/hook_interflow.proto

package interflows

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ReceiveRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReceiveRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReceiveRequestMultiError,
// or nil if none found.
func (m *ReceiveRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTopic()) < 1 {
		err := ReceiveRequestValidationError{
			field:  "Topic",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Value

	if len(errors) > 0 {
		return ReceiveRequestMultiError(errors)
	}

	return nil
}

// ReceiveRequestMultiError is an error wrapping multiple validation errors
// returned by ReceiveRequest.ValidateAll() if the designated constraints
// aren't met.
type ReceiveRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveRequestMultiError) AllErrors() []error { return m }

// ReceiveRequestValidationError is the validation error returned by
// ReceiveRequest.Validate if the designated constraints aren't met.
type ReceiveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveRequestValidationError) ErrorName() string { return "ReceiveRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReceiveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveRequestValidationError{}

// Validate checks the field values on ReceiveResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReceiveResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReceiveResponseMultiError, or nil if none found.
func (m *ReceiveResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if len(errors) > 0 {
		return ReceiveResponseMultiError(errors)
	}

	return nil
}

// ReceiveResponseMultiError is an error wrapping multiple validation errors
// returned by ReceiveResponse.ValidateAll() if the designated constraints
// aren't met.
type ReceiveResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveResponseMultiError) AllErrors() []error { return m }

// ReceiveResponseValidationError is the validation error returned by
// ReceiveResponse.Validate if the designated constraints aren't met.
type ReceiveResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveResponseValidationError) ErrorName() string { return "ReceiveResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReceiveResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveResponseValidationError{}
