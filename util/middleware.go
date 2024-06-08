package util

import "context"

func GetCtxString(ctx context.Context, ctxKey interface{}) string {
	value := ctx.Value(ctxKey)
	if value != nil {
		valStr, ok := value.(string)
		if ok {
			return valStr
		}
	}

	return ""
}
