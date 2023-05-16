package tracer

type sConfig struct {
	isEnabled         bool
	sampler           bool
	useStdout         bool
	jaegerHost        string
	jaegerPort        string
	serviceId         int
	serviceName       string
	serviceNamespace  string
	serviceInstanceId string
	serviceVersion    string
	serviceMode       string
	serviceCommitId   string
}
