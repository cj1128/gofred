package gofred

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimalOutput(t *testing.T) {
	res := New()
	res.AddItem(&Item{
		Title:    "title",
		Subtitle: "subtitle",
	})

	json, err := res.JSON()
	assert.Nil(t, err)
	assert.Equal(t, `{
  "items": [
    {
      "title": "title",
      "subtitle": "subtitle"
    }
  ]
}`, json)
}

func TestFullOutput(t *testing.T) {
	res := New()
	res.Rerun = 1
	res.Variables = map[string]string{
		"key": "value",
	}
	res.AddItem(&Item{
		Uid:          "uid",
		Type:         "type",
		Title:        "title",
		Subtitle:     "subtitle",
		Arg:          "arg",
		Autocomplete: "autocomplete",
		Icon: &Icon{
			Type: "type",
			Path: "path",
		},
		Valid:        true,
		Quicklookurl: "quicklookurl",
		Mods: Mods{
			CmdKey: {
				Valid:    true,
				Arg:      "arg",
				Subtitle: "subtitle",
				Icon: &Icon{
					Type: "cmdtype",
					Path: "cmdpath",
				},
				Variables: map[string]string{
					"cmdkey": "cmdvalue",
				},
			},
			CtrlKey: {
				Valid:    false,
				Arg:      "arg",
				Subtitle: "subtitle",
			},
			ShiftKey: {
				Valid:    true,
				Arg:      "arg",
				Subtitle: "subtitle",
			},
			FnKey: {
				Valid:    true,
				Arg:      "arg",
				Subtitle: "subtitle",
			},
			AltKey: {
				Valid:    true,
				Arg:      "arg",
				Subtitle: "subtitle",
			},
		},
		Text: &Text{
			Copy:      "copy",
			Largetype: "largetype",
		},
	})

	json, err := res.JSON()
	assert.Nil(t, err)

	// NOTE: `mods` order is decided by runtime
	assert.Equal(t, `{
  "rerurn": 1,
  "variables": {
    "key": "value"
  },
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
          "subtitle": "subtitle",
          "icon": {
            "type": "cmdtype",
            "path": "cmdpath"
          },
          "variables": {
            "cmdkey": "cmdvalue"
          }
        },
        "ctrl": {
          "valid": false,
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
}`, json)
}
