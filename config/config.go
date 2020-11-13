// Package config is an interface for dynamic configuration.
package config

import (
	"context"

	"github.com/panovateam/go-micro/config/loader"
	"github.com/panovateam/go-micro/config/reader"
	"github.com/panovateam/go-micro/config/source"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// Init the config
	Init(opts ...Option) error
	// Options in the config
	Options() Options
	// Load value from config
	Load(path ...string) (reader.Value, error)
	// Watch a value for changes
	Watch(path ...string) (Watcher, error)
	// Force a source changeset sync
	Sync() error
	// Stop the config loader/watcher
	Close() error
}

// Watcher is the config watcher
type Watcher interface {
	Next() (reader.Value, error)
	Stop() error
}

type Options struct {
	Loader loader.Loader
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context
}

type Option func(o *Options)
