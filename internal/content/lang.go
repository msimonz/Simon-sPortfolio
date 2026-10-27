package content

import (
	"fmt"
	"time"
)

// Lang is a language the site is published in. Each one becomes its own page:
// English at the root, Spanish under es/.
type Lang string

const (
	English Lang = "en"
	Spanish Lang = "es"
)

// Langs lists every published language, in the order the switch shows them.
var Langs = []Lang{English, Spanish}

// Path is where the language's page lives, relative to the site root: "" for
// English, "es/" for Spanish.
func (l Lang) Path() string {
	if l == English {
		return ""
	}
	return string(l) + "/"
}

// pick returns the text for this language. Every string on the site goes
// through it, so the English and the Spanish sit side by side in data.go.
func (l Lang) pick(en, es string) string {
	if l == Spanish {
		return es
	}
	return en
}

// date formats the "updated" stamp in the footer. time.Format only knows
// English month names, so Spanish gets its own table.
func (l Lang) date(t time.Time) string {
	if l != Spanish {
		return t.Format("2 Jan 2006")
	}
	months := [...]string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sep", "oct", "nov", "dic"}
	return fmt.Sprintf("%d %s %d", t.Day(), months[t.Month()-1], t.Year())
}

// T picks the text for the page being rendered. It is what the template uses
// for its own labels: {{.T "Contact" "Contacto"}}.
func (s Site) T(en, es string) string {
	return s.Lang.pick(en, es)
}

// Root is the relative path from this page back to the site root, so assets
// and the language switch resolve the same on localhost, on GitHub Pages under
// a repository subpath, or anywhere else.
func (s Site) Root() string {
	if s.Lang.Path() == "" {
		return "./"
	}
	return "../"
}
