# Todo App

Приложение для управления задачами (Todo), построенное на Go с использованием PostgreSQL в качестве базы данных. Проект запускается в Docker-контейнерах.

## Требования

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Запуск проекта

### 1. Клонируйте репозиторий

```bash
git clone https://github.com/Shammaaaa/TODO-APP.git
cd todo-app
```

### 2.Настройка переменных окружения
#### Создайте файл .env в корне проекта и заполните его следующими переменными:
```
POSTGRES_USER=user
POSTGRES_PASSWORD=pass
POSTGRES_DB=Tasks
DB_HOST=db
DB_PORT=5432
```
#### Пример для config.yml уже есть
### 3.Запуск контейнеров
#### Запустите контейнеры с помощью Docker Compose:
```bash
docker-compose up --build
```

#### После запуска приложение будет доступно на http://localhost:3003

### 4. Проверка
#### Используйте API для создания, чтения, обновления и удаление задач.
## API Endpoints
### Создание задачи
### Метод: POST
### URL: /tasks
### Тело запроса:
```json
{
  "title": "задача №1",
  "description": "Описание задачи",
  "status": "done"
}
```
### Ответ:
```json
{
  "id": 1,
  "title": "задача №1",
  "description": "Описание задачи",
  "status": "done",
  "created_at": "2023-02-22T12:00:00Z",
  "updated_at": "2025-02-22T12:00:00Z"
}
```
### Получение задач
### Метод: GET
### URL: /tasks

### Ответ:
```json
{
  "id": 1,
  "title": "задача №1",
  "description": "Описание задачи",
  "status": "done",
  "created_at": "2023-02-22T12:00:00Z",
  "updated_at": "2025-02-22T12:00:00Z"
}
```
### Обновление задачи
### Метод: PUT
### URL: /tasks/:id
### Тело запроса:
```json
{
  "title": "задача №1",
  "description": "Описание задачи",
  "status": "done"
}
```
### Ответ:
```json
{
  "id": 1,
  "title": "задача №1",
  "description": "Описание задачи",
  "status": "done",
  "created_at": "2023-02-22T12:00:00Z",
  "updated_at": "2025-02-22T12:00:00Z"
}
```
### Удаление задачи
### Метод: DELETE
### URL: /tasks/:id

### Ответ: 204 No Content


