// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authorization.proto

package v1

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

// Validate checks the field values on AuthorizeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthorizeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthorizeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthorizeRequestMultiError, or nil if none found.
func (m *AuthorizeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthorizeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for Method

	if len(errors) > 0 {
		return AuthorizeRequestMultiError(errors)
	}

	return nil
}

// AuthorizeRequestMultiError is an error wrapping multiple validation errors
// returned by AuthorizeRequest.ValidateAll() if the designated constraints
// aren't met.
type AuthorizeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthorizeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthorizeRequestMultiError) AllErrors() []error { return m }

// AuthorizeRequestValidationError is the validation error returned by
// AuthorizeRequest.Validate if the designated constraints aren't met.
type AuthorizeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorizeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorizeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorizeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorizeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorizeRequestValidationError) ErrorName() string { return "AuthorizeRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthorizeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthorizeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorizeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorizeRequestValidationError{}

// Validate checks the field values on AuthorizeResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthorizeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthorizeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthorizeResponseMultiError, or nil if none found.
func (m *AuthorizeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthorizeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AuthorizeResponseMultiError(errors)
	}

	return nil
}

// AuthorizeResponseMultiError is an error wrapping multiple validation errors
// returned by AuthorizeResponse.ValidateAll() if the designated constraints
// aren't met.
type AuthorizeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthorizeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthorizeResponseMultiError) AllErrors() []error { return m }

// AuthorizeResponseValidationError is the validation error returned by
// AuthorizeResponse.Validate if the designated constraints aren't met.
type AuthorizeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorizeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorizeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorizeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorizeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorizeResponseValidationError) ErrorName() string {
	return "AuthorizeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthorizeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthorizeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorizeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorizeResponseValidationError{}