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
			jaegerModel.String("key2", "value2"),
		},
		Process: &jaegerModel.Process{
			ServiceName: "serviceA",
			Tags: jaegerModel.KeyValues{
				jaegerModel.String("tenant", "tenant1"),
				jaegerModel.String("key3", "value3"),
			},
		},
	}

	event := NewEvent(span)
	assert.Equal(t, "tenant1", event.Tenant())
	assert.Equal(t, "jaegerspan", event.Type())
	assert.Equal(t, time.Unix(0, 1000), event.Timestamp())

	len := 8
	assert.Len(t, event.Fields(), len)

	tests := []struct {
		fieldName  string
		fieldType  model.FieldType
		fieldValue []byte
	}{
		{fieldName: "traceId", fieldType: model.StringType, fieldValue: []byte(span.TraceID.String())},
		{fieldName: "spanId", fieldType: model.StringType, fieldValue: []byte(span.SpanID.String())},
		{fieldName: "service", fieldType: model.StringType, fieldValue: []byte(span.Process.ServiceName)},
		{fieldName: "operation", fieldType: model.StringType, fieldValue: []byte(span.OperationName)},
		{fieldName: "duration", fieldType: model.Int64Type, fieldValue: []byte(strconv.FormatInt(span.Duration.Nanoseconds(), 10))},
		{fieldName: "key1", fieldType: model.StringType, fieldValue: []byte("value1")},
		{fieldName: "key2", fieldType: model.StringType, fieldValue: []byte("value2")},
		{fieldName: "key3", fieldType: model.StringType, fieldValue: []byte("value3")},
	}

	// Confirm a test entry exists for each field
	assert.Len(t, tests, len)

	for _, test := range tests {
		found := false
		for _, field := range event.Fields() {
			if field.Name() != test.fieldName {
				continue
			}
			assert.Equal(t, test.fieldType, field.Type())
			assert.Equal(t, test.fieldValue, field.Value())
			found = true
			break
		}
		if !found {
			assert.Fail(t, "Field not found", "Field name %s", test.fieldName)
		}
	}
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

	event := NewEvent(span)
	assert.Equal(t, "", event.Tenant())
}
