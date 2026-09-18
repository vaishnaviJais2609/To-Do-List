# Todo List API

A REST API for managing Todo items built using Go.

## Base URL

```text
https://to-do-list-2-ekcw.onrender.com
```

## Todo Model

```json
{
  "id": 1,
  "title": "Complete Task",
  "completed": false
}
```

## API Endpoints

### 1\. Create Todo

**Method:** `POST`

**Endpoint:**

```text
/todos
```

**Full URL:**

```text
https://to-do-list-2-ekcw.onrender.com/todos
```

**Required headers:**

```text
Content-Type: application/json
```

**Request body:**

```json
{
  "title": "Complete Task",
  "completed": false
}
```

**Required parameter:**

* `title` — string, required
* `completed` — boolean

**Example response:**

```json
{
  "id": 1,
  "title": "Complete Task",
  "completed": false
}
```

\---

### 2\. Get All Todos

**Method:** `GET`

**Endpoint:**

```text
/todos
```

**Full URL:**

```text
https://to-do-list-2-ekcw.onrender.com/todos
```

**Parameters:** None

**Request body:** None

**Example response:**

```json
\[
  {
    "id": 1,
    "title": "Complete Task",
    "completed": false
  }
]
```

\---

### 3\. Get Todo by ID

**Method:** `GET`

**Endpoint:**

```text
/todos/:id
```

**Example URL:**

```text
https://to-do-list-2-ekcw.onrender.com/todos/1
```

**Required parameter:**

* `id` — Todo ID in the URL

**Request body:** None

\---

### 4\. Update Todo

**Method:** `PUT`

**Endpoint:**

```text
/todos/:id
```

**Example URL:**

```text
https://to-do-list-2-ekcw.onrender.com/todos/1
```

**Required parameter:**

* `id` — Todo ID in the URL

**Required headers:**

```text
Content-Type: application/json
```

**Request body:**

```json
{
  "title": "Complete Updated Task",
  "completed": true
}
```

\---

### 5\. Delete Todo

**Method:** `DELETE`

**Endpoint:**

```text
/todos/:id
```

**Example URL:**

```text
https://to-do-list-2-ekcw.onrender.com/todos/1
```

**Required parameter:**

* `id` — Todo ID in the URL

**Request body:** None



