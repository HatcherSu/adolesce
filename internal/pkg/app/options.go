package app

import "os"

// Option defines optional parameters for initializing the application structure
type Option func(*App)

func WithShort(short string) Option {
	return func(a *App) {
		a.shortDesc = short
	}
}

func WithLong(long string) Option {
	return func(a *App) {
		a.LongDesc = long
	}
}

func WithRunFunc(f RunFunc) Option {
	return func(a *App) {
		a.runFunc = f
	}
}

// WithSignal Signal with exit signals.
func WithSignal(sigs ...os.Signal) Option {
	return func(a *App) {
		a.sigs = sigs
	}
}
