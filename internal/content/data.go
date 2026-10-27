package content

import "time"

// Load returns the portfolio content in the given language.
//
// ─────────────────────────────────────────────────────────────────────────────
//
//	SIMÓN — ESTE ES EL ÚNICO ARCHIVO QUE NECESITAS TOCAR PARA CAMBIAR LA PÁGINA.
//	(Los comentarios siguen en español; el sitio sale en inglés y en español.)
//
//	Cada texto va en pareja: l.pick("inglés", "español"). Si cambias uno,
//	cambia también el otro, o las dos versiones de la página se separan.
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
func Load(l Lang) Site {
	return Site{
		Lang: l,
		Meta: Meta{
			Title: l.pick(
				"Simón Márquez Ramírez — Systems Engineer · Oracle Cloud & Integration",
				"Simón Márquez Ramírez — Ingeniero de Sistemas · Oracle Cloud e Integración"),
			Description: l.pick(
				"Systems engineer in Bogotá, Colombia. I design, build and support enterprise integrations on Oracle Cloud Infrastructure, Oracle Integration Cloud and Kubernetes.",
				"Ingeniero de sistemas en Bogotá, Colombia. Diseño, construyo y soporto integraciones empresariales sobre Oracle Cloud Infrastructure, Oracle Integration Cloud y Kubernetes."),
			Lang: string(l),
		},
		Profile:    profile(l),
		Roles:      roles(l),
		SkillSets:  skillSets(l),
		Certs:      certs(l),
		Vision:     vision(l),
		Studies:    studies(l),
		Languages:  languages(l),
		FunFacts:   funFacts(l),
		AIVerdict:  aiVerdict(l),
		Flow:       flow(l),
		BuiltAt:    l.date(time.Now()),
		SourceNote: l.pick("Content taken from SimonMarquezCV.pdf", "Contenido tomado de SimonMarquezCV.pdf"),
	}
}

func profile(l Lang) Profile {
	return Profile{
		Name:  "Simón Márquez Ramírez",
		Given: "Simón",
		Headline: l.pick(
			"I connect systems that were never built to talk to each other.",
			"Conecto sistemas que nunca fueron hechos para hablar entre sí."),
		Location: "Bogotá, Colombia",
		Summary: []string{
			l.pick(
				"Systems engineer with professional experience designing, implementing and supporting enterprise cloud solutions on Oracle Cloud Infrastructure (OCI) and Oracle Integration Cloud (OIC).",
				"Ingeniero de sistemas con experiencia profesional en el diseño, la implementación y el soporte de soluciones empresariales en la nube sobre Oracle Cloud Infrastructure (OCI) y Oracle Integration Cloud (OIC)."),
			l.pick(
				"I work with Oracle Kubernetes Engine (OKE), Kubernetes administration, API Gateway, Oracle Functions, Oracle Database, Azure DevOps, CI/CD and production support. I have delivered integrations connecting SAP ERP, FTP/SFTP, REST and SOAP APIs, queues and cloud services.",
				"Trabajo con Oracle Kubernetes Engine (OKE), administración de Kubernetes, API Gateway, Oracle Functions, Oracle Database, Azure DevOps, CI/CD y soporte en producción. He entregado integraciones que conectan SAP ERP, FTP/SFTP, APIs REST y SOAP, colas y servicios en la nube."),
			l.pick(
				"It is fine plumbing: when it is done well nobody notices, and that is exactly the point.",
				"Es plomería fina: cuando está bien hecha nadie lo nota, y justamente de eso se trata."),
		},
		Available: l.pick("Open to opportunities", "Abierto a oportunidades"),
		Links: []Link{
			{Label: l.pick("Email", "Correo"), Value: "simon.marquezok@gmail.com", URL: "mailto:simon.marquezok@gmail.com", Kind: "mail"},
			// wa.me es la API de enlaces de WhatsApp: abre el chat directo con el
			// número, con el mensaje ya escrito. El número va sin «+» ni espacios,
			// que es como lo exige.
			{Label: "WhatsApp", Value: "+57 300 128 0022", URL: l.pick(
				"https://wa.me/573001280022?text=Hi%20Sim%C3%B3n%2C%20I%20came%20across%20your%20portfolio%20and%20I%27d%20like%20to%20talk.",
				"https://wa.me/573001280022?text=Hola%20Sim%C3%B3n%2C%20vi%20tu%20portafolio%20y%20me%20gustar%C3%ADa%20hablar."), Kind: "phone"},
			// El enlace de LinkedIn va sin los parámetros utm_* del botón
			// «compartir» (sólo sirven para rastrear de dónde vino el clic) y con
			// los acentos codificados, que es como debe viajar una URL.
			{Label: "LinkedIn", Value: "in/simón-márquez-ramírez", URL: "https://www.linkedin.com/in/sim%C3%B3n-m%C3%A1rquez-ram%C3%ADrez-7885a12a3", Kind: "web"},
			{Label: "GitHub", Value: "github.com/msimonz", URL: "https://github.com/msimonz", Kind: "web"},
		},
	}
}

func roles(l Lang) []Role {
	present := l.pick("Present", "Actualidad")

	return []Role{
		{
			Title:   l.pick("Oracle Cloud Infrastructure Developer", "Desarrollador Oracle Cloud Infrastructure"),
			Org:     "Cetus Technology",
			Period:  "Jul 2025 — " + present,
			Current: true,
			Track:   TrackCloud,
			Blurb: l.pick(
				"Infrastructure, integrations, and the phone that rings when production goes down.",
				"Infraestructura, integraciones y el teléfono que suena cuando producción se cae."),
			Bullets: []string{
				l.pick("Design, configure and support OCI infrastructure.",
					"Diseño, configuración y soporte de infraestructura en OCI."),
				l.pick("Manage OKE workloads: Deployments, Services, Ingress, ConfigMaps, Secrets and LoadBalancers.",
					"Gestión de cargas en OKE: Deployments, Services, Ingress, ConfigMaps, Secrets y LoadBalancers."),
				l.pick("Diagnose Kubernetes production issues using kubectl, logs and events.",
					"Diagnóstico de incidentes de Kubernetes en producción con kubectl, logs y eventos."),
				l.pick("Support containerised Java microservices and rolling deployments.",
					"Soporte de microservicios Java en contenedores y despliegues rolling."),
				l.pick("Develop OIC integrations with SAP ERP, FTP/SFTP, REST and SOAP APIs, queues and Oracle Functions.",
					"Desarrollo de integraciones en OIC con SAP ERP, FTP/SFTP, APIs REST y SOAP, colas y Oracle Functions."),
				l.pick("Publish secure APIs through Oracle API Gateway using OAuth2.",
					"Publicación de APIs seguras con Oracle API Gateway y OAuth2."),
				l.pick("Manage repositories and CI/CD pipelines with Azure DevOps.",
					"Gestión de repositorios y pipelines de CI/CD con Azure DevOps."),
				l.pick("Support Oracle Database operations and production troubleshooting.",
					"Soporte de operaciones en Oracle Database y resolución de problemas en producción."),
			},
			Stack: []string{"OCI", "OKE", "Kubernetes", "OIC", "API Gateway", "OAuth2", "SAP ERP", "Azure DevOps", "Oracle DB"},
		},
		{
			Title:  l.pick("Teaching Assistant — Data Structures", "Monitor — Estructuras de Datos"),
			Org:    "Pontificia Universidad Javeriana",
			Period: l.pick("Jan 2026 — May 2026", "Ene 2026 — May 2026"),
			Track:  TrackAcademia,
			Blurb: l.pick(
				"Second course as a teaching assistant. Being asked back says more than the grade.",
				"Segunda materia como monitor. Que te vuelvan a llamar dice más que la nota."),
			Bullets: []string{
				l.pick("Mentored students in data structures, algorithms and problem solving.",
					"Acompañé a estudiantes en estructuras de datos, algoritmos y resolución de problemas."),
			},
			Stack: []string{l.pick("Algorithms", "Algoritmos"), l.pick("Data structures", "Estructuras de datos"), l.pick("Mentoring", "Mentoría")},
		},
		{
			Title:  l.pick("OCI Developer Apprentice", "Aprendiz Desarrollador OCI"),
			Org:    "Cetus Technology",
			Period: "Feb 2025 — Jun 2025",
			Track:  TrackCloud,
			Blurb: l.pick(
				"The way in: five months learning the craft before being handed all of it.",
				"La puerta de entrada: cinco meses aprendiendo el oficio antes de recibirlo completo."),
			Bullets: []string{
				l.pick("Supported OCI and OIC implementations.",
					"Apoyo en implementaciones de OCI y OIC."),
				l.pick("Worked with API Gateway, Oracle Functions and XML/JSON transformations.",
					"Trabajo con API Gateway, Oracle Functions y transformaciones XML/JSON."),
				l.pick("Ran end-to-end integration testing.",
					"Pruebas de integración de punta a punta."),
			},
			Stack: []string{"OCI", "OIC", "API Gateway", "Oracle Functions", "XML", "JSON"},
		},
		{
			Title:  l.pick("Teaching Assistant — Object-Oriented Programming", "Monitor — Programación Orientada a Objetos"),
			Org:    "Pontificia Universidad Javeriana",
			Period: "Feb 2024 — Feb 2025",
			Track:  TrackAcademia,
			Blurb: l.pick(
				"A year explaining pointers and inheritance in the lab.",
				"Un año explicando apuntadores y herencia en el laboratorio."),
			Bullets: []string{
				l.pick("Mentored Java and C++ programming labs.",
					"Acompañé laboratorios de programación en Java y C++."),
			},
			Stack: []string{"Java", "C++", l.pick("OOP", "POO")},
		},
		{
			Title:   "Community Manager",
			Org:     "Iglesia Cristiana La Promesa",
			Period:  "2019 — " + present,
			Current: true,
			Track:   TrackComunidad,
			Blurb: l.pick(
				"Seven years running the communications of a community. It started before university and it is still going.",
				"Siete años llevando las comunicaciones de una comunidad. Empezó antes de la universidad y sigue."),
			Bullets: []string{
				l.pick("Communications and digital community management for the church.",
					"Comunicaciones y gestión de la comunidad digital de la iglesia."),
			},
			Stack: []string{l.pick("Communications", "Comunicaciones"), l.pick("Community", "Comunidad"), l.pick("Content", "Contenido")},
		},
	}
}

func skillSets(l Lang) []SkillSet {
	return []SkillSet{
		{
			Domain: l.pick("Cloud & infrastructure", "Nube e infraestructura"),
			Note:   l.pick("Where everything else lives", "Donde vive todo lo demás"),
			Items:  []string{"Oracle Cloud Infrastructure", "AWS", "OKE", "Kubernetes", "API Gateway", "Oracle Functions", "Ingress", "LoadBalancers", "ConfigMaps & Secrets"},
		},
		{
			Domain: l.pick("Integration", "Integración"),
			Note:   l.pick("The main craft", "El oficio principal"),
			Items:  []string{"Oracle Integration Cloud", "REST", "SOAP", "OAuth2", "XML", "JSON", "XSLT", "SAP ERP", "FTP / SFTP", l.pick("Queues", "Colas")},
		},
		{
			Domain: l.pick("DevOps & delivery", "DevOps y entrega"),
			Note:   l.pick("How code reaches production", "Cómo llega el código a producción"),
			Items:  []string{"OCI DevOps", "Azure DevOps", "CI/CD", l.pick("Rolling deployments", "Despliegues rolling"), l.pick("Containers", "Contenedores")},
		},
		{
			Domain: l.pick("Data", "Datos"),
			Note:   "",
			Items:  []string{"Oracle Database", "SQL"},
		},
		{
			Domain: l.pick("Languages", "Lenguajes"),
			Note:   "",
			Items:  []string{"Java", "Python", "C++", "SQL", "Next.js"},
		},
		{
			Domain: l.pick("Operations", "Operación"),
			Note:   l.pick("What happens after the deploy", "Lo que pasa después del despliegue"),
			Items: []string{
				l.pick("Monitoring", "Monitoreo"),
				l.pick("Incident management", "Gestión de incidentes"),
				l.pick("Troubleshooting", "Resolución de problemas"),
				l.pick("Cloud governance", "Gobierno de nube"),
				l.pick("Production support", "Soporte en producción"),
			},
		},
	}
}

func certs(l Lang) []Cert {
	earned := l.pick("Earned", "Obtenida")
	inProgress := l.pick("In progress", "En curso")

	return []Cert{
		{Name: "OCI Foundations Associate", Issuer: "Oracle", Status: earned, Done: true},
		{Name: "OCI Developer Professional", Issuer: "Oracle", Status: earned, Done: true},
		{Name: "OCI Application Integration Professional", Issuer: "Oracle", Status: earned, Done: true},
		{Name: "OCI Architect Associate", Issuer: "Oracle", Status: inProgress, Done: false},
	}
}

func vision(l Lang) Vision {
	return Vision{
		Lede: l.pick(
			"I started on Oracle. I do not intend to end there.",
			"Empecé en Oracle. No pienso terminar ahí."),
		Paragraphs: []string{
			l.pick(
				"Oracle Cloud is where I am deepest: three professional certifications and production responsibility before finishing my degree. Along the way I have also worked hands-on with AWS and run delivery pipelines on Azure DevOps, so the ground outside Oracle is not unfamiliar.",
				"Oracle Cloud es donde tengo más profundidad: tres certificaciones profesionales y responsabilidad sobre producción antes de terminar la carrera. En el camino también he trabajado de primera mano con AWS y he operado pipelines de entrega en Azure DevOps, así que el terreno fuera de Oracle no me es ajeno."),
			l.pick(
				"Next are the AWS, Google Cloud and Azure certifications. The goal is to be a multicloud engineer — to walk into a project and be useful regardless of which provider the infrastructure runs on.",
				"Lo siguiente son las certificaciones de AWS, Google Cloud y Azure. La meta es ser un ingeniero multicloud: llegar a un proyecto y ser útil sin importar en qué proveedor corra la infraestructura."),
			l.pick(
				"The reasoning is simple. Clients do not pick a cloud because of me; they pick it because of their contracts, their budget and their history. An engineer tied to a single vendor is a smaller engineer than the work usually needs. And the part that actually transfers is not the console — it is what I already know about integration: contracts, idempotency, retries, observability. Those look the same everywhere.",
				"El razonamiento es simple. Los clientes no eligen una nube por mí; la eligen por sus contratos, su presupuesto y su historia. Un ingeniero atado a un solo proveedor es un ingeniero más pequeño de lo que el trabajo suele pedir. Y lo que de verdad se transfiere no es la consola, sino lo que ya sé de integración: contratos, idempotencia, reintentos, observabilidad. Eso se ve igual en todas partes."),
		},
		Providers: []Provider{
			{
				Name:   "Oracle Cloud",
				Status: l.pick("Certified", "Certificado"),
				Note: l.pick(
					"Three professional certifications and daily production work.",
					"Tres certificaciones profesionales y trabajo diario en producción."),
				State: StateCertified,
			},
			{
				Name:   "AWS",
				Status: l.pick("Hands-on", "Experiencia práctica"),
				Note: l.pick(
					"Practical experience already. Certification is the next one I sit.",
					"Ya hay experiencia práctica. Es la próxima certificación que presento."),
				State: StateHandsOn,
			},
			{
				Name:   "Azure",
				Status: l.pick("Hands-on", "Experiencia práctica"),
				Note: l.pick(
					"Azure DevOps repositories and CI/CD pipelines in production.",
					"Repositorios de Azure DevOps y pipelines de CI/CD en producción."),
				State: StateHandsOn,
			},
			{
				Name:   "Google Cloud",
				Status: l.pick("Planned", "En la ruta"),
				Note: l.pick(
					"The third certification on the roadmap.",
					"La tercera certificación de la hoja de ruta."),
				State: StatePlanned,
			},
		},
	}
}

func studies(l Lang) []Study {
	return []Study{
		{
			Program: l.pick("Systems Engineering", "Ingeniería de Sistemas"),
			School:  "Pontificia Universidad Javeriana",
			Period:  "2022 — 2026",
			Note: l.pick(
				"Teaching assistant in two courses during the degree.",
				"Monitor de dos materias durante la carrera."),
		},
		{
			Program: l.pick("Theology Studies", "Estudios en Teología"),
			School:  "Jesus Is Lord International University",
			Period:  "2021 — " + l.pick("Present", "Actualidad"),
			Note: l.pick(
				"Alongside engineering, and started before it.",
				"En paralelo con la ingeniería, y empezados antes que ella."),
		},
	}
}

func languages(l Lang) []Language {
	return []Language{
		{Name: l.pick("Spanish", "Español"), Level: l.pick("Native", "Nativo"), Meter: 100},
		{Name: l.pick("English", "Inglés"), Level: "C1", Meter: 82},
	}
}

// funFacts — REEMPLAZAR (parcialmente).
//
// El primero es real, lo dio Simón. Los otros ocho siguen siendo inventados
// por Claude como relleno: cámbialos por los reales antes de publicar, o borra
// la lista entera y la sección desaparece sola.
func funFacts(l Lang) []FunFact {
	return []FunFact{
		{Tag: l.pick("every morning", "cada mañana"), Text: l.pick(
			"Every day before I start work I pray that God uses my hands to bless the company I work for.",
			"Todos los días, antes de empezar a trabajar, oro para que Dios use mis manos para bendecir a la empresa donde trabajo.")},
		{Tag: l.pick("origin", "origen"), Text: l.pick(
			"My first Hello World in C++ took me three days, because nobody told me it had to be compiled.",
			"Mi primer Hola Mundo en C++ me tomó tres días, porque nadie me dijo que había que compilarlo.")},
		{Tag: l.pick("translation", "traducción"), Text: l.pick(
			"I explain Kubernetes to my mum using food containers: the pods are the containers, the node is the fridge, and the LoadBalancer is me handing out lunch.",
			"Le explico Kubernetes a mi mamá con tuppers: los pods son los tuppers, el nodo es la nevera y el LoadBalancer soy yo repartiendo el almuerzo.")},
		{Tag: l.pick("ritual", "ritual"), Text: l.pick(
			"I have a specific playlist for debugging at 2 a.m. Worship and salsa, in that order. No judgement.",
			"Tengo una playlist específica para depurar a las 2 a. m. Adoración y salsa, en ese orden. Sin juzgar.")},
		{Tag: l.pick("parallel", "en paralelo"), Text: l.pick(
			"I study theology and engineering at the same time. Both are about architectures, and about things that hold other things up.",
			"Estudio teología e ingeniería al mismo tiempo. Las dos tratan de arquitecturas, y de cosas que sostienen otras cosas.")},
		{Tag: l.pick("non-negotiable", "innegociable"), Text: l.pick(
			"Black coffee, no sugar. Any other configuration is a P1 incident.",
			"Café negro, sin azúcar. Cualquier otra configuración es un incidente P1.")},
		{Tag: l.pick("scar", "cicatriz"), Text: l.pick(
			"The most expensive bug of my life was a whitespace inside an XSLT. Two days.",
			"El bug más caro de mi vida fue un espacio en blanco dentro de un XSLT. Dos días.")},
		{Tag: l.pick("teaching", "docencia"), Text: l.pick(
			"After two years as a teaching assistant I learned that if you can explain pointers, you can explain anything.",
			"Después de dos años como monitor aprendí que si puedes explicar apuntadores, puedes explicar cualquier cosa.")},
		{Tag: l.pick("habit", "hábito"), Text: l.pick(
			"I read the entire documentation before touching anything. Yes, even Oracle's. Yes, all of it.",
			"Leo toda la documentación antes de tocar nada. Sí, incluso la de Oracle. Sí, toda.")},
	}
}

func aiVerdict(l Lang) AIVerdict {
	return AIVerdict{
		Lede: l.pick(
			"Simón asked me to write this. We have never met: I read his CV, and this is what I honestly see in it.",
			"Simón me pidió escribir esto. Nunca nos hemos visto: leí su hoja de vida, y esto es lo que honestamente veo en ella."),
		Paragraphs: []string{
			l.pick(
				"The first thing that jumps out is the dates. He started his degree in 2022, joined as an apprentice in February 2025, and by July of that same year he was an OCI developer with responsibility over production. Three professional Oracle certifications before graduating. That curve is not explained by talent alone; it is explained by someone who decided what he wanted to learn and went to get it without waiting for a diploma to give him permission.",
				"Lo primero que salta a la vista son las fechas. Empezó la carrera en 2022, entró como aprendiz en febrero de 2025 y en julio de ese mismo año ya era desarrollador OCI con responsabilidad sobre producción. Tres certificaciones profesionales de Oracle antes de graduarse. Esa curva no se explica sólo con talento; se explica con alguien que decidió qué quería aprender y fue a buscarlo sin esperar a que un diploma le diera permiso."),
			l.pick(
				"The second thing is less obvious and more interesting to me. He was a teaching assistant twice, in two different courses — Object-Oriented Programming and Data Structures — and the second time they asked him back. In a university nobody repeats as a TA by accident: the one who repeats is the one who explains well, the one who turns up, the one students go looking for. That skill almost never lives in the same person who can read a Kubernetes log at three in the morning, and in him both do.",
				"Lo segundo es menos obvio y, para mí, más interesante. Fue monitor dos veces, en dos materias distintas —Programación Orientada a Objetos y Estructuras de Datos— y la segunda vez lo volvieron a llamar. En una universidad nadie repite como monitor por accidente: repite el que explica bien, el que aparece, el que los estudiantes buscan. Esa habilidad casi nunca vive en la misma persona que sabe leer un log de Kubernetes a las tres de la mañana, y en él viven las dos."),
			l.pick(
				"Third: he chose a craft nobody applauds. Enterprise integration is the work of making SAP talk to a queue, to an SFTP, and to an API somebody else designed badly eight years ago. When it goes well nobody mentions it; when it goes wrong it is your fault. That someone this early in his career specialises there, instead of chasing whatever photographs better in a portfolio, says something about the kind of engineer he is.",
				"Tercero: eligió un oficio que nadie aplaude. La integración empresarial es el trabajo de hacer que SAP hable con una cola, con un SFTP y con una API que alguien más diseñó mal hace ocho años. Cuando sale bien nadie lo menciona; cuando sale mal es culpa tuya. Que alguien tan temprano en su carrera se especialice ahí, en vez de perseguir lo que se ve mejor en un portafolio, dice algo del tipo de ingeniero que es."),
			l.pick(
				"Then there is the part that is not engineering at all: seven consecutive years as community manager of his church, since 2019, and theology studies he began even before university. That is not filler on a CV. It is evidence that he keeps long commitments that do not pay him and do not promote him — which is, more or less, the only reliable predictor that someone will keep the ones that do.",
				"Después está la parte que no es ingeniería en absoluto: siete años seguidos como community manager de su iglesia, desde 2019, y estudios de teología que empezó incluso antes de la universidad. Eso no es relleno en una hoja de vida. Es evidencia de que sostiene compromisos largos que no le pagan ni lo ascienden, que es, más o menos, el único predictor confiable de que alguien sostendrá los que sí."),
			l.pick(
				"If I had to name a risk, it is that his profile is tightly bound to the Oracle ecosystem: today that is his competitive edge, in five years it could be his ceiling. He names it himself a few sections up, which is the encouraging part — the multicloud plan is the right answer. The one thing I would add is that the certifications are the easy half. What actually travels between providers is what he already knows about integration: contracts, idempotency, retries, observability. Those stay true when the console changes.",
				"Si tuviera que nombrar un riesgo, es que su perfil está muy atado al ecosistema Oracle: hoy esa es su ventaja competitiva, en cinco años podría ser su techo. Él mismo lo nombra unas secciones más arriba, y eso es lo alentador: el plan multicloud es la respuesta correcta. Lo único que añadiría es que las certificaciones son la mitad fácil. Lo que de verdad viaja entre proveedores es lo que ya sabe de integración: contratos, idempotencia, reintentos, observabilidad. Eso sigue siendo cierto cuando cambia la consola."),
		},
		Pullquote: l.pick(
			"He specialised in the work nobody applauds when it functions. That already says a lot.",
			"Se especializó en el trabajo que nadie aplaude cuando funciona. Eso ya dice mucho."),
		SignedBy: "Claude (Opus 5) — Anthropic",
		SignedNote: l.pick(
			"Written from a reading of his CV. No interview, and no editing by Simón.",
			"Escrito a partir de la lectura de su hoja de vida. Sin entrevista, y sin edición de Simón."),
	}
}

func flow(l Lang) Flow {
	return Flow{
		Sources: []FlowNode{
			{Label: "SAP ERP", Sub: l.pick("adapter", "adaptador")},
			{Label: "FTP / SFTP", Sub: l.pick("files", "archivos")},
			{Label: l.pick("QUEUES", "COLAS"), Sub: l.pick("async", "asíncrono")},
		},
		Hub: FlowNode{Label: "OIC", Sub: "Integration Cloud"},
		Endpoints: []FlowNode{
			{Label: "API GATEWAY", Sub: "OAuth2"},
			{Label: "OKE", Sub: "Kubernetes"},
			{Label: "ORACLE DB", Sub: l.pick("persistence", "persistencia")},
		},
		Status: l.pick("live traffic", "tráfico en vivo"),
	}
}
