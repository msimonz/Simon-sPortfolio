# Portafolio — Simón Márquez Ramírez

Sitio personal escrito **enteramente en Go**, sólo con la librería estándar
(`net/http`, `html/template`, `embed`). Cero dependencias, cero `node_modules`.

La página está **en inglés**; el código y estos comentarios siguen en español.

## Correrlo

```sh
go run .              # http://localhost:8080
go run . -port 3000   # otro puerto
```

## Publicarlo

El mismo binario genera el sitio como archivos planos, listos para GitHub
Pages, Netlify, Vercel o un bucket de OCI Object Storage:

```sh
go run . -build dist
```

Eso escribe `dist/index.html` y `dist/static/`. No hace falta servidor Go en
producción: el servidor es sólo para desarrollo.

## Estructura

```
main.go                     servidor + generador estático
internal/content/
  types.go                  los tipos del contenido
  data.go                   ← TODO EL TEXTO DEL SITIO ESTÁ AQUÍ
web/templates/index.html    la página
web/static/css/site.css     estilos
web/static/js/site.js       rail de navegación + medición de los cables del diagrama
```

Los assets van embebidos en el binario con `go:embed`, así que el ejecutable
compilado es un solo archivo sin nada alrededor.

## Qué falta por llenar

En `internal/content/data.go`, buscar la palabra **REEMPLAZAR**:

1. **`profile()` → `Links`** — las URL de LinkedIn y GitHub están vacías. Un
   enlace sin URL no se pinta, así que el sitio no muestra links rotos; en
   cuanto pongas la URL aparecen solos en el hero y en Contacto.
2. **`funFacts()`** — el primero (la oración de cada mañana) es real, lo diste
   tú. Los otros ocho siguen siendo **inventados**, de relleno. Cámbialos por
   los reales. Si dejas la lista vacía, la sección desaparece sola; y si queda
   un número impar de datos, la última tarjeta se estira a todo el ancho para
   no dejar un hueco.

Todo lo demás salió de `SimonMarquezCV.pdf`, salvo la sección `vision()`, que
salió de lo que contaste: el plan multicloud y la experiencia con AWS.

**Ojo con AWS:** no aparece en el PDF, lo agregué a `skillSets()` y a
`vision()` porque lo mencionaste. Si querías decir algo más liviano que
«experiencia», bájale el tono ahí.

## Decisiones de diseño

- **Paleta:** `#020031` como tinta dominante sobre papel azulado frío
  (`#EEF1F8`). El verde `#6CCB64` está reservado para *señal viva* —tráfico en
  el diagrama, estado activo, certificación obtenida— y nunca se usa como
  adorno. Para texto verde sobre fondo claro se usa `#2C7A2A`, porque el verde
  de marca no alcanza el contraste mínimo legible.
- **Tipografía:** Archivo Expanded para titulares (aire de rótulo de plano
  técnico), IBM Plex Sans para cuerpo, y **San Francisco** —la tipografía de
  sistema de Apple— para etiquetas, chips y rótulos del diagrama. SF se pide
  por nombre de sistema (`-apple-system`), no se descarga: en Mac y iPhone es
  la real y en otras plataformas cae en la fuente nativa. Como SF es
  ópticamente más ligera que una monoespaciada al mismo tamaño, todas las
  etiquetas llevan peso explícito (600 en versalitas, 500 en el resto).
- **El diagrama del hero** no es decoración: es la integración que Simón
  construye a diario (SAP ERP, SFTP y colas → OIC → API Gateway, OKE, Oracle
  DB), con ruteo ortogonal como en una consola de iPaaS y paquetes viajando por
  los cables.
- **Las vías de la experiencia** (cloud / academia / comunidad) se codifican en
  el filete izquierdo de cada rol, no con números `01 / 02 / 03`.
- Sin fundidos al hacer scroll: con secciones tan altas dejaban el texto a
  medio pintar. `prefers-reduced-motion` detiene los paquetes y el trazado.
- Hay hoja de estilos de impresión: la página se imprime como hoja de vida.
