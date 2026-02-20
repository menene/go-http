# 06 - POSTS

En esta etapa el servidor deja de ser únicamente un generador de vistas y comienza a **recibir información desde el cliente**.

Introducimos formularios HTML y el método `POST`, permitiendo interacción real entre navegador y servidor.

Este es el punto donde el backend pasa de ser estático a procesar datos.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Cómo funciona un formulario HTML
* Qué diferencia hay entre `GET` y `POST`
* Qué ocurre cuando el navegador envía datos al servidor
* Cómo usar `r.ParseForm()`
* Cómo usar `r.FormValue()`
* Cómo renderizar una respuesta basada en datos enviados por el usuario

---

## 📁 Estructura del proyecto

```
.
├── main.go
├── src/
│   ├── css/
│   │   └── styles.css
│   ├── assets/
│   │   └── gopher.png
│   └── templates/
│       ├── layout.html
│       ├── index.html
│       ├── about.html
│       └── form.html
├── Dockerfile
└── docker-compose.yml
```

Se agrega un nuevo template: `form.html`.

---

## 🧠 Qué cambió respecto a la rama anterior

Antes:

* El servidor solo renderizaba vistas
* No recibía datos del usuario

Ahora:

* Existe un formulario HTML
* El navegador envía una petición `POST`
* El servidor procesa datos enviados en el body
* El contenido renderizado depende del input del usuario

El flujo ahora es bidireccional:

Cliente → Servidor → Respuesta dinámica

---

## 🧩 Flujo completo

1. El usuario visita `/form`
2. El servidor responde con el formulario
3. El usuario envía el formulario
4. El navegador envía una petición `POST` a `/form`
5. El servidor ejecuta:

```go
r.ParseForm()
name := r.FormValue("name")
```

6. El servidor vuelve a renderizar la vista mostrando el resultado

---

## 🔎 Manejo de métodos HTTP

En esta rama se introduce control explícito por método:

```go
switch r.Method {
case http.MethodGet:
case http.MethodPost:
default:
}
```

Esto permite definir comportamientos distintos según el tipo de request.

---

## 🧠 Qué hace `ParseForm()`

`r.ParseForm()` analiza:

* Parámetros en la URL
* Datos enviados en el body (formularios POST)

Después de ejecutarlo, los valores quedan disponibles para ser leídos.

---

## 🧠 Qué hace `FormValue()`

`r.FormValue("name")` devuelve el primer valor asociado a la clave indicada.

Es una forma conveniente de acceder a datos del formulario.

---

## 🐳 Ejecución

El servidor sigue escuchando en el puerto 80 dentro del contenedor.

En `docker-compose.yml` se mapea:

```yaml
ports:
  - "8080:80"
```

Acceder desde el navegador:

```
http://localhost:8080/form
```

---

## 📌 Qué estamos aprendiendo realmente

En esta etapa introducimos:

* Lectura del body de una petición HTTP
* Manejo de métodos distintos (GET vs POST)
* Interacción real entre cliente y servidor
* Generación de vistas basadas en datos enviados por el usuario

Este es el momento donde el backend comienza a procesar información, no solo a servir contenido.
