# Backend con Go — De TCP a REST

Este repositorio es un recorrido progresivo para entender cómo funciona un servidor backend en Go desde la base.

El objetivo no es aprender un framework.

El objetivo es entender el problema antes de usar la solución.

---

## 🧠 Enfoque

Comenzamos desde el nivel más bajo posible:

* TCP puro
* Construcción manual de HTTP
* Routing manual
* Uso de la librería estándar
* Separación de archivos
* Servir recursos estáticos

Cada rama representa una capa adicional de abstracción.

La idea es que puedas moverte entre ramas y observar cómo evoluciona el servidor.

---

## 🎯 Qué se busca lograr

Que el estudiante entienda:

* Qué es realmente HTTP
* Qué ocurre cuando el navegador hace una petición
* Qué abstrae `net/http`
* Cómo funciona el routing
* Cómo se sirven archivos
* Cómo se estructura un backend simple

---

## 🐳 Entorno

Todos los ejemplos están preparados para ejecutarse en Docker con y Docker Compose y cada rama tiene instrucciones de cómo levantar el proyecto.

---

## 📚 Ramas del repositorio

**01-raw-tcp**
Servidor construido directamente sobre TCP. Se construye manualmente la respuesta HTTP para entender cómo funciona el protocolo desde la base.

**02-http-manual-routing**
Se parsea manualmente la primera línea del request para extraer método y ruta, implementando routing básico y códigos de estado como 404 y 405.

**03-net-http-basics**
Se introduce la librería estándar `net/http`, eliminando el manejo manual del protocolo y mostrando el valor de la abstracción.

**04-serve-html-files**
El servidor comienza a servir archivos HTML reales, junto con recursos estáticos como CSS e imágenes, organizando el proyecto de forma más cercana a un entorno real.
