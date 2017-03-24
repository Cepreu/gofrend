package IVRParser

type Prompts struct {
	TTSes                        []TTSPrompt   `xml:"ttsPrompt"`
	Files                        []FilePrompt  `xml:"filePrompt"`
	Pauses                       []PausePrompt `xml:"pausePrompt"`
	Interruptible                bool          `xml:"interruptible"`
	CanChangeInterruptableOption bool          `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool          `xml:"ttsEnumed"`
	ExitModuleOnException        bool          `xml:"exitModuleOnException"`
}

type FilePrompt struct {
	PromptDirectly     bool   `xml:"promptData>promptSelected"`
	PromptId           int32  `xml:"promptData>prompt>id"`
	PromptName         string `xml:"promptData>prompt>name"`
	PromptVariableName string `xml:"promptData>promptVariableName"`
	IsRecordedMessage  bool   `xml:"promptData>isRecordedMessage"`
}
type PausePrompt struct {
	Timeout int32 `xml:"timeout"`
}
type TTSPrompt struct {
	TtsPromptXML    string `xml:"xml"`
	PromptTTSEnumed bool   `xml:"promptTTSEnumed"`
}

type VivrPrompts struct {
	VPrompt                      []VivrPrompt `xml:"vivrPrompt"`
	Interruptible                bool         `xml:"interruptible"`
	CanChangeInterruptableOption bool         `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool         `xml:"ttsEnumed"`
	ExitModuleOnException        bool         `xml:"exitModuleOnException"`
}

type VivrHeader struct {
	VPrompt                      VivrPrompt `xml:"vivrPrompt"`
	Interruptible                bool       `xml:"interruptible"`
	CanChangeInterruptableOption bool       `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool       `xml:"ttsEnumed"`
	ExitModuleOnException        bool       `xml:"exitModuleOnException"`
}

type VivrPrompt struct {
	VivrXML string `xml:"xml"`
}

type multilingualPrompts struct {
	Entries []MLPromptEntry `xml:"entry"`
}
type MLPromptEntry struct {
	Key          string    `xml:"key"`
	Id           string    `xml:"value>promptId"`
	Name         string    `xml:"value>name"`
	Description  string    `xml:"value>description"`
	Type         string    `xml:"value>type"`
	Prompts      MLPrompts `xml:"value>prompts"`
	DefLanguage  string    `xml:"defaultLanguage"`
	IsPersistent bool      `xml:"isPersistent"`
}
type MLPrompts struct {
	Entries []MLanguageEntry `xml:"entry"`
}
type MLanguageEntry struct {
	Language string        `xml:"key,attr"`
	TTSes    []TTSPrompt   `xml:"ttsPrompt"`
	Files    []FilePrompt  `xml:"filePrompt"`
	Pauses   []PausePrompt `xml:"pausePrompt"`
}
