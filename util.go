package ginzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func zapFieldStringsByStringMap(key string, m map[string][]string) zapcore.Field {
	return zap.Object(key, stringsByStringMarshaler(m))
}

func stringsByStringMarshaler(m map[string][]string) zapcore.ObjectMarshalerFunc {
	return func(inner zapcore.ObjectEncoder) error {
		for k, values := range m {
			err := inner.AddArray(k, zapcore.ArrayMarshalerFunc(func(inner zapcore.ArrayEncoder) error {
				for _, v := range values {
					inner.AppendString(v)
				}
				return nil
			}))
			if err != nil {
				return err
			}
		}
		return nil
	}
}
