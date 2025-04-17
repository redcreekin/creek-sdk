package telemetry

const (
	HostAttribute = "http.host"
)

type Configuration struct {
	TracingEnabled bool `toml:"tracing_enabled"`
}
