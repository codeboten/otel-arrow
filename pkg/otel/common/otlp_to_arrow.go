package common

import (
	commonpb "otel-arrow-adapter/api/go.opentelemetry.io/proto/otlp/common/v1"
	resourcepb "otel-arrow-adapter/api/go.opentelemetry.io/proto/otlp/resource/v1"
	"otel-arrow-adapter/pkg/otel/constants"
	"otel-arrow-adapter/pkg/rbb"
	"otel-arrow-adapter/pkg/rbb/field_value"
)

func NewAttributes(attributes []*commonpb.KeyValue) *field_value.Field {
	if attributes == nil || len(attributes) == 0 {
		return nil
	}

	attributeFields := make([]field_value.Field, 0, len(attributes))

	for _, attribute := range attributes {
		value := OtlpAnyValueToValue(attribute.Value)
		if value != nil {
			attributeFields = append(attributeFields, field_value.Field{
				Name:  attribute.Key,
				Value: value,
			})
		}
	}
	if len(attributeFields) > 0 {
		attrs := field_value.MakeStructField(constants.ATTRIBUTES, field_value.Struct{
			Fields: attributeFields,
		})
		return &attrs
	}
	return nil
}

func AddResource(record *rbb.Record, resource *resourcepb.Resource) {
	resourceField := ResourceField(resource)
	if resourceField != nil {
		// ToDo check optimization for when fields are always pointers or interfaces instead of structs as today.
		record.AddField(*resourceField)
	}
}

func ResourceField(resource *resourcepb.Resource) *field_value.Field {
	var resourceFields []field_value.Field

	attributes := NewAttributes(resource.Attributes)
	if attributes != nil {
		resourceFields = append(resourceFields, *attributes)
	}

	if resource.DroppedAttributesCount > 0 {
		resourceFields = append(resourceFields, field_value.MakeU32Field(constants.DROPPED_ATTRIBUTES_COUNT, resource.DroppedAttributesCount))
	}
	if len(resourceFields) > 0 {
		field := field_value.MakeStructField(constants.RESOURCE, field_value.Struct{
			Fields: resourceFields,
		})
		return &field
	} else {
		return nil
	}
}

func AddScope(record *rbb.Record, scopeKey string, scope *commonpb.InstrumentationScope) {
	scopeField := ScopeField(scopeKey, scope)
	if scopeField != nil {
		// ToDo check optimization for when fields are always pointers or interfaces instead of structs as today.
		record.AddField(*scopeField)
	}
}

func ScopeField(scopeKey string, scope *commonpb.InstrumentationScope) *field_value.Field {
	var fields []field_value.Field

	fields = append(fields, field_value.MakeStringField(constants.NAME, scope.Name))
	fields = append(fields, field_value.MakeStringField(constants.VERSION, scope.Version))
	attributes := NewAttributes(scope.Attributes)
	if attributes != nil {
		fields = append(fields, *attributes)
	}
	if scope.DroppedAttributesCount > 0 {
		fields = append(fields, field_value.MakeU32Field(constants.DROPPED_ATTRIBUTES_COUNT, scope.DroppedAttributesCount))
	}

	field := field_value.MakeStructField(scopeKey, field_value.Struct{
		Fields: fields,
	})
	return &field
}

func OtlpAnyValueToValue(value *commonpb.AnyValue) field_value.Value {
	if value != nil {
		switch value.Value.(type) {
		case *commonpb.AnyValue_BoolValue:
			return &field_value.Bool{Value: value.GetBoolValue()}
		case *commonpb.AnyValue_IntValue:
			return &field_value.I64{Value: value.GetIntValue()}
		case *commonpb.AnyValue_DoubleValue:
			return &field_value.F64{Value: value.GetDoubleValue()}
		case *commonpb.AnyValue_StringValue:
			return &field_value.String{Value: value.GetStringValue()}
		case *commonpb.AnyValue_BytesValue:
			return &field_value.Binary{Value: value.GetBytesValue()}
		case *commonpb.AnyValue_ArrayValue:
			values := value.GetArrayValue()
			fieldValues := make([]field_value.Value, 0, len(values.Values))
			for _, value := range values.Values {
				v := OtlpAnyValueToValue(value)
				if v != nil {
					fieldValues = append(fieldValues, v)
				}
			}
			return &field_value.List{Values: fieldValues}
		case *commonpb.AnyValue_KvlistValue:
			values := value.GetKvlistValue()
			if values == nil || len(values.Values) == 0 {
				return nil
			} else {
				fields := make([]field_value.Field, 0, len(values.Values))
				for _, kv := range values.Values {
					v := OtlpAnyValueToValue(kv.Value)
					if v != nil {
						fields = append(fields, field_value.Field{
							Name:  kv.Key,
							Value: v,
						})
					}
				}
				return &field_value.Struct{Fields: fields}
			}
		default:
			return nil
		}
	} else {
		return nil
	}
}