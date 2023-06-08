// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: order/order.proto

package order

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

// Validate checks the field values on SubmitReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SubmitReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubmitReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubmitReqMultiError, or nil
// if none found.
func (m *SubmitReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SubmitReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetSkuIds()) < 1 {
		err := SubmitReqValidationError{
			field:  "SkuIds",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAddressId() <= 0 {
		err := SubmitReqValidationError{
			field:  "AddressId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CouponId

	// no validation rules for Note

	if len(errors) > 0 {
		return SubmitReqMultiError(errors)
	}

	return nil
}

// SubmitReqMultiError is an error wrapping multiple validation errors returned
// by SubmitReq.ValidateAll() if the designated constraints aren't met.
type SubmitReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubmitReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubmitReqMultiError) AllErrors() []error { return m }

// SubmitReqValidationError is the validation error returned by
// SubmitReq.Validate if the designated constraints aren't met.
type SubmitReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitReqValidationError) ErrorName() string { return "SubmitReqValidationError" }

// Error satisfies the builtin error interface
func (e SubmitReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitReqValidationError{}

// Validate checks the field values on SkuSubmitReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SkuSubmitReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SkuSubmitReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SkuSubmitReqMultiError, or
// nil if none found.
func (m *SkuSubmitReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SkuSubmitReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSkuId() <= 0 {
		err := SkuSubmitReqValidationError{
			field:  "SkuId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAddressId() <= 0 {
		err := SkuSubmitReqValidationError{
			field:  "AddressId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CouponId

	// no validation rules for Note

	if m.GetNum() <= 0 {
		err := SkuSubmitReqValidationError{
			field:  "Num",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SkuSubmitReqMultiError(errors)
	}

	return nil
}

// SkuSubmitReqMultiError is an error wrapping multiple validation errors
// returned by SkuSubmitReq.ValidateAll() if the designated constraints aren't met.
type SkuSubmitReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SkuSubmitReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SkuSubmitReqMultiError) AllErrors() []error { return m }

// SkuSubmitReqValidationError is the validation error returned by
// SkuSubmitReq.Validate if the designated constraints aren't met.
type SkuSubmitReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SkuSubmitReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SkuSubmitReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SkuSubmitReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SkuSubmitReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SkuSubmitReqValidationError) ErrorName() string { return "SkuSubmitReqValidationError" }

// Error satisfies the builtin error interface
func (e SkuSubmitReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSkuSubmitReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SkuSubmitReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SkuSubmitReqValidationError{}

// Validate checks the field values on ListReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ListReqMultiError, or nil if none found.
func (m *ListReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Page

	if len(errors) > 0 {
		return ListReqMultiError(errors)
	}

	return nil
}

// ListReqMultiError is an error wrapping multiple validation errors returned
// by ListReq.ValidateAll() if the designated constraints aren't met.
type ListReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReqMultiError) AllErrors() []error { return m }

// ListReqValidationError is the validation error returned by ListReq.Validate
// if the designated constraints aren't met.
type ListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReqValidationError) ErrorName() string { return "ListReqValidationError" }

// Error satisfies the builtin error interface
func (e ListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReqValidationError{}

// Validate checks the field values on PayNotifyReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PayNotifyReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayNotifyReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PayNotifyReqMultiError, or
// nil if none found.
func (m *PayNotifyReq) ValidateAll() error {
	return m.validate(true)
}

func (m *PayNotifyReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPayAmount() <= 0 {
		err := PayNotifyReqValidationError{
			field:  "PayAmount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPayType() <= 0 {
		err := PayNotifyReqValidationError{
			field:  "PayType",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetOrderNo()); l < 16 || l > 30 {
		err := PayNotifyReqValidationError{
			field:  "OrderNo",
			reason: "value length must be between 16 and 30 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetTradeNo()); l < 16 || l > 64 {
		err := PayNotifyReqValidationError{
			field:  "TradeNo",
			reason: "value length must be between 16 and 64 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for TransHash

	if len(errors) > 0 {
		return PayNotifyReqMultiError(errors)
	}

	return nil
}

// PayNotifyReqMultiError is an error wrapping multiple validation errors
// returned by PayNotifyReq.ValidateAll() if the designated constraints aren't met.
type PayNotifyReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayNotifyReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayNotifyReqMultiError) AllErrors() []error { return m }

// PayNotifyReqValidationError is the validation error returned by
// PayNotifyReq.Validate if the designated constraints aren't met.
type PayNotifyReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayNotifyReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayNotifyReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayNotifyReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayNotifyReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayNotifyReqValidationError) ErrorName() string { return "PayNotifyReqValidationError" }

// Error satisfies the builtin error interface
func (e PayNotifyReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayNotifyReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayNotifyReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayNotifyReqValidationError{}

// Validate checks the field values on RefundReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RefundReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefundReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RefundReqMultiError, or nil
// if none found.
func (m *RefundReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RefundReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetOrderId() <= 0 {
		err := RefundReqValidationError{
			field:  "OrderId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 5 || l > 255 {
		err := RefundReqValidationError{
			field:  "Content",
			reason: "value length must be between 5 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RefundReqMultiError(errors)
	}

	return nil
}

// RefundReqMultiError is an error wrapping multiple validation errors returned
// by RefundReq.ValidateAll() if the designated constraints aren't met.
type RefundReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefundReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefundReqMultiError) AllErrors() []error { return m }

// RefundReqValidationError is the validation error returned by
// RefundReq.Validate if the designated constraints aren't met.
type RefundReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefundReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefundReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefundReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefundReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefundReqValidationError) ErrorName() string { return "RefundReqValidationError" }

// Error satisfies the builtin error interface
func (e RefundReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefundReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefundReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefundReqValidationError{}

// Validate checks the field values on CommentReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommentReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommentReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommentReqMultiError, or
// nil if none found.
func (m *CommentReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CommentReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetStar() <= 0 {
		err := CommentReqValidationError{
			field:  "Star",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetOrderId() <= 0 {
		err := CommentReqValidationError{
			field:  "OrderId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 5 || l > 255 {
		err := CommentReqValidationError{
			field:  "Content",
			reason: "value length must be between 5 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Resources

	if len(m.GetSkuIds()) < 1 {
		err := CommentReqValidationError{
			field:  "SkuIds",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CommentReqMultiError(errors)
	}

	return nil
}

// CommentReqMultiError is an error wrapping multiple validation errors
// returned by CommentReq.ValidateAll() if the designated constraints aren't met.
type CommentReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommentReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommentReqMultiError) AllErrors() []error { return m }

// CommentReqValidationError is the validation error returned by
// CommentReq.Validate if the designated constraints aren't met.
type CommentReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentReqValidationError) ErrorName() string { return "CommentReqValidationError" }

// Error satisfies the builtin error interface
func (e CommentReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentReqValidationError{}

// Validate checks the field values on OrderIDReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderIDReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderIDReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderIDReqMultiError, or
// nil if none found.
func (m *OrderIDReq) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderIDReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := OrderIDReqValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OrderIDReqMultiError(errors)
	}

	return nil
}

// OrderIDReqMultiError is an error wrapping multiple validation errors
// returned by OrderIDReq.ValidateAll() if the designated constraints aren't met.
type OrderIDReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderIDReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderIDReqMultiError) AllErrors() []error { return m }

// OrderIDReqValidationError is the validation error returned by
// OrderIDReq.Validate if the designated constraints aren't met.
type OrderIDReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderIDReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderIDReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderIDReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderIDReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderIDReqValidationError) ErrorName() string { return "OrderIDReqValidationError" }

// Error satisfies the builtin error interface
func (e OrderIDReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderIDReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderIDReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderIDReqValidationError{}

// Validate checks the field values on ListReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListReplyMultiError, or nil
// if none found.
func (m *ListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListReplyValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListReplyMultiError(errors)
	}

	return nil
}

// ListReplyMultiError is an error wrapping multiple validation errors returned
// by ListReply.ValidateAll() if the designated constraints aren't met.
type ListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReplyMultiError) AllErrors() []error { return m }

// ListReplyValidationError is the validation error returned by
// ListReply.Validate if the designated constraints aren't met.
type ListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReplyValidationError) ErrorName() string { return "ListReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReplyValidationError{}

// Validate checks the field values on OrderIDReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderIDReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderIDReplyMultiError, or
// nil if none found.
func (m *OrderIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return OrderIDReplyMultiError(errors)
	}

	return nil
}

// OrderIDReplyMultiError is an error wrapping multiple validation errors
// returned by OrderIDReply.ValidateAll() if the designated constraints aren't met.
type OrderIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderIDReplyMultiError) AllErrors() []error { return m }

// OrderIDReplyValidationError is the validation error returned by
// OrderIDReply.Validate if the designated constraints aren't met.
type OrderIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderIDReplyValidationError) ErrorName() string { return "OrderIDReplyValidationError" }

// Error satisfies the builtin error interface
func (e OrderIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderIDReplyValidationError{}

// Validate checks the field values on OrderInfoReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderInfoReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderInfoReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderInfoReplyMultiError,
// or nil if none found.
func (m *OrderInfoReply) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderInfoReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderInfoReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderInfoReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderInfoReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OrderInfoReplyMultiError(errors)
	}

	return nil
}

// OrderInfoReplyMultiError is an error wrapping multiple validation errors
// returned by OrderInfoReply.ValidateAll() if the designated constraints
// aren't met.
type OrderInfoReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderInfoReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderInfoReplyMultiError) AllErrors() []error { return m }

// OrderInfoReplyValidationError is the validation error returned by
// OrderInfoReply.Validate if the designated constraints aren't met.
type OrderInfoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderInfoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderInfoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderInfoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderInfoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderInfoReplyValidationError) ErrorName() string { return "OrderInfoReplyValidationError" }

// Error satisfies the builtin error interface
func (e OrderInfoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderInfoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderInfoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderInfoReplyValidationError{}
