// Package content holds every piece of text that appears on the portfolio.
// Nothing here touches HTTP or HTML — it is plain data, so editing the site
// never means editing a template.
package content

// Site is the whole portfolio, ready to hand to a template.
type Site struct {
	Meta       Meta
	Profile    Profile
	Roles      []Role
	SkillSets  []SkillSet
	Certs      []Cert
	Vision     Vision
	Studies    []Study
	Languages  []Language
	FunFacts   []FunFact
	AIVerdict  AIVerdict
	Flow       Flow
	BuiltAt    string
	SourceNote string
}

// Meta drives <head> and social cards.
type Meta struct {
	Title       string
	Description string
	Lang        string
}

// Profile is the person: who they are and how to reach them.
type Profile struct {
	Name      string
	Given     string
	Headline  string
	Location  string
	Summary   []string
	Available string
	Links     []Link
}

// Link is a contact method. Empty URL means the link is skipped at render time,
// so a half-filled field never ships a dead anchor.
type Link struct {
	Label string
	Value string
	URL   string
	Kind  string // mail | phone | web
}

// Track separates the three parallel careers visible in the CV: the cloud job,
// the university teaching, and the church community work. The timeline colours
// and labels each role by its track rather than by an arbitrary number.
type Track string

const (
	TrackCloud     Track = "cloud"
	TrackAcademia  Track = "academia"
	TrackComunidad Track = "comunidad"
)

// Role is one job.
type Role struct {
	Title   string
	Org     string
	Period  string
	Current bool
	Track   Track
	Blurb   string
	Bullets []string
	Stack   []string
}

// SkillSet groups competencies by the problem they solve.
type SkillSet struct {
	Domain string
	Note   string
	Items  []string
}

// Cert is a certification. Status is "obtenida" or "en curso"; the page shows a
// live green dot for the former and a hollow one for the latter.
type Cert struct {
	Name   string
	Issuer string
	Status string
	Done   bool
}

// Vision is the multicloud plan: where the Oracle track is headed next.
type Vision struct {
	Lede       string
	Paragraphs []string
	Providers  []Provider
}

// ProviderState is how far along a cloud provider is. It reuses the same three
// steps the certification list uses, so a filled dot means the same thing
// everywhere on the page.
type ProviderState string

const (
	// StateCertified — certified and shipping on it.
	StateCertified ProviderState = "done"
	// StateHandsOn — real work done on it, certification still ahead.
	StateHandsOn ProviderState = "active"
	// StatePlanned — on the roadmap, not started.
	StatePlanned ProviderState = "planned"
)

// Provider is one cloud vendor on the roadmap.
type Provider struct {
	Name   string
	Status string
	Note   string
	State  ProviderState
}

// Study is a degree or programme.
type Study struct {
	Program string
	School  string
	Period  string
	Note    string
}

// Language is a spoken language and its level.
type Language struct {
	Name  string
	Level string
	Meter int // 0-100, drives the bar width
}

// FunFact is one entry of "Datos curiosos sobre mí".
type FunFact struct {
	Text string
	Tag  string
}

// AIVerdict is the "¿Qué piensa la IA sobre mí?" section: written by Claude,
// signed by Claude, and honest about what it is based on.
type AIVerdict struct {
	Lede       string
	Paragraphs []string
	Pullquote  string
	SignedBy   string
	SignedNote string
}

// Flow is the hero diagram: the integration Simón actually builds for a living.
type Flow struct {
	Sources   []FlowNode
	Hub       FlowNode
	Endpoints []FlowNode
	Status    string
}

// FlowNode is one box on the diagram.
type FlowNode struct {
	Label string
	Sub   string
}
