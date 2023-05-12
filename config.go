package tracer

type SConfig struct {
	IsEnabled      bool
	Host           string
	Port           string
	serviceId      int
	serviceName    string
	serviceVersion string
	serviceMode    string
	Sampler        bool
	UseStdout      bool
}
