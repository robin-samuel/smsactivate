package smsactivate

import "context"

type ContextKey string

const SMSACTIVATE_KEY ContextKey = "smsactivate"

func WithContext(ctx context.Context, client *Client) context.Context {
	return context.WithValue(ctx, SMSACTIVATE_KEY, client)
}

func FromContext(ctx context.Context) *Client {
	client, ok := ctx.Value(SMSACTIVATE_KEY).(*Client)
	if !ok {
		return nil
	}
	return client
}
