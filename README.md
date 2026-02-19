# 05 - Templates

En esta etapa dejamos de servir archivos HTML estáticos directamente y comenzamos a utilizar **plantillas (templates)** con `html/template`.

El objetivo no es agregar lógica dinámica todavía.

El objetivo es entender cómo funciona el renderizado del lado del servidor y cómo reutilizar una estructura común (layout).

---

## 🎯 Objetivo de esta etapa

Comprender:

* Qué es `html/template`
* Cómo separar layout y contenido
* Cómo renderizar vistas desde el backend
* Cómo reutilizar estructura HTML sin duplicación
* Cómo funciona la composición de templates en Go

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
│       └── about.html
├── Dockerfile
└── docker-compose.yml
```

Ahora el HTML vive dentro de la carpeta `templates`.

---

## 🧠 Qué cambió respecto a la rama anterior

Antes:

* Servíamos archivos HTML directamente con `http.ServeFile`
* El servidor solo entregaba archivos

Ahora:

* El servidor **renderiza** vistas usando plantillas
* Existe un `layout.html` compartido
* Cada página define su bloque de contenido

El backend ya no solo entrega archivos.
Ahora genera la vista final.

---

## 🧩 Cómo funciona el renderizado

Para cada request:

1. Se parsea `layout.html`
2. Se parsea el template específico de la página
3. El layout incluye el bloque `content`
4. Se ejecuta el template resultante

Ejemplo simplificado en Go:

```go
tmpl, _ := template.ParseFiles(
    "layout.html",
    "index.html",
)

tmpl.Execute(w, nil)
```

El layout define dónde se inserta el contenido:

```html
{{ template "content" . }}
```

Y cada página define ese bloque:

```html
{{ define "content" }}
<h1>Home</h1>
{{ end }}
```

---

## 🔐 ¿Por qué usamos `html/template`?

Porque:

* Escapa automáticamente HTML
* Previene vulnerabilidades XSS
* Está diseñada para renderizar contenido web seguro

No usamos `text/template` porque no tiene estas protecciones.

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
http://localhost:8080/
http://localhost:8080/about
```

---

## 📌 Qué estamos aprendiendo realmente

En esta etapa introducimos:

* Renderizado del lado del servidor
* Reutilización de layout
* Organización de vistas
* Separación estructural entre contenido y estructura

Este es el punto donde el backend deja de ser solo un servidor de archivos y se convierte en un generador de vistas.
