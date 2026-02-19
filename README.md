# 02 - HTTP Manual Routing

En esta etapa seguimos trabajando con **TCP puro**, sin usar `net/http`.

La diferencia con la rama anterior es que ahora comenzamos a entender y manipular la estructura real del protocolo HTTP.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Qué contiene realmente la primera línea de un request HTTP
* Qué es el método (`GET`, `POST`, etc.)
* Qué es el path (`/`, `/about`)
* Cómo funciona el routing internamente
* Cómo devolver códigos de estado correctos (200, 404, 405)

Aquí todavía no hay abstracciones.
Seguimos trabajando directamente sobre TCP.

---

## 🧠 Qué cambió respecto a 01

Antes:

* El servidor respondía siempre lo mismo.

Ahora:

* Leemos la primera línea del request
* La dividimos en partes
* Extraemos método y ruta
* Decidimos qué respuesta enviar según la ruta

Ejemplo real de request:

```
GET /about HTTP/1.1
Host: localhost:8080
User-Agent: ...
```

Estamos parseando manualmente:

* `GET`
* `/about`
* `HTTP/1.1`

---

## 🧩 Routing manual

El routing ahora es simplemente lógica condicional:

* Si la ruta es `/` → devolver Home
* Si la ruta es `/about` → devolver About
* Si no existe → devolver 404
* Si el método no es GET → devolver 405

Nada mágico.
Así funcionan los frameworks internamente.

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
http://localhost:8080/no-existe
```

También puedes probar:

```bash
curl -X POST http://localhost:8080/
```

Deberías recibir un `405 Method Not Allowed`.

---

## 📌 Limitaciones de esta etapa

* Solo soporta método GET
* No parsea el body
* No procesa headers avanzados
* No soporta keep-alive
* No hay manejo robusto de errores
