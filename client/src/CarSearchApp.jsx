import React, { useState } from "react";

export function CarSearchApp() {
  const [query, setQuery] = useState("");
  const [field, setField] = useState("brand");
  const [cars, setCars] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const [hasSearched, setHasSearched] = useState(false);

  const searchCars = async () => {
    const trimmedQuery = query.trim();
    if (!trimmedQuery) return;

    setHasSearched(true);
    setLoading(true);
    setError("");
    setCars([]);

    const fetchUrl = `/api/search?q=${encodeURIComponent(trimmedQuery)}&field=${field}`;

    try {
      const res = await fetch(fetchUrl);

      if (!res.ok) {
        throw new Error((await res.text()) || `Ошибка сети: ${res.status}`);
      }

      let data = await res.json();

      if (!data) {
        data = [];
      }

      if (!Array.isArray(data)) {
        console.error("API response format error:", data);
        throw new Error(
          "Некорректный формат ответа от сервера. Ожидается массив.",
        );
      }

      setCars(data);
    } catch (err) {
      setError(
        err.message || "Не удалось выполнить поиск. Проверьте соединение.",
      );
      setCars([]);
    } finally {
      setLoading(false);
    }
  };

  const handleKeyDown = (e) => {
    if (e.key === "Enter" && query.trim() && !loading) {
      searchCars();
    }
  };

  return (
    <div className="app-container">
      <h1 className="md-headline">🔍 Поиск автомобилей</h1>

      <div className="search-card">
        <div className="search-controls">
          <div className="md-select-container">
            <select
              value={field}
              onChange={(e) => setField(e.target.value)}
              className="md-select"
            >
              <option value="brand">Марка</option>
              <option value="model">Модель</option>
            </select>
          </div>

          <div className="md-input-container">
            <input
              type="text"
              placeholder="Введите запрос..."
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              onKeyDown={handleKeyDown}
              className="md-input"
            />
            <label className="md-input-label">Поиск</label>
          </div>
        </div>

        <button
          onClick={searchCars}
          className="md-button md-button-filled"
          // Кнопка отключена, если идет загрузка или запрос пуст
          disabled={loading || !query.trim()}
        >
          {loading ? "Поиск..." : "Искать"}
        </button>
      </div>

      {/* Отображение ошибки: сработает при 400, 500 или ошибке сети */}
      {error && <p className="md-error-message">Ошибка: {error}</p>}

      {/* Отображение результатов */}
      {cars.length > 0 && (
        <div className="md-data-table-container">
          <table className="md-data-table">
            <thead>
              <tr>
                <th>Марка</th>
                <th>Модель</th>
                <th>Год</th>
                <th>Цена (тыс.₽)</th>
              </tr>
            </thead>
            <tbody>
              {cars.map((car, idx) => (
                <tr key={idx} className="md-table-row">
                  <td>{car.brand}</td>
                  <td>{car.model}</td>
                  <td>{car.year}</td>
                  <td>{car.price_thousands}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {hasSearched && cars.length === 0 && !loading && !error && (
        <p className="md-info-message">
          Нет результатов, удовлетворяющих вашему запросу.
        </p>
      )}

      {!hasSearched && !loading && !error && (
        <p className="md-info-message">
          Введите марку или модель автомобиля в поле выше и нажмите{" "}
          <b>Искать</b> или <b>Enter</b>, чтобы увидеть список доступных машин.
        </p>
      )}
    </div>
  );
}
