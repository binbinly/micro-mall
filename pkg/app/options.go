package app

import (
	"pkg/wrap"
)

type Option func(o *options)

type options struct {
	client   bool
	name     string
	version  string
	migrate  func()
	authFunc wrap.AuthFunc
}

// WithMigrate 设置数据库迁移
func WithMigrate(migrate func()) Option {
	return func(o *options) {
		o.migrate = migrate
	}
}

// WithAuthFunc 设置身份验证
func WithAuthFunc(authFunc wrap.AuthFunc) Option {
	return func(o *options) {
		o.authFunc = authFunc
	}
}

// WithName with
func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

// WithVersion with
func WithVersion(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

// WithClient with
func WithClient() Option {
	return func(o *options) {
		o.client = true
	}
}
