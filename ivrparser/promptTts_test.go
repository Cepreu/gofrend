package ivrparser

import (
	"fmt"
	"strings"
	"testing"
)

var expected = &ttsNode{
	nodeType: "speakElement",
	attributes: []ttsAttr{
		{name: "xml:lang", value: "en-US"},
		{name: "xml:lang", value: "en-US"},
	},
	children: []*ttsNode{
		&ttsNode{
			nodeType: "sayAsElement",
			children: []*ttsNode{
				&ttsNode{
					nodeType: "textElement",
					body:     "Please name your favorite fruit.",
				},
			},
		},
		&ttsNode{
			nodeType: "paragraphElement",
			children: []*ttsNode{
				&ttsNode{
					nodeType: "sayAsElement",
					children: []*ttsNode{
						&ttsNode{
							nodeType: "textElement",
							body:     "It can be:",
						},
					},
				},
				&ttsNode{
					nodeType:   "breakElement",
					attributes: []ttsAttr{{name: "strength", value: "medium"}},
				},
				&ttsNode{
					nodeType:   "emphasisElement",
					attributes: []ttsAttr{{name: "level", value: "strong"}},
					children: []*ttsNode{
						&ttsNode{
							nodeType: "sayAsElement",
							children: []*ttsNode{
								&ttsNode{
									nodeType: "textElement",
									body:     "cucumber",
								},
							},
						},
					},
				},
				&ttsNode{
					nodeType: "prosodyElement",
					attributes: []ttsAttr{
						{name: "rate", value: "slow"},
						{name: "volume", value: "loud"},
					},
					children: []*ttsNode{
						&ttsNode{
							nodeType: "sayAsElement",
							attributes: []ttsAttr{
								{name: "interpret-as", value: "name"},
							},
							children: []*ttsNode{
								&ttsNode{
									nodeType: "textElement",
									body:     "apple",
								},
							},
						},
						&ttsNode{
							nodeType: "sayAsElement",
							children: []*ttsNode{
								&ttsNode{
									nodeType: "textElement",
									body:     "persimone",
								},
							},
						},
					},
				},
				&ttsNode{
					nodeType: "voiceElement",
					attributes: []ttsAttr{
						{name: "xml:lang", value: "en-US"},
						{name: "gender", value: "male"},
						{name: "name", value: "Tom"},
						{name: "variant", value: "1"},
					},
					children: []*ttsNode{
						&ttsNode{
							nodeType: "sayAsElement",
							children: []*ttsNode{
								&ttsNode{
									nodeType: "textElement",
									body:     "appricot",
								},
							},
						},
						&ttsNode{
							nodeType: "sayAsElement",
							children: []*ttsNode{
								&ttsNode{
									nodeType: "textElement",
									body:     "plum",
								},
							},
						},
						&ttsNode{
							nodeType: "breakElement",
							attributes: []ttsAttr{{
								name:  "strength",
								value: "medium"},
							},
						},
					},
				},
				&ttsNode{
					nodeType: "sayAsElement",
					children: []*ttsNode{
						&ttsNode{
							nodeType: "textElement",
							body:     "or even",
						},
					},
				},
				&ttsNode{
					nodeType: "variableElement",
					attributes: []ttsAttr{
						{name: "format", value: "default"},
						{name: "interpret-as", value: "duration"},
					},
					variableName: "__TIME__",
				},
				&ttsNode{
					nodeType: "sayAsElement",
					children: []*ttsNode{
						&ttsNode{
							nodeType: "textElement",
							body:     "of",
						},
					},
				},
				&ttsNode{
					nodeType: "variableElement",
					attributes: []ttsAttr{
						{name: "format", value: "dm"},
						{name: "interpret-as", value: "date"},
					},
					variableName: "__DATE__",
				},
			},
		},
		&ttsNode{
			nodeType: "sentenceElement",
			attributes: []ttsAttr{
				{name: "xml:lang", value: "en-US"},
			},
			children: []*ttsNode{
				&ttsNode{
					nodeType: "sayAsElement",
					children: []*ttsNode{
						&ttsNode{
							nodeType: "textElement",
							body:     "Enjoy, my friend!",
						},
					},
				},
				&ttsNode{
					nodeType: "variableElement",
					attributes: []ttsAttr{
						{name: "interpret-as", value: "ordinal"},
					},
					variableName: "MyIntVar",
				},
			},
		},
	},
}

func TestTtsPromptParsing(t *testing.T) {
	var xmlData = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<speakElement>
	<attributes>
		<langAttr>
			<name>xml:lang</name>
			<attributeValueBase value="en-US"/>
		</langAttr>
		<langAttr>
			<name>xml:lang</name>
			<attributeValueBase value="en-US"/>
		</langAttr>
	</attributes>	
	<items>		
		<sayAsElement>
			<attributes/>
			<items>
			<textElement>
				<attributes/>
				<items/>
				<body>Please name your favorite fruit.</body>
			</textElement>
			</items>
		</sayAsElement>		
		<paragraphElement>
			<attributes/>
			<items>
				<sayAsElement>
					<attributes/>
					<items>
						<textElement>
							<attributes/>
							<items/>
							<body>It can be:</body>
						</textElement>
					</items>
				</sayAsElement>
				
				<breakElement>
					<attributes>
						<strengthAttr>
							<name>strength</name>
							<strengthAttrValue value="medium"/>
						</strengthAttr>
					</attributes>
					<items/>
				</breakElement>
				
				<emphasisElement>
					<attributes>
						<levelAttr>
							<name>level</name>
							<levelAttributeValue value="strong"/>
						</levelAttr>
					</attributes>
					<items>
						<sayAsElement>
							<attributes/>
							<items>
								<textElement>
									<attributes/>
									<items/>
									<body>cucumber</body>
								</textElement>
							</items>
						</sayAsElement>
					</items>
				</emphasisElement>
				
				<prosodyElement>
					<attributes>
						<rateAttr>
							<name>rate</name>
							<rateAttributeValue value="slow"/>
						</rateAttr>
						<volumeAttr>
							<name>volume</name>
							<volumeAttributeValue value="loud"/>
						</volumeAttr>
					</attributes>
					<items>
						<sayAsElement>
							<attributes>
							<interpretAsAttr>
								<name>interpret-as</name>
								<sayAsAttributeValue value="name"/>
							</interpretAsAttr>
							</attributes>
							<items>
								<textElement>
									<attributes/>
									<items/>
									<body>apple</body>
								</textElement>
							</items>
						</sayAsElement>
						<sayAsElement>
							<attributes/>
							<items>
								<textElement>
									<attributes/>
									<items/>
									<body>persimone</body>
								</textElement>
							</items>
						</sayAsElement>
					</items>
				</prosodyElement>
				
				<voiceElement>
					<attributes>
						<langAttr>
							<name>xml:lang</name>
							<attributeValueBase value="en-US"/>
						</langAttr>
						<genderAttr>
							<name>gender</name>
							<genderAttributeValue value="male"/>
						</genderAttr>
						<voiceNameAttr>
							<name>name</name>
							<attributeValueBase value="Tom"/>
						</voiceNameAttr>
						<variantAttr>
							<name>variant</name>
							<attributeValueBase value="1"/>
						</variantAttr>
					</attributes>
					<items>
						<sayAsElement>
							<attributes/>
							<items>
								<textElement>
									<attributes/>
									<items/>
									<body>appricot</body>
								</textElement>
							</items>
						</sayAsElement>
						<sayAsElement>
							<attributes/>
							<items>
								<textElement>
									<attributes/>
									<items/>
									<body>plum</body>
								</textElement>
							</items>
						</sayAsElement>
						<breakElement>
							<attributes>
								<strengthAttr>
									<name>strength</name>
									<strengthAttrValue value="medium"/>
								</strengthAttr>
							</attributes>
							<items/>
						</breakElement>
					</items>
				</voiceElement>
				
				<sayAsElement>
					<attributes/>
					<items>
						<textElement>
							<attributes/>
							<items/>
							<body>or even</body>
						</textElement>
					</items>
				</sayAsElement>

				<variableElement>
					<attributes>
						<interpretAsAttr>
							<name>interpret-as</name>
							<sayAsAttributeValue value="duration"/>
							<formatAsAttr>
								<name>format</name>
								<sayAsAttributeValue value="default"/>
								<parentAttributeValue>duration</parentAttributeValue>
							</formatAsAttr>
						</interpretAsAttr>
					</attributes>
					<items>
						<textElement>
							<attributes/>
							<items/>
							<body>12:30AM</body>
						</textElement>
					</items>
					<variableName>__TIME__</variableName>
				</variableElement>

				<sayAsElement>
					<attributes/>
					<items>
						<textElement>
							<attributes/>
							<items/>
							<body>of</body>
						</textElement>
					</items>
				</sayAsElement>

				<variableElement>
					<attributes>
						<interpretAsAttr>
							<name>interpret-as</name>
							<sayAsAttributeValue value="date"/>
							<formatAsAttr>
								<name>format</name>
								<sayAsAttributeValue value="dm"/>
								<parentAttributeValue>date</parentAttributeValue>
							</formatAsAttr>
						</interpretAsAttr>
					</attributes>
					<items>
						<textElement>
							<attributes/>
							<items/>
							<body>08/02</body>
						</textElement>
					</items>
					<variableName>__DATE__</variableName>
				</variableElement>
			
			</items>			
		</paragraphElement>
	
		<sentenceElement>
			<attributes>
				<langAttr>
					<name>xml:lang</name>
					<attributeValueBase value="en-US"/>
				</langAttr>
			</attributes>
			<items>
				<sayAsElement>
					<attributes/>
					<items>
						<textElement>
							<attributes/>
							<items/>
							<body>Enjoy, my friend!</body>
						</textElement>
					</items>
				</sayAsElement>
				<variableElement>
					<attributes>
						<interpretAsAttr>
							<name>interpret-as</name>
							<sayAsAttributeValue value="ordinal"/>
						</interpretAsAttr>
					</attributes>
					<items>
						<textElement>
							<attributes/>
							<items/>
							<body>2</body>
						</textElement>
					</items>
					<variableName>MyIntVar</variableName>
				</variableElement>
			</items>
		</sentenceElement>			
	</items>
</speakElement>
`

	res, err := parseTtsPrompt(strings.NewReader(xmlData))
	if err != nil {
		t.Fatal("Tts prompt wasn't parsed...")
	}

	if expected.walk(0) != res.walk(0) {
		//		if false == reflect.DeepEqual(expected, res) {
		t.Errorf("\nTTS Prompt: \nExpected:\n\n%s \n\nIn reality: \n%s", expected.walk(0), res.walk(0))
	}
}

// walk is a recursive function that calls itself for every child
func (n *ttsNode) walk(depth int) string {
	s := n.visit(depth)
	depth++
	for _, child := range n.children {
		s = s + child.walk(depth)
	}
	return s
}

// visit will get called on every node in the tree.
func (n *ttsNode) visit(depth int) string {
	d := ""
	for i := 0; i <= depth; i++ {
		d = d + "`~~"
	}
	return fmt.Sprintf("%s  \"%s\", attributes %v, variable: %s,body: %s\n", d, n.nodeType, n.attributes, n.variableName, n.body)
}

// func TestTtsPromptToAI(t *testing.T) {
// 	if s := expected.transformToAI(); s != "" {
// 		t.Fatal(s)
// 	}
// }
