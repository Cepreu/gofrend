package ivr

import "strings"

// Module - parsed ivr module
type Module interface {
	GetID() ModuleID
	GetDescendant() ModuleID
	SetGeneralInfo(string, ModuleID, []ModuleID, ModuleID, ModuleID, string, string)
}

// ModuleID - the IVR module's ID, string
type ModuleID string

// Lower returns lower case string representation of ID
func (ID ModuleID) Lower() string {
	return strings.ToLower(string(ID))
}

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

func (m *module) SetGeneralInfo(name string, id ModuleID,
	ascendants []ModuleID, descendant ModuleID, exceptionalDesc ModuleID,
	dispo string, collapsible string) {
	if id != "" {
		m.ID = id
	}
	if len(ascendants) > 0 {
		m.Ascendants = append(m.Ascendants, ascendants...)
	}
	if descendant != "" {
		m.Descendant = descendant
	}
	if exceptionalDesc != "" {
		m.ExceptionalDesc = exceptionalDesc
	}
	if name != "" {
		m.Name = name
	}
	if dispo != "" {
		m.Dispo = dispo
	}
	if collapsible == "true" {
		m.Collapsible = true
	}
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

//ForeignScriptModule - Struct representing IVR's Foreign Script module
type ForeignScriptModule struct {
	module
	IvrScript        string
	PassCRM          bool
	ReturnCRM        bool
	Parameters       []KeyValue
	ReturnParameters []KeyValue
	IsConsistent     bool
}

func (*ForeignScriptModule) transformToAI() string {
	return ""
}

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

//HangupModule - Describes IVR's Hangup module
type HangupModule struct {
	module
	Return2Caller bool
	ErrCode       VariableID
	ErrDescr      VariableID
	OverwriteDisp bool
}

type IfElseModule struct {
	module
	BranchIf   OutputBranch
	BranchElse OutputBranch
}

func (*IfElseModule) transformToAI() string {
	return ""
}

type IncomingCallModule struct {
	module
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

		Parameters       []KeyValue
		ReturnValues     []KeyValue
		URLParts         []VariableID
		RequestInfo      RequestInfo
		Headers          []KeyValue
		RequestBodyType  string
		SaveStatusCode   bool
		SaveReasonPhrase bool
		ResponseInfos    []*ResponseInfo
	}

	KeyValue struct {
		Key   string
		Value VariableID
	}

	ResponseInfo struct {
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
	RequestInfo struct {
		Template     string
		Base64       string
		Replacements []*Replacement
	}
	Replacement struct {
		Position     int
		VariableName string
	}
)

type (
	SetVariableModule struct {
		module
		Exprs []*Expression
	}
	Expression struct {
		Lval string
		Rval FuncInvocation
	}
)

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
