package ivr

const (
	CONF_REQUIRED     string = "REQUIRED"
	CONF_NOT_REQUIRED string = "NOT_REQUIRED"
)

type ConfirmData struct {
	ConfirmRequired      string
	RequiredConfidence   int
	MaxAttemptsToConfirm int
	NoInputTimeout       int
	VoicePromptIDs       ModulePrompts
	Events               []*RecoEvent
}
