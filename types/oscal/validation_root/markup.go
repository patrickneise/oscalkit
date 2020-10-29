package validation_root

import (
	"encoding/json"
	"regexp"
	"strings"
)

// Markup ...
type Markup struct {
	Raw       string `xml:",innerxml" yaml:"raw,omitempty"`
	PlainText string `xml:"-"`
}

// Short-hand method to build "markup-line" text
func ML(plain string) *Markup {
	plain = strings.ReplaceAll(plain, "&", "&amp;")
	plain = strings.ReplaceAll(plain, "<", "&lt;")
	plain = strings.ReplaceAll(plain, "<", "&gt;")
	return &Markup{
		Raw: plain,
	}
}

// Short-hand method to build "markup-multiline" text
func MML(plain string) *Markup {
	plain = strings.ReplaceAll(plain, "&", "&amp;")
	plain = strings.ReplaceAll(plain, "<", "&lt;")
	plain = strings.ReplaceAll(plain, "<", "&gt;")
	return &Markup{
		Raw: "<p>" + plain + "</p>",
	}
}

// Deprecated: Use validation_root.MML instead.
func MarkupFromPlain(plain string) *Markup {
	return MML(plain)
}

// PlainText representation
func (m *Markup) PlainString() string {
	if m.PlainText != "" {
		return m.PlainText
	}
	return m.xmlToPlain()
}

func (m *Markup) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &m.PlainText); err != nil {
		return err
	}
	return nil
}

func (m *Markup) xmlToPlain() string {
	re := regexp.MustCompile(`</p>`)
	s := re.ReplaceAllString(m.Raw, "")

	re = regexp.MustCompile(`<p>`)
	s = re.ReplaceAllString(s, "\t")
	return s
}
