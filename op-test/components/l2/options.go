package l2

import (
	"github.com/ethereum-optimism/optimism/op-test/components/superchain"
	"github.com/ethereum-optimism/optimism/op-test/test"
)

type Settings struct {
	Kind test.BackendKind

	Superchain superchain.Superchain
}

type Option interface {
	Apply(settings *Settings) error
}

type OptionFn func(settings *Settings) error

func (fn OptionFn) Apply(settings *Settings) error {
	return fn(settings)
}

func Kind(kind test.BackendKind) Option {
	return OptionFn(func(settings *Settings) error {
		settings.Kind = kind
		return nil
	})
}

func Superchain(ch superchain.Superchain) Option {
	return OptionFn(func(settings *Settings) error {
		settings.Superchain = ch
		return nil
	})
}

// TODO active hardfork option

// TODO scheduled hardforks option

// TODO feature-toggle option (FPs, Interop

// TODO L2 contract settings / constraints options

// TODO genesis account funding option

// TODO reservation type option (similar to L1 chain reservation type option)
// hands-off: work against any system, no config changes
// temp: temporary config changes
// controlled: breaking setup or operation changes