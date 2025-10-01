#!/bin/bash

# Входной JSON
INPUT_JSON="cars.json"
OUTPUT_SQL="test_data.sql"

# Диапазон цен (в тысячах)
MIN_PRICE=500
MAX_PRICE=100000

# Текущий год
CURRENT_YEAR=$(date +%Y)

# 1. Вставляем бренды
jq -r '.results[].Make' "$INPUT_JSON" | sort -u | while read brand; do
    echo "INSERT INTO brands (name) VALUES ('$brand') ON CONFLICT (name) DO NOTHING;" >> "$OUTPUT_SQL"
done

# 2. Вставляем модели
jq -r '.results[] | "\(.Make)|\(.Model)"' "$INPUT_JSON" | sort -u | while IFS="|" read brand model; do
    echo "INSERT INTO models (brand_id, name) VALUES ((SELECT id FROM brands WHERE name='$brand'), '$model') ON CONFLICT DO NOTHING;" >> "$OUTPUT_SQL"
done

# 3. Для каждой уникальной модели создаём 10 автомобилей
jq -r '.results[] | "\(.Make)|\(.Model)|\(.Year)"' "$INPUT_JSON" | sort -u | while IFS="|" read brand model year; do
    for i in {1..100}; do
        # Рандомный год от year до текущего
        rand_year=$(( year + (RANDOM % (CURRENT_YEAR - year + 1)) ))
        # Рандомная цена
        price=$(( (RANDOM % (MAX_PRICE - MIN_PRICE + 1)) + MIN_PRICE ))
        echo "INSERT INTO cars (model_id, year, price_thousands) VALUES ((SELECT id FROM models WHERE name='$model' AND brand_id=(SELECT id FROM brands WHERE name='$brand')), $rand_year, $price);" >> "$OUTPUT_SQL"
    done
done

echo "✅ SQL-файл сгенерирован: $OUTPUT_SQL"
