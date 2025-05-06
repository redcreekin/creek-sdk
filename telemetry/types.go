package telemetry

import (
	"time"

	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace/propagation"
)

const (
	HostAttribute      = "http.host"
	MethodAttribute    = "http.method"
	PathAttribute      = "http.path"
	StatusAttribute    = "http.status_code"
	SizeAttribute      = "http.size"
	DurationAttribute  = "http.duration"
	ServiceAttribute   = "service"
	ProjectAttribute   = "project"
	WorkflowAttribute  = "workflow"
	UserAgentAttribute = "http.user_agent"
)

var DefaultFormat propagation.HTTPFormat = &b3.HTTPFormat{}

type Configuration struct {
	TracingEnabled bool `toml:"tracingEnabled" validate:"omitempty"`
	Exporters      struct {
		Jaeger struct {
			ServiceName         string  `toml:"serviceName" validate:"omitempty" default:""`
			CollectorEndpoint   string  `toml:"collectorEndpoint" validate:"omitempty" default:"http://localhost:14268/api/traces"`
			SamplingProbability float64 `toml:"samplingProbability" validate:"omitempty" default:"0.1"`
		}
	}
}

var (
	// DefaultSizeDistribution 25k, 100k, 250k, 500k, 1M, 2M, 5M, 10M
	DefaultSizeDistribution = view.Distribution(25*1024, 100*1024, 250*1024, 500*1024, 1*1024*1024, 2*1024*1024, 5*1024*1024, 10*1024*1024)
	// DefaultLatencyDistribution 1ms, 5ms, 10ms, 25ms, 50ms, 100ms, 250ms, 500ms, 1s, 2s
	DefaultLatencyDistribution = view.Distribution(100, 200, 300, 400, 500, 750, 1000, 2000, 5000)
)

type ExposedView struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Dimension   string   `json:"dimension"`
	Aggregation string   `json:"aggregagtion"`
}

type contextKey int

const (
	contextTraceExporter contextKey = iota
	contextStatsExporter
)

// B3 headers that OpenCensus understands.
const (
	TraceIdHeader  = "X-B3-TraceId"
	SpanIdHeader   = "X-B3-SpanId"
	SampledHeader  = "X-B3-Sampled"
	FlagsHeader    = "X-B3-Flags"
	ParentIdHeader = "X-B3-ParentSpanId"

	ContextTraceIDHeader contextKey = iota
	ContextSpanIDHeader
	ContextSampledHeader
	ContextMainSpan
)

type HTTPExporterView struct {
	Name  string            `json:"name"`
	Tags  map[string]string `json:"tags"`
	Value float64           `json:"value"`
	Date  time.Time         `json:"date"`
}
