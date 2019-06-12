package ivr

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

func parseTtsPrompt(ttsxml io.Reader) (root *ttsNode, err error) {
	decoder := xml.NewDecoder(ttsxml)
	decoder.CharsetReader = charset.NewReaderLabel

	for {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}
		switch v := t.(type) {
		case xml.StartElement:
			return parseTtsNode(decoder, v.Name.Local, nil)
		}
	}
	return nil, errors.New("Incorrect TTS element")
}

func parseTtsNode(decoder *xml.Decoder, nodeType string, parent *ttsNode) (*ttsNode, error) {
	var (
		node         = &ttsNode{nodeType: nodeType, parent: parent}
		immersion    = 1
		inItems      = false
		inAttributes = false
		inBody       = false
		inVarName    = false
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "attributes" {
				inAttributes = true
			} else if v.Name.Local == "items" {
				inItems = true
			} else if v.Name.Local == "body" {
				inBody = true
			} else if v.Name.Local == "variableName" {
				inVarName = true
			} else if inItems && node.variableName == "" {
				if child, err := parseTtsNode(decoder, v.Name.Local, node); err == nil {
					node.children = append(node.children, child)
				}
				immersion--
			} else if inAttributes {
				if attr, err := parseTtsAttribute(decoder); err == nil {
					node.attributes = append(node.attributes, attr...)
				}
				immersion--
			}
		case xml.EndElement:
			immersion--
			if v.Name.Local == "items" {
				inItems = false
			} else if v.Name.Local == "attributes" {
				inAttributes = false
			} else if v.Name.Local == "body" {
				inBody = false
			} else if v.Name.Local == "variableName" {
				inVarName = false
			}

		case xml.CharData:
			if inVarName {
				node.variableName = string(v)
				node.children = nil // remove node - "test value"
			} else if inBody {
				node.body = string(v)
			}
		}
	}
	return node, nil
}

func parseTtsAttribute(decoder *xml.Decoder) (attributes []ttsAttr, err error) {
	var (
		immersion = 1
		inName    = false
		theAttr   ttsAttr
	)

	for immersion > 0 {
		t, err := decoder.Token()
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			immersion++
			if v.Name.Local == "name" {
				inName = true
			} else if v.Name.Local == "parentAttributeValue" {
				// just skip it
			} else {
				if len(v.Attr) > 0 {
					theAttr.value = v.Attr[0].Value
				} else {
					if childAttr, err := parseTtsAttribute(decoder); err == nil {
						attributes = append(attributes, childAttr...)
					}
					immersion--
				}
			}
		case xml.CharData:
			if inName {
				theAttr.name = string(v)
			}

		case xml.EndElement:
			immersion--
			if v.Name.Local == "name" {
				inName = false
			}
		}
	}
	return append(attributes, theAttr), nil
}
