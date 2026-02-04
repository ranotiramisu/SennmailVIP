package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

// GetIPRotationConfigReq Request to get IP rotation configuration
type GetIPRotationConfigReq struct {
	g.Meta        `path:"/ip_rotation/config" method:"get" tags:"IPRotation" summary:"Get IP rotation configuration"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
}

// GetIPRotationConfigRes Response for IP rotation configuration
type GetIPRotationConfigRes struct {
	api_v1.StandardRes
	Data struct {
		Enabled         bool     `json:"enabled" dc:"Whether IP rotation is enabled"`
		IPs             []string `json:"ips" dc:"List of IPs for rotation"`
		EmailsPerIP     int      `json:"emails_per_ip" dc:"Number of emails to send before rotating IP"`
		CurrentIPIndex  int      `json:"current_ip_index" dc:"Current IP index in rotation"`
		TotalEmailsSent int64    `json:"total_emails_sent" dc:"Total emails sent"`
	} `json:"data"`
}

// SetIPRotationConfigReq Request to set IP rotation configuration
type SetIPRotationConfigReq struct {
	g.Meta        `path:"/ip_rotation/config" method:"post" tags:"IPRotation" summary:"Set IP rotation configuration"`
	Authorization string   `json:"authorization" dc:"Authorization" in:"header"`
	Enabled       bool     `json:"enabled" dc:"Enable or disable IP rotation"`
	IPs           []string `json:"ips" v:"required-if:Enabled,true" dc:"List of IPs for rotation"`
	EmailsPerIP   int      `json:"emails_per_ip" dc:"Number of emails per IP before rotation (default: 3000)"`
}

// SetIPRotationConfigRes Response for setting IP rotation configuration
type SetIPRotationConfigRes struct {
	api_v1.StandardRes
}

// GetIPRotationStatsReq Request to get IP rotation statistics
type GetIPRotationStatsReq struct {
	g.Meta        `path:"/ip_rotation/stats" method:"get" tags:"IPRotation" summary:"Get IP rotation statistics"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
}

// GetIPRotationStatsRes Response for IP rotation statistics
type GetIPRotationStatsRes struct {
	api_v1.StandardRes
	Data struct {
		Enabled             bool   `json:"enabled" dc:"Whether IP rotation is enabled"`
		TotalIPs            int    `json:"total_ips" dc:"Total number of IPs configured"`
		CurrentIPIndex      int    `json:"current_ip_index" dc:"Current IP index"`
		CurrentIP           string `json:"current_ip" dc:"Current IP being used"`
		EmailsPerIP         int    `json:"emails_per_ip" dc:"Emails per IP before rotation"`
		EmailsSentCurrent   int64  `json:"emails_sent_current" dc:"Emails sent with current IP"`
		EmailsUntilRotate   int    `json:"emails_until_rotate" dc:"Emails remaining before next rotation"`
		LastRotation        string `json:"last_rotation" dc:"Last rotation timestamp"`
	} `json:"data"`
}

// DetectSystemIPsReq Request to detect system IPs
type DetectSystemIPsReq struct {
	g.Meta        `path:"/ip_rotation/detect" method:"get" tags:"IPRotation" summary:"Detect available system IPs"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
}

// DetectSystemIPsRes Response for detected system IPs
type DetectSystemIPsRes struct {
	api_v1.StandardRes
	Data struct {
		IPs []string `json:"ips" dc:"Detected system IPs"`
	} `json:"data"`
}
