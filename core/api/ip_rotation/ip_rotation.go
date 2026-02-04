package ip_rotation

import (
	v1 "billionmail-core/api/ip_rotation/v1"
	"context"
)

type IIPRotationV1 interface {
	// GetIPRotationConfig Get IP rotation configuration
	GetIPRotationConfig(ctx context.Context, req *v1.GetIPRotationConfigReq) (*v1.GetIPRotationConfigRes, error)
	// SetIPRotationConfig Set IP rotation configuration
	SetIPRotationConfig(ctx context.Context, req *v1.SetIPRotationConfigReq) (*v1.SetIPRotationConfigRes, error)
	// GetIPRotationStats Get IP rotation statistics
	GetIPRotationStats(ctx context.Context, req *v1.GetIPRotationStatsReq) (*v1.GetIPRotationStatsRes, error)
	// DetectSystemIPs Detect available system IPs
	DetectSystemIPs(ctx context.Context, req *v1.DetectSystemIPsReq) (*v1.DetectSystemIPsRes, error)
}
