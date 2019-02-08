package dialogflow

import (
	"testing"

	ivr "github.com/Cepreu/gofrend/ivrparser"
)

func TestMenu(t *testing.T) {
	Menu := &ivr.MenuModule{
		GeneralInfo: ivr.GeneralInfo{
			ID:              "0B19BD7CEE3E4B85BA6C6631F1CCA222",
			Ascendants:      []ivr.ModuleID{"35EBEC8BD6294BA2B6CF4C96D54BE72D"},
			ExceptionalDesc: "C4BA62EF0F3F406A819ADE2CABD1669C",
			Name:            "Copy of Menu5",
			Dispo:           "Caller Disconnected",
			Collapsible:     false,
		},
		VoicePromptIDs: ivr.ModulePrompts{
			ivr.AttemptPrompts{
				LangPrArr: []ivr.LanguagePrompts{
					{
						PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_V__T_21"},
						Language: "Default",
					}, {
						PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_V__T_21"},
						Language: "en-US",
					}, {
						PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_V__T_21"},
						Language: "uk",
					},
				},
				Count: 1,
			},
			ivr.AttemptPrompts{
				LangPrArr: []ivr.LanguagePrompts{
					{
						PrArr:    []ivr.PromptID{"FCD392D339FB4A20B95105FAFF8AC7ED_en-US_T_36"},
						Language: "Default",
					}, {
						PrArr:    []ivr.PromptID{"FCD392D339FB4A20B95105FAFF8AC7ED_en-US_T_36"},
						Language: "en-US",
					}, {
						PrArr:    []ivr.PromptID{},
						Language: "uk",
					},
				},
				Count: 2,
			},
		},
		Branches: []*ivr.OutputBranch{
			{
				Key: "No Match",
				Value: struct {
					Name       string
					Descendant ivr.ModuleID
				}{"No Match", "C4BA62EF0F3F406A819ADE2CABD1669C"},
			}, {
				Key: "apples",
				Value: struct {
					Name       string
					Descendant ivr.ModuleID
				}{"apples", "C4BA62EF0F3F406A819ADE2CABD1669C"},
			}, {
				Key: "plums",
				Value: struct {
					Name       string
					Descendant ivr.ModuleID
				}{"plums", "C4BA62EF0F3F406A819ADE2CABD1669C"},
			}, {
				Key: "appricots",
				Value: struct {
					Name       string
					Descendant ivr.ModuleID
				}{"appricots", "C4BA62EF0F3F406A819ADE2CABD1669C"},
			}, {
				Key: "tutti frutti",
				Value: struct {
					Name       string
					Descendant ivr.ModuleID
				}{"tutti frutti", "C4BA62EF0F3F406A819ADE2CABD1669C"},
			},
		},
		Items: []*ivr.MenuItem{
			{
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_24"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_24"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_24"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "apples"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_25"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_25"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_25"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "appricots"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_26"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_26"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_26"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "plums"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"FFBEE18B77834E10B034889F5457DEB4_en-US_A"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"FFBEE18B77834E10B034889F5457DEB4_en-US_A"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"FFBEE18B77834E10B034889F5457DEB4_uk_A"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "tutti frutti"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"7D0A23161B1B47CEA640C61D490D6FEF_en-US_A"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"7D0A23161B1B47CEA640C61D490D6FEF_en-US_A"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"7D0A23161B1B47CEA640C61D490D6FEF_uk_A"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "appricots"},
			}, {
				Prompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"3EF9450C730E462FB97E9ADA7A0E37AE_en-US_A"},
							Language: "Default",
						}, {
							PrArr:    []ivr.PromptID{"3EF9450C730E462FB97E9ADA7A0E37AE_en-US_A"},
							Language: "en-US",
						}, {
							PrArr:    []ivr.PromptID{"3EF9450C730E462FB97E9ADA7A0E37AE_uk_A"},
							Language: "uk",
						},
					},
					Count: 1,
				},
				ShowInVivr: true,
				MatchExact: false,
				Dtmf:       "DTMF_AUTO",
				Action: struct {
					Type ivr.ActionType
					Name string
				}{"BRANCH", "plums"},
			},
		},
		UseASR:  true,
		UseDTMF: true,
		Events: []*ivr.RecoEvent{
			{
				Event:  "NO_MATCH",
				Action: "CONTINUE",
				CountAndPrompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222__RE__T_19"},
							Language: "Default",
						},
					},
					Count: 1,
				},
			}, {
				Event:  "NO_INPUT",
				Action: "REPROMPT",
				CountAndPrompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222__RE__T_20"},
							Language: "Default",
						},
					},
					Count: 1,
				},
			}, {
				Event:  "NO_INPUT",
				Action: "REPROMPT",
				CountAndPrompt: ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"DAA16853152C4EBFA4DEEDAB00227C2A"},
							Language: "Default",
						},
					},
					Count: 2,
				},
			},
		},
		ConfData: &ivr.ConfirmData{
			ConfirmRequired:      "REQUIRED",
			RequiredConfidence:   75,
			MaxAttemptsToConfirm: 3,
			NoInputTimeout:       3,
			VoicePromptIDs: ivr.ModulePrompts{
				ivr.AttemptPrompts{
					LangPrArr: []ivr.LanguagePrompts{
						{
							PrArr:    []ivr.PromptID{"87F5501915CA4ACC874D8767B8C4369E"},
							Language: "Default",
						},
					},
					Count: 1,
				},
			},
			Events: []*ivr.RecoEvent{
				{
					Event:  "NO_MATCH",
					Action: "REPROMPT",
					CountAndPrompt: ivr.AttemptPrompts{
						LangPrArr: []ivr.LanguagePrompts{
							{
								PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_CD___RE__T_22"},
								Language: "Default",
							},
						},
						Count: 1,
					},
				}, {
					Event:  "NO_INPUT",
					Action: "REPROMPT",
					CountAndPrompt: ivr.AttemptPrompts{
						LangPrArr: []ivr.LanguagePrompts{
							{
								PrArr:    []ivr.PromptID{"0B19BD7CEE3E4B85BA6C6631F1CCA222_CD___RE__T_23"},
								Language: "Default",
							},
						},
						Count: 1,
					},
				}, {
					Event:  "HELP",
					Action: "REPROMPT",
					CountAndPrompt: ivr.AttemptPrompts{
						LangPrArr: []ivr.LanguagePrompts{
							{
								PrArr:    []ivr.PromptID{},
								Language: "Default",
							},
						},
						Count: 1,
					},
				},
			},
		},
		RecoParams: struct {
			SpeechCompleteTimeout int
			MaxTimeToEnter        int
			NoInputTimeout        int
		}{1, 5, 5},
	}

	Prompts := ivr.ScriptPrompts{
		"073DF289D7EA415BB6A09B115F438229_en-US_T_30": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eIncorrect input.\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_24": ivr.TtsPrompt{
			TTSPromptXML: "\n\u0026lt;?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n\t\u003cattributes\u003e\n\t\t\u003clangAttr\u003e\n\t\t\t\u003cname\u003exml:lang\u003c/name\u003e\n\t\t\t\u003cattributeValueBase value=\"en-US\"/\u003e\n\t\t\u003c/langAttr\u003e\n\t\u003c/attributes\u003e\n\t\u003citems\u003e\n\t\t\u003csayAsElement\u003e\n\t\t\t\u003cattributes/\u003e\n\t\t\t\u003citems\u003e\n\t\t\t\t\u003ctextElement\u003e\n\t\t\t\t\t\u003cattributes/\u003e\n\t\t\t\t\t\u003citems/\u003e\n\t\t\t\t\t\u003cbody\u003eapple\u003c/body\u003e\n\t\t\t\t\u003c/textElement\u003e\n\t\t\t\u003c/items\u003e\n\t\t\u003c/sayAsElement\u003e\n\t\u003c/items\u003e\n\u003c/speakElement\u003e",
		},
		"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_25": ivr.TtsPrompt{
			TTSPromptXML: "\u0026lt;?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n          \u003cattributeValueBase value=\"en-US\"/\u003e\n       \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n     \t\t\t\u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n            \u003cvariableName\u003e__BUFFER__\u003c/variableName\u003e\n        \u003c/variableElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e",
		},
		"0B19BD7CEE3E4B85BA6C6631F1CCA222_A_26": ivr.TtsPrompt{
			TTSPromptXML: "\u0026lt;?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n          \u003cattributeValueBase value=\"en-US\"/\u003e\n       \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n     \t\t\t\u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n            \u003cvariableName\u003e35EBEC8BD6294BA2B6CF4C96D54BE72D:__BUFFER__\u003c/variableName\u003e\n        \u003c/variableElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e",
		},
		"0B19BD7CEE3E4B85BA6C6631F1CCA222_CD___RE__T_22": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eIncorrect input\u003c/body\u003e\n                \u003c/textElement\u003e\n           \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"0B19BD7CEE3E4B85BA6C6631F1CCA222_CD___RE__T_23": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eI cannot hear you\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"0B19BD7CEE3E4B85BA6C6631F1CCA222_V__T_21": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003ePlease name your favorite fruit.\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"0B19BD7CEE3E4B85BA6C6631F1CCA222__RE__T_19": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eIncorrect input\u003c/body\u003e\n                \u003c/textElement\u003e\n           \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"0B19BD7CEE3E4B85BA6C6631F1CCA222__RE__T_20": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eIncorrect input\u003c/body\u003e\n                \u003c/textElement\u003e\n           \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"2008C4FD8BCB4A9AA2F58E7C2839B01C_en-US_T_31": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eYou entered: \u003c/body\u003e\n                \u003c/textElement\u003e\n         \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems/\u003e\n            \u003cvariableName\u003e__SWI_LITERAL__\u003c/variableName\u003e\n        \u003c/variableElement\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n    \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n              \u003citems/\u003e\n                    \u003cbody\u003e. Press \"1\" if this is correct. Otherwise press \"2\"\u003c/body\u003e\n\u003c/textElement\u003e\n            \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"2008C4FD8BCB4A9AA2F58E7C2839B01C_uk_T_32": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"es-MX\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eТи мав на думцi \u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e\u003c/body\u003e\n                \u003c/textElement\u003e\n       \u003c/items\u003e\n            \u003cvariableName\u003e__SWI_LITERAL__\u003c/variableName\u003e\n        \u003c/variableElement\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e?\u003c/body\u003e\n                \u003c/textElement\u003e\n \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"3EF9450C730E462FB97E9ADA7A0E37AE_en-US_A": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n            \u003cvariableName\u003eQuery3:__BUFFER__\u003c/variableName\u003e\n        \u003c/variableElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"3EF9450C730E462FB97E9ADA7A0E37AE_uk_A": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"es-MX\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n            \u003cvariableName\u003eQuery3:__BUFFER__\u003c/variableName\u003e\n        \u003c/variableElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"7D0A23161B1B47CEA640C61D490D6FEF_en-US_A": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n            \u003cvariableName\u003e__BUFFER__\u003c/variableName\u003e\n        \u003c/variableElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"7D0A23161B1B47CEA640C61D490D6FEF_en-US_V": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003croot\u003e\n    \u003cattributes/\u003e\n    \u003citems\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems/\u003e\n            \u003cvariable\u003e__BUFFER__\u003c/variable\u003e\n        \u003c/variableElement\u003e\n    \u003c/items\u003e\n\u003c/root\u003e\n",
		},
		"7D0A23161B1B47CEA640C61D490D6FEF_uk_A": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"es-MX\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n            \u003cvariableName\u003e__BUFFER__\u003c/variableName\u003e\n        \u003c/variableElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"7D0A23161B1B47CEA640C61D490D6FEF_uk_V": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003croot\u003e\n    \u003cattributes/\u003e\n    \u003citems\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems/\u003e\n            \u003cvariable\u003e__BUFFER__\u003c/variable\u003e\n        \u003c/variableElement\u003e\n    \u003c/items\u003e\n\u003c/root\u003e\n",
		},
		"87F5501915CA4ACC874D8767B8C4369E_en-US_T_33": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eYou entered: \u003c/body\u003e\n                \u003c/textElement\u003e\n         \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n        \u003cvariableElement\u003e\n            \u003cattributes/\u003e\n            \u003citems/\u003e\n            \u003cvariableName\u003e__SWI_LITERAL__\u003c/variableName\u003e\n        \u003c/variableElement\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n    \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n              \u003citems/\u003e\n                    \u003cbody\u003e. Say \"Yes\" or press \"1\" if this is correct. Otherwise say \"No\" or press \"2\"\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"87F5501915CA4ACC874D8767B8C4369E_uk_T_34": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"es-MX\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e Га? Що ти сказав?\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"DAA16853152C4EBFA4DEEDAB00227C2A_en-US_T_35": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eI cannot hear you.\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"FCD392D339FB4A20B95105FAFF8AC7ED_en-US_T_36": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eTo be or not to be\u003c/body\u003e\n                \u003c/textElement\u003e\n            \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"FFBEE18B77834E10B034889F5457DEB4_en-US_A": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"en-US\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003eCherry\u003c/body\u003e\n                \u003c/textElement\u003e\n  \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
		"FFBEE18B77834E10B034889F5457DEB4_uk_A": ivr.TtsPrompt{
			TTSPromptXML: "\u003c?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?\u003e\n\u003cspeakElement\u003e\n    \u003cattributes\u003e\n        \u003clangAttr\u003e\n            \u003cname\u003exml:lang\u003c/name\u003e\n       \u003cattributeValueBase value=\"es-MX\"/\u003e\n        \u003c/langAttr\u003e\n    \u003c/attributes\u003e\n    \u003citems\u003e\n        \u003csayAsElement\u003e\n            \u003cattributes/\u003e\n            \u003citems\u003e\n                \u003ctextElement\u003e\n                    \u003cattributes/\u003e\n                    \u003citems/\u003e\n                    \u003cbody\u003e  Черешнi\u003c/body\u003e\n                \u003c/textElement\u003e\n     \u003c/items\u003e\n        \u003c/sayAsElement\u003e\n    \u003c/items\u003e\n\u003c/speakElement\u003e\n",
		},
	}
	intent, _ := menu2intents(Menu, Prompts)

	t.Errorf("\nMenuToIntents conversation: \n%v \nwas expected, in reality: \n%v", Menu, intent)
}
