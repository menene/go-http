# 11 - REST API

En esta etapa se introduce un diseño REST completo utilizando rutas basadas en recursos.

Se deja de utilizar query parameters para acceder a elementos individuales y se adopta el uso de path parameters.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Qué significa realmente REST
* Cómo modelar recursos en una API
* Cómo utilizar path parameters (`/api/teams/1`)
* Cómo manejar múltiples métodos HTTP sobre un mismo endpoint
* Cómo estructurar operaciones CRUD

---

## 📁 Estructura del proyecto

```
.
├── main.go
├── data/
│   └── teams.json
├── Dockerfile
└── docker-compose.yml
```

La estructura se mantiene igual que en la rama anterior.

---

## 🧠 Qué cambió respecto a la rama anterior

Antes:

* Se accedía a un recurso usando query parameters (`?id=`)

Ahora:

* Se accede mediante rutas:

```
GET /api/teams/1
```

Además:

* Un mismo endpoint maneja múltiples métodos HTTP
* Se implementan operaciones CRUD completas

---

## 🧩 Endpoints disponibles

```
GET    /api/teams
GET    /api/teams/1
POST   /api/teams
PUT    /api/teams/1
DELETE /api/teams/1
```

---

## 🔎 Conceptos introducidos

* Path parameters utilizando `strings.Split`
* Diseño basado en recursos
* Manejo de múltiples métodos en un mismo handler
* Eliminación de elementos en slices
* Actualización de datos en memoria

---

## 🐳 Ejecución

El servidor continúa utilizando Docker sin cambios.

Probar con:

```bash
curl http://localhost:8080/api/teams
curl http://localhost:8080/api/teams/1
```

---

## 🧪 Pruebas con Postman

Se incluye una colección de Postman en el repositorio para facilitar las pruebas de los endpoints.

Ruta:

```
postman/collection.json
```

### Cómo usarla

1. Abrir Postman.
2. Hacer clic en **Import**.
3. Seleccionar **Upload Files**.
4. Elegir el archivo `postman/collection.json`.

Esto cargará automáticamente todos los endpoints listos para probar.

---

### Alternativa

También se pueden crear las peticiones manualmente o utilizar herramientas como:

* [https://hoppscotch.io](https://hoppscotch.io)

---

## 📌 Qué estamos aprendiendo realmente

En esta etapa se comprende que:

* REST no es una librería, es una forma de diseñar APIs
* Los recursos se representan mediante rutas
* Los métodos HTTP definen la acción sobre esos recursos
