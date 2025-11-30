package waiter

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/grafvonb/c8volt/config"
	d "github.com/grafvonb/c8volt/internal/domain"
	"github.com/grafvonb/c8volt/internal/services"
)

type PIWaiter interface {
	GetProcessInstance(ctx context.Context, key string, opts ...services.CallOption) (d.ProcessInstance, error)
	GetProcessInstanceStateByKey(ctx context.Context, key string, opts ...services.CallOption) (d.State, d.ProcessInstance, error)
}

func WaitForProcessInstanceState(ctx context.Context, s PIWaiter, cfg *config.Config, log *slog.Logger, key string, desired d.States, opts ...services.CallOption) (d.State, d.ProcessInstance, error) {
	_ = services.ApplyCallOptions(opts)
	backoff := cfg.App.Backoff
	start := time.Now()
	if backoff.Timeout > 0 {
		deadline := time.Now().Add(backoff.Timeout)
		if dl, ok := ctx.Deadline(); !ok || deadline.Before(dl) {
			var cancel context.CancelFunc
			ctx, cancel = context.WithDeadline(ctx, deadline)
			defer cancel()
		}
	}

	attempts := 0
	delay := backoff.InitialDelay
	for {
		if errInDelay := ctx.Err(); errInDelay != nil {
			{
				elapsed := time.Since(start)
				log.Info(fmt.Sprintf("stopped waiting for process instance %s after %d attempts in %s (context error: %v)", key, attempts, elapsed, errInDelay))
			}
			return "", d.ProcessInstance{}, errInDelay
		}
		attempts++
		log.Debug(fmt.Sprintf("attempt #%d to fetch state for process instance %s", attempts, key))
		got, pi, errInDelay := s.GetProcessInstanceStateByKey(ctx, key)
		if errInDelay == nil {
			if stateIn(got, desired) {
				if attempts == 1 {
					{
						elapsed := time.Since(start)
						log.Info(fmt.Sprintf("process instance %s is already in one of the desired state(s) [%s] (current: %s); attempts=%d, elapsed=%s", key, desired, got, attempts, elapsed))
					}
					log.Debug(fmt.Sprintf("process instance %s is already in one of the desired state(s) [%s] (current: %s)", key, desired, got))
					return got, pi, nil
				}
				{
					elapsed := time.Since(start)
					log.Info(fmt.Sprintf("process instance %s reached one of the desired state(s) [%s] (current: %s) after %d checks in %s", key, desired, got, attempts, elapsed))
				}
				log.Debug(fmt.Sprintf("process instance %s reached one of the desired state(s) [%s] (current: %s) after %d checks", key, desired, got, attempts))
				return got, pi, nil
			}
			log.Info(fmt.Sprintf("process instance %s currently in state %s; waiting... (attempt %d)", key, got, attempts))
		} else if errInDelay != nil {
			if strings.Contains(errInDelay.Error(), "404") {
				got = d.StateAbsent
				if stateIn(got, desired) {
					{
						elapsed := time.Since(start)
						log.Info(fmt.Sprintf("process instance %s reached one of the desired state(s) [%s] (current: %s) after %d checks in %s", key, desired, got, attempts, elapsed))
					}
					log.Debug(fmt.Sprintf("process instance %s reached one of the desired state(s) [%s] (current: %s) after %d checks", key, desired, got, attempts))
					return got, d.ProcessInstance{}, nil
				}
				log.Debug(fmt.Sprintf("process instance %s is absent (not found); waiting... (attempt %d)", key, attempts))
			} else {
				log.Error(fmt.Sprintf("fetching state for %q failed: %v (will NOT retry)", key, errInDelay))
				{
					elapsed := time.Since(start)
					log.Info(fmt.Sprintf("stopped waiting for process instance %s after %d attempts in %s due to error", key, attempts, elapsed))
				}
				return "", d.ProcessInstance{}, fmt.Errorf("fetching state for %q failed: %w", key, errInDelay)
			}
		}
		if backoff.MaxRetries > 0 && attempts >= backoff.MaxRetries {
			{
				elapsed := time.Since(start)
				log.Info(fmt.Sprintf("exceeded max_retries (%d) waiting for state %q of process instance %s after %d attempts in %s", backoff.MaxRetries, desired, key, attempts, elapsed))
			}
			return "", d.ProcessInstance{}, fmt.Errorf("exceeded max_retries (%d) waiting for state %q", backoff.MaxRetries, desired)
		}
		select {
		case <-time.After(delay):
			delay = backoff.NextDelay(delay)
		case <-ctx.Done():
			{
				elapsed := time.Since(start)
				log.Info(fmt.Sprintf("stopped waiting for process instance %s after %d attempts in %s (context done: %v)", key, attempts, elapsed, ctx.Err()))
			}
			return "", d.ProcessInstance{}, fmt.Errorf("%w: %s", d.ErrGatewayTimeout, ctx.Err().Error())
		}
	}
}

func stateIn(st d.State, set d.States) bool {
	for _, x := range set {
		if st.EqualsIgnoreCase(x) {
			return true
		}
	}
	return false
}
