package web

type GenericInputModule struct {
	TargetVariable string              `json:"targetVariable"`
	Restriction    GenericRestriction  `json:"restriction"`
	Descendant     *GenericInputModule `json:"descendant"`
}
type Input struct {
	ModuleDescription
	GenericInputModule
}
type GetDigits struct {
	ModuleDescription
	GenericInputModule
}
type VoiceInput struct {
	ModuleDescription
	GenericInputModule
}

///////////////////////////////////////
type GenericRestriction struct {
	Type     string `json:"type"`
	E164     bool   `json:"isE164"`
	Language string `json:"language"`
}
type Digits struct {
	GenericRestriction
	MaxLength int32 `json:"maxLength"`
	MinLength int32 `json:"minLength"`
}
type Number struct {
	GenericRestriction
	MaxAllowedDouble float64 `json:"maxAllowedDouble"`
	MinAllowedDouble float64 `json:"maxAllowedDouble"`
}
type Currency struct {
	GenericRestriction
	MaxAllowedDouble float64 `json:"maxAllowedDouble"`
	MinAllowedDouble float64 `json:"maxAllowedDouble"`
}
type IntNumber struct {
	GenericRestriction
	MaxAllowedInteger int32 `json:"maxAllowedInteger"`
	MinAllowedInteger int32 `json:"maxAllowedInteger"`
}
type Time struct {
	GenericRestriction
	MaxAllowedInteger int32 `json:"maxAllowedInteger"`
	MinAllowedInteger int32 `json:"maxAllowedInteger"`
}
type Date struct {
	GenericRestriction
	MaxAllowedInteger int32 `json:"maxAllowedInteger"`
	MinAllowedInteger int32 `json:"maxAllowedInteger"`
}
type CCExpDate struct {
	GenericRestriction
	ReferenceDate string `json:"referenceDate"`
}
