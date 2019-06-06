package ivr

type recoevent string

const (
	eventNoMatch recoevent = "NO_MATCH"
	eventNoInput recoevent = "NO_INPUT"
	eventHelp    recoevent = "HELP"
)

type recoaction string

const (
	actionContinue recoaction = "CONTINUE"
	actionReprompt recoaction = "REPROMPT"
	actionExit     recoaction = "EXIT"
)

// RecoEvent - recognition event
type RecoEvent struct {
	Event          recoevent
	Action         recoaction
	CountAndPrompt AttemptPrompts
}
