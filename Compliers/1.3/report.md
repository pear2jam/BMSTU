% Лабораторная работа № 1.3. Объектно-ориентированный лексический анализатор
% 29 апреля 2024 г.
% Наумов Сергей ИУ9-62Б

# Цель работы
Целью данной работы является приобретение навыка реализации
лексического анализатора на объектно-ориентированном языке без применения 
каких-либо средств автоматизации решения задачи лексического анализа.

# Индивидуальный вариант
Комментарии: начинаются с «(*» или «{», заканчиваются на «*)» или «}» и могут пересекать границы строк текста.\
Целочисленные литералы: последовательности десятичных цифр.\
Дробные литералы: строки вида «digits/digits», где «digits» — последовательность десятичных цифр.\
Атрибут (для лабораторных работ 1.3 и 1.5) дробного числа — пара целых чисел (числитель и знаменатель).\

# Реализация
Будем считывать символ за символом поддерживая текущий считываемый тип лексемы.\
Реализуем несколько веток ветвеления оператора if для каждого типа текущей считываемой лексемы.\
Таким образом тело считывающего цикла программы должно отображать кортеж (текущий тип, символ) в 
новый считываемый тип\
Рассмотрим следущие типы:\
`NONE`: Означает отсутствие текущей считываемой лексемы.\
`COMMENT` Означает комментарий\
`DIGIT` Означает целое число\
`FRAC` Означает дробь\
`ERROR` Означает ошибочную лексему, не относящуюся ни к одному типу\

Для каждого отдельного типа выделим отдельный блок ветвления, в котором в зависимости от
очередного считывающего символа тип лексемы может остаться, а может установиться новый\
Также поддерживаем номер строки и диапазон символов относящихся к данной лексеме, а также
накапливаем значение самой лексемы

В рамках реализации лабораторной будем пользоваться принципами ООП

Выделим два класса используемых для решения:
`class Token`
Реализует токен, как набор полей описывающий его в тексте программы: тип, явное значение, аттрибут,
а также координаты в считываемой программе.
Также реализуем строковое представление токена для удобного вывода информации о нем.
```python
class Token:
    def __init__(self, tok_type, val, attr, start, end):
        self.tok_type = tok_type
        self.val = val
        self.attr = attr
        
        self.start, self.end = start, end
    
    def __str__(self):
        return f"{self.tok_type} | Val = {self.val} | Attr = {self.attr} | Start = ({self.start[0]}, {self.start[1]}) | Finish = ({self.end[0]}, {self.end[1]})"
```

`class Lexer`
Реализует лексический анализатор работающий по принципу описанному выше
Анализатор реализован однопроходным: метод лексического анализа представляет из себя
генератор, который не накапливает содержимое считываемой программы и раз за разом по мере
считывания программы выдает новый токен

Также имеется генератор readfile считывающий посимвольно код программы
Для визуального удобства перевод строки и пробельные символы возвращаются
как специальные строки

```python
class Lexer:
    def __init__(self, filename):
        self.filename = filename
        
        self.DIGITS = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
        self.pos = 0
        self.line = 1

    def read_file(filename):
        with open(filename, 'r') as file:
            for line in file:
                for char in line:
                    if char == '\n':
                        yield 'LINEEND'
                    elif char == ' ':
                        yield 'SPACE'
                    else:
                        yield char

    def analyze(self):
        cur = 'NONE'
        cur_token = None
        for c in self.read_file(self.filename):
            # обновляем self.pos и self.line
            if c == 'LINEEND':
                self.line += 1
                self.pos = 0
            else:
                self.pos += 1

            if cur == 'NONE' and (c == 'SPACE' or c == 'LINEEND'): continue

            # Если пока нет считываемого токена
            if cur == 'NONE':
                if c == '{':
                    cur = 'COMMENT'
                    cur_token = Token('COMMENT', '{', '', [self.line, self.pos], 0)
                elif c == '(':
                    cur = 'COMMENT'
                    cur_token = Token('COMMENT', '(', '', [self.line, self.pos], 0)
                elif c in self.DIGITS:
                    cur = 'DIGIT'
                    cur_token = Token('DIGIT', c, c, [self.line, self.pos], 0)
                else:
                    cur = 'ERROR'
                    cur_token = Token('ERROR', c, c, [self.line, self.pos], 0)

            # Если считываем комментарий
            elif cur == 'COMMENT':
                if cur_token.val == '(' and c != '*':
                    cur = 'ERROR'
                    cur_token = Token('ERROR', cur_token.val + c, cur_token.val, cur_token.start, cur_token.end)
                elif cur_token.val[0] == '{' and c == '}':
                    cur_token.val += c
                    cur_token.end = [self.line, self.pos]
                    yield cur_token
                    cur = 'NONE'
                elif len(cur_token.val) >= 3 and cur_token.val[0:2] == '(*' and cur_token.val[-1] == '*' and c == ')':
                    cur_token.val += c
                    cur_token.attr = cur_token.attr[1:-2]
                    cur_token.end = [self.line, self.pos]
                    yield cur_token
                    cur = 'NONE'
                elif c == 'SPACE':
                    cur_token.val += ' '
                    cur_token.attr += ' '
                elif c == 'LINEEND':
                    cur_token.val += '\\n'
                    cur_token.attr += '\\n'
                else:
                    cur_token.val += c
                    cur_token.attr += c

            # Если считываем число
            elif cur == 'DIGIT':
                if c == 'SPACE' or c == 'self.LINEEND':
                    cur_token.end = [self.line, self.pos]
                    yield cur_token
                    cur = 'NONE'
                elif c == '/':
                    cur = 'FRAC'
                    cur_token = Token('FRAC', cur_token.val + '/', [cur_token.val, ''], cur_token.start, cur_token.end)
                elif c not in self.DIGITS:
                    cur = 'ERROR'
                    cur_token = Token('ERROR', cur_token.val + c, cur_token.val + c, cur_token.start, cur_token.end)
                else:
                    cur_token.val += c
                    cur_token.attr += c
            

            # Если считываем дробь
            elif cur == 'FRAC':
                if c == 'SPACE' or c == 'self.LINEEND':
                    if cur_token.attr[1] == '':
                        cur_token = Token('ERROR', cur_token.val, cur_token.val, cur_token.start, cur_token.end)
                    cur_token.end = [self.line, self.pos]
                    yield cur_token
                    cur = 'NONE'
                elif c in self.DIGITS:
                    cur_token.val += c
                    cur_token.attr[1] += c
                else:
                    cur_token = Token('ERROR', cur_token.val + c, cur_token.val + c, cur_token.start, cur_token.end)

            # Если считываем ошибку
            elif cur == 'ERROR':
                if c == 'SPACE' or c == 'self.LINEEND':
                    cur_token.end = [self.line, self.pos]
                    yield cur_token
                    cur = 'NONE'
                else:
                    cur_token.val += c
                    cur_token.attr += c

        yield Token('EOF', '', '', [self.line, self.pos], [self.line, self.pos])
```

После чего программу можно применить к тексту программы
```python
lex = Lexer('input.txt')
tokens = lex.analyze()

for l in tokens:
    print(l)
```

# Вывод
В результате выполнения лабораторной работы, я углубил свои знания в области лексического анализа 
и объектно-ориентированного программирования. 
В частности, я научился реализовывать однопроходный лексический анализатор, который способен распознавать
различные типы лексем, такие как комментарии, целочисленные и дробные литералы.

Интересные моменты работы включали в себя изучение принципов работы лексического анализатора,
разработку эффективного алгоритма обработки текста программы и использование 
объектно-ориентированного подхода для создания классов Token и Lexer.

В результате работы я создал лексический анализатор, способный выделять различные
типы лексем в программном коде, что позволяет последующим модулям программы работать 
с этими лексемами для выполнения различных задач, таких как синтаксический анализ или выполнение программы.






