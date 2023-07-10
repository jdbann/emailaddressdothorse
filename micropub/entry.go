package micropub

import "net/url"

// Entry is the representation of a post in micropub's schema.
type Entry struct {
	Content    string
	Categories []string
	Photo      string
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

func entryFromJSONValues(val *createRequest) Entry {
	return Entry{
		Content:    val.Properties.Content[0],
		Categories: val.Properties.Categories,
	}
}
