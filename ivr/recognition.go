package ivr

type RecoType string

const (
	eventNoMatch RecoType = "NO_MATCH"
	eventNoInput RecoType = "NO_INPUT"
	eventHelp    RecoType = "HELP"
)

type RecoAction string

const (
	actionContinue RecoAction = "CONTINUE"
	actionReprompt RecoAction = "REPROMPT"
	actionExit     RecoAction = "EXIT"
)

// RecoEvent - recognition event
type RecoEvent struct {
	Event          RecoType
	Action         RecoAction
	CountAndPrompt AttemptPrompts
}
