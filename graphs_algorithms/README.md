# Лабораторная работа №3: Алгоритмы поиска в графах

## Введение
В данной лабораторной работе реализованы и проанализированы основные алгоритмы обхода графов:

- **Обход в ширину (Breadth-First Search, BFS)** — алгоритм для последовательного исследования всех узлов графа на определенном расстоянии от начального узла.
- **Обход в глубину (Depth-First Search, DFS)** — алгоритм для исследования графа с углублением до самого дальнего узла перед возвратом.

Работа включает:
- Построение направленных и неориентированных графов.
- Реализацию алгоритмов BFS и DFS.
- Вычисление суммы значений соседей для заданного узла.
- Анализ сложности и оптимальных сценариев применения.

---

### Особенности представления графов в памяти
#### Направленные графы:
В направленных графах узел `A` может быть соседом для узла `B`, но узел `B` не обязательно является соседом для узла `A`.  
В памяти это отражается следующим образом:
- Граф представлен в виде структуры, где каждая вершина содержит список соседей.  
  Например:
  ```plaintext
  A → [B, C]
  B → []
  C → [D]
  D → []
  ```

#### Неориентированные графы:
В неориентированных графах узлы A и B всегда являются соседями друг друга.
В памяти это отражается так:
- Каждое ребро добавляется в список соседей для обоих узлов.
  Например:
  ```plaintext
  A → [B]
  B → [A, C]
  C → [B]
  ```

---

## Реализация алгоритмов

### 1. Построение графов
#### Типы графов:
1. **Направленный граф (Directed Graph):**
    - Узлы соединены дугами, которые имеют направление.
    - Пример: граф, где A → B, A → C, C → D.

2. **Неориентированный граф (Undirected Graph):**
    - Узлы соединены ребрами без указания направления.
    - Пример: граф, где A ↔ B, B ↔ C.

---

### 2. Алгоритм обхода в ширину (BFS)
#### Описание:
- Начинает с заданного узла и посещает всех его соседей.
- Постепенно исследует соседей соседей, расширяя круг.

#### Временная сложность:
- **O(V + E)**, где `V` — количество узлов, `E` — количество ребер.

#### Применение:
- Поиск кратчайшего пути в невзвешенных графах.
- Выявление компонент связности.

#### Пример обхода:
Для графа: A → B, A → C, B → D  
_Порядок обхода:_ `A → B → C → D`

---

### 3. Алгоритм обхода в глубину (DFS)
#### Описание:
- Углубляется в граф до самого дальнего узла, затем возвращается и исследует другие ветви.

#### Временная сложность:
- **O(V + E)**, где `V` — количество узлов, `E` — количество ребер.

#### Применение:
- Поиск путей и циклов.
- Топологическая сортировка.

#### Пример обхода:
Для графа: A → B, A → C, B → D  
_Порядок обхода:_ `A → B → D → C`

---

### 4. Вычисление суммы соседей
#### Описание:
Для любого заданного узла алгоритм вычисляет сумму значений его соседей.

#### Пример:
Для узла `A` с соседями `B (2)` и `C (3)`  
_Сумма соседей:_ `2 + 3 = 5`

#### Временная сложность:
- **O(k)**, где `k` — количество соседей.

---

## Преимущества и ограничения алгоритмов

### Обход в ширину (BFS)
#### Преимущества:
- Эффективен для поиска кратчайшего пути.
- Работает во всех типах графов.

#### Ограничения:
- Требует больше памяти из-за хранения очереди.

---

### Обход в глубину (DFS)
#### Преимущества:
- Прост в реализации с использованием рекурсии.
- Эффективен для поиска циклов и глубокой проверки.

#### Ограничения:
- Риск переполнения стека в графах с глубокой вложенностью.

---

## Заключение
### Выводы:
1. Алгоритмы BFS и DFS имеют свои сильные стороны и применяются в зависимости от требований задачи.
2. BFS подходит для задач с кратчайшим путём или исследованиями уровня.
3. DFS удобен для анализа структуры графа, поиска циклов и глубокой проверки.
4. Алгоритмы были протестированы на нескольких графах (направленных и неориентированных).
5. Результаты обхода и вычисления суммы соседей корректны.