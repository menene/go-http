# 07 - JSON API

En esta etapa el servidor deja completamente de renderizar HTML y pasa a ser una **API pura que devuelve JSON**.

Se elimina cualquier rastro de vistas, templates o archivos estáticos.

El backend ahora actúa como un servicio que expone datos a través de endpoints HTTP.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Que un backend no necesariamente devuelve HTML
* Que HTTP es solo el medio de transporte
* Cómo devolver datos estructurados en formato JSON
* Cómo usar el paquete estándar `encoding/json`
* Cómo establecer correctamente headers y códigos de estado

---

## 📁 Estructura del proyecto

```
.
├── main.go
├── Dockerfile
└── docker-compose.yml
```

El proyecto ahora es mínimo: solo servidor y lógica de API.

---

## 🧠 Qué cambió respecto a la rama anterior

Antes:

* El servidor renderizaba vistas HTML
* Existían templates y archivos estáticos

Ahora:

* El servidor solo expone endpoints bajo `/api/`
* Las respuestas son JSON
* No existe capa de presentación

Se produce una separación clara entre backend y frontend.

---

## 🧩 Ejemplo de endpoint

Un endpoint típico:

```
GET /api/hello
```

Respuesta:

```json
{"message":"Hello from pure JSON API"}
```

---

## 🔎 Conceptos introducidos

* Uso de `encoding/json` para serializar structs
* Uso de struct tags para controlar el formato JSON
* Manejo explícito de `Content-Type: application/json`
* Envío de códigos de estado HTTP apropiados

---

## 🐳 Ejecución

El servidor escucha en el puerto 80 dentro del contenedor.

En `docker-compose.yml` se mapea:

```yaml
ports:
  - "8080:80"
```

Probar con:

```bash
curl http://localhost:8080/api/hello
```

---

## 📌 Qué estamos aprendiendo realmente

En esta etapa entendemos que:

* Una API es simplemente un conjunto de endpoints HTTP
* JSON es solo un formato de serialización
* El backend puede existir sin interfaz gráfica

Este es el paso previo antes de modelar recursos y construir una API REST completa.

---

Autor: Curso Backend con Go
