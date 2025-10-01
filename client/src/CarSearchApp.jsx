import { useState } from "react";

export function CarSearchApp() {
  const [query, setQuery] = useState("");
  const [field, setField] = useState("brand");
  const [cars, setCars] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const searchCars = async () => {
    if (!query) return;
    setLoading(true);
    setError("");
    try {
      const res = await fetch(
        `/api/search?q=${encodeURIComponent(query)}&field=${field}`,
      );
      if (!res.ok) {
        throw new Error(await res.text());
      }
      const data = await res.json();
      setCars(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 p-6 flex flex-col items-center">
      <h1 className="text-2xl font-bold mb-6">🔍 Поиск автомобилей</h1>

      <div className="flex gap-2 mb-6">
        <select
          value={field}
          onChange={(e) => setField(e.target.value)}
          className="border rounded p-2"
        >
          <option value="brand">Марка</option>
          <option value="model">Модель</option>
        </select>

        <input
          type="text"
          placeholder="Введите запрос..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          className="border rounded p-2"
        />

        <button
          onClick={searchCars}
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        >
          Искать
        </button>
      </div>

      {loading && <p>Загрузка...</p>}
      {error && <p className="text-red-600">{error}</p>}

      {cars.length > 0 && (
        <table className="border-collapse border border-gray-400 w-full max-w-3xl bg-white shadow">
          <thead>
            <tr className="bg-gray-200">
              <th className="border border-gray-300 p-2">Марка</th>
              <th className="border border-gray-300 p-2">Модель</th>
              <th className="border border-gray-300 p-2">Год</th>
              <th className="border border-gray-300 p-2">Цена (тыс.$)</th>
            </tr>
          </thead>
          <tbody>
            {cars.map((car, idx) => (
              <tr key={idx} className="hover:bg-gray-100">
                <td className="border border-gray-300 p-2">{car.brand}</td>
                <td className="border border-gray-300 p-2">{car.model}</td>
                <td className="border border-gray-300 p-2">{car.year}</td>
                <td className="border border-gray-300 p-2">
                  {car.price_thousands}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
