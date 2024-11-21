# Лабораторная работа №6: Анализ алгоритмов сортировки

В данной лабораторной работе были реализованы и протестированы четыре алгоритма сортировки:
- Сортировка выбором (Selection Sort)
- Пузырьковая сортировка (Bubble Sort)
- Быстрая сортировка (Quick Sort)
- Сортировка слиянием (Merge Sort)

**Цель анализа:** оценка производительности каждого алгоритма в зависимости от размера массива, порядка элементов и количества операций (сравнений и перестановок), сравнение практической сложности с теоретической и определение оптимальных сценариев применения.

---

## 1. Зависимость времени выполнения от количества и порядка элементов

### 1.1 Сортировка выбором (Selection Sort)
**Описание:** Итеративно находит минимальный элемент из оставшихся и помещает его в начало массива.

**Теоретическая сложность:**
- Время выполнения: `O(n^2)` (два вложенных цикла)
- Затраты памяти: `O(1)` (не требует дополнительной памяти)

**Анализ результатов:**

| Размер массива | Порядок      | Среднее количество сравнений | Среднее количество перестановок | Среднее время (нс) |
|----------------|--------------|------------------------------|---------------------------------|--------------------|
| 1000           | Ascending    | 486,991                      | 0                               | 682,259            |
| 1000           | Descending   | 483,648                      | 491                             | 629,910            |
| 1000           | Random       | 482,866                      | 977                             | 727,312            |
| 10,000         | Ascending    | 37,251,556                   | 0                               | 53,093,247         |
| 10,000         | Descending   | 37,258,498                   | 4,316                           | 48,023,052         |
| 10,000         | Random       | 37,243,715                   | 8,622                           | 53,215,278         |
| 50,000         | Ascending    | 329,283,148                  | 0                               | 453,639,825        |
| 50,000         | Descending   | 328,005,900                  | 12,806                          | 429,853,091        |
| 50,000         | Random       | 328,518,916                  | 25,624                          | 468,834,102        |
| 1,000,000      | Ascending    | 536,821,761                  | 0                               | 728,027,467        |
| 1,000,000      | Descending   | 536,821,761                  | 16,383                          | 679,392,004        |
| 1,000,000      | Random       | 536,821,761                  | 32,755                          | 733,608,587        |

**Выводы:**
- Время выполнения растет квадратично с увеличением размера массива, независимо от порядка элементов.
- Алгоритм эффективно обрабатывает уже отсортированные массивы (минимальное количество перестановок).

### 1.2 Пузырьковая сортировка (Bubble Sort)
**Описание:** Попарно сравнивает элементы массива, перемещая большие элементы к концу.

**Теоретическая сложность:**
- Время выполнения: `O(n^2)` (два вложенных цикла)
- Затраты памяти: `O(1)` (не требует дополнительной памяти)

**Анализ результатов:**

| Размер массива | Порядок      | Среднее количество сравнений | Среднее количество перестановок | Среднее время (нс) |
|----------------|--------------|------------------------------|---------------------------------|--------------------|
| 1000           | Ascending    | 484,426                      | 0                               | 672,244            |
| 1000           | Descending   | 484,237                      | 484,237                         | 803,865            |
| 1000           | Random       | 480,889                      | 242,305                         | 1,088,055          |
| 10,000         | Ascending    | 37,129,173                   | 0                               | 46,780,507         |
| 10,000         | Descending   | 37,172,251                   | 37,172,251                      | 61,206,515         |
| 10,000         | Random       | 37,324,246                   | 18,709,765                      | 103,011,344        |
| 50,000         | Ascending    | 329,663,574                  | 0                               | 415,329,896        |
| 50,000         | Descending   | 328,918,142                  | 328,918,142                     | 570,990,656        |
| 50,000         | Random       | 328,894,401                  | 164,095,902                     | 1,200,617,711      |
| 1,000,000      | Ascending    | 536,821,761                  | 0                               | 676,147,444        |
| 1,000,000      | Descending   | 536,821,761                  | 536,821,761                     | 882,592,526        |
| 1,000,000      | Random       | 536,821,761                  | 268,824,238                     | 1,985,997,135      |

**Выводы:**
- Алгоритм показывает худшую производительность при случайном или убывающем порядке элементов.
- На уже отсортированных массивах количество перестановок равно нулю, что улучшает производительность.

### 1.3 Быстрая сортировка (Quick Sort)
**Описание:** Рекурсивно делит массив на подмассивы вокруг опорного элемента.

**Теоретическая сложность:**
- Время выполнения: `O(n log n)` (в среднем)
- Затраты памяти: `O(log n)`

**Анализ результатов:**

| Размер массива | Порядок      | Среднее количество сравнений | Среднее количество перестановок | Среднее время (нс) |
|----------------|--------------|------------------------------|---------------------------------|--------------------|
| 1000           | Ascending    | 8,381                        | 7,870                           | 225,464            |
| 1000           | Descending   | 8,337                        | 7,826                           | 199,442            |
| 1000           | Random       | 11,225                       | 10,571                          | 391,164            |
| 10,000         | Ascending    | 100,355                      | 95,821                          | 2,355,649          |
| 10,000         | Descending   | 100,119                      | 95,602                          | 2,066,120          |
| 10,000         | Random       | 134,594                      | 128,866                         | 3,369,944          |
| 50,000         | Ascending    | 342,865                      | 326,482                         | 6,780,317          |
| 50,000         | Descending   | 342,465                      | 326,082                         | 7,307,897          |
| 50,000         | Random       | 456,338                      | 439,257                         | 11,348,813         |
| 1,000,000      | Ascending    | 442,369                      | 425,986                         | 8,737,927          |
| 1,000,000      | Descending   | 442,369                      | 425,986                         | 9,853,186          |
| 1,000,000      | Random       | 600,451                      | 578,602                         | 17,544,783         |

**Выводы:**
- Производительность быстрой сортировки зависит от выбора опорного элемента: случайный порядок элементов может увеличить количество сравнений и перестановок.
- Наиболее стабильная производительность достигается для уже упорядоченных массивов.

### 1.4 Сортировка слиянием (Merge Sort)
**Описание:** Рекурсивно делит массив на половины, сортирует их и объединяет.

**Теоретическая сложность:**
- Время выполнения: `O(n log n)`
- Затраты памяти: `O(n)`

**Анализ результатов:**

| Размер массива | Порядок      | Среднее количество сравнений | Среднее количество перестановок | Среднее время (нс) |
|----------------|--------------|------------------------------|---------------------------------|--------------------|
| 1000           | Ascending    | 4,836                        | 9,824                           | 84,088             |
| 1000           | Descending   | 4,983                        | 9,811                           | 74,345             |
| 1000           | Random       | 8,560                        | 9,822                           | 124,328            |
| 10,000         | Ascending    | 55,484                       | 112,919                         | 864,288            |
| 10,000         | Descending   | 57,609                       | 113,225                         | 764,370            |
| 10,000         | Random       | 101,780                      | 112,838                         | 1,433,657          |
| 50,000         | Ascending    | 183,734                      | 377,968                         | 2,639,658          |
| 50,000         | Descending   | 193,819                      | 377,219                         | 2,679,289          |
| 50,000         | Random       | 343,265                      | 377,100                         | 4,589,993          |
| 1,000,000      | Ascending    | 245,745                      | 491,504                         | 2,995,481          |
| 1,000,000      | Descending   | 245,759                      | 491,504                         | 3,082,242          |
| 1,000,000      | Random       | 450,073                      | 491,504                         | 5,597,090          |

**Выводы:**
- Алгоритм стабильно демонстрирует производительность с логарифмической зависимостью от размера массива.
- Убывающий и случайный порядок увеличивают время выполнения из-за более сложного объединения.

---

## 2. Сравнение теоретической и практической сложности

| Алгоритм           | Теоретическая сложность | Практическая сложность (по времени)               |
|--------------------|-------------------------|---------------------------------------------------|
| **Selection Sort** | O(n^2)                  | Соответствует теоретической                       |
| **Bubble Sort**    | O(n^2)                  | Соответствует теоретической                       |
| **Quick Sort**     | O(n log n)              | Приближается к O(n log n), но зависит от порядка  |
| **Merge Sort**     | O(n log n)              | Соответствует теоретической                       |

---

## 3. Сравнение алгоритмов и их применение

   ### Selection Sort
   - **Преимущества**: Простота реализации. Эффективен для небольших массивов.
   - **Недостатки**: Низкая производительность на больших массивах.
   - **Оптимальное применение**: Небольшие массивы, где требуется минимальное использование памяти.

   ### Bubble Sort
   - **Преимущества**: Простота. Позволяет выявлять уже отсортированные массивы.
   - **Недостатки**: Низкая производительность для больших и случайных массивов.
   - **Оптимальное применение**: Маленькие массивы с почти упорядоченными данными.

   ### Quick Sort
   - **Преимущества**: Высокая производительность в среднем. Низкие затраты памяти.
   - **Недостатки**: Чувствителен к выбору опорного элемента.
   - **Оптимальное применение**: Большие массивы, особенно при среднем или случайном порядке данных.

   ### Merge Sort
   - **Преимущества:** Гарантированная производительность O(n log n) для любого массива.
   - **Недостатки**: Требует дополнительной памяти.
   - **Оптимальное применение**: Большие массивы, где требуется стабильная производительность независимо от порядка данных.

---

## 4. Анализ зависимости времени выполнения от входных данных

1. **Алгоритм.** Алгоритмы `Selection Sort` и `Bubble Sort` имеют квадратичную сложность `O(n^2)`. Это приводит к значительному увеличению времени выполнения для больших массивов.  
   `Quick Sort` и `Merge Sort`, обладающие сложностью `O(n log n)`, показывают стабильную производительность даже на больших массивах.

2. **Размер массива.** Время выполнения всех алгоритмов увеличивается с ростом размера массива. Для `Selection Sort` и `Bubble Sort` зависимость близка к квадратичной, тогда как для `Quick Sort` и `Merge Sort` наблюдается логарифмический рост.

3. **Расположение элементов.**
   - **`Selection Sort`:** Зависит от числа перестановок, минимальное время наблюдается для отсортированных массивов.
   - **`Bubble Sort`:** Наилучшее время достигается для отсортированных массивов (отсутствие перестановок). Худшее — для убывающих массивов.
   - **`Quick Sort`:** Производительность зависит от выбора опорного элемента. Случайный порядок может увеличивать количество сравнений.
   - **`Merge Sort`:** Независим от расположения элементов и демонстрирует стабильное время выполнения.

---

## 5. Сравнение алгоритмов со сложностью `O(n log n)`

| Алгоритм       | Плюсы                                                               | Минусы                                                                                  |
|----------------|---------------------------------------------------------------------|-----------------------------------------------------------------------------------------|
| **Quick Sort** | - Высокая производительность на средних и случайных данных.        | - Чувствителен к выбору опорного элемента.                                             |
|                | - Низкие затраты памяти.                                           | - В худшем случае может деградировать до `O(n^2)` (для уже отсортированных данных).     |
| **Merge Sort** | - Гарантированная сложность `O(n log n)` для любых данных.         | - Требует дополнительной памяти `O(n)`.                                                |
|                | - Стабильность, независимость от порядка входных данных.           | - Сложнее в реализации.                                                                |

**Вывод:**
- Быстрая сортировка предпочтительна для массивов с случайным порядком или данными среднего размера, где можно оптимизировать выбор опорного элемента.
- Сортировка слиянием стабильно работает на всех типах данных, но её использование оправдано, если дополнительные затраты памяти не являются проблемой.

---

## 6. Заключение

1. Алгоритмы `Selection Sort` и `Bubble Sort` подходят только для небольших массивов из-за их квадратичной сложности.
2. `Quick Sort` показывает лучшую производительность на больших массивах, но требует оптимального выбора опорного элемента.
3. `Merge Sort` универсальна, подходит для всех типов данных и обеспечивает стабильную производительность.
4. Выбор алгоритма зависит от размера массива, расположения данных и требований к ресурсам памяти.