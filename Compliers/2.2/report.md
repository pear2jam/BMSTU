% Лабораторная работа № 2.2. Абстрактные синтаксические деревья
% 19 июня 2024 г.
% Наумов Сергей ИУ9-62Б

# Цель работы

Целью данной работы является приобретение навыков составления 
грамматик и проектирования синтаксических деревьев.

# Вариант
Язык L4.
Вложенные комментарии реализовывать не нужно.

# Реализация

## Описание языка
Для начала опишем грамматику языка L4

### Лексические домены

INT_CONST  = `(([A-Za-z0-9]+{\d+})|\d+)`
CHAR_CONST  = `\"\p{L}?\"`
STRING_CONST  = `\'\.*\'`
VarName = `[_|!|@|.|#][\p{L}]*`
FunctionName = `[\p{L}]*`


### Основная структура программы

Программа состоит из последовательности определений функций, 
обозначаемых как `Program : Functions*`, функций может быть ноль или более.

#### Определение функции

Каждая функция включает в себя заголовок и тело:

- `Functions : FunctionHeader FunctionBody`

#### Заголовок функции

Заголовок функции определяет её тип и параметры. Существует 
два типа заголовков:

1. Для функций, возвращающих значение:
   
   - `FunctionHeader : (Type [FunctionName FuncParams])`
     
2. Для функций, не возвращающих значение:
   
   - `FunctionHeader : [FunctionName FuncParams]`

#### Параметры функции

Параметры функции — это объявления переменных, которые функция принимает:

- `FuncParams : Var*`

Параметров может быть ноль или более.

#### Тело функции

Тело функции содержит последовательность операторов и завершается 
специальным символом `%%`:

- `FunctionBody : Stmts %%`

#### Переменные

Переменные представлены в виде пары типа и имени:

- `Var : (Type VarName)`


#### Операторы

Операторы — это инструкции в теле функции, которые могут включать 
присваивание, вызовы функций, условные конструкции и циклы:

- `Stmts : Statement , … , Statement | eps`

`eps` означает пустую последовательность операторов.

#### Оператор присваивания

- `Var := Expr`

#### Оператор вызова функции

- `[VarName (VarName)*]`

#### Операторы выбора

В языке существует два оператора выбора

- Короткий `(? Expr) Stmts %`
- Полный `(? Expr) Stmts +++ Stmts %`

#### Циклы

В языке существует два вариант задать цикл:

- Короткий `(& Expr) Stmts %`
- Полный ```((Type VarName : Expr, Expr, INT_CONST) |
         | (Type VarName : Expr, Expr)) Stmts %```

#### Оператор завершения функции:

- `^ Expr`

#### Оператор-предупреждение

- `\ Expr`

#### Выражения

Выражения могут быть переменными, константами, унарными или бинарными операциями:

```
Expr : VarName
     | SpecOp
     | Const
     | Expr BinOp Expr
     | UnaryOp Expr
```

#### Типы данных

Язык поддерживает следующие типы данных:

- Примитивные типы: целые числа (Int), символы Unicode/ASCII
  (Char) и булевские значения (Bool)
- Ссылочные типы массивы: указатель на массив: Array  

- `Type : Int | Char | Bool | Array`

#### Специальные операции и константы

Специальные операции и константы определены следующим образом:

1. Операции
   
```
SpecOp : ARRAY_ACCESS | FUNC_CALL | NEW
BinOp : + | - | * | / | AND | OR | XOR | GT |
        LE | GE | EQ | NE | LT
UnaryOp :  - | NOT
```

2. Константы
   
`Const : INT_CONST | CHAR_CONST | STRING_CONST | REF_CONST | TRUE | FALSE`

### Опишем финальную грамматику

```
Prog : Prog FuncDeclare | FuncDeclare
FuncDeclare :  FuncTitle FunctionBody
FuncTitle : (Type [FunctionName FuncParams]) | [FunctionName FuncParams]
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

Cycle : (CycleVar : Expr, Expr, INT_CONST) | (CycleVar : Expr, Expr)
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

Const :  INT_CONST | CHAR_CONST | STRING_CONST | TRUE_CONST | FALSE_CONST
```

## Код реализации

Код реализации представяет из себя два файла

Файл с объявлениями классов различных сущностей языка `classes.py`:
Здесь мы используем тип Any из модуля typing для аннотации переменных,
когда точный тип не может быть указан.
Это делается для того, чтобы решить проблему, когда функция может 
принимать аргументы различных типов,
или когда тип переменной неизвестен во время написания кода.

ABC или Abstract Base Class из модуля abc используется для создания
абстрактных классов в Python.
Абстрактные классы служат шаблоном для других классов и содержат один
или несколько абстрактных методов,
которые должны быть реализованы в наследуемых классах.
Это будет помогает обеспечить соблюдение определенного интерфейса
или поведения в наборе классов.

```
from abc import ABC as Abstract
from typing import Any
from dataclasses import dataclass

# Класс выражения из операторов
class Stmt(Abstract): pass
# Класс типа
class Type(Abstract): pass
# Класс заголовка функции
class FunctionT(Abstract): pass
# Класс тела функции
class Function(Abstract): pass
# Класс выражения
class Expression(Abstract): pass
# Класс цикла for
class ForCycle(Abstract): pass

# Класс базового типа
class SimpleType:
    Char, Integer, Boolean = "char", "int", "bool"

# Класс типа массива
@dataclass
class ArrType:
    type: Type

# Класс переменной базового типа
@dataclass
class SimpleVar:
    var_name: Any
    type: SimpleType or Type
    
# Класс последовательности символов
@dataclass
class SeqChar(Type):
    value: str

# Класс объявления в цикле
@dataclass
class CycleVar:
    type: Type or SimpleType
    var_name: Any

# Класс полного заголовка функции
@dataclass
class FuncTLong(FunctionT):
    type: Type or SimpleType
    func_name: Any
    func_params: list[SimpleVar]

# Класс короткого заголовка функции
@dataclass
class FuncTShort(FunctionT):
    func_name: Any
    func_params: list[SimpleVar]

# Класс функции целиком
@dataclass
class Func:
    header: FunctionT
    body: Function

# Класс условного оператора
@dataclass
class IfSingle(Expression):
    condition: Expression
    then_branch: Stmt

# Класс присвоения значений массиву
@dataclass
class VarArrayElems(Expression):
    array: Expression
    index: Expression

# Класс назначения переменной
@dataclass
class Assign(Stmt):
    variable: Any
    expr: Expression

# Класс вызова функции
@dataclass
class FunctionCall(Expression):
    func_name: Any
    args: Expression

# Класс выделения памяти
@dataclass
class NewAlloc(Stmt):
    type: Type
    alloc_size: Any

# Класс объявления переменной
@dataclass
class NewVar(Stmt):
    variable: Any

# Класс полного условного оператора
@dataclass
class StmtFull(Stmt):
    then_part: Stmt
    else_part: Stmt
    head: Expression

# Класс условия цикла While
@dataclass
class WhileStmt(Stmt):
    condition: Expression
    body: Stmt

# Класс полного заголовка цикла for
@dataclass
class ForHeaderFull(ForCycle):
    for_value: CycleVar
    start: Expression
    end: Expression
    step: Any

# Класс короткого заголовка цикла for
@dataclass
class ForHeaderShort(ForCycle):
    for_val: CycleVar
    start: Expression
    end: Expression

# Класс условия цикла for
@dataclass
class ForStatement(Stmt):
    header: ForCycle
    body: Stmt

# Класс возвращаемого выражения
@dataclass
class ReturnStmt(Stmt):
    expr: Expression

# Класс предупреждения
@dataclass
class WarningStmt(Stmt):
    expr: Expression

# Класс выражения для переменной
@dataclass
class Var(Expression):
    var_name: Any

# Класс выражения для константы
@dataclass
class ConstExpr(Expression):
    value: Any
    type: Type or SimpleType

# Класс унарного выражения
@dataclass
class UnaryExpr(Expression):
    expr: Expression
    op: str

# Класс бинарного выражения
@dataclass
class BinExpr(Expression):
    left: Expression
    right: Expression
    op: str
```

В следущем файле main.py мы зададим лексические домены, грамматику,
лексические домены для пропуска и передадим это парсеру edsl

```
import parser_edsl.parser_edsl as pe
from classes import *
# Определим лексические домены 

# Числа
NUM = pe.Terminal("NUM", r"(\d+|([A-Za-z0-9]+{\d+}))", str, priority=7)

# Здесь используем \p{L} - шаблон, который соответствует любой букве
# (включая буквы различных алфавитов и символы Unicode, обозначаемые \
  как буквы)

# Переменные
VAR = pe.Terminal("VAR", r"[@.#_!][\p{L}\p{N}]*", str)
# Символы
CHAR = pe.Terminal("CHAR", r'\"\p{L}?\"', str)
# Функции
FUNC = pe.Terminal("FUNC", r"[\p{L}]*", str)
# И строки
STR = pe.Terminal("STR", r'\".*\"', str)


# Задаём грамматику

arr = ["Stmt","Expression", "Cycle", "CycleVar", "Args", "LogicExpr",\
 "CompareExpr", "Function"]
Stmt, Expr, CycleFull, CycleFullVar, Args, NLogicExpr, \
    NCompareExpr, FuncBody = [pe.NonTerminal(i) for i in arr]

arr = ["CompOp", "ArithmExpr", "ComplexVar", "AddOp", "PowExpr", "Term", \
"MulOp", "Factor", "Spec", "Const", "Var"]
CompOper, NArithmExpr,ComplexVar, AddOp, PowExpr, NTerm, NMulOp,\
    Factor, Spec, Const, Variable = [pe.NonTerminal(i) for i in arr]

arr = ["TRUE", "FALSE", "MOD", "OR", "AND", "NOT", "XOR"]
TRUE_CONST, FALSE_CONST, MOD_CONST, OR_CONST, AND_CONST, NOT_CONST, \
XOR_CONST  = [pe.NonTerminal(i) for i in arr]

arr = ["Program","Func", "FunctionT", "FuncParams", "SimpleVar", "Type",\
 "SimpleType", "Stmts", "FunctionCall"]
Prog, FuncD, FuncT, FuncParams, SimpleVariable, Type, \
    NewBasicType, Stmts, FuncCall = [pe.NonTerminal(i) for i in arr]

INT_CONST, CHAR_CONST, BOOL_CONST = [pe.Terminal(i, i, lambda\
    name: None, priority=10) for i in ['char', 'int', 'bool']]

Prog |= FuncD, lambda statement: [statement]
Prog |= Prog, FuncD, lambda fncs, fn: fncs + [fn]

FuncD |= FuncT, FuncBody, Func

FuncT |= "[", FUNC, FuncParams, "]", FuncTShort
FuncT |= "(", Type, "[", FUNC, FuncParams, "]", ")", \
    FuncTLong

FuncBody |= Stmts, "%%"

FuncParams |= SimpleVariable, lambda statement: [statement]
FuncParams |= FuncParams, SimpleVariable, \
    lambda vars, var: vars + [var]

SimpleVariable |= "(", Type, VAR, ")", SimpleVar

Type |= NewBasicType
Type |= "<", Type, ">", ArrType

NewBasicType |= INT_CONST, lambda: SimpleType.Integer
NewBasicType |= CHAR_CONST, lambda: SimpleType.Char
NewBasicType |= BOOL_CONST, lambda: SimpleType.Boolean

Stmts |= Stmt, lambda statement: [statement]
Stmts |= Stmts, ",", Stmt, lambda \
    sts, statement: sts + [statement]

Stmt |= "^", Expr, ReturnStmt
Stmt |= "\\", Expr, WarningStmt
Stmt |= ComplexVar, ":=", Expr, Assign
Stmt |= "[", FUNC, Args, "]", FunctionCall
Stmt |= SimpleVariable, NewVar
Stmt |= "(", "?", Expr, ")", Stmts, "+++", \
    Stmts, "%", StmtFull
Stmt |= "(", "?", Expr, ")", Stmts, "%", \
    IfSingle
Stmt |= "(", "&", Expr, ")", Stmts, "%", \
    WhileStmt
Stmt |= CycleFull, Stmts, "%", ForStatement

CycleFull |= "(", CycleFullVar, ":", Expr, ",", Expr, ",", \
    NUM, ")", ForHeaderFull
CycleFull |= "(", CycleFullVar, ":", Expr, ",", Expr, ")", \
    ForHeaderShort
CycleFullVar |= Type, VAR, CycleVar

Args |= VAR, lambda vn: [vn]
Args |= Args, VAR, lambda args, arg: args + [arg]

for oper in ["_gt_", "_le_", "_ge_", "_eq_", "_ne_", "_lt_"]:\
    CompOper |= oper, lambda: oper

Expr |= NLogicExpr
Expr |= Expr, OR_CONST, NLogicExpr, BinExpr
Expr |= Expr, XOR_CONST, NLogicExpr, BinExpr

NLogicExpr |= NCompareExpr
NLogicExpr |= NLogicExpr, AND_CONST, NCompareExpr, BinExpr

NCompareExpr |= NArithmExpr
NCompareExpr |= NArithmExpr, CompOper, NArithmExpr, BinExpr

NArithmExpr |= PowExpr
NArithmExpr |= PowExpr, AddOp, PowExpr, BinExpr

AddOp |= "+", lambda: "+"
AddOp |= "-", lambda: "-"

PowExpr |= NTerm, "_pow_", PowExpr, lambda p, f: BinExpr\
    (p, "_pow_", f)
PowExpr |= NTerm

NTerm |= Factor
NTerm |= Factor, NMulOp, NTerm, BinExpr

FuncCall |= "[", FUNC, Args, "]", FunctionCall

NMulOp |= "/", lambda: "/"
NMulOp |= MOD_CONST, lambda: "mod"
NMulOp |= "*", lambda: "*"

Factor |= "-", Spec, lambda t: UnaryExpr("-", t)
Factor |= Spec
Factor |= NOT_CONST, Spec, lambda p: UnaryExpr("not", p)

Spec |= FuncCall
Spec |= Const
Spec |= Variable
Spec |= "new_", Type, VAR, NewAlloc
Spec |= "new_", Type, NUM, NewAlloc
Spec |= "(", Expr, ")"

Variable |= VAR, Var
Variable |= "<", Spec, Expr, ">", VarArrayElems

ComplexVar |= SimpleVariable
ComplexVar |= Variable

Const |= NUM, lambda v: ConstExpr(v, SimpleType.Integer)
Const |= CHAR, lambda v: ConstExpr(v, SimpleType.Char)
Const |= STR, SeqChar
Const |= TRUE_CONST, lambda: ConstExpr(True, SimpleType.Boolean)
Const |= FALSE_CONST, lambda: ConstExpr(False, SimpleType.Boolean)

parser = pe.Parser(Prog)

# Пропускаем комментарии и пробельные символы
parser.add_skipped_domain("\\{[^\\}]*\\}")
parser.add_skipped_domain("\s")

try:
    with open('test.txt') as f:
        tree = parser.parse(f.read())
        print(tree)
except pe.Error as err:
    print(f"Error: {err.message} in {err.pos}")
```

# Тестирование

Напишем тестовую программу на языке L4:

```
(<int> [DoubleArrayElements (<<int>> !array)])
    (int #length) := [length !array],
    (<int> #i : 0, #length - 1)
        <!array #i> := <!array #i> * 2
    %,
    ^ !array
%%

[PrintArray (<<int>> !array)]
    (int #length) := [length !array],
    (<int> #i : 0, #length - 1)
        [Print !array #i]
    %
%%
```

Программа строит абстрактное дерево

```
[Func(header=FuncTLong(type=ArrType(type='char'),
  func_name='DoubleArrayElements',
  func_params=[SimpleVar(var_name=ArrType(type=ArrType(type='char')),
    type='!array')]),
  body=[Assign(variable=SimpleVar(var_name='char', type='#length'),      
    expr=FunctionCall(func_name='length', args=['!array'])),  
  ForStatement(header=ForHeaderShort(for_val=\
  \CycleVar(type=ArrType(type='char'),
      var_name='#i'),
    start=ConstExpr(value='0',    
      type='int'),
    end=BinExpr(left=Var(var_name='#length'),
      right='-',        
      op=ConstExpr(value='1',
        type='int'))),
  body=[Assign(variable=VarArrayElems(array=Var(var_name='!array'),
    index=Var(var_name='#i')),
  expr=BinExpr(left=VarArrayElems(array=Var(var_name='!array'),
    index=Var(var_name='#i')),
    right='*',
  op=ConstExpr(value='2',   
    type='int')))]),
  ReturnStmt(expr=Var(var_name='!array'))]),
 Func(header=FuncTShort(func_name='PrintArray',
  func_params=[SimpleVar(var_name=ArrType(type=ArrType(type='char')),
  type='!array')]),
  body=[Assign(variable=SimpleVar(var_name='char', type='#length'),      
  expr=FunctionCall(func_name='length', args=['!array'])),  
  ForStatement(header=ForHeaderShort(for_val=CycleVar(type=\
  \ArrType(type='char'),
    var_name='#i'),
    start=ConstExpr(value='0',    
      type='int'),
    end=BinExpr(left=Var(var_name='#length'),
    right='-',
    op=ConstExpr(value='1',
      type='int'))),
  body=[FunctionCall(func_name='Print',
    args=['!array', '#i'])])])]
```

# Вывод
В ходе выполнения лабораторной работы были достигнуты следующие результаты:

1. **Понимание грамматик и синтаксических деревьев:**
   
   - Получены навыки составления грамматик для языка программирования L4.
   - Изучена структура абстрактных синтаксических деревьев (AST) и их роль в
     представлении структуры программы.

3. **Разработка грамматики языка L4:**
   
   - Сформулированы правила для основных элементов программы, включая функции,
     переменные, операторы и выражения.
   - Определены лексические домены

4. **Проектирование классов для элементов языка:**
   
   - Созданы классы для различных сущностей языка с использованием
     декоратора `@dataclass`

5. **Использование типа `Any` и абстрактных классов:**
   
   - Применение типа `Any` позволило обработать ситуации с
     неопределенными типами данных.
   - Абстрактные классы использовались для определения общего интерфейса
     сущностей

6. **Тестирование парсера:**
   
   - Проведено тестирование грамматики на примере тестовой программы. Парсер
     успешно и правильно по заданной грамматике построил AST дерево для программы

7. **Интересные моменты:**
    
   - В ходе поиска решения проблемы в интернете узнал об операции `|=`:
     операции присваивания-или, несмотря на то что я прекрасно знал о
     `+=`, `-=`... с этой операцией не встречался до этого, и не видел
     такого использования

8. **Заключение:**
   - В результате работы углублены знания в области компиляторостроения
   - и синтаксического анализа.
   - Хорошо познакомился с AST-деревьями, получил опыт построения
   - грамматики для простого языка.
   - Полученный опыт будет полезен при дальнейшем изучении языков
   -  программирования и их компиляторов. 
