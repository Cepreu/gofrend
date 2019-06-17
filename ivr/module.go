package ivr

import "github.com/Cepreu/gofrend/utils"

// Module - parsed ivr module
type Module interface {
	GetID() ModuleID
	GetDescendant() ModuleID
	normalize(*IVRScript) error
}

// ModuleID - the IVR module's ID, string
type ModuleID string

type module struct {
	ID              ModuleID
	Ascendants      []ModuleID
	Descendant      ModuleID
	ExceptionalDesc ModuleID
	Name            string
	Dispo           string
	Collapsible     bool
}

// GetID - returns ID if the module
func (m *module) GetID() ModuleID {
	return m.ID
}

// GetDescendant - returns ID if the module's descendant
func (m *module) GetDescendant() ModuleID {
	return m.Descendant
}

const (
	cIncomingCall  string = "incomingCall"
	cStartOnHangup string = "startOnHangup"
	cHangup        string = "hangup"
	cGetDigits     string = "getDigits"
	cPlay          string = "play"
	cInput         string = "input"
	cVoiceInput    string = "recording"
	cMenu          string = "menu"
	cQuery         string = "query"
	cSetVariables  string = "setVariable"
	cIfElse        string = "ifElse"
	cCase          string = "case"
	cForeignScript string = "foreignScript"

	cPrompt      string = "prompt"
	cConfirmData string = "ConfirmData"
	cMenuItems   string = "items"
)

type CaseModule struct {
	module
	Branches []*OutputBranch
}

func (*CaseModule) transformToAI() string {
	return ""
}

func (*CaseModule) normalize(*IVRScript) error { return nil }

type ForeignScriptModule struct {
	module
	IvrScript        string
	PassCRM          bool
	ReturnCRM        bool
	Parameters       []keyValueParametrized
	ReturnParameters []keyValue
	IsConsistent     bool
}

func (*ForeignScriptModule) transformToAI() string {
	return ""
}

func (*ForeignScriptModule) normalize(*IVRScript) error { return nil }

type GetDigitsModule struct {
	module
	VoicePromptIDs     ModulePrompts
	VisualPromptIDs    []PromptID
	TextPromptIDs      []PromptID
	TargetVariableName string
	InputInfo          struct {
		NumberOfDigits   int
		TerminateDigit   string
		ClearDigitBuffer bool
		MaxTime          int
		MaxSilence       int
		Format           string
	}
}

func (module *GetDigitsModule) normalize(s *IVRScript) error {
	return s.normalizePrompt(module.VoicePromptIDs)
}

type HangupModule struct {
	module
	Return2Caller bool

	ErrCode       parametrized
	ErrDescr      parametrized
	OverwriteDisp bool
}

func (*HangupModule) normalize(*IVRScript) error {
	return nil
}

type groupingType string //"ALL", "ANY", or "CUSTOM"

type IfElseModule struct {
	module
	BranchIf   OutputBranch
	BranchElse OutputBranch
}

func (*IfElseModule) transformToAI() string {
	return ""
}

func (*IfElseModule) normalize(*IVRScript) error { return nil }

type IncomingCallModule struct {
	module
}

func (*IncomingCallModule) normalize(*IVRScript) error {
	return nil
}

type xInputGrammar struct {
	Type           string           `xml:"type,attr"`
	MRVname        string           `xml:"mainReturnValue>name"`
	MRVtype        string           `xml:"mainReturnValue>type"`
	MRVvariable    string           `xml:"mainReturnValue>varName"`
	StrProperties  []xGrammPropStr  `xml:"stringProperty"`
	ListProperties []xGrammPropList `xml:"listProperty"`
	ARVname        string           `xml:"additionalReturnValues>name"`
	ARVtype        string           `xml:"additionalReturnValues>type"`
}
type xGrammPropStr struct {
	PropType    string `xml:"type"`
	PropValue   string `xml:"value"`
	PropEnabled bool   `xml:"enabled"`
}
type xGrammPropList struct {
	PropType    string   `xml:"type"`
	PropValEnum []string `xml:"list"`
	PropValue   string   `xml:"value"`
	PropEnabled bool     `xml:"enabled"`
}

type InputModule struct {
	module
	VoicePromptIDs ModulePrompts
	//	VisualPromptIDs  modulePrompt
	//	TextPromptIDs    modulePrompt
	Grammar  xInputGrammar
	Events   []*RecoEvent
	ConfData *ConfirmData
}

func (module *InputModule) normalize(s *IVRScript) error {
	return s.normalizePrompt(module.VoicePromptIDs)
}

//MenuModule - Menu module definition
type MenuModule struct {
	module

	VoicePromptIDs ModulePrompts
	//	VisualPromptIDs  []PromptID
	//	TextPromptIDs    []PromptID

	Branches []*OutputBranch
	Items    []*MenuItem

	UseASR  bool
	UseDTMF bool
	//	RecordUserInput bool

	Events   []*RecoEvent
	ConfData *ConfirmData

	RecoParams struct {
		SpeechCompleteTimeout int
		MaxTimeToEnter        int
		NoInputTimeout        int
	}
}

func (module *MenuModule) normalize(s *IVRScript) error {
	s.normalizePrompt(module.VoicePromptIDs)
	for i := range module.Items {
		s.normalizeAttemptPrompt(&module.Items[i].Prompt, false)
	}
	return nil
}

// ActionType - Menu module item's action
type ActionType string

type PlayModule struct {
	module
	VoicePromptIDs ModulePrompts
	//	VisualPromptIDs modulePrompt
	//	TextPromptIDs   modulePrompt
	AuxInfo struct {
		NumberOfDigits   int
		TerminateDigit   string
		ClearDigitBuffer bool
	}
}

func (module *PlayModule) normalize(s *IVRScript) error {
	return s.normalizePrompt(module.VoicePromptIDs)
}

type (
	QueryModule struct {
		module
		VoicePromptIDs ModulePrompts
		//	VisualPromptIDs modulePrompt
		//	TextPromptIDs   modulePrompt

		URL                                  string
		Method                               string
		FetchTimeout                         int
		StoreNumberOfArrayElementsInVariable bool

		Parameters       []keyValueParametrized
		ReturnValues     []keyValue
		URLParts         []*parametrized
		RequestInfo      requestInfo
		Headers          []keyValueParametrized
		RequestBodyType  string
		SaveStatusCode   bool
		SaveReasonPhrase bool
		ResponseInfos    []*responseInfo
	}

	keyValue struct {
		Key   string
		Value string
	}

	responseInfo struct {
		HTTPCodeFrom  int
		HTTPCodeTo    int
		ParsingMethod string
		Regexp        struct {
			RegexpBody  string
			RegexpFlags int
		}
		Function struct {
			Name       string
			ReturnType string
			Arguments  string
		}
		TargetVariables []string
	}
	requestInfo struct {
		Template     string
		base64       string
		Replacements []*replacement
	}
	replacement struct {
		Position     int
		VariableName string
	}
)

func (module *QueryModule) normalize(s *IVRScript) error {
	err := s.normalizePrompt(module.VoicePromptIDs)
	if err != nil {
		return err
	}
	module.RequestInfo.Template, err = utils.CmdUnzip(module.RequestInfo.base64)
	return err
}

type (
	SetVariableModule struct {
		module
		Exprs []*expression
	}

	expression struct {
		Lval   string
		IsFunc bool
		Rval   assigner
	}

	assigner struct {
		P *parametrized
		F *ivrFuncInvocation
	}

	ivrFunc struct {
		Name       string
		ReturnType string
		ArgTypes   []string
	}

	ivrFuncInvocation struct {
		FuncDef ivrFunc
		Params  []*parametrized
	}
)

func (module *SetVariableModule) normalize(s *IVRScript) (err error) {
	return nil
}

type VoiceInputModule struct {
	module
	VoicePromptIDs ModulePrompts
	//	VisualPromptIDs  modulePrompt
	//	TextPromptIDs    modulePrompt
	Events   []*RecoEvent
	ConfData *ConfirmData

	RecordingParams struct {
		MaxTime                           int
		FinalSilence                      int
		DTMFtermination                   bool
		ThrowExceptionIfMaxSilenceReached bool
		MaxAttempts                       int
	}
	PostRecording struct {
		VarToAccessRecording    string
		RecordingDurationVar    string
		TerminationCharacterVar string
	}
	AsGreeting struct {
		UseRecordingAsGreeting bool
		IsVarSelected          bool
		StringValue            string
		VariableName           string
	}
}

func (module *VoiceInputModule) normalize(s *IVRScript) error {
	return s.normalizePrompt(module.VoicePromptIDs)
}
