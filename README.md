# 01 - Raw TCP HTTP Server

Esta rama representa el **primer paso** en la construcción progresiva del servidor.

En este punto NO usamos `net/http`.
El objetivo es entender cómo funciona HTTP realmente sobre TCP.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Qué es un socket TCP
* Cómo un servidor acepta conexiones
* Cómo el navegador envía una petición HTTP
* Cómo se construye manualmente una respuesta HTTP válida
* La estructura real del protocolo HTTP

Aquí no hay frameworks.
Aquí no hay abstracciones.
Solo TCP y texto.

---

## 🧠 Qué está pasando realmente

El servidor:

1. Escucha en el puerto 80 dentro del contenedor
2. Acepta una conexión TCP
3. Lee línea por línea la petición HTTP
4. Detecta el fin de los headers (línea vacía)
5. Escribe manualmente una respuesta HTTP
6. Cierra la conexión

Ejemplo simplificado de respuesta enviada:

```
HTTP/1.1 200 OK
Content-Type: text/html
Connection: close

<html>
  <body>
    <h1>Hello from raw TCP</h1>
  </body>
</html>
```

Ese bloque de texto ES HTTP.
No hay magia.

---

## 🐳 Ejecución con Docker

El contenedor expone el puerto 80 internamente.

En `docker-compose.yml` se mapea:

```yaml
ports:
  - "8080:80"
```

Esto significa:

* Tu máquina → [http://localhost:8080](http://localhost:8080)
* Contenedor → puerto 80

---

## ▶️ Cómo ejecutar

Desde el directorio del proyecto:

```bash
docker compose up --build
```

Luego visitar:

```
http://localhost:8080
```

---

## 🔎 Qué debes observar

En la terminal verás algo como:

```
Received: GET / HTTP/1.1
Received: Host: localhost:8080
Received: User-Agent: ...
```

Eso es el navegador enviando texto plano.

HTTP es texto sobre TCP.

---

## 📌 Limitaciones de esta etapa

* No hay routing real
* No hay manejo de métodos (GET/POST)
* No hay parsing formal
* No hay manejo de errores
* No hay soporte de keep-alive
* No hay templates
