package IVRParser

type Prompts struct {
	TTSes                        []TTSPrompt `xml:"ttsPrompt"`
	Interruptible                bool        `xml:"interruptible"`
	CanChangeInterruptableOption bool        `xml:"canChangeInterruptableOption"`
	TtsEnumed                    bool        `xml:"ttsEnumed"`
	ExitModuleOnException        bool        `xml:"exitModuleOnException"`
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
