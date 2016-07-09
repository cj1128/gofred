/*
* @Author: CJ Ting
* @Date:   2016-07-05 11:15:57
* @Last Modified by:   CJ Ting
* @Last Modified time: 2016-07-09 09:59:35
 */

package alfred

import (
	"encoding/json"
	"encoding/xml"
)

type Response struct {
	XMLName xml.Name `json:"-" xml:"items"`
	Items   []*Item  `json:"items"`
}

type Item struct {
	XMLName      xml.Name       `json:"-" xml:"item"`
	Uid          string         `json:"uid,omitempty" xml:"uid,attr"`
	Type         string         `json:"type,omitempty" xml:"type,attr"`
	Title        string         `json:"title" xml:"title"`
	Subtitle     string         `json:"subtitle" xml:"subittle"`
	Arg          string         `json:"arg" xml:"arg,attr"`
	Autocomplete string         `json:"autocomplete" xml:"autocomplete,attr"`
	Icon         Icon           `json:"icon" xml:"icon"`
	Valid        bool           `json:"valid,omitempty" xml:"valid,attr"`
	Quicklookurl string         `json:"quicklookurl,omitempty" xml:"quicklookurl"`
	Mods         map[string]Mod `json:"mods,omitempty"`
	Text         Text           `json:"text"`
}

type Icon struct {
	Type string `json:"type" xml:"type,attr"`
	Path string `json:"path" xml:",chardata"`
}

type Mod struct {
	// Key      string `json:"-" xml:"key,attr"`
	Valid    bool   `json:"valid" xml:"valid,attr"`
	Arg      string `json:"arg" xml:"arg,attr"`
	Subtitle string `json:"subtitle" xml:"subtitle,attr"`
}

type Text struct {
	Copy      string `json:"copy"`
	Largetype string `json:"largetype"`
}

func (i *Item) AddMod(key string, item Mod) {
	if i.Mods == nil {
		i.Mods = make(map[string]Mod)
	}
	i.Mods[key] = item
}

var defaultResponse = Response{}

// add Item to response
func (r *Response) AddItem(item Item) {
	r.Items = append(r.Items, &item)
}

func (r *Response) ItemLength() int {
	return len(r.Items)
}

// clear all items
func (r *Response) ClearItems() {
	r.Items = nil
}

func (r *Response) JSON() (string, error) {
	bytes, err := json.Marshal(*r)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// top level functions
func AddItem(item Item) {
	defaultResponse.AddItem(item)
}

func ItemLength() int {
	return defaultResponse.ItemLength()
}

func ClearItems() {
	defaultResponse.ClearItems()
}

func JSON() (string, error) {
	return defaultResponse.JSON()
}

// func XML() (string, error) {
// 	response.Items = items
// 	bytes, err := xml.Marshal(response)
// 	if err != nil {
// 		return "", nil
// 	}
// 	return xml.Header + string(bytes), nil
// }
