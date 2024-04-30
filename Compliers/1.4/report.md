% Лабораторная работа № 1.4. Лексический распознаватель
% 29 апреля 2024 г.
% Наумов Сергей ИУ9-62Б

# Цель работы
Целью данной работы является изучение использования детерминированных конечных автоматов
с размеченными заключительными состояниями (лексических распознавателей) для решения задачи
лексического анализа.

# Индивидуальный вариант
real, longreal, >=, :=, комментарии начинаются с :: и продолжаются до конца строки.

# Реализация
В этой лабораторной работе нужно построить лексический анализатор, работающий
на ДКА с размеченными вершинами.

Опишем распознаваемые лексемы с помощью регулярных выражений:
- **Идентификаторы**: [a-zA-Z][a-zA-Z0-9]*\
Это выражение распознает непустые последовательности латинских букв и десятичных цифр,
начинающиеся с буквы.
- **Целочисленные литералы**: \d+\
Это выражение распознает непустые последовательности десятичных цифр.
- **Слова "real" и "longreal"**: real|longreal\
Это выражение распознает слова "real" и "longreal".
- **Знаки операций**: >= и :=: >=|:=\
Это выражение распознает знаки операций "присваивание" (:=) и "больше или равно" (>=).
- **Комментарии**: ::.*$\
Это выражение распознает комментарии, начинающиеся с "::" и продолжающиеся до конца строки.

Пропустим этап написания НДКА и сразу перейдем к написанию ДКА, так как это несложная задача

```mermaid
graph TD;
    A("0") --> |"l"| B("1")
    B --> |"o"| C("2")
    C --> |"n"| D("3")
    D --> |"g"| E("4")
    E --> |"r"| F("5")
    F --> |"e"| G("6")
    G --> |"a"| H("7")
    H --> |"l"| I("8 | longreal")

    A --> |"a-qst-z"| W("19 | Ident")

    A --> |"r"| J("9")
    J --> |"e"| K("10")
    K --> |"a"| L("11")
    L --> |"l"| M("12 | real")

    A --> |">"| N("13")
    N --> |"="| O("14 | >=")

    A --> |"\d+"| P("20 | Number")
    P --> |"\d+"| P

    A --> |":"| Q("15")
    Q --> |"="| R("16 | :=")
    Q --> |":"| S("17")
    S --> |"^newline"| S("7")
    S --> |"newline"| T("18 | Comment")

    U("S") --> |"char"| V("S+1")
    U --> |"^char"| W
    W --> |"^space newline"| W

    X("Error") --> |"^space newline"| X
    X --> |"space newline"| A
  ```
Пояснение устройства автомата:\
Вершина S относится к ветке считывания real или longreal и если мы
отклоняемся от ветки считывания этих двух слов, то переходим в ветку
считывания Ident

Вершина Error обозначает состояния автомата, при котором автомат
работает в ветке считывания ошибки, если считываемая лексема не может
быть отнесена к ветке определенного типа литерала

Логика работы автомата: если при считывании очередного символа он есть
в переходе из текущей вершины, мы обновляем вершину согласно заданным переходам.\

Если перехода из текущей вершины не сущесвутет, однако мы оказались в
размеченной вершине, автомат возвращает метку вершины как тип считанной лексемы

Если перехода из текущей вершины не существует и текущая вершина не размечена 
мы переходим в вершину Error

Код реализации:
```python
import re

class Automata:
    def __init__(self, start, finishes_data : list, transitions : dict) -> None:
        """
        start (int): Start state
        finishes_data (list(tuple)): List of finish states with their
        correspoding lexeme names - [(2, 'str'), (7, 'comment') ...]
        transitions: dict(list(tuple)): dictionary of automata transitions
        in format - dict[state] = [(state, regex_rule)]
        """

        self.start = start
        self.finishes = [i[0] for i in finishes_data]
        self.finish_lexemes = [i[1] for i in finishes_data]
        self.transitions = transitions
        if self.start in self.finishes: raise Exception('Finish states can not\
        contain Start state')
        self.s = self.start
        self.attr = ''
    
    def reset(self) -> None: 
        # Setting automata in start condition
        self.s = self.start
        self.attr = ''

    def update(self, symbol:str) -> str:
        #print(self.s, symbol)
        """
        Updating automata by reading a symbol and returning lexema name
        corresponding to reached state: ('NONE' if not finial state)
        symbol (str): symbol to update an automata
        Return (str): lexema name
        """
        if self.s == self.start: self.reset()
        self.attr += symbol

        if self.s in self.transitions:
            rules = self.transitions[self.s]
            for rule in rules:
                if rule[1].match(symbol):
                    self.s = rule[0]
                    return 'NONE', self.attr

        if self.s in self.finishes:
            res = self.finish_lexemes[self.finishes.index(self.s)]
            attr = self.attr
            self.reset()
            return res, attr[:-1]

        self.s = -1 # transition into ERROR state (-1)
        return 'NONE', ''


class Symbolizer:
    def __init__(self, text):
        self.text = text

    def next_symbol(self):
        with open(self.text, 'r') as file:
            for line in file:
                for char in line: yield char
            yield '\n'

class Lexer:
    def __init__(self, text, automata):
        self.aut = automata
        self.symbolizer = Symbolizer(text)

    def next_token(self):
        line, pos = 1, 0
        for char in self.symbolizer.next_symbol():
            pos += 1
            r = self.aut.update(char)
            if r[0] != 'NONE':
                self.aut.reset()
                self.aut.update(char)
                yield r[0], r[1], line, pos-len(r[1]), pos
            if char == '\n':
                line += 1
                pos = 0
                

automata = Automata(
    start = 0,
    finishes_data = [(8, 'LONGREAL'), (12, 'REAL'), (14, '>='), (16, ':='), \
    (17, 'COMMENT'), (19, 'IDENT'), (21, 'NUM')],
    transitions = {
        0:  [(0, re.compile('[\s\r]')), (1, re.compile('l')), (9, re.compile('r')),\
     (13, re.compile('>')), (15, re.compile(':')), (19, re.compile('[a-qst-z]')), (21, re.compile('\d+'))],
        1:  [(2, re.compile('o')), (19, re.compile('[b-z0-9]'))],
        2:  [(3, re.compile('n')), (19, re.compile('[a-mo-z0-9]'))],
        3:  [(4, re.compile('g')), (19, re.compile('[a-fh-z0-9]'))],
        4:  [(5, re.compile('r')), (19, re.compile('[a-qs-z0-9]'))],
        5:  [(6, re.compile('e')), (19, re.compile('[a-df-z0-9]'))],
        6:  [(7, re.compile('a')), (19, re.compile('[b-z0-9]'))],
        7:  [(8, re.compile('l')), (19, re.compile('[a-km-z0-9]'))],
        8:  [(19, re.compile('[a-z0-9]'))],

        9:  [(10, re.compile('e')), (19, re.compile('[a-df-z0-9]'))],
        10: [(11, re.compile('a')), (19, re.compile('[b-z0-9]'))],
        11: [(12, re.compile('l')), (19, re.compile('[a-km-z0-9]'))],
        12: [(19, re.compile('[a-z0-9]'))],

        13: [(14, re.compile('='))],

        15: [(16, re.compile('=')), (17, re.compile(':'))],

        17: [(17, re.compile('[^\n]'))],

        19: [(19, re.compile('[A-Za-z0-9]+'))],
        
        21: [(21, re.compile('\d+'))],
        
        -1: [(-1, re.compile('^\s'))],
        -1: [(0, re.compile('\s|\n'))]
    }
)


lex = Lexer('lab1-4/input.txt', automata)
#for i in lex.next_token():
#    print(i)
for i in lex.next_token(): print(f'{i[0]}{" "*(8-len(i[0]))} line: {i[2]}, pos:\
     {i[3]}-{i[4]}{" "*(5-len(str(i[3]))-len(str(i[4])))} val: {i[1]}')
```

Аналогично, как и в предыдущей лабораторной работе лексический 
анализатор является однопроходным

# Тестирование
Входная программа
```
real:=1234>=hello::!@
>= p123
reality
```
Разбор
```
REAL     line: 1, pos: 1-5    val: real
:=       line: 1, pos: 5-7    val: :=
NUM      line: 1, pos: 7-11   val: 1234
>=       line: 1, pos: 11-13  val: >=
IDENT    line: 1, pos: 13-18  val: hello
COMMENT  line: 1, pos: 18-22  val: ::!@
>=       line: 2, pos: 1-3    val: >=
IDENT    line: 2, pos: 4-8    val: p123
IDENT    line: 3, pos: 1-8    val: reality
```
# Вывод

В этой лабораторной работе я рассмотрел создание лексического анализатора,
который является первым этапом в процессе создания компилятора или интерпретатора.

Для реализации лексического анализатора я использовал конечный автомат,
состояния которого представляли ветки считывания разного типа лексем, а переходы между 
состояниями осуществлялись в зависимости от символов и правил перехода. 
Каждое состояние автомата было размечено на конечные и неконечные. 
Конечные состояния соответствовали окончанию лексемы, а неконечные продолжению
считывания символов.

В процессе выполнения лабораторной работы я также научился эффективно
тестировать лексический анализатор на различных входных данных, 
проверяя его корректность и работоспособность.

В итоге, выполнение этой лабораторной работы помогло мне лучше понять
принципы работы лексических анализаторов, освоить методы и инструменты
разработки компиляторов и интерпретаторов, а также улучшить навыки
работы с конечными автоматами.


