package propraven

import "context"

// HealthService binds to /v1/health.
type HealthService struct {
	client *Client
}

// Check returns the API's current health. Useful as a connectivity probe in
// startup checks and as a cheap target for keep-alive pings.
func (s *HealthService) Check(ctx context.Context) (*HealthStatus, error) {
	var h HealthStatus
	if err := s.client.do(ctx, "GET", "/v1/health", nil, nil, &h); err != nil {
		return nil, err
	}
	return &h, nil
}
