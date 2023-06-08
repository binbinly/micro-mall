package trace

var (
	service         = "trace-demo"
	defaultEnv      = "production"
	defaultEndpoint = "http://127.0.0.1:14268/api/traces"
)

// Option is a function that sets some option on the client.
type Option func(c *Options)

// Options control behavior of the client.
type Options struct {
	serviceName   string // The name of this service
	endpoint      string
	samplingRatio float64
	environment   string
}

func applyOptions(options ...Option) Options {
	opts := Options{
		serviceName:   service,
		endpoint:      defaultEndpoint,
		samplingRatio: 1,
		environment:   defaultEnv,
	}
	for _, option := range options {
		option(&opts)
	}

	return opts
}

//WithServiceName with name
func WithServiceName(name string) Option {
	return func(c *Options) {
		c.serviceName = name
	}
}

//WithEndpoint with endpoint
func WithEndpoint(endpoint string) Option {
	return func(c *Options) {
		c.endpoint = endpoint
	}
}

//WithEnv with env
func WithEnv(env string) Option {
	return func(c *Options) {
		c.environment = env
	}
}
