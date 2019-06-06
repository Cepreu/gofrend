package ivr

const (
	menuActionTypeEVENT  string = "EVENT"
	menuActionTypeBRANCH string = "BRANCH"
)

const (
	menuEventHELP    string = "HELP"
	menuEventNOINPUT string = "NO_INPUT"
	menuEventNOMATCH string = "NO_MATCH"
)

type MenuItem struct {
	Prompt     AttemptPrompts
	ShowInVivr bool
	MatchExact bool
	Dtmf       string
	Action     struct {
		Type ActionType
		Name string
	}
}
