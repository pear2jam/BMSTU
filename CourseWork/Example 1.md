Q = ```Сегодня я хочу позаниматься курсовой```

Запрос к LLM:
```
Постановка задачи:
Определи - это изменение планов, или получение информации

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
```Изменение```

Запрос к LLM:
```
Постановка задачи:
Составь SQL запрос для следущей задачи:

# Информация о календаре

Эмбеддинги:
type1: 4
type2: 25
type3: 93

Пример 1:
Запрос:
################
Ответ:
################

Пример 2:
Запрос:
################
Ответ:
################

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