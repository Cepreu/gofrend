package IVRParser

type xPrompts struct {
	TTSes                        []xTTSPrompt   `xml:"ttsPrompt"`
	Files                        []xFilePrompt  `xml:"filePrompt"`
	Pauses                       []xPausePrompt `xml:"pausePrompt"`
	Interruptible                bool           `xml:"interruptible"`
	CanChangeInterruptableOption bool           `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool           `xml:"ttsEnumed"`
	ExitModuleOnException        bool           `xml:"exitModuleOnException"`
}
type xFilePrompt struct {
	PromptDirectly     bool   `xml:"promptData>promptSelected"`
	PromptID           int32  `xml:"promptData>prompt>id"`
	PromptName         string `xml:"promptData>prompt>name"`
	PromptVariableName string `xml:"promptData>promptVariableName"`
	IsRecordedMessage  bool   `xml:"promptData>isRecordedMessage"`
}
type xPausePrompt struct {
	Timeout int32 `xml:"timeout"`
}
type xTTSPrompt struct {
	TtsPromptXML    string `xml:"xml"`
	PromptTTSEnumed bool   `xml:"promptTTSEnumed"`
}
type xVivrPrompts struct {
	VPrompt                      []xVivrPrompt `xml:"vivrPrompt"`
	Interruptible                bool          `xml:"interruptible"`
	CanChangeInterruptableOption bool          `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool          `xml:"ttsEnumed"`
	ExitModuleOnException        bool          `xml:"exitModuleOnException"`
}
type xVivrHeader struct {
	VPrompt                      xVivrPrompt `xml:"vivrPrompt"`
	Interruptible                bool        `xml:"interruptible"`
	CanChangeInterruptableOption bool        `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool        `xml:"ttsEnumed"`
	ExitModuleOnException        bool        `xml:"exitModuleOnException"`
}
type xVivrPrompt struct {
	VivrXML string `xml:"xml"`
}
type xMultilingualPrompts struct {
	Entries []xMLPromptEntry `xml:"entry"`
}
type xMLPromptEntry struct {
	Key          string     `xml:"key"`
	ID           string     `xml:"value>promptId"`
	Name         string     `xml:"value>name"`
	Description  string     `xml:"value>description"`
	Type         string     `xml:"value>type"`
	Prompts      xMLPrompts `xml:"value>prompts"`
	DefLanguage  string     `xml:"defaultLanguage"`
	IsPersistent bool       `xml:"isPersistent"`
}
type xMLPrompts struct {
	Entries []xMLanguageEntry `xml:"entry"`
}
type xMLanguageEntry struct {
	Language string         `xml:"key,attr"`
	TTSes    []xTTSPrompt   `xml:"ttsPrompt"`
	Files    []xFilePrompt  `xml:"filePrompt"`
	Pauses   []xPausePrompt `xml:"pausePrompt"`
}
