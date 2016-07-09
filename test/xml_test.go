/*
* @Author: CJ Ting
* @Date:   2016-07-09 10:22:00
* @Last Modified by:   CJ Ting
* @Last Modified time: 2016-07-09 12:03:31
 */

package alfred_test

import (
	"testing"

	"github.com/fate-lovely/go-alfred"
	"github.com/stretchr/testify/assert"
)

func TestXML_Response(t *testing.T) {
	response := alfred.Response{}
	item := alfred.Item{}
	response.AddItem(item)
	_, err := response.XML()
	if err != nil {
		t.Error(err)
	}
}

func TestXML_SampleResponse(t *testing.T) {
	response := alfred.Response{}
	item := alfred.Item{
		Uid:          "uid",
		Type:         "type",
		Title:        "title",
		Subtitle:     "subtitle",
		Arg:          "arg",
		Autocomplete: "autocomplete",
		Icon: alfred.Icon{
			Type: "type",
			Path: "path",
		},
		Valid:        true,
		Quicklookurl: "quicklookurl",
		Mods: alfred.Mods{
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
		Text: alfred.Text{
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
