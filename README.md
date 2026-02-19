# 03 - net/http Basics

En esta etapa dejamos de trabajar directamente con TCP y comenzamos a usar la librería estándar de Go: `net/http`.

El objetivo es entender qué problemas resuelve esta abstracción y cuánto código desaparece cuando la utilizamos.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Qué es `net/http`
* Cómo simplifica la creación de servidores HTTP
* Cómo funciona el routing básico con `HandleFunc`
* Cómo manejar métodos HTTP correctamente
* Cómo enviar respuestas y códigos de estado sin construir manualmente el protocolo

---

## 🧠 Qué cambia respecto a la rama anterior

Antes (02-http-manual-routing):

* Parseábamos manualmente la primera línea del request
* Extraíamos método y ruta
* Construíamos manualmente la respuesta HTTP
* Escribíamos headers y status line a mano

Ahora:

* `net/http` parsea automáticamente el request
* El routing se define con `http.HandleFunc`
* Los códigos de estado se manejan con `http.Error`
* No escribimos manualmente la estructura HTTP

Gran parte de la complejidad desaparece.

---

## 🧩 Routing con net/http

El routing ahora se define así:

```go
http.HandleFunc("/", homeHandler)
http.HandleFunc("/about", aboutHandler)
```

Cada función recibe:

```go
func(w http.ResponseWriter, r *http.Request)
```

Donde:

* `r` contiene toda la información del request
* `w` permite escribir la respuesta

---

## 🐳 Ejecución

El servidor escucha en el puerto 80 dentro del contenedor.

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

También puedes probar métodos no permitidos:

```bash
curl -X POST http://localhost:8080/
```

Deberías recibir un `405 Method Not Allowed`.

---

## 📌 Qué estamos aprendiendo realmente

`net/http` no es un framework externo.

Es la implementación oficial y robusta del protocolo HTTP en Go.

Nos permite:

* Evitar errores al construir respuestas
* Manejar correctamente headers
* Tener soporte para keep-alive
* Tener un servidor concurrente listo para producción
