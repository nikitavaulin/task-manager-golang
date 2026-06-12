package domain

import "context"

const usernameContextKey string = "usernameContextKey"

func UsernameToContext(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, usernameContextKey, username)
}

func UsernameFromContext(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(usernameContextKey).(string)
	return username, ok
}
