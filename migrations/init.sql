-- Создание таблицы марок автомобилей
CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Создание таблицы моделей автомобилей
CREATE TABLE models (
    id SERIAL PRIMARY KEY,
    brand_id INT NOT NULL REFERENCES brands(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    UNIQUE (brand_id, name)  -- уникальная модель в рамках одной марки
);

-- Создание таблицы автомобилей
CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    model_id INT NOT NULL REFERENCES models(id) ON DELETE CASCADE,
    year INT NOT NULL CHECK (year > 1900 AND year <= EXTRACT(YEAR FROM CURRENT_DATE)),
    price_thousands NUMERIC(10,2) NOT NULL CHECK (price_thousands > 0)
);
