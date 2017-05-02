package web

type SkillTransfer struct {
	ModuleDescription
	Priority                   int32 `json:"priority"`
	CallbackAllowInternational bool  `json:"callbackAllowInternational"`
	IsTcpaConsentEnabled       bool  `json:"isTcpaConsentEnabled"`
	IsCallbackEnabled          bool  `json:"isCallbackEnabled"`
	IsChatEnabled              bool  `json:"isChatEnabled"`
	IsEmailEnabled             bool  `json:"isEmailEnabled"`
	IsVideoEnabled             bool  `json:"isVideoEnabled"`
	TCPAConsentText            bool  `json:"tcpaConsentText"`
}
