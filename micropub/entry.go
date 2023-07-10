package micropub

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// Entry is the representation of a post in micropub's schema.
type Entry struct {
	Content     string
	ContentHTML string
	Categories  []string
	Photo       string
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
	Content    []contentProperty `json:"content"`
	Categories []string          `json:"category"`
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
	return Entry{
		Content:     props.Content[0].Plain,
		ContentHTML: props.Content[0].HTML,
		Categories:  props.Categories,
	}
}
