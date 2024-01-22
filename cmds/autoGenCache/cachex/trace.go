package cachex

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var sqlAttributeKey = attribute.Key("cache.method")

// spanName is used to identify the span name for the SQL execution.
const spanName = "cache"

func StartSpan(ctx context.Context, method string) (context.Context, oteltrace.Span) {
	tracer := trace.TracerFromContext(ctx)
	start, span := tracer.Start(ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	span.SetAttributes(sqlAttributeKey.String(method))

	return start, span
}

func EndSpan(span oteltrace.Span, err error) {
	defer span.End()

	if err == nil || errors.Is(err, sql.ErrNoRows) {
		span.SetStatus(codes.Ok, "")
		return
	}

	span.SetStatus(codes.Error, err.Error())
	span.RecordError(err)
}
