/*
* @Author: CJ Ting
* @Date:   2016-07-08 17:28:53
* @Last Modified by:   CJ Ting
* @Last Modified time: 2016-07-18 17:47:03
 */

package alfred

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON_AddItem(t *testing.T) {
	response := Response{}
	item := Item{}
	response.AddItem(item)
	assert.Equal(t, 1, response.ItemLength(), "ItemLength should return items' length")
}

func TestJSON_ClearItems(t *testing.T) {
	response := Response{}
	item := Item{}
	response.AddItem(item)
	response.ClearItems()
	assert.Equal(t, 0, response.ItemLength(), "ClearItems should clear all items")
}

func TestJSON_Response(t *testing.T) {
	response := Response{}
	item := Item{}
	response.AddItem(item)

	_, err := response.JSON()
	if err != nil {
		t.Error(err)
	}
}

// if uid is present but empty, alfred will adjust the order
func TestJSON_UIDOmitEmpty(t *testing.T) {
	response := Response{}
	item := Item{}
	item.Title = "title"
	response.AddItem(item)
	str, _ := response.JSON()
	strings.Index(str, "uid")
	assert.Equal(t, -1, strings.Index(str, "uid"))
}

func TestJSON_SampleResponse(t *testing.T) {
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

	expectedResponse := `{
  "items": [
    {
      "uid": "uid",
      "type": "type",
      "title": "title",
      "subtitle": "subtitle",
      "arg": "arg",
      "autocomplete": "autocomplete",
      "icon": {
        "type": "type",
        "path": "path"
      },
      "valid": true,
      "quicklookurl": "quicklookurl",
      "mods": {
        "alt": {
          "valid": true,
          "arg": "arg",
          "subtitle": "subtitle"
        },
        "cmd": {
          "valid": true,
          "arg": "arg",
          "subtitle": "subtitle"
        },
        "ctrl": {
          "valid": true,
          "arg": "arg",
          "subtitle": "subtitle"
        },
        "fn": {
          "valid": true,
          "arg": "arg",
          "subtitle": "subtitle"
        },
        "shift": {
          "valid": true,
          "arg": "arg",
          "subtitle": "subtitle"
        }
      },
      "text": {
        "copy": "copy",
        "largetype": "largetype"
      }
    }
  ]
}`

	jsonStr, err := response.JSON()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expectedResponse, jsonStr)
}
