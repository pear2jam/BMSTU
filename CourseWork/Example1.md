Q = ```Сегодня я хочу позаниматься курсовой```

Запрос к LLM:
```
Изменение планов - это намерение, желание сделать что то, планировка событий
Получение информации - вопрос о прошлом, настоящем или будущем

Определи - это изменение планов, или получение информации ориентируясь на примеры и запрос и выбери ТОЛЬКО букву, БЕЗ ПОЯСНЕНИЙ:
a) Изменение планов
b) Получение информации

Пример 1:
Запрос:
Сегодня буду делать дз
Ответ:
Изменение

Пример 2:
Запрос:
Сегодня надо сделать диффуры
Ответ:
Изменение

Запрос:
Сегодня я хочу позаниматься курсовой
Ответ:
```
Ответ LLM:
```a) Изменение планов```

Запрос к LLM:
```
Постановка задачи:
Составь SQL запрос для следущей задачи:

# Информация о календаре

Пример 1:

Эмбеддинги:
type1: 3
type2: 21
type3: 83

Запрос:
Сегодня хочу сделать домашнее задание
Ответ:
INSERT INTO table_name (column1, column2, column3, ...)
VALUES (value1, value2, value3, ...);

Пример 2:

Эмбеддинги:
type1: 3
type2: 23
type3: 88

Запрос:
Сегодня хочу сделать домашнюю работу
Ответ:
INSERT INTO table_name (column1, column2, column3, ...)
VALUES (value1, value2, value3, ...);

Эмбеддинги:
type1: 4
type2: 25
type3: 93

Запрос:
Сегодня я хочу позаниматься курсовой
Ответ:
```

Ответ LLM:
```
INSERT INTO table_name (column1, column2, column3, ...)
VALUES (value1, value2, value3, ...);
```

Ответ LLM -> Query DB
