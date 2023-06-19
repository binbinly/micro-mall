// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: seckill/seckill.proto

package gateway

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

// Validate checks the field values on KillReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *KillReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KillReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in KillReqMultiError, or nil if none found.
func (m *KillReq) ValidateAll() error {
	return m.validate(true)
}

func (m *KillReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSkuId() <= 0 {
		err := KillReqValidationError{
			field:  "SkuId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAddressId() <= 0 {
		err := KillReqValidationError{
			field:  "AddressId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetNum() <= 0 {
		err := KillReqValidationError{
			field:  "Num",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetKey()) != 32 {
		err := KillReqValidationError{
			field:  "Key",
			reason: "value length must be 32 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return KillReqMultiError(errors)
	}

	return nil
}

// KillReqMultiError is an error wrapping multiple validation errors returned
// by KillReq.ValidateAll() if the designated constraints aren't met.
type KillReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KillReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KillReqMultiError) AllErrors() []error { return m }

// KillReqValidationError is the validation error returned by KillReq.Validate
// if the designated constraints aren't met.
type KillReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KillReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KillReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KillReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KillReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KillReqValidationError) ErrorName() string { return "KillReqValidationError" }

// Error satisfies the builtin error interface
func (e KillReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKillReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KillReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KillReqValidationError{}

// Validate checks the field values on SessionIdReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SessionIdReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SessionIdReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SessionIdReqMultiError, or
// nil if none found.
func (m *SessionIdReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SessionIdReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSessionId() <= 0 {
		err := SessionIdReqValidationError{
			field:  "SessionId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SessionIdReqMultiError(errors)
	}

	return nil
}

// SessionIdReqMultiError is an error wrapping multiple validation errors
// returned by SessionIdReq.ValidateAll() if the designated constraints aren't met.
type SessionIdReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SessionIdReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SessionIdReqMultiError) AllErrors() []error { return m }

// SessionIdReqValidationError is the validation error returned by
// SessionIdReq.Validate if the designated constraints aren't met.
type SessionIdReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SessionIdReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SessionIdReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SessionIdReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SessionIdReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SessionIdReqValidationError) ErrorName() string { return "SessionIdReqValidationError" }

// Error satisfies the builtin error interface
func (e SessionIdReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSessionIdReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SessionIdReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SessionIdReqValidationError{}

// Validate checks the field values on KillReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *KillReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KillReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in KillReplyMultiError, or nil
// if none found.
func (m *KillReply) ValidateAll() error {
	return m.validate(true)
}

func (m *KillReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	// no validation rules for Data

	if len(errors) > 0 {
		return KillReplyMultiError(errors)
	}

	return nil
}

// KillReplyMultiError is an error wrapping multiple validation errors returned
// by KillReply.ValidateAll() if the designated constraints aren't met.
type KillReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KillReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KillReplyMultiError) AllErrors() []error { return m }

// KillReplyValidationError is the validation error returned by
// KillReply.Validate if the designated constraints aren't met.
type KillReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KillReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KillReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KillReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KillReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KillReplyValidationError) ErrorName() string { return "KillReplyValidationError" }

// Error satisfies the builtin error interface
func (e KillReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKillReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KillReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KillReplyValidationError{}

// Validate checks the field values on SessionsReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SessionsReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SessionsReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SessionsReplyMultiError, or
// nil if none found.
func (m *SessionsReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SessionsReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SessionsReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SessionsReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SessionsReplyValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SessionsReplyMultiError(errors)
	}

	return nil
}

// SessionsReplyMultiError is an error wrapping multiple validation errors
// returned by SessionsReply.ValidateAll() if the designated constraints
// aren't met.
type SessionsReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SessionsReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SessionsReplyMultiError) AllErrors() []error { return m }

// SessionsReplyValidationError is the validation error returned by
// SessionsReply.Validate if the designated constraints aren't met.
type SessionsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SessionsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SessionsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SessionsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SessionsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SessionsReplyValidationError) ErrorName() string { return "SessionsReplyValidationError" }

// Error satisfies the builtin error interface
func (e SessionsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSessionsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SessionsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SessionsReplyValidationError{}

// Validate checks the field values on SkusReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SkusReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SkusReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SkusReplyMultiError, or nil
// if none found.
func (m *SkusReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SkusReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SkusReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SkusReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SkusReplyValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SkusReplyMultiError(errors)
	}

	return nil
}

// SkusReplyMultiError is an error wrapping multiple validation errors returned
// by SkusReply.ValidateAll() if the designated constraints aren't met.
type SkusReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SkusReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SkusReplyMultiError) AllErrors() []error { return m }

// SkusReplyValidationError is the validation error returned by
// SkusReply.Validate if the designated constraints aren't met.
type SkusReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SkusReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SkusReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SkusReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SkusReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SkusReplyValidationError) ErrorName() string { return "SkusReplyValidationError" }

// Error satisfies the builtin error interface
func (e SkusReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSkusReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SkusReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SkusReplyValidationError{}

// Validate checks the field values on SkuReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SkuReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SkuReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SkuReplyMultiError, or nil
// if none found.
func (m *SkuReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SkuReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SkuReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SkuReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SkuReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SkuReplyMultiError(errors)
	}

	return nil
}

// SkuReplyMultiError is an error wrapping multiple validation errors returned
// by SkuReply.ValidateAll() if the designated constraints aren't met.
type SkuReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SkuReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SkuReplyMultiError) AllErrors() []error { return m }

// SkuReplyValidationError is the validation error returned by
// SkuReply.Validate if the designated constraints aren't met.
type SkuReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SkuReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SkuReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SkuReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SkuReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SkuReplyValidationError) ErrorName() string { return "SkuReplyValidationError" }

// Error satisfies the builtin error interface
func (e SkuReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSkuReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SkuReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SkuReplyValidationError{}