/*
* @Author: CJ Ting
* @Date:   2016-07-08 17:28:53
* @Last Modified by:   CJ Ting
* @Last Modified time: 2016-07-08 19:48:48
 */

package alfred_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fate-lovely/go-alfred"
)

const sampleJSON = `
{"items": [
  {
    "uid": "desktop",
    "type": "file",
    "title": "Desktop",
    "subtitle": "~/Desktop",
    "arg": "~/Desktop",
    "autocomplete": "Desktop",
    "icon": {
      "type": "fileicon",
      "path": "~/Desktop"
    }
  },

  {
    "valid": false,
    "uid": "flickr",
    "title": "Flickr",
    "icon": {
      "path": "flickr.png"
    }
  },

  {
    "uid": "image",
    "type": "file",
    "title": "My holiday photo",
    "subtitle": "~/Pictures/My holiday photo.jpg",
    "autocomplete": "My holiday photo",
    "icon": {
      "type": "filetype",
      "path": "public.jpeg"
    }
  },

  {
    "valid": false,
    "uid": "alfredapp",
    "title": "Alfred Website",
    "subtitle": "https://www.alfredapp.com/",
    "arg": "alfredapp.com",
    "autocomplete": "Alfred Website",
    "quicklookurl": "https://www.alfredapp.com/",
    "mods": {
      "alt": {
        "valid": true,
        "arg": "alfredapp.com/powerpack",
        "subtitle": "https://www.alfredapp.com/powerpack/"
      },
      "cmd": {
        "valid": true,
        "arg": "alfredapp.com/powerpack/buy/",
        "subtitle": "https://www.alfredapp.com/powerpack/buy/"
      }
    },
    "text": {
      "copy": "https://www.alfredapp.com/ (text here to copy)",
      "largetype": "https://www.alfredapp.com/ (text here for large type)"
    }
  }
]}
`

// var fullItem = alfred.Item{
// 	Uid:          "uid",
// 	Type:         "type",
// 	Title:        "title",
// 	Subtitle:     "subtitle",
// 	Arg:          "arg",
// 	Autocomplete: "autocomplete",
// 	Icon:         alfred.Icon{},
// 	Valid:        true,
// 	Quicklookurl: "quicklookurl",
// 	Mods: map[string]alfred.Mod{
// 		"cmd": {
// 			Valid:    true,
// 			Arg:      "arg",
// 			Subtitle: "subtitle",
// 		},
// 		"ctrl": {
// 			Valid:    true,
// 			Arg:      "arg",
// 			Subtitle: "subtitle",
// 		},
// 		"shift": {
// 			Valid:    true,
// 			Arg:      "arg",
// 			Subtitle: "subtitle",
// 		},
// 		"fn": {
// 			Valid:    true,
// 			Arg:      "arg",
// 			Subtitle: "subtitle",
// 		},
// 		"alt": {
// 			Valid:    true,
// 			Arg:      "arg",
// 			Subtitle: "subtitle",
// 		},
// 	},
// 	Text: alfred.Text{
// 		Copy:      "copy",
// 		Largetype: "largetype",
// 	},
// }

func TestJSON_AddItem(t *testing.T) {
	item := alfred.Item{}
	alfred.AddItem(item)
	assert.Equal(t, 1, alfred.ItemLength(), "ItemLength should return items' length")
	alfred.ClearItems()
}

func TestJSON_ClearItems(t *testing.T) {
	item := alfred.Item{}
	alfred.AddItem(item)
	alfred.ClearItems()
	assert.Equal(t, 0, alfred.ItemLength(), "ClearItems should clear all items")
}

func TestJSON_JSONResponse(t *testing.T) {
	response := alfred.Response{}
	item := alfred.Item{}
	response.AddItem(item)

	_, err := response.JSON()
	if err != nil {
		t.Error(err)
	}
}

func TestJSON_SampleResponse(t *testing.T) {
	response := alfred.Response{}
	item := alfred.Item{}

	item.Uid = "desktoppp"
	item.Type = "file"
	item.Title = "Desktop"
	item.Subtitle = "~/Desktop"
	item.Arg = "~/Desktop"
	item.Autocomplete = "Desktop"
	item.Icon = alfred.Icon{
		Type: "fileicon",
		Path: "~/Desktop/",
	}

	response.AddItem(item)

	var sampleResponse alfred.Response
	err := json.Unmarshal([]byte(sampleJSON), &sampleResponse)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(response, sampleResponse) {
		t.Errorf("expected %v to eq %v", response, sampleResponse)
	}
}
