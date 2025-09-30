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

-- Демо-данные

-- Марки
INSERT INTO brands (name) VALUES
('Toyota'),
('BMW'),
('Lada');

-- Модели
INSERT INTO models (brand_id, name) VALUES
((SELECT id FROM brands WHERE name = 'Toyota'), 'Camry'),
((SELECT id FROM brands WHERE name = 'Toyota'), 'Corolla'),
((SELECT id FROM brands WHERE name = 'BMW'), 'X5'),
((SELECT id FROM brands WHERE name = 'BMW'), 'M3'),
((SELECT id FROM brands WHERE name = 'Lada'), 'Vesta'),
((SELECT id FROM brands WHERE name = 'Lada'), 'Granta');

-- Автомобили
INSERT INTO cars (model_id, year, price_thousands) VALUES
((SELECT id FROM models WHERE name = 'Camry'), 2020, 2500),
((SELECT id FROM models WHERE name = 'Corolla'), 2019, 1500),
((SELECT id FROM models WHERE name = 'X5'), 2021, 5500),
((SELECT id FROM models WHERE name = 'M3'), 2018, 4200),
((SELECT id FROM models WHERE name = 'Vesta'), 2022, 900),
((SELECT id FROM models WHERE name = 'Granta'), 2021, 600);
