// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api_comment.proto

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

// Validate checks the field values on CommonCommentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommonCommentResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonCommentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonCommentResponseMultiError, or nil if none found.
func (m *CommonCommentResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonCommentResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetStatusCode() < 1 {
		err := CommonCommentResponseValidationError{
			field:  "StatusCode",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetMessage()) < 2 {
		err := CommonCommentResponseValidationError{
			field:  "Message",
			reason: "value length must be at least 2 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CommonCommentResponseMultiError(errors)
	}

	return nil
}

// CommonCommentResponseMultiError is an error wrapping multiple validation
// errors returned by CommonCommentResponse.ValidateAll() if the designated
// constraints aren't met.
type CommonCommentResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonCommentResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonCommentResponseMultiError) AllErrors() []error { return m }

// CommonCommentResponseValidationError is the validation error returned by
// CommonCommentResponse.Validate if the designated constraints aren't met.
type CommonCommentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonCommentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonCommentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonCommentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonCommentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonCommentResponseValidationError) ErrorName() string {
	return "CommonCommentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CommonCommentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonCommentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonCommentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonCommentResponseValidationError{}

// Validate checks the field values on CommentInput with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommentInput) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommentInput with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommentInputMultiError, or
// nil if none found.
func (m *CommentInput) ValidateAll() error {
	return m.validate(true)
}

func (m *CommentInput) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserID

	// no validation rules for Content

	if m.PostID != nil {
		// no validation rules for PostID
	}

	if m.ExerciseID != nil {
		// no validation rules for ExerciseID
	}

	if len(errors) > 0 {
		return CommentInputMultiError(errors)
	}

	return nil
}

// CommentInputMultiError is an error wrapping multiple validation errors
// returned by CommentInput.ValidateAll() if the designated constraints aren't met.
type CommentInputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommentInputMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommentInputMultiError) AllErrors() []error { return m }

// CommentInputValidationError is the validation error returned by
// CommentInput.Validate if the designated constraints aren't met.
type CommentInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentInputValidationError) ErrorName() string { return "CommentInputValidationError" }

// Error satisfies the builtin error interface
func (e CommentInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentInputValidationError{}

// Validate checks the field values on CreateCommentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCommentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCommentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCommentRequestMultiError, or nil if none found.
func (m *CreateCommentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCommentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetComment() == nil {
		err := CreateCommentRequestValidationError{
			field:  "Comment",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetComment()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCommentRequestValidationError{
					field:  "Comment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCommentRequestValidationError{
					field:  "Comment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetComment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCommentRequestValidationError{
				field:  "Comment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateCommentRequestMultiError(errors)
	}

	return nil
}

// CreateCommentRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCommentRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCommentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCommentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCommentRequestMultiError) AllErrors() []error { return m }

// CreateCommentRequestValidationError is the validation error returned by
// CreateCommentRequest.Validate if the designated constraints aren't met.
type CreateCommentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCommentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCommentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCommentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCommentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCommentRequestValidationError) ErrorName() string {
	return "CreateCommentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCommentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCommentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCommentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCommentRequestValidationError{}

// Validate checks the field values on CreateCommentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCommentResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCommentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCommentResponseMultiError, or nil if none found.
func (m *CreateCommentResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCommentResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetResponse() == nil {
		err := CreateCommentResponseValidationError{
			field:  "Response",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCommentResponseValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCommentResponseValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCommentResponseValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateCommentResponseMultiError(errors)
	}

	return nil
}

// CreateCommentResponseMultiError is an error wrapping multiple validation
// errors returned by CreateCommentResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateCommentResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCommentResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCommentResponseMultiError) AllErrors() []error { return m }

// CreateCommentResponseValidationError is the validation error returned by
// CreateCommentResponse.Validate if the designated constraints aren't met.
type CreateCommentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCommentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCommentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCommentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCommentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCommentResponseValidationError) ErrorName() string {
	return "CreateCommentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCommentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCommentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCommentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCommentResponseValidationError{}