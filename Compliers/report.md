% Лабораторная работа № 3.3 «Семантический анализ»
% 21 июня 2024 г.
% Наумов Сергей ИУ9-62Б

# Цель работы
Целью данной работы является получение навыков 
выполнения семантического анализа.

# Реализация
Возьмём ранее полученную в лабораторной 2.2 грамматику языка L4
```
Prog : Prog FuncDeclare | FuncDeclare
FuncDeclare :  FuncTitle FunctionBody
FuncTitle : (Type [FunctionName FuncParams]) |
  [FunctionName FuncParams]
FuncParams : FuncParams SimpleVariable | SimpleVariable
SimpleType : INT_CONST | CHAT_CONST | BOOL_CONST

FunctionBody : Stmts %%
SimpleVariable : (Type VarName)
Type : SimpleType | <Type>

Expr : LogicExpr
     | Expr OR LogicExpr
     | Expr XOR LogicExpr
     
LogicExpr : CompareExpr | LogicExpr AND CompareExpr

Stmts : Stmt | Stmts , Stmt

Stmt :     (? Expr) Stmts % |(? Expr) Stmts +++ Stmts % |
          (& Expr) Stmts % | Cycle Stmts % |
           | ComplexVar := Expr | [VarName Args] |
            VarName := Expr | Variable | 
           | ^ Expr | \ Expr

Cycle : (CycleVar : Expr, Expr, INT_CONST) |
  (CycleVar : Expr, Expr)
CycleVar : Type VarName                 
Args : VarName | Args VarName

CompareExpr : ArithmExpr | ArithmExpr CompOp ArithmExpr      
CompOp : _gt_ | _le_ | _ge_ | _eq_ | _ne_ | _lt_
AddOp : + | -
MulOp : * | / | MOD_CONST

ArithmExpr : PowExpr | ArithmExpr AddOp PowExpr
PowExpr : Term | Term _pow_ PowExpr

Term : Factor | Term MulOp Factor

Factor : NOT Spec | - Spec | Spec
Spec : [VarName Args] | new_ Type VarName | new_ Type INT_CONST \
\| Const | Variable

ComplexVar : SimpleVariable | Variable
Variable : <Spec Expr> | VarName

Const :  INT_CONST | CHAR_CONST | STRING_CONST |
  TRUE_CONST | FALSE_CONST
```

Подключим необходимые библиотеки

```python
import abc # Для работы с абстрактными классами и методами
from abc import ABC as Abstract
from dataclasses import dataclass
from typing import Any
import enum
import re
import string
import sys
import os


import parser_edsl_sema.parser_edsl as pe
```

В рамках семантического анализа в этой лабораторной, мы
хотим отлавливать семантические ошибки которые могут возникать
в программе, например такие как:

1) Повторное объявление переменных
   ```python
   class MultipleVariablesError(SemError):
    def __init__(self, pos, variable_name):
        self.variable_name = variable_name
        self.pos = pos
        
    @property
    def exception_text(self):
        return f'Error: Re-declaring a variable\
    {self.variable_name}'

    ```
   
2) Некорректный тип при присвоении
  ```python
  class VarWrongTypeError(SemError):
    def __init__(self, pos, lt, rt):
        self.lt = lt
        self.rt = rt
        self.pos = pos
        
    @property
    def exception_text(self):
        return f'{self.lt} is mismached with type \
{self.rt}'
  ```

3) Необъявленная переменная
   ```python
   class VarUnknownError(SemError):
    def __init__(self, pos, variable_name):
        self.variable_name = variable_name
        self.pos = pos

    @property
    def exception_text(self):
        return f'Unknown Variable: {self.variable_name}'
   ```

4) Необъявленная функция
  ```python
  class UndefinedFunctionError(SemError):
    def __init__(self, pos, funcname):
        self.pos = pos
        self.funcname = funcname
        
        
    @property
    def exception_text(self):
        return f'Undefined function {self.funcname}'

  ```

5) Повторное объявление функции
   ```python
   class FunctionRedeclarationError(SemError):
    def __init__(self, pos, funcname):
       self.funcname = funcname
        self.pos = pos

    @property
    def exception_text(self):
        return f"Function Redeclaration {self.funcname}"

   ```
  

  6) Неверное количество аргументов для вызова функции
     
   ```python
    class FunctionArgumentCountError(SemError):
    def __init__(self, pos, fun, need, actual):
        self.need = need
        self.actual = actual
        self.pos = pos
        self.fun = fun
        
    @property
    def exception_text(self):
        return (f'For call function {self.fun} needed {self.need}\
 arguments,'
                f' but have recieved {self.actual}')

   ```
  Также в определены классы для ошибок:

  1. Неправильный тип аргумента для вызова функции
  2. Вместо char - получили int
  3. Ожидался int получили не int
  4. Ожидался тип массива - получили не тип массива
  5. Типы несовместимы для данной операции
  6. Возвращаемое значение неправльного типа
  7. Несовместимый тип

В процессе разбора программы по грамматике, мы поддерживаем
статистики объявленных переменных, функций, текущий тип,
текущие принимаемые типы, необходимое кол-во переменных и так
далее, и если при проверки мы находим несоответсвие семантике
мы возвращаем ошибку при разборе.

Приведем примеры классов разбора программы, особенно тех, в
которыых возвращается семантическая ошибка


1. `BasicVar` - переменная простого типа
Класс `BasicVar`, используется для создания и валидации простых
переменных в программе. 
`create` - Это статический метод, декорированный как
 `@pe.ExAction`, что предполагает использование в парсере
или интерпретаторе.
Метод принимает атрибуты и координаты, из которых извлекает тип 
переменной и её имя, а затем создаёт новый экземпляр BasicVar 
с этими данными.

`validate` - Метод проверяет, не было ли уже объявлено переменной
с таким же именем в текущей области видимости (`vars_set`). 
Если переменная с таким именем уже существует, метод генерирует
ошибку `MultipleVariablesError`, указывая на позицию имени переменной 
и её имя. В противном случае, он добавляет новую переменную в набор 
переменных (`vars_set`), связывая имя переменной с её типом.

```python
@dataclass
class BasicVar:
    type: Type or BasicType
    name: str
    name_position: pe.Position

    @pe.ExAction
    def create(attrs, positions, res_pos):
        type_attr, name = attrs
        _, _, name_coord, _ = positions
        return BasicVar(type_attr, name, name_coord.start)

    def validate(self, funcs_set, exp_ret_type, vars_set):
        if self.name in vars_set:
            raise MultipleVariablesError(self.name_position, \
self.name)
        vars_set[self.name] = self.type
```

2.Класс `CycleNewVariable`, представляет собой структуру данных для 
управления переменными, используемыми в циклах.

`validate` - Метод выполняет валидацию переменной. Сначала он 
проверяет, соответствует ли тип переменной одному из допустимых типов
(`BasicType.Int` или `BasicType.Char`). Если тип не соответствует, метод 
генерирует ошибку `InvalidCycleVariableTypeError`, передавая позицию
типа и строковое представление типа. Затем метод проверяет, существует 
ли уже переменная с таким же именем в наборе переменных (`vars_set`).
Если да, то генерируется ошибка `MultipleVariablesError`. 
Если переменная уникальна, она добавляется в `vars_set` с её типом.

```python
@dataclass
class CycleNewVariable:
    var_type: BasicType
    type_position: pe.Position
    name: str
    name_position: pe.Position

    @pe.ExAction
    def create(attrs, coords, res_coord):

        type_attr, name = attrs
        ctype, cvariable = coords
        return CycleNewVariable(type_attr, ctype.start, name, \
cvariable.start)

    def validate(self, vars_set):
        if self.type not in (BasicType.Int, BasicType.Char):
            res_type = type_to_str(self.var_type)
            raise InvalidCycleVariableTypeError(self.type_position, \
res_type)

        if self.name in vars_set:
            raise MultipleVariablesError(self.name_position,\
 self.name)

        vars_set[self.name] = self.type
```

3. Класс UnOpExpr, представляет собой выражение с унарной
операцией в программе. Он содержит следующие атрибуты:

op: строка, представляющая оператор.
op_type: тип оператора, который может быть Type или BasicType.
coords_op: координаты оператора в исходном коде, представленные 
через pe.Position.
expr: выражение, к которому применяется унарная операция.
Метод create является статическим и создает действие (action), 
которое используется для создания экземпляра UnOpExpr во время 
парсинга кода. Это действие принимает атрибуты и координаты, 
распаковывает их и возвращает новый экземпляр UnOpExpr.

Метод validate выполняет валидацию унарной операции, проверяя
следующее:

Вызывает метод validate для валидации вложенного выражения expr.
Если оператор coords_op равен 'not_', то проверяет, что тип
expr.op_type является BasicType.Int или BasicType.Char. 
Если нет, то вызывается исключение UnaryOperationTypeError.
В противном случае устанавливает op_type как BasicType.Bool.
Если оператор coords_op равен '-', то выполняет аналогичную 
проверку и устанавливает op_type как BasicType.Int.

```python
@dataclass
class UnOpExpr(Expr):
    op: str
    op_type: Union[Type, BasicType]
    coords_op: Position
    expr: Expr
    
    @staticmethod
    def create(operation):
        @pe.ExAction
        def action(attrs, coords, res_coords):
            expr, = attrs
            coords_op, _ = coords
            return UnOpExpr(operation, expr.op_type, \
coords_op.start, expr)

        return action

    def validate(self, funcs_set, ret_type, vars_set):
        self.expr.validate(self.coords_op, ret_type, vars_set)
        if self.op == 'not_':
            if self.expr.op_type not in (BasicType.Int, \
BasicType.Char):
                raise UnaryOperationTypeError(self.coords_op, \
self.op, self.expr.op_type)
            else:
                self.op_type = BasicType.Bool
        elif self.op == '-':
            if self.expr.op_type not in (BasicType.Int, \
BasicType.Char):
                raise UnaryOperationTypeError(self.coords_op,\
 self.op, self.expr.op_type)
            else:
                self.op_type = BasicType.Int
```

4. Класс AssignStmt, представляет собой оператор присваивания
   в программе. Он содержит следующие атрибуты:

variable: переменная, которой может быть присвоено значение.
Это может быть элемент массива (ArrayElemByVar), обычная переменная
(Var) или базовая переменная (BasicVar).
var_coord: координаты переменной в исходном коде, представленные
через pe.Position.
expression: выражение, значение которого присваивается переменной.
Метод create используется для создания экземпляра AssignStmt 
с использованием атрибутов и координат, полученных в результате
парсинга кода. Этот метод возвращает новый экземпляр AssignStmt
с начальной позицией оператора присваивания.

Метод validate выполняет валидацию оператора присваивания, 
проверяя следующее:

Если expression является элементом массива (ArrayElemByVar), 
то сначала валидируется expression, а затем variable.
В противном случае сначала валидируется variable, а затем expression.
Проверяется соответствие типов variable и expression. Если оба 
типа являются типами массивов (ArrType), то дальнейшая 
проверка не требуется.
Если тип expression является BasicType.Char и тип variable 
является BasicType.Int, то дальнейшая проверка также не требуется.
Если типы variable и expression не совпадают, то проверяется, 
не является ли левый тип (variable.type) символом (<char>) и 
правый тип (expression.type) строковой константой (const_str). 
Если это не так, то вызывается исключение VarWrongTypeError

```python
@dataclass
class AssignStmt(Statement):
    variable: Union[ArrayElemByVar, \
Var, BasicVar]
    var_coord: Position
    expression: Expr

    @pe.ExAction
    def create_assignment(attrs, coords,\
res_coord):
        variable, expression = attrs
        _, coord_assign, _ = coords
        return AssignStmt(variable, \
coord_assign.start, expression)

    def validate_assignment(self, funcs_set, \
return_type, vars_set):
        # Валидация выражения и переменной
        if isinstance(self.expression, ArrayElemByVar):
            self.expression.validate(funcs_set,\
 return_type, vars_set)
        self.variable.validate(funcs_set, \
return_type, vars_set)
        self.expression.validate(funcs_set, \
return_type, vars_set)

        # Проверка соответствия типов
        if not (isinstance(self.variable.type, \
ArrayType) and isinstance(self.expression.type, \
ArrayType)):
            if self.variable.type != self.expression.type:
                left_type = type_to_str\
(self.variable.type)
                right_type = type_to_str\
(self.expression.type)
                if not (left_type == "<char>" \
and right_type == "const_str"):
                    raise VarWrongTypeError\
(self.var_coord, left_type, right_type)
```

5. Класс Program представляет собой структуру данных,
которая содержит список функций, определенных в программе.
Метод validate этого класса выполняет две основные задачи:

Первый проход: Он проверяет, не была ли какая-либо функция
объявлена более одного раза. Это делается путем итерации 
по списку функций self.functions и добавления имени каждой
функции в словарь functions. Если имя функции уже присутствует
в словаре, это означает, что функция была объявлена повторно,
и вызывается исключение FunctionRedeclarationError.

Второй проход: После проверки на повторное объявление, метод
итерирует по списку функций еще раз, на этот раз вызывая метод
validate каждой функции. Этот метод validate функции проверяет
корректность самой функции, используя словарь functions, который
содержит информацию обо всех других функциях, и vars_set, который
представляет собой множество переменных, используемых в функции.

Эти шаги обеспечивают, что каждая функция в программе уникальна 
и что каждая функция сама по себе корректна в 
соответствии с правилами, определенными в ее методе validate.

```python
@dataclass
class Program:
    functions: list[Func]
    
    def validate(self):
        functions = {}
        # Первый проход: проверка \
на повторное объявление функций
        for func in self.functions:
            func_name = func.header.func_name
            func_name_pos = func.header.\
func_name_pos
            if func.header.func_name in functions:
                raise FunctionRedeclarationError\
(func_name_pos, func_name)
            functions[func.header.func_name] = \
func.header
        for func in self.functions:
            vars_set = {}
            func.validate(functions, vars_set)
```

Далее формируется грамматика

Начинаем с класса Prog и дальше двигаемся к функциям
и их составляющих
```python
Prog |= Funcs, Program

Funcs |= FuncD, lambda statement: [statement]
Funcs |= Funcs, FuncD, lambda fun_list, fun: /
fun_list + [fun]

FuncD |= FuncT, NewFuncBody, Func

FuncT |= '(', Type, '[', FUNC, FuncParams, ']', ')'
FuncT |= '[', FUNC, FuncParams, ']', FunctionHeadShort.create

NewFuncBody |= Stmts, '%%', FuncBody

FuncParams |= SimpleVariable, lambda statement:/
 [statement]
FuncParams |= FuncParams, SimpleVariable, lambda/
 vars_set, var: vars_set + [var]
```

Подробное описание задание грамматики для разбора
можно найти в отчете к лабораторной 2.2

После чего мы иннициализируем парсер, и считавая
входной файл проводим семнатический анализ

```python
parser = pe.Parser(Prog)

# Добавляем в парсер домены для пропуска
parser.add_skipped_domain('\\{.*?\\}') # Комментарии
parser.add_skipped_domain('\s') # Пробелы
try:
    with open('test.txt') as f:
        tree = parser.parse(f.read())
        tree.validate()
        print('No semantic errors have found')
except pe.Error as err:
    print(f'Error {err.exception_text} in {err.pos}')
```

# Тестирование

Протестируем программу на семантические ошибки

1. Без ошибок
```
   (int [Print (<<int>> !args)] )
^ 0
%%

(<<int>> [DoubleArrayElements (<<int>> !array)])
    (int #length) := 1,
    (int #i : 0, #length - 1)
        <!array #i> := <!array #i>
    %,
    ^ !array
%%

[PrintArray (<<int>> !array)]
    (int #length) := 3,
    (int #i : 0, #length - 1)
        [Print !array]
    %
%%
```
> No semantic errors have found

2. w
```
(int [Print (<<char>> !args)] )
```
>Error (16, 16): Для вызова функции Print необходим
 аргумент типа <<char>>, но получен аргумент типа <<int>>

# Вывод
