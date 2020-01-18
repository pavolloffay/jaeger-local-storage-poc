package jaeger

import (
	"strconv"
	"time"

	jaegerModel "github.com/jaegertracing/jaeger/model"
	"github.com/pavolloffay/jaeger-local-storage-poc/pkg/model"
)

// Event The jaeger span version of the event
type Event struct {
	tenant    string
	timestamp time.Time
	fields    []*model.Field
}

// NewEvent creates an event from a Jaeger span
func NewEvent(span *jaegerModel.Span) model.Event {
	fields := make([]*model.Field, 0)
	fields = append(fields, model.NewField("traceId", model.StringType, []byte(span.TraceID.String())))
	fields = append(fields, model.NewField("spanId", model.StringType, []byte(span.SpanID.String())))
	fields = append(fields, model.NewField("service", model.StringType, []byte(span.Process.ServiceName)))
	fields = append(fields, model.NewField("operation", model.StringType, []byte(span.OperationName)))
	fields = append(fields, model.NewField("duration", model.Int64Type, []byte(strconv.FormatInt(span.Duration.Nanoseconds(), 10))))

	// Iterate through process tags - intercept tenant separately
	tenant := ""

	for _, kv := range span.Process.Tags {
		if kv.GetKey() == "tenant" {
			tenant = kv.GetVStr()
			continue
		}
		if field := convertTag(&kv); field != nil {
			fields = append(fields, field)
		}
	}

	// Iterate through span tags
	for _, kv := range span.Tags {
		if field := convertTag(&kv); field != nil {
			fields = append(fields, field)
		}
	}

	return &Event{
		tenant:    tenant,
		timestamp: span.GetStartTime(),
		fields:    fields,
	}
}

func convertTag(kv *jaegerModel.KeyValue) *model.Field {
	switch kv.GetVType() {
	case jaegerModel.ValueType_STRING:
		return model.NewField(kv.GetKey(), model.StringType, []byte(kv.GetVStr()))
	case jaegerModel.ValueType_BOOL:
		//fields = append(fields, model.NewField(kv.GetKey(), model.BooleanType, kv.GetVBool())
	case jaegerModel.ValueType_INT64:
		return model.NewField(kv.GetKey(), model.Int64Type, []byte(strconv.FormatInt(kv.GetVInt64(), 10)))
	case jaegerModel.ValueType_FLOAT64:
		//fields = append(fields, model.NewField(kv.GetKey(), model.Float64Type, []byte(kv.GetVStr())))
	case jaegerModel.ValueType_BINARY:
		return model.NewField(kv.GetKey(), model.BinaryType, kv.GetVBinary())
	default:
	}

	return nil
}

// Tenant The optional tenant associated with the event
func (event *Event) Tenant() string {
	return event.tenant
}

// Type The optional type associated with the event
func (event *Event) Type() string {
	return "jaegerspan"
}

// Timestamp The optional timestamp, will use current time if not defined
func (event *Event) Timestamp() time.Time {
	return event.timestamp
}

// Fields The list of fields to be indexed for the event
func (event *Event) Fields() []*model.Field {
	return event.fields
}
