package jaeger

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	jaegerModel "github.com/jaegertracing/jaeger/model"

	"github.com/pavolloffay/jaeger-local-storage-poc/pkg/model"
)

func TestJaegerSpan(t *testing.T) {
	span := &jaegerModel.Span{
		TraceID:       jaegerModel.NewTraceID(0, 123),
		SpanID:        jaegerModel.NewSpanID(567),
		OperationName: "op1",
		StartTime:     time.Unix(0, 1000),
		Duration:      5000,
		Tags: jaegerModel.KeyValues{
			jaegerModel.String("key1", "value1"),
			jaegerModel.Bool("key2", true),
		},
		Process: &jaegerModel.Process{
			ServiceName: "serviceA",
			Tags: jaegerModel.KeyValues{
				jaegerModel.String("tenant", "tenant1"),
				jaegerModel.String("key3", "value3"),
			},
		},
	}

	event, err := NewEvent(span)
	assert.NoError(t, err)
	assert.Equal(t, "tenant1", event.Tenant())
	assert.Equal(t, "jaegerspan", event.Type())
	assert.Equal(t, time.Unix(0, 1000), event.Timestamp())

	len := 9
	assert.Len(t, event.Fields(), len)

	restrue := byte(1)
	spanData, _ := span.Marshal()

	fields := make([]*model.Field, 0)
	fields = append(fields, model.NewField("__traceId", model.StringType, []byte(span.TraceID.String())),
		model.NewField("__spanId", model.StringType, []byte(span.SpanID.String())),
		model.NewField("__service", model.StringType, []byte(span.Process.ServiceName)),
		model.NewField("__operation", model.StringType, []byte(span.OperationName)),
		model.NewField("__duration", model.Int64Type, []byte(strconv.FormatInt(span.Duration.Nanoseconds(), 10))),
		model.NewField("key3", model.StringType, []byte("value3")),
		model.NewField("key1", model.StringType, []byte("value1")),
		model.NewField("key2", model.BooleanType, []byte{restrue}),
		model.NewField("__span", model.BinaryType, spanData))
	assert.Equal(t, fields, event.Fields())
}

func TestJaegerSpanNoTenant(t *testing.T) {
	span := &jaegerModel.Span{
		TraceID:       jaegerModel.NewTraceID(0, 123),
		SpanID:        jaegerModel.NewSpanID(567),
		OperationName: "op1",
		StartTime:     time.Unix(0, 1000),
		Duration:      5000,
		Process: &jaegerModel.Process{
			ServiceName: "serviceA",
		},
	}

	event, err := NewEvent(span)
	assert.NoError(t, err)
	assert.Equal(t, "", event.Tenant())
}
