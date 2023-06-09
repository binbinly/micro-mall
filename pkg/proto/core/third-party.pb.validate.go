// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: core/third-party.proto

package core

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

// Validate checks the field values on SendSMSReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendSMSReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendSMSReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SendSMSReqMultiError, or
// nil if none found.
func (m *SendSMSReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SendSMSReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetPhone()) != 11 {
		err := SendSMSReqValidationError{
			field:  "Phone",
			reason: "value length must be 11 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return SendSMSReqMultiError(errors)
	}

	return nil
}

// SendSMSReqMultiError is an error wrapping multiple validation errors
// returned by SendSMSReq.ValidateAll() if the designated constraints aren't met.
type SendSMSReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendSMSReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendSMSReqMultiError) AllErrors() []error { return m }

// SendSMSReqValidationError is the validation error returned by
// SendSMSReq.Validate if the designated constraints aren't met.
type SendSMSReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendSMSReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendSMSReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendSMSReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendSMSReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendSMSReqValidationError) ErrorName() string { return "SendSMSReqValidationError" }

// Error satisfies the builtin error interface
func (e SendSMSReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendSMSReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendSMSReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendSMSReqValidationError{}

// Validate checks the field values on VCodeReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VCodeReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VCodeReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VCodeReqMultiError, or nil
// if none found.
func (m *VCodeReq) ValidateAll() error {
	return m.validate(true)
}

func (m *VCodeReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPhone() <= 0 {
		err := VCodeReqValidationError{
			field:  "Phone",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCode()) != 6 {
		err := VCodeReqValidationError{
			field:  "Code",
			reason: "value length must be 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return VCodeReqMultiError(errors)
	}

	return nil
}

// VCodeReqMultiError is an error wrapping multiple validation errors returned
// by VCodeReq.ValidateAll() if the designated constraints aren't met.
type VCodeReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VCodeReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VCodeReqMultiError) AllErrors() []error { return m }

// VCodeReqValidationError is the validation error returned by
// VCodeReq.Validate if the designated constraints aren't met.
type VCodeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VCodeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VCodeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VCodeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VCodeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VCodeReqValidationError) ErrorName() string { return "VCodeReqValidationError" }

// Error satisfies the builtin error interface
func (e VCodeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVCodeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VCodeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VCodeReqValidationError{}

// Validate checks the field values on ETHPayReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ETHPayReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ETHPayReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ETHPayReqMultiError, or nil
// if none found.
func (m *ETHPayReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ETHPayReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := ETHPayReqValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAddress()) < 3 {
		err := ETHPayReqValidationError{
			field:  "Address",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetOrderNo()) < 3 {
		err := ETHPayReqValidationError{
			field:  "OrderNo",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ETHPayReqMultiError(errors)
	}

	return nil
}

// ETHPayReqMultiError is an error wrapping multiple validation errors returned
// by ETHPayReq.ValidateAll() if the designated constraints aren't met.
type ETHPayReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ETHPayReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ETHPayReqMultiError) AllErrors() []error { return m }

// ETHPayReqValidationError is the validation error returned by
// ETHPayReq.Validate if the designated constraints aren't met.
type ETHPayReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ETHPayReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ETHPayReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ETHPayReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ETHPayReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ETHPayReqValidationError) ErrorName() string { return "ETHPayReqValidationError" }

// Error satisfies the builtin error interface
func (e ETHPayReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sETHPayReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ETHPayReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ETHPayReqValidationError{}

// Validate checks the field values on CodeReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CodeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CodeReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CodeReplyMultiError, or nil
// if none found.
func (m *CodeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CodeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	if len(errors) > 0 {
		return CodeReplyMultiError(errors)
	}

	return nil
}

// CodeReplyMultiError is an error wrapping multiple validation errors returned
// by CodeReply.ValidateAll() if the designated constraints aren't met.
type CodeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CodeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CodeReplyMultiError) AllErrors() []error { return m }

// CodeReplyValidationError is the validation error returned by
// CodeReply.Validate if the designated constraints aren't met.
type CodeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CodeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CodeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CodeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CodeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CodeReplyValidationError) ErrorName() string { return "CodeReplyValidationError" }

// Error satisfies the builtin error interface
func (e CodeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCodeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CodeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CodeReplyValidationError{}
