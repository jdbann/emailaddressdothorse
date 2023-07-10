package micropub

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// Entry is the representation of a post in micropub's schema.
type Entry struct {
	Content       string
	ContentHTML   string
	Categories    []string
	Photo         string
	NestedObjects map[string]json.RawMessage
}

func entryFromFormValues(form url.Values) Entry {
	categories := form["category[]"]

	if category, ok := form["category"]; ok {
		categories = append(categories, category...)
	}

	return Entry{
		Content:    form.Get("content"),
		Categories: categories,
		Photo:      form.Get("photo"),
	}
}

type entryProperties struct {
	Content       []contentProperty
	Categories    []string
	Photos        []string
	NestedObjects map[string]json.RawMessage
}

func (p *entryProperties) UnmarshalJSON(b []byte) error {
	all := map[string]json.RawMessage{}
	if err := json.Unmarshal(b, &all); err != nil {
		return err
	}

	for k, v := range all {
		switch k {
		case "content":
			if err := json.Unmarshal(v, &p.Content); err != nil {
				return err
			}
		case "category":
			if err := json.Unmarshal(v, &p.Categories); err != nil {
				return err
			}
		case "photo":
			if err := json.Unmarshal(v, &p.Photos); err != nil {
				return err
			}
		default:
			if p.NestedObjects == nil {
				p.NestedObjects = make(map[string]json.RawMessage)
			}
			p.NestedObjects[k] = v
		}
	}

	return nil
}

type contentProperty struct {
	Plain string
	HTML  string
}

type htmlContentProperty struct {
	HTML string `json:"html"`
}

func (c *contentProperty) UnmarshalJSON(b []byte) error {
	if bytes.HasPrefix(b, []byte("{")) {
		prop := &htmlContentProperty{}
		if err := json.Unmarshal(b, prop); err != nil {
			return err
		}
		c.HTML = prop.HTML
	} else {
		var prop string
		if err := json.Unmarshal(b, &prop); err != nil {
			return err
		}
		c.Plain = prop
	}

	return nil
}

func entryFromJSONValues(props entryProperties) Entry {
	e := Entry{
		Content:       props.Content[0].Plain,
		ContentHTML:   props.Content[0].HTML,
		Categories:    props.Categories,
		NestedObjects: props.NestedObjects,
	}

	if len(props.Photos) > 0 {
		e.Photo = props.Photos[0]
	}

	return e
}
