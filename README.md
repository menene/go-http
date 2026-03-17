# 12 - REST DB API

En esta etapa la API deja de utilizar almacenamiento en memoria y pasa a utilizar una base de datos real (PostgreSQL).

Los datos ahora son persistentes, lo que significa que sobreviven reinicios del servidor.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Cómo conectar una aplicación Go a una base de datos PostgreSQL
* Cómo ejecutar queries SQL desde Go
* Cómo mapear resultados de la base de datos a structs
* Cómo mantener persistencia real de datos
* Cómo integrar múltiples servicios con Docker Compose

---

## 📁 Estructura del proyecto

```
.
├── main.go
├── db/
│   └── init.sql
├── docker-compose.yml
├── Dockerfile
└── postman/
    └── collection.json
```

Se agrega una carpeta `db/` que contiene el script de inicialización de la base de datos.

---

## 🧠 Qué cambió respecto a la rama anterior

Antes:

* Los datos se almacenaban en memoria
* Se perdían al reiniciar el servidor

Ahora:

* Los datos se almacenan en PostgreSQL
* Se mantienen entre reinicios
* La API interactúa con la base de datos usando SQL

---

## 🧩 Base de datos

Se utiliza PostgreSQL ejecutándose en un contenedor Docker.

El archivo `init.sql`:

* Crea la tabla `teams`
* Inserta datos iniciales (La Liga)

Este script se ejecuta automáticamente al iniciar el contenedor por primera vez.

---

## 🔎 Endpoints disponibles

```
GET    /api/teams
GET    /api/teams/1
POST   /api/teams
PUT    /api/teams/1
DELETE /api/teams/1
```

Todos los endpoints funcionan igual que en la rama anterior, pero ahora operan sobre la base de datos.

---

## 🐳 Ejecución

Levantar los servicios con:

```bash
docker compose up --build
```

Esto iniciará:

* La API en Go
* PostgreSQL

---

## 🧪 Pruebas

Se puede utilizar la colección incluida:

```
postman/collection.json
```

O probar manualmente:

```bash
curl http://localhost:8080/api/teams
```

---

## 🌐 CORS

Al consumir la API desde un frontend que corre en otro puerto (por ejemplo `http://localhost:3000`), el navegador puede bloquear las peticiones debido a la política de seguridad conocida como CORS (Cross-Origin Resource Sharing).

Un *origin* está compuesto por:

```
protocolo + host + puerto
```

Por ejemplo:

* `http://localhost:3000`
* `http://localhost:8080`

Aunque ambos estén en `localhost`, son considerados orígenes distintos.

---

### 🔎 Qué ocurre sin CORS

Si el frontend intenta hacer una petición:

```javascript
fetch("http://localhost:8080/api/teams")
```

El navegador puede bloquearla con un error como:

```
Blocked by CORS policy
```

Esto sucede aunque el backend esté funcionando correctamente.

---

### ✅ Solución

El backend debe incluir headers que indiquen que permite peticiones desde otros orígenes:

```go
w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
```

Además, para métodos como `POST`, `PUT` o `DELETE`, el navegador envía una petición previa (`OPTIONS`) que también debe ser manejada.

---

## 📌 Qué estamos aprendiendo realmente

En esta etapa se comprende que:

* Un backend real necesita persistencia
* Las bases de datos se integran mediante drivers
* SQL sigue siendo una parte fundamental del desarrollo backend
* Docker permite orquestar múltiples servicios fácilmente
* El frontend y backend requieren configuración adicional para comunicarse (CORS)
