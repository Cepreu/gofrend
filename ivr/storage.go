package ivr

import (
	"encoding/json"
	"fmt"
	"log"
)

type StorageScript struct {
	Domain             string
	Properties         string
	ModuleKeys         []ModuleID
	ModuleVals         []*StorageModule
	ModuleOnHangupKeys []ModuleID
	ModuleOnHangupVals []*StorageModule
	PromptKeys         []PromptID
	PromptVals         []*StoragePrompt
	//MLChoices          []*MultilanguageMenuChoice
	VariableKeys []VariableID
	VariableVals []*Variable
	Input        []VariableID
	Output       []VariableID
	Languages    []Language
	Functions    []*Function
	Menus        []ModuleID
}

type StorageModuleType int

const (
	cHangupModule StorageModuleType = iota
	cIfElseModule
	cCaseModule
	cIncomingCallModule
	cInputModule
	cMenuModule
	cPlayModule
	cQueryModule
	cSetVariableModule
	cSkillTransferModule
)

type StorageModule struct {
	Type          StorageModuleType
	Hangup        *HangupModule
	IfElse        *IfElseModule
	Case          *CaseModule
	IncomingCall  *IncomingCallModule
	Input         *InputModule
	Menu          *MenuModule
	Play          *PlayModule
	Query         *QueryModule
	SetVariable   *SetVariableModule
	SkillTransfer *SkillTransferModule
}

func (module *StorageModule) GetModule() Module {
	switch module.Type {
	case cHangupModule:
		return module.Hangup
	case cIfElseModule:
		return module.IfElse
	case cCaseModule:
		return module.Case
	case cIncomingCallModule:
		return module.IncomingCall
	case cInputModule:
		return module.Input
	case cMenuModule:
		return module.Menu
	case cPlayModule:
		return module.Play
	case cQueryModule:
		return module.Query
	case cSetVariableModule:
		return module.SetVariable
	case cSkillTransferModule:
		return module.SkillTransfer
	default:
		return nil
	}
}

func MakeStorageModule(module Module) *StorageModule {
	ret := &StorageModule{}
	switch v := module.(type) {
	case *HangupModule:
		ret.Type = cHangupModule
		ret.Hangup = v
	case *IfElseModule:
		ret.Type = cIfElseModule
		ret.IfElse = v
	case *CaseModule:
		ret.Type = cCaseModule
		ret.Case = v
	case *IncomingCallModule:
		ret.Type = cIncomingCallModule
		ret.IncomingCall = v
	case *InputModule:
		ret.Type = cInputModule
		ret.Input = v
	case *MenuModule:
		ret.Type = cMenuModule
		ret.Menu = v
	case *PlayModule:
		ret.Type = cPlayModule
		ret.Play = v
	case *QueryModule:
		ret.Type = cQueryModule
		ret.Query = v
	case *SetVariableModule:
		ret.Type = cSetVariableModule
		ret.SetVariable = v
	case *SkillTransferModule:
		ret.Type = cSkillTransferModule
		ret.SkillTransfer = v
	default:
		log.Panicf("Unsupported module type: %T", v)
	}
	return ret
}

type StoragePromptType int

const (
	cTtsPrompt StoragePromptType = iota
	cXFilePrompt
	cXPausePrompt
)

type StoragePrompt struct {
	Type         StoragePromptType
	TtsPrompt    *TtsPrompt
	XFilePrompt  *XFilePrompt
	XPausePrompt *XPausePrompt
}

func (prompt *StoragePrompt) GetPrompt() prompt {
	switch prompt.Type {
	case cTtsPrompt:
		return prompt.TtsPrompt
	case cXFilePrompt:
		return prompt.XFilePrompt
	case cXPausePrompt:
		return prompt.XPausePrompt
	default:
		return nil
	}
}

func MakeStoragePrompt(p prompt) *StoragePrompt {
	ret := &StoragePrompt{}
	switch v := p.(type) {
	case *TtsPrompt:
		ret.Type = cTtsPrompt
		ret.TtsPrompt = v
	case *XFilePrompt:
		ret.Type = cXFilePrompt
		ret.XFilePrompt = v
	case *XPausePrompt:
		ret.Type = cXPausePrompt
		ret.XPausePrompt = v
	default:
		return nil
	}
	return ret
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func MakeStorageScript(script *IVRScript) *StorageScript {
	storageModuleKeys := make([]ModuleID, 0)
	storageModuleVals := make([]*StorageModule, 0)
	for id, module := range script.Modules {
		storageModuleKeys = append(storageModuleKeys, id)
		storageModuleVals = append(storageModuleVals, MakeStorageModule(module))
	}
	storageModuleOnHangupKeys := make([]ModuleID, 0)
	storageModuleOnHangupVals := make([]*StorageModule, 0)
	for id, module := range script.ModulesOnHangup {
		storageModuleOnHangupKeys = append(storageModuleOnHangupKeys, id)
		storageModuleOnHangupVals = append(storageModuleOnHangupVals, MakeStorageModule(module))
	}
	storagePromptKeys := make([]PromptID, 0)
	storagePromptVals := make([]*StoragePrompt, 0)
	for id, p := range script.Prompts {
		storagePromptKeys = append(storagePromptKeys, id)
		storagePromptVals = append(storagePromptVals, MakeStoragePrompt(p))
	}
	variableKeys := make([]VariableID, 0)
	variableVals := make([]*Variable, 0)
	for id, variable := range script.Variables {
		variableKeys = append(variableKeys, id)
		variableVals = append(variableVals, variable)
	}
	return &StorageScript{
		Domain:             script.Domain,
		Properties:         script.Properties,
		ModuleKeys:         storageModuleKeys,
		ModuleVals:         storageModuleVals,
		ModuleOnHangupKeys: storageModuleOnHangupKeys,
		ModuleOnHangupVals: storageModuleOnHangupVals,
		PromptKeys:         storagePromptKeys,
		PromptVals:         storagePromptVals,
		//MLChoices:          script.MLChoices,
		VariableKeys: variableKeys,
		VariableVals: variableVals,
		Input:        script.Input,
		Output:       script.Output,
		Languages:    script.Languages,
		Functions:    script.Functions,
		Menus:        script.Menus,
	}
}

func (storageScript *StorageScript) GetScript() *IVRScript {
	modules := make(map[ModuleID]Module)
	for i := range storageScript.ModuleKeys {
		modules[storageScript.ModuleKeys[i]] = storageScript.ModuleVals[i].GetModule()
	}
	modulesOnHangup := make(map[ModuleID]Module)
	for i := range storageScript.ModuleOnHangupKeys {
		modulesOnHangup[storageScript.ModuleOnHangupKeys[i]] = storageScript.ModuleOnHangupVals[i].GetModule()
	}
	scriptPrompts := make(map[PromptID]prompt)
	for i := range storageScript.PromptKeys {
		scriptPrompts[storageScript.PromptKeys[i]] = storageScript.PromptVals[i].GetPrompt()
	}
	variables := make(map[VariableID]*Variable)
	for i := range storageScript.VariableKeys {
		variables[storageScript.VariableKeys[i]] = storageScript.VariableVals[i]
	}
	return &IVRScript{
		Domain:          storageScript.Domain,
		Properties:      storageScript.Properties,
		Modules:         modules,
		ModulesOnHangup: modulesOnHangup,
		Prompts:         scriptPrompts,
		//MLChoices:       storageScript.MLChoices,
		Variables: variables,
		Input:     storageScript.Input,
		Output:    storageScript.Output,
		Languages: storageScript.Languages,
		Functions: storageScript.Functions,
		Menus:     storageScript.Menus,
	}
}
