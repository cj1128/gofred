/*
* @Author: CJ Ting
* @Date:   2016-07-09 10:22:00
* @Last Modified by:   CJ Ting
* @Last Modified time: 2016-07-18 17:47:01
 */

package alfred

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXML_Response(t *testing.T) {
	response := Response{}
	item := Item{}
	response.AddItem(item)
	_, err := response.XML()
	if err != nil {
		t.Error(err)
	}
}

// if uid is present but empty, alfred will adjust the order
func TestXML_UIDOmitEmpty(t *testing.T) {
	response := Response{}
	item := Item{}
	response.AddItem(item)
	str, _ := response.XML()
	assert.Equal(t, -1, strings.Index(str, "uid"))
}

func TestXML_SampleResponse(t *testing.T) {
	response := Response{}
	item := Item{
		Uid:          "uid",
		Type:         "type",
		Title:        "title",
		Subtitle:     "subtitle",
		Arg:          "arg",
		Autocomplete: "autocomplete",
		Icon: Icon{
			Type: "type",
			Path: "path",
		},
		Valid:        true,
		Quicklookurl: "quicklookurl",
		Mods: Mods{
			"cmd": {
				Valid:    true,
				Arg:      "arg",
				Subtitle: "subtitle",
			},
			"ctrl": {
				Valid:    true,
				Arg:      "arg",
				Subtitle: "subtitle",
			},
			"shift": {
				Valid:    true,
				Arg:      "arg",
				Subtitle: "subtitle",
			},
			"fn": {
				Valid:    true,
				Arg:      "arg",
				Subtitle: "subtitle",
			},
			"alt": {
				Valid:    true,
				Arg:      "arg",
				Subtitle: "subtitle",
			},
		},
		Text: Text{
			Copy:      "copy",
			Largetype: "largetype",
		},
	}
	response.AddItem(item)

	const expectedResponse = `<?xml version="1.0" encoding="UTF-8"?>
<items>
  <item uid="uid" type="type" arg="arg" autocomplete="autocomplete" valid="yes">
    <title>title</title>
    <subtitle>subtitle</subtitle>
    <icon type="type">path</icon>
    <quicklookurl>quicklookurl</quicklookurl>
    <mod key="cmd" subtitle="subtitle" valid="yes" arg="arg"></mod>
    <mod key="ctrl" subtitle="subtitle" valid="yes" arg="arg"></mod>
    <mod key="shift" subtitle="subtitle" valid="yes" arg="arg"></mod>
    <mod key="fn" subtitle="subtitle" valid="yes" arg="arg"></mod>
    <mod key="alt" subtitle="subtitle" valid="yes" arg="arg"></mod>
    <text type="copy">copy</text>
    <text type="largetype">largetype</text>
  </item>
</items>`

	xmlStr, _ := response.XML()
	assert.Equal(t, expectedResponse, xmlStr)
}
