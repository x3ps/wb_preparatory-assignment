<p align="center">
  <img src="https://www.svgrepo.com/show/490914/car.svg" alt="Иконка автомобиля" width="100"/>
</p>

# 🚗 Проект "Автосалон" (WB Preparatory Assignment)

Веб-приложение для поиска автомобилей, разработанное в рамках задания для студентов.
Цель проекта — знакомство с технологическим стеком WB (**Go**, **PostgreSQL**, **React**).

Приложение предоставляет простой интерфейс поиска автомобилей по маркам и моделям.
Бэкенд на Go обрабатывает запросы, извлекает данные из PostgreSQL и возвращает результат в формате JSON.

---
## Скришноты интерфейса

<img width="2560" height="1305" alt="image" src="https://github.com/user-attachments/assets/094990c9-39c9-4bb7-9d06-a01190d8c369" />
<img width="2560" height="1305" alt="image" src="https://github.com/user-attachments/assets/f289bd2a-6279-4b21-96d6-bab00abda453" />

---
## 🛠 Технологический стек

- **Бэкенд:** Go (`net/http`)
- **Фронтенд:** React, HTML, CSS, JS
- **База данных:** PostgreSQL
- **Развертывание:** Docker, Docker Compose, Nginx

---

## 📂 Структура проекта

```bash
.
├── client/          # фронтенд на React
├── server/          # бэкенд на Go
├── deploy/          # Docker и конфигурации
├── migrations/      # SQL-миграции
└── utils/           # генерация тестовых данных
````

---

## 🚀 Начало работы

### Предварительные требования

* Docker
* Docker Compose
* Веб-браузер

### Запуск проекта

1. Клонируйте репозиторий:

   ```sh
   git clone <URL-репозитория>
   cd <папка-репозитория>
   ```

2. Настройте переменные окружения:

   ```sh
   cp .env.example .env
   ```

3. Соберите и запустите контейнеры:

   ```sh
   cd deploy
   docker-compose -f podman-compose.yml up --build
   ```

4. Откройте приложение в браузере:

   ```
   http://localhost:8080
   ```

---

## 🗄 Структура базы данных

### Таблица `brands`

| Колонка | Тип    | Описание                 |
| ------- | ------ | ------------------------ |
| id      | SERIAL | Уникальный идентификатор |
| name    | TEXT   | Название марки           |

### Таблица `models`

| Колонка  | Тип  | Описание                   |
| -------- | ---- | -------------------------- |
| id       | INT  | Уникальный идентификатор   |
| name     | TEXT | Название модели            |
| brand_id | INT  | Внешний ключ к `brands.id` |

### Таблица `cars`

| Колонка         | Тип     | Описание                   |
| --------------- | ------- | -------------------------- |
| id              | SERIAL  | Уникальный идентификатор   |
| model_id        | INT     | Внешний ключ к `models.id` |
| year            | INT     | Год выпуска                |
| price_thousands | NUMERIC | Цена в тысячах рублей      |

---

## 🧪 Тестовые данные и генерация

### Источники

* [Car Make/Model Dataset — Back4App](https://www.back4app.com/database/back4app/car-make-model-dataset)
* Иконка: [SVGRepo](https://www.svgrepo.com/svg/490914/car)

### Скрипт генерации

Файл: `utils/test_data-generator.sh`
Назначение: генерация SQL-скрипта `migrations/test_data.sql` из исходного JSON `cars.json`.

Запуск:

```sh
bash utils/test_data-generator.sh
```

После выполнения появится файл:

```
test_data.sql
```

---

## 🔌 API

### `GET /api/search`

Поиск автомобилей по базе.

**Параметры:**

* `q` (string, **обязательный**) — строка поиска
* `field` (string, **обязательный**) — `brand` или `model`

**Пример ответа (`200 OK`):**

```json
[
  {
    "brand": "Toyota",
    "model": "Camry",
    "year": 2021,
    "price_thousands": 2500.00
  },
  {
    "brand": "Toyota",
    "model": "RAV4",
    "year": 2022,
    "price_thousands": 2800.00
  }
]
```

---

## 📄 Лицензия

* Данные предоставлены [Back4App](https://www.back4app.com/)
* Иконка предоставлена [SVGRepo](https://www.svgrepo.com/)
* Код проекта предоставляется под лицензией **[Unlicense](https://unlicense.org/)**
---
