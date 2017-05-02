package web

type EWT struct {
	DomainID      int32 `json:"domain_id"`
	SessionID     int32 `json:"session_id"`
	LastExpiresAt int32 `json:"last_expires_at"`
}
