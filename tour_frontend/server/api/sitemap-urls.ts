interface Tour {
  id: number;
  name: string;
  updatedAt: string; // или Date, в зависимости от формата
}

export default defineEventHandler(async () => {
  try {
    const response = await fetch('http://localhost:8080/api/tours');
    if (!response.ok) {
      throw new Error(`API responded with status ${response.status}`);
    }

    const tours: Tour[] = await response.json();

    return tours.map((tour: Tour) => ({
      loc: `/tour/${tour.id}`, // Шаблон маршрута для страниц туров
      lastmod: tour.updatedAt,  // Дата последнего обновления
    }));
  } catch (error) {
    console.error('Error fetching tour data:', error);
    return []; // Возвращаем пустой массив или статус ошибки, если нужно
  }
});
