# 08 - File DB

En esta etapa la API deja de devolver datos estáticos definidos en el código y comienza a leer información desde un archivo JSON.

Simulamos una base de datos utilizando un archivo local, introduciendo una capa de datos separada de la lógica HTTP.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Cómo leer archivos en Go usando `os.ReadFile`
* Cómo deserializar JSON con `json.Unmarshal`
* Cómo almacenar datos en memoria
* Cómo separar la capa de datos de la capa HTTP
* Cómo estructurar una API respaldada por datos persistentes

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

El archivo `teams.json` actúa como una base de datos simple.

---

## 🧠 Qué cambió respecto a la rama anterior

Antes:

* Los datos se definían directamente en el código
* Las respuestas eran completamente estáticas

Ahora:

* Los datos se cargan desde un archivo externo
* El servidor mantiene los datos en memoria
* La API responde usando información leída desde disco

Esto introduce una separación clara entre datos y transporte.

---

## 🧩 Flujo de datos

1. El servidor arranca
2. Se ejecuta `loadTeams()`
3. Se lee el archivo `data/teams.json`
4. Se convierte el JSON en structs de Go
5. Los endpoints devuelven esos datos

---

## 🔎 Endpoint principal

```
GET /api/teams
```

Respuesta:

```json
[
  { "id": 1, "name": "Barcelona" },
  { "id": 2, "name": "Atlético Madrid" }
]
```

---

## 🧠 Conceptos introducidos

* `os.ReadFile` para leer archivos
* `json.Unmarshal` para convertir JSON en structs
* Uso de slices como almacenamiento en memoria
* Inicialización de datos al iniciar el servidor

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
curl http://localhost:8080/api/ping
curl http://localhost:8080/api/teams
```

---

## 📌 Qué estamos aprendiendo realmente

En esta etapa entendemos que:

* Una API puede estar respaldada por datos externos
* No siempre se necesita una base de datos real para modelar persistencia
* La separación entre capa de datos y capa HTTP es fundamental

Este es el paso previo antes de introducir filtrado, parámetros y modelado más avanzado de recursos.
