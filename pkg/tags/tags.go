package tags

// List all available tags.
// Listing all available tags makes it easier to create a standardized tag list.
const (
	// App describes application name.
	App = "app"
	// Env describes current application environment.
	Env = "env"
	// Time describes a timestamp.
	Time = "time"
	// Method describes process method name.
	Method = "method"
	// MetricStatus describes process resulting status.
	// Available value is "attempt", "success", "error", or "expected_error".
	// Refer to errorx.MetricStatus for more detail.
	MetricStatus = "metric_status"
	// Code describes current process resulting code.
	// Refer to errorx.Code for more detail.
	Code = "code"
	// Error describes current process error.
	Error = "err"
	// Request describes current request payload.
	Request = "request"
	// Response describes current response payload.
	Response = "response"
	// Latency describes the duration to execute the operation.
	Latency = "latency"
	// RequestID describes current request id from context.
	// Each operation should have one unique request id.
	RequestID = "request_id"
	// ContextID describes current context id from context.
	// Context id is not unique but based on current operation context.
	// For instance, operation get subscription detail,
	// context would be subscription, and context id would be subscription id.
	ContextID = "context_id"
	// ErrorLine describes the original line where the error was received or created.
	ErrorLine = "err_line"
	// StackTrace describes the stack trace of an operation.
	StackTrace = "stack_trace"
	// Panic describes the panic trigger.
	Panic = "panic"
	// Message describes the user message.
	Message = "message"
	// Detail describes the operation detail usually combination request, payload, and latency.
	Detail = "detail"
	// Fields describes the errorx.Fields.
	Fields = "fields"
	// User describes the username.
	User = "user"
	// Hash describes current build hash number.
	Hash = "hash"
	// Build describes current build detail, usually combination build number, and timestamp.
	Build = "build"
	// Ops describes list of internal operations that is being called.
	Ops = "ops"
	// Topic describes consumer or publisher topic name.
	Topic = "topic"
	// Channel describes consumer channel name.
	Channel = "channel"
	// AddMention describes is slack message should add here mention.
	AddMention = "add_mention"
	// Template describes template name.
	Template = "template"
	// Database describes database name.
	Database = "database"
	// Setting describes setting or configuration value.
	Setting = "setting"
	// UUID
	UUID = "uuid"
	// FilterType
	FilterType = "filter_type"
	// TotalData
	TotalData = "total_data"
	// Limit
	Limit = "limit"
	// Offset
	Offset = "offset"
	// WithoutCache describes current operation return data from db directly, not from cache.
	WithoutCache = "without_cache"
	// Shared describes current result is result from a single group operation.
	Shared = "shared"
	// Key
	Key = "key"
	// Address describes the ip address.
	Address = "address"
	// Hostname describes current machine hostname.
	Hostname = "hostname"
	// Kind
	Kind = "kind"
	// Count
	Count = "count"
	// Duration
	Duration = "duration"
	// Mode
	Mode = "mode"
	// URL
	URL = "url"
	// Event
	Event = "event"
	// OpenConnections
	OpenConnections = "open_connections"
	// OpenConnectionAlertThreshold
	OpenConnectionAlertThreshold = "open_connection_alert_threshold"
	// MaxOpenConnections
	MaxOpenConnections = "max_open_connections"
	// InUse
	InUse = "in_use"
	// Idle
	Idle = "idle"
	// WaitCount
	WaitCount = "wait_count"
	// WaitDuration
	WaitDuration = "wait_duration"
	// MaxIdleClosed
	MaxIdleClosed = "max_idle_closed"
	// MaxLifetimeClosed
	MaxLifetimeClosed = "max_lifetime_closed"
	// Metadata
	Metadata = "metadata"
	// Cache
	Cache = "cache"
	// Client
	Client = "client"
)
