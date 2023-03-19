package gchat

type Option func(opts *Options)

const (
	Redis = iota
	File
)

type Options struct {
	// 是否持久化
	Persistence bool
	// 持久化方式
	PersistenceType int
	// 钩子
	Hook Hook
	// 自定义日记
	Logger Logger
}

func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

// WithOptions accepts the whole options' config.
func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

// WithPersistence 是否持久化
func WithPersistence(persistence bool) Option {
	return func(opts *Options) {
		opts.Persistence = persistence
	}
}

// WithHook Define hooks
func WithHook(hook Hook) Option {
	return func(opts *Options) {
		opts.Hook = hook
	}
}
