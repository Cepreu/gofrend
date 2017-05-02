package web

type Idle struct {
	DomainID  int32      `json:"domain_id"`
	SessionID int32      `json:"session_id"`
	args      ScriptArgs `json:"args"`
}
