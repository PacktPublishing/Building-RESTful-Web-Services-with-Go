package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

// InstrumentingMiddleware is a struct representing middleware
type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           EncryptService
}

func (mw InstrumentingMiddleware) Encrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "encrypt", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Encrypt(ctx, key, text)
	return
}

func (mw InstrumentingMiddleware) Decrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "decrypt", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}

