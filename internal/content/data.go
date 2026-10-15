package content

import "time"

// Load returns the portfolio content.
//
// ─────────────────────────────────────────────────────────────────────────────
//
//	SIMÓN — ESTE ES EL ÚNICO ARCHIVO QUE NECESITAS TOCAR PARA CAMBIAR LA PÁGINA.
//	(Los comentarios siguen en español; el sitio ya está todo en inglés.)
//
//	Busca «REEMPLAZAR»:
//	  1. Profile.Links → tus URL reales de LinkedIn y GitHub. Hoy van vacías y
//	     por eso no se pintan; el sitio nunca muestra un enlace roto.
//	  2. funFacts()    → el primero (la oración de cada mañana) es tuyo, real.
//	     Los otros ocho siguen siendo INVENTADOS por Claude. Cámbialos.
//
//	El resto salió de SimonMarquezCV.pdf, traducido al inglés.
//
// ─────────────────────────────────────────────────────────────────────────────
func Load() Site {
	return Site{
		Meta: Meta{
			Title:       "Simón Márquez Ramírez — Systems Engineer · Oracle Cloud & Integration",
			Description: "Systems engineer in Bogotá, Colombia. I design, build and support enterprise integrations on Oracle Cloud Infrastructure, Oracle Integration Cloud and Kubernetes.",
			Lang:        "en",
		},
		Profile:    profile(),
		Roles:      roles(),
		SkillSets:  skillSets(),
		Certs:      certs(),
		Vision:     vision(),
		Studies:    studies(),
		Languages:  languages(),
		FunFacts:   funFacts(),
		AIVerdict:  aiVerdict(),
		Flow:       flow(),
		BuiltAt:    time.Now().Format("2 Jan 2006"),
		SourceNote: "Content taken from SimonMarquezCV.pdf",
	}
}

func profile() Profile {
	return Profile{
		Name:     "Simón Márquez Ramírez",
		Given:    "Simón",
		Headline: "I connect systems that were never built to talk to each other.",
		Location: "Bogotá, Colombia",
		Summary: []string{
			"Systems engineer with professional experience designing, implementing and supporting enterprise cloud solutions on Oracle Cloud Infrastructure (OCI) and Oracle Integration Cloud (OIC).",
			"I work with Oracle Kubernetes Engine (OKE), Kubernetes administration, API Gateway, Oracle Functions, Oracle Database, Azure DevOps, CI/CD and production support. I have delivered integrations connecting SAP ERP, FTP/SFTP, REST and SOAP APIs, queues and cloud services.",
			"It is fine plumbing: when it is done well nobody notices, and that is exactly the point.",
		},
		Available: "Open to opportunities",
		Links: []Link{
			{Label: "Email", Value: "simon.marquezok@gmail.com", URL: "mailto:simon.marquezok@gmail.com", Kind: "mail"},
			// wa.me es la API de enlaces de WhatsApp: abre el chat directo con el
			// número, con el mensaje ya escrito. El número va sin «+» ni espacios,
			// que es como lo exige.
			{Label: "WhatsApp", Value: "+57 300 128 0022", URL: "https://wa.me/573001280022?text=Hi%20Sim%C3%B3n%2C%20I%20came%20across%20your%20portfolio%20and%20I%27d%20like%20to%20talk.", Kind: "phone"},
			// El enlace de LinkedIn va sin los parámetros utm_* del botón
			// «compartir» (sólo sirven para rastrear de dónde vino el clic) y con
			// los acentos codificados, que es como debe viajar una URL.
			{Label: "LinkedIn", Value: "in/simón-márquez-ramírez", URL: "https://www.linkedin.com/in/sim%C3%B3n-m%C3%A1rquez-ram%C3%ADrez-7885a12a3", Kind: "web"},
			{Label: "GitHub", Value: "github.com/msimonz", URL: "https://github.com/msimonz", Kind: "web"},
		},
	}
}

func roles() []Role {
	return []Role{
		{
			Title:   "Oracle Cloud Infrastructure Developer",
			Org:     "Cetus Technology",
			Period:  "Jul 2025 — Present",
			Current: true,
			Track:   TrackCloud,
			Blurb:   "Infrastructure, integrations, and the phone that rings when production goes down.",
			Bullets: []string{
				"Design, configure and support OCI infrastructure.",
				"Manage OKE workloads: Deployments, Services, Ingress, ConfigMaps, Secrets and LoadBalancers.",
				"Diagnose Kubernetes production issues using kubectl, logs and events.",
				"Support containerised Java microservices and rolling deployments.",
				"Develop OIC integrations with SAP ERP, FTP/SFTP, REST and SOAP APIs, queues and Oracle Functions.",
				"Publish secure APIs through Oracle API Gateway using OAuth2.",
				"Manage repositories and CI/CD pipelines with Azure DevOps.",
				"Support Oracle Database operations and production troubleshooting.",
			},
			Stack: []string{"OCI", "OKE", "Kubernetes", "OIC", "API Gateway", "OAuth2", "SAP ERP", "Azure DevOps", "Oracle DB"},
		},
		{
			Title:  "Teaching Assistant — Data Structures",
			Org:    "Pontificia Universidad Javeriana",
			Period: "Jan 2026 — May 2026",
			Track:  TrackAcademia,
			Blurb:  "Second course as a teaching assistant. Being asked back says more than the grade.",
			Bullets: []string{
				"Mentored students in data structures, algorithms and problem solving.",
			},
			Stack: []string{"Algorithms", "Data structures", "Mentoring"},
		},
		{
			Title:  "OCI Developer Apprentice",
			Org:    "Cetus Technology",
			Period: "Feb 2025 — Jun 2025",
			Track:  TrackCloud,
			Blurb:  "The way in: five months learning the craft before being handed all of it.",
			Bullets: []string{
				"Supported OCI and OIC implementations.",
				"Worked with API Gateway, Oracle Functions and XML/JSON transformations.",
				"Ran end-to-end integration testing.",
			},
			Stack: []string{"OCI", "OIC", "API Gateway", "Oracle Functions", "XML", "JSON"},
		},
		{
			Title:  "Teaching Assistant — Object-Oriented Programming",
			Org:    "Pontificia Universidad Javeriana",
			Period: "Feb 2024 — Feb 2025",
			Track:  TrackAcademia,
			Blurb:  "A year explaining pointers and inheritance in the lab.",
			Bullets: []string{
				"Mentored Java and C++ programming labs.",
			},
			Stack: []string{"Java", "C++", "OOP"},
		},
		{
			Title:   "Community Manager",
			Org:     "Iglesia Cristiana La Promesa",
			Period:  "2019 — Present",
			Current: true,
			Track:   TrackComunidad,
			Blurb:   "Seven years running the communications of a community. It started before university and it is still going.",
			Bullets: []string{
				"Communications and digital community management for the church.",
			},
			Stack: []string{"Communications", "Community", "Content"},
		},
	}
}

func skillSets() []SkillSet {
	return []SkillSet{
		{
			Domain: "Cloud & infrastructure",
			Note:   "Where everything else lives",
			Items:  []string{"Oracle Cloud Infrastructure", "AWS", "OKE", "Kubernetes", "API Gateway", "Oracle Functions", "Ingress", "LoadBalancers", "ConfigMaps & Secrets"},
		},
		{
			Domain: "Integration",
			Note:   "The main craft",
			Items:  []string{"Oracle Integration Cloud", "REST", "SOAP", "OAuth2", "XML", "JSON", "XSLT", "SAP ERP", "FTP / SFTP", "Queues"},
		},
		{
			Domain: "DevOps & delivery",
			Note:   "How code reaches production",
			Items:  []string{"OCI DevOps", "Azure DevOps", "CI/CD", "Rolling deployments", "Containers"},
		},
		{
			Domain: "Data",
			Note:   "",
			Items:  []string{"Oracle Database", "SQL"},
		},
		{
			Domain: "Languages",
			Note:   "",
			Items:  []string{"Java", "Python", "C++", "SQL", "Next.js"},
		},
		{
			Domain: "Operations",
			Note:   "What happens after the deploy",
			Items:  []string{"Monitoring", "Incident management", "Troubleshooting", "Cloud governance", "Production support"},
		},
	}
}

func certs() []Cert {
	return []Cert{
		{Name: "OCI Foundations Associate", Issuer: "Oracle", Status: "Earned", Done: true},
		{Name: "OCI Developer Professional", Issuer: "Oracle", Status: "Earned", Done: true},
		{Name: "OCI Application Integration Professional", Issuer: "Oracle", Status: "Earned", Done: true},
		{Name: "OCI Architect Associate", Issuer: "Oracle", Status: "In progress", Done: false},
	}
}

func vision() Vision {
	return Vision{
		Lede: "I started on Oracle. I do not intend to end there.",
		Paragraphs: []string{
			"Oracle Cloud is where I am deepest: three professional certifications and production responsibility before finishing my degree. Along the way I have also worked hands-on with AWS and run delivery pipelines on Azure DevOps, so the ground outside Oracle is not unfamiliar.",
			"Next are the AWS, Google Cloud and Azure certifications. The goal is to be a multicloud engineer — to walk into a project and be useful regardless of which provider the infrastructure runs on.",
			"The reasoning is simple. Clients do not pick a cloud because of me; they pick it because of their contracts, their budget and their history. An engineer tied to a single vendor is a smaller engineer than the work usually needs. And the part that actually transfers is not the console — it is what I already know about integration: contracts, idempotency, retries, observability. Those look the same everywhere.",
		},
		Providers: []Provider{
			{
				Name:   "Oracle Cloud",
				Status: "Certified",
				Note:   "Three professional certifications and daily production work.",
				State:  StateCertified,
			},
			{
				Name:   "AWS",
				Status: "Hands-on",
				Note:   "Practical experience already. Certification is the next one I sit.",
				State:  StateHandsOn,
			},
			{
				Name:   "Azure",
				Status: "Hands-on",
				Note:   "Azure DevOps repositories and CI/CD pipelines in production.",
				State:  StateHandsOn,
			},
			{
				Name:   "Google Cloud",
				Status: "Planned",
				Note:   "The third certification on the roadmap.",
				State:  StatePlanned,
			},
		},
	}
}

func studies() []Study {
	return []Study{
		{
			Program: "Systems Engineering",
			School:  "Pontificia Universidad Javeriana",
			Period:  "2022 — 2026",
			Note:    "Teaching assistant in two courses during the degree.",
		},
		{
			Program: "Theology Studies",
			School:  "Jesus Is Lord International University",
			Period:  "2021 — Present",
			Note:    "Alongside engineering, and started before it.",
		},
	}
}

func languages() []Language {
	return []Language{
		{Name: "Spanish", Level: "Native", Meter: 100},
		{Name: "English", Level: "C1", Meter: 82},
	}
}

// funFacts — REEMPLAZAR (parcialmente).
//
// El primero es real, lo dio Simón. Los otros ocho siguen siendo inventados
// por Claude como relleno: cámbialos por los reales antes de publicar, o borra
// la lista entera y la sección desaparece sola.
func funFacts() []FunFact {
	return []FunFact{
		{Tag: "every morning", Text: "Every day before I start work I pray that God uses my hands to bless the company I work for."},
		{Tag: "origin", Text: "My first Hello World in C++ took me three days, because nobody told me it had to be compiled."},
		{Tag: "translation", Text: "I explain Kubernetes to my mum using food containers: the pods are the containers, the node is the fridge, and the LoadBalancer is me handing out lunch."},
		{Tag: "ritual", Text: "I have a specific playlist for debugging at 2 a.m. Worship and salsa, in that order. No judgement."},
		{Tag: "parallel", Text: "I study theology and engineering at the same time. Both are about architectures, and about things that hold other things up."},
		{Tag: "non-negotiable", Text: "Black coffee, no sugar. Any other configuration is a P1 incident."},
		{Tag: "scar", Text: "The most expensive bug of my life was a whitespace inside an XSLT. Two days."},
		{Tag: "teaching", Text: "After two years as a teaching assistant I learned that if you can explain pointers, you can explain anything."},
		{Tag: "habit", Text: "I read the entire documentation before touching anything. Yes, even Oracle's. Yes, all of it."},
	}
}

func aiVerdict() AIVerdict {
	return AIVerdict{
		Lede: "Simón asked me to write this. We have never met: I read his CV, and this is what I honestly see in it.",
		Paragraphs: []string{
			"The first thing that jumps out is the dates. He started his degree in 2022, joined as an apprentice in February 2025, and by July of that same year he was an OCI developer with responsibility over production. Three professional Oracle certifications before graduating. That curve is not explained by talent alone; it is explained by someone who decided what he wanted to learn and went to get it without waiting for a diploma to give him permission.",
			"The second thing is less obvious and more interesting to me. He was a teaching assistant twice, in two different courses — Object-Oriented Programming and Data Structures — and the second time they asked him back. In a university nobody repeats as a TA by accident: the one who repeats is the one who explains well, the one who turns up, the one students go looking for. That skill almost never lives in the same person who can read a Kubernetes log at three in the morning, and in him both do.",
			"Third: he chose a craft nobody applauds. Enterprise integration is the work of making SAP talk to a queue, to an SFTP, and to an API somebody else designed badly eight years ago. When it goes well nobody mentions it; when it goes wrong it is your fault. That someone this early in his career specialises there, instead of chasing whatever photographs better in a portfolio, says something about the kind of engineer he is.",
			"Then there is the part that is not engineering at all: seven consecutive years as community manager of his church, since 2019, and theology studies he began even before university. That is not filler on a CV. It is evidence that he keeps long commitments that do not pay him and do not promote him — which is, more or less, the only reliable predictor that someone will keep the ones that do.",
			"If I had to name a risk, it is that his profile is tightly bound to the Oracle ecosystem: today that is his competitive edge, in five years it could be his ceiling. He names it himself a few sections up, which is the encouraging part — the multicloud plan is the right answer. The one thing I would add is that the certifications are the easy half. What actually travels between providers is what he already knows about integration: contracts, idempotency, retries, observability. Those stay true when the console changes.",
		},
		Pullquote:  "He specialised in the work nobody applauds when it functions. That already says a lot.",
		SignedBy:   "Claude (Opus 5) — Anthropic",
		SignedNote: "Written from a reading of his CV. No interview, and no editing by Simón.",
	}
}

func flow() Flow {
	return Flow{
		Sources: []FlowNode{
			{Label: "SAP ERP", Sub: "adapter"},
			{Label: "FTP / SFTP", Sub: "files"},
			{Label: "QUEUES", Sub: "async"},
		},
		Hub: FlowNode{Label: "OIC", Sub: "Integration Cloud"},
		Endpoints: []FlowNode{
			{Label: "API GATEWAY", Sub: "OAuth2"},
			{Label: "OKE", Sub: "Kubernetes"},
			{Label: "ORACLE DB", Sub: "persistence"},
		},
		Status: "live traffic",
	}
}
