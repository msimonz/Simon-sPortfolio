// Portafolio — dos comportamientos, nada más:
//   1. medir cada cable del diagrama para que se dibuje solo,
//   2. encender el nodo del rail correspondiente a la sección visible.
//
// No hay aparición al hacer scroll a propósito: con secciones tan altas el
// fundido deja el texto a medio pintar mientras se baja, y el movimiento del
// diagrama ya es suficiente.

(function () {
  "use strict";

  // 1. Cada trazo se anima con su propia longitud, así los seis cables
  //    terminan de dibujarse juntos en vez de a destiempo.
  document.querySelectorAll(".flow .wire").forEach(function (wire) {
    wire.style.setProperty("--len", Math.ceil(wire.getTotalLength()));
  });

  // 2. Rail. Se resuelve por posición y no con IntersectionObserver: las
  //    secciones aquí son mucho más altas que la ventana, y comparar razones
  //    de intersección deja encendida la sección grande equivocada.
  var nodes = Array.prototype.slice.call(document.querySelectorAll("[data-node]"));
  var sections = nodes
    .map(function (node) { return document.querySelector(node.getAttribute("href")); })
    .filter(Boolean);

  if (!sections.length) return;

  var current = null;

  function resolve() {
    // La línea de lectura: un tercio bajo el borde superior de la ventana.
    var line = window.innerHeight * 0.33;
    var active = sections[0];

    sections.forEach(function (section) {
      if (section.getBoundingClientRect().top <= line) active = section;
    });

    // Al pie del documento gana la última sección: Contacto es demasiado corta
    // para llegar nunca a cruzar la línea de lectura por sí sola.
    var bottom = window.innerHeight + window.scrollY;
    if (bottom >= document.documentElement.scrollHeight - 4) {
      active = sections[sections.length - 1];
    }

    if (active.id === current) return;
    current = active.id;

    nodes.forEach(function (node) {
      node.classList.toggle("is-active", node.getAttribute("href") === "#" + current);
    });
  }

  // Se resuelve directo en el evento, sin encolar en requestAnimationFrame: en
  // una pestaña en segundo plano el frame no llega nunca y el rail se quedaba
  // congelado en la sección que estuviera marcada al perder el foco.
  window.addEventListener("scroll", resolve, { passive: true });
  window.addEventListener("resize", resolve, { passive: true });
  resolve();
})();
