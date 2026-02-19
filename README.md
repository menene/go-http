# 04 - Serve HTML Files

En esta etapa dejamos de generar HTML directamente desde el código Go y comenzamos a servir archivos HTML reales desde el sistema de archivos.

Además, introducimos el manejo de archivos estáticos como CSS e imágenes.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Cómo servir archivos HTML usando `http.ServeFile`
* Cómo servir archivos estáticos con `http.FileServer`
* Qué es `StripPrefix` y por qué es necesario
* Cómo organizar un proyecto web de manera más realista
* Separación básica entre backend y frontend

---

## 📁 Nueva estructura del proyecto

```
.
├── main.go
├── src/
│   ├── index.html
│   ├── about.html
│   ├── css/
│   │   └── styles.css
│   └── assets/
│       └── gopher.png
├── Dockerfile
└── docker-compose.yml
```

Ahora el HTML ya no está embebido en el código Go.

---

## 🧠 Qué cambió respecto a la rama anterior

Antes:

* Las respuestas HTML se generaban con `fmt.Fprint`
* Todo el contenido estaba dentro del archivo `main.go`

Ahora:

* Usamos `http.ServeFile` para enviar archivos HTML
* Usamos `http.FileServer` para servir directorios estáticos
* El CSS y las imágenes viven en carpetas separadas

El servidor ahora se comporta más como un servidor web real.

---

## 🧩 Servir archivos específicos

Para servir un archivo HTML:

```go
http.ServeFile(w, r, "./src/index.html")
```

Esto envía el archivo directamente al cliente.

---

## 📦 Servir archivos estáticos

Para servir CSS e imágenes usamos `FileServer`:

```go
css := http.FileServer(http.Dir("./src/css"))
http.Handle("/css/", http.StripPrefix("/css/", css))
```

Lo mismo para `/assets/`.

### ¿Por qué usamos `StripPrefix`?

Cuando el navegador solicita:

```
/assets/gopher.png
```

Si no eliminamos el prefijo, Go intentaría buscar:

```
./src/assets/assets/gopher.png
```

Con `StripPrefix` logramos que el path interno coincida correctamente con el sistema de archivos.

---

## 🖼 Carga de imágenes

En `about.html` ahora podemos usar:

```html
<img src="/assets/gopher.png" alt="Gopher">
```

El servidor entrega la imagen desde el directorio `src/assets`.

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

* Separación de responsabilidades
* Organización de archivos
* Manejo básico de recursos estáticos

Este es el paso previo antes de introducir plantillas dinámicas.

Ahora el servidor ya no construye HTML.
Solo lo entrega.
