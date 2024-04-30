% Лабораторная работа № 2.2. Абстрактные синтаксические деревья
% Наумов Сергей, ИУ9-62Б

# Цель работы
Целью данной работы является получение навыков составления грамматик и проектирования синтаксических деревьев.

# Индивидуальный вариант
Язык L4.
Вложенные комментарии реализовывать не нужно.

# Реализация
Лексическая структура языка следущая:

```
VARNAME ::= '[_|!|@|.|#][\p{L}]*'
FUNCNAME ::= '[\p{L}]*'
REF_CONST ::= 'nothing'
INT_CONST ::= '(([A-Za-z0-9]+{\d+})|\d+)'
CHAR_CONST ::= '\"\p{L}?\"'
STRING_CONST ::= '\'.*\''
```

Абстрактный синтаксис:

Программа представляет собой набор функциональных определений.\
Функция состоит из заголовка и тела.\
Заголовок функции может быть с указанием типа результата или без него.\
Параметры функции - это ноль или более объявлений аргументов.\
Тело функции состоит из последовательности инструкций и завершается специальным символом "%%".\
Типы данных могут быть целыми, символьными, логическими, массивами или двумерными массивами.\
Последовательность инструкций может включать операторы присваивания, \
вызовы функций, объявления переменных, условные операторы и циклы.\
Лексическая структура и конкретный синтаксис были уточнены следующим образом:\

Программа может содержать несколько функций или одну функцию.\
Заголовок функции может содержать тип возвращаемого результата и список аргументов.\
Аргументы функции могут быть любого типа данных.\
Тело функции состоит из последовательности операций, разделенных запятыми.\
Операции могут включать в себя присваивание переменной, вызов функции, условные операторы и циклы.\
Выражения могут быть логическими, арифметическими или специальными.\
Лексическая структура определена следующим образом:\

Имена переменных и функций могут содержать буквы, цифры и специальные символы, \
начиная с определенных символов.
Константы могут быть целыми числами, символами, строками, ссылками или\
логическими значениями.

```
NProgram → NProgram NFunc | NFunc
NFunc →  NFuncHeader NFuncBody
NFuncHeader → (NType [FUNCNAME NFuncParams]) | [FUNCNAME NFuncParams]
NFuncParams → NFuncParams NBasicVar | NBasicVar
NBasicType → KW_INT | KW_CHAR | KW_BOOL
NPointerType → NBasicType | <NBasicType>

NFuncBody → NStatements %%
NBasicVar → (NType VARNAME)
NType → NBasicType | <NType>


NCycle → (NCycleVar : NExpr, NExpr, INT_CONST) | (NCycleVar : NExpr, NExpr)
NCycleVar → NType VARNAME                 
NArgs → VARNAME | NArgs VARNAME

NExpr → NLogicalExpr
     | NExpr OR NLogicalExpr
     | NExpr XOR NLogicalExpr
     
NLogicalExpr → NCompareExpr | NLogicalExpr AND NCompareExpr

NStatements → NStatement | NStatements , NStatement

NStatement → (? NExpr) NStatements +++ NStatements % | (? NExpr) NStatements % | (& NExpr) NStatements % | NCycle NStatements % |
           | NExtendedVar := NExpr | [VARNAME NArgs] | VARNAME := NExpr | NVar | 
           | \ NExpr | ^ NExpr

NCompareExpr → NArithmExpr | NArithmExpr NCmpOp NArithmExpr      
NCmpOp → _eq_ | _ne_ | _lt_ | _gt_ | _le_ | _ge_

NArithmExpr → NPowExpr | NArithmExpr NAddOp NPowExpr
NAddOp → + | - 

NPowExpr → NTerm | NTerm _pow_ NPowExpr

NTerm → NFactor | NTerm NMulOp NFactor
NMulOp → * | / | KW_MOD

NFactor → NOT NSpec | - NSpec | NSpec

NExtendedVar → NBasicVar | NVar
NVar → <NSpec NExpr> | VARNAME

NSpec → [VARNAME NArgs] | new_ NType VARNAME | new_ NType INT_CONST | NConst | NVar

NConst →  KW_TRUE | KW_FALSE | INT_CONST | CHAR_CONST | STRING_CONST | REF_CONST
```

## Реализация на Python

Опишем нетерминалы из грамматики с помощью абстрактных dataclasses
```python
from dataclasses import dataclass
import abc
from typing import Any


class FuncHeader(abc.ABC):
    pass

class FuncBody(abc.ABC):
    pass

class Statement(abc.ABC):
    pass

class Type(abc.ABC):
    pass

class BasicType(enum.Enum):
    Char, Integer, Boolean = 'char', 'int', 'bool'

class Expr(abc.ABC):
    pass

class ForHeader(abc.ABC):
    pass

@dataclass
class CycleVar:
    type: Type or BasicType
    var_name: Any

@dataclass
class FuncHeaderFull(FuncHeader):
    type: Type or BasicType
    func_name: Any
    func_params: list[BasicVar]

@dataclass
class CharSequenceType(Type):
    value: str

@dataclass
class ArrayType:
    type: Type


@dataclass
class BasicVar:
    type: Type or BasicType
    var_name: Any


@dataclass
class FuncHeaderShort(FuncHeader):
    func_name: Any
    func_params: list[BasicVar]


@dataclass
class Func:
    header: FuncHeader
    body: FuncBody


@dataclass
class AssignStatement(Statement):
    variable: Any
    expr: Expr


@dataclass
class FuncCallStatement(Expr):
    func_name: Any
    args: Expr


@dataclass
class NewVarStatement(Statement):
    variable: Any


@dataclass
class IfStatementFull(Statement):
    condition: Expr
    then_branch: Statement
    else_branch: Statement


@dataclass
class IfStatementOneBranch(Expr):
    condition: Expr
    then_branch: Statement


@dataclass
class ArrayElemByVar(Expr):
    array: Expr
    index: Expr


@dataclass
class WhileStatement(Statement):
    condition: Expr
    body: Statement

@dataclass
class ForHeaderFull(ForHeader):
    for_val: CycleVar
    start: Expr
    end: Expr
    step: Any


@dataclass
class ForHeaderShort(ForHeader):
    for_val: CycleVar
    start: Expr
    end: Expr


@dataclass
class ForStatement(Statement):
    header: ForHeader
    body: Statement


@dataclass
class ReturnStatement(Statement):
    expr: Expr


@dataclass
class Var(Expr):
    var_name: Any


@dataclass
class ConstExpr(Expr):
    value: Any
    type: Type or BasicType


@dataclass
class BinOpExpr(Expr):
    left: Expr
    op: str
    right: Expr


@dataclass
class UnOpExpr(Expr):
    op: str
    expr: Expr
```
Опишем лексическую структуру

```python
VARNAME = pe.Terminal('VARNAME', '[_|!|@|.|#][\p{L}]*', str)
FUNCNAME = pe.Terminal('FUNCNAME', '[\p{L}]*', str)

REF_CONST = ('REF_CONST', 'nothing', str)

INTEGER = pe.Terminal('INTEGER', '(([A-Za-z0-9]+{\d+})|\d+)', str, priority=7)
CHAR = pe.Terminal('CHAR', '\"\p{L}?\"', str)
STRING = pe.Terminal('STRING', '\'.*\'', str)
```

Теперь определим синтаксис, вспомогательные функции и тело для разбора программы
```python
make_op_lambda = lambda op: lambda: op

def make_keyword(image):
    return pe.Terminal(image, image, lambda name: None, priority=10)


KW_INT, KW_CHAR, KW_BOOL = \
    map(make_keyword, 'int char bool'.split())

KW_XOR, KW_OR, KW_MOD, KW_AND, KW_NOT, KW_TRUE, KW_FALSE = \
    map(make_keyword, 'XOR OR MOD AND NOT TRUE FALSE'.split())

NProgram, NFunc, NFuncHeader, NFuncParams, NBasicVar, NType, NBasicType, NStatements, NFuncCall = \
    map(pe.NonTerminal, 'Program Func FuncHeader FuncParams BasicVar Type BasicType Statements FuncCall'.split())

NStatement, NExpr, NCycle, NCycleVar, NArgs, NLogicalExpr, NCompareExpr, NFuncBody = \
    map(pe.NonTerminal, 'Statement Expr Cycle CycleVar Args LogicalExpr CompareExpr FuncBody'.split())

NCmpOp, NArithmExpr,NExtendedVar, NAddOp, NPowExpr, NTerm, NMulOp, NFactor, NSpec, NConst, NVar = \
    map(pe.NonTerminal, 'CmpOp ArithmExpr NExtendedVar AddOp PowExpr Term MulOp Factor Spec Const Var'.split())

# грамматика
NProgram |= NFunc, lambda st: [st]
NProgram |= NProgram, NFunc, lambda fncs, fn: fncs + [fn]

NFunc |= NFuncHeader, NFuncBody, Func

NFuncHeader |= '(', NType, '[', FUNCNAME, NFuncParams, ']', ')', FuncHeaderFull
NFuncHeader |= '[', FUNCNAME, NFuncParams, ']', FuncHeaderShort

NFuncBody |= NStatements, '%%'

NFuncParams |= NBasicVar, lambda st: [st]
NFuncParams |= NFuncParams, NBasicVar, lambda vars, var: vars + [var]

NBasicVar |= '(', NType, VARNAME, ')', BasicVar

NType |= NBasicType
NType |= '<', NType, '>', ArrayType

NBasicType |= KW_INT, lambda: BasicType.Integer
NBasicType |= KW_CHAR, lambda: BasicType.Char
NBasicType |= KW_BOOL, lambda: BasicType.Boolean

NStatements |= NStatement, lambda st: [st]
NStatements |= NStatements, ',', NStatement, lambda sts, st: sts + [st]

NStatement |= '^', NExpr, ReturnStatement
NStatement |= '\\', NExpr, WarningStatement
NStatement |= NExtendedVar, ':=', NExpr, AssignStatement
NStatement |= '[', FUNCNAME, NArgs, ']', FuncCallStatement
NStatement |= NBasicVar, NewVarStatement
NStatement |= '(', '?', NExpr, ')', NStatements, '+++', NStatements, '%', IfStatementFull
NStatement |= '(', '?', NExpr, ')', NStatements, '%', IfStatementOneBranch
NStatement |= '(', '&', NExpr, ')', NStatements, '%', WhileStatement
NStatement |= NCycle, NStatements, '%', ForStatement

NCycle |= '(', NCycleVar, ':', NExpr, ',', NExpr, ',', INTEGER, ')', ForHeaderFull
NCycle |= '(', NCycleVar, ':', NExpr, ',', NExpr, ')', ForHeaderShort
NCycleVar |= NType, VARNAME, CycleVar

NArgs |= VARNAME, lambda vn: [vn]
NArgs |= NArgs, VARNAME, lambda args, arg: args + [arg]


for op in ('_eq_', '_ne_', '_lt_', '_gt_', '_le_', '_ge_'):
    NCmpOp |= op, make_op_lambda(op)

NExpr |= NLogicalExpr
NExpr |= NExpr, KW_OR, NLogicalExpr, BinOpExpr
NExpr |= NExpr, KW_XOR, NLogicalExpr, BinOpExpr

NLogicalExpr |= NCompareExpr
NLogicalExpr |= NLogicalExpr, KW_AND, NCompareExpr, BinOpExpr

NCompareExpr |= NArithmExpr
NCompareExpr |= NArithmExpr, NCmpOp, NArithmExpr, BinOpExpr

NArithmExpr |= NPowExpr
NArithmExpr |= NPowExpr, NAddOp, NPowExpr, BinOpExpr

NAddOp |= '+', lambda: '+'
NAddOp |= '-', lambda: '-'

NPowExpr |= NTerm
NPowExpr |= NTerm, '_pow_', NPowExpr, lambda p, f: BinOpExpr(p, '_pow_', f)

NTerm |= NFactor
NTerm |= NFactor, NMulOp, NTerm, BinOpExpr

NMulOp |= '*', lambda: '*'
NMulOp |= '/', lambda: '/'
NMulOp |= KW_MOD, lambda: 'mod'

NFactor |= KW_NOT, NSpec, lambda p: UnOpExpr('not', p)
NFactor |= '-', NSpec, lambda t: UnOpExpr('-', t)
NFactor |= NSpec

NFuncCall |= '[', FUNCNAME, NArgs, ']', FuncCallStatement

NSpec |= NFuncCall
NSpec |= 'new_', NType, VARNAME, NewAllocStatement
NSpec |= 'new_', NType, INTEGER, NewAllocStatement
NSpec |= NConst
NSpec |= NVar
NSpec |= '(', NExpr, ')'

NVar |= VARNAME, Var
NVar |= '<', NSpec, NExpr, '>', ArrayElemByVar

NExtendedVar |= NBasicVar
NExtendedVar |= NVar

NConst |= INTEGER, lambda v: ConstExpr(v, BasicType.Integer)
NConst |= CHAR, lambda v: ConstExpr(v, BasicType.Char)
NConst |= STRING, CharSequenceType
NConst |= KW_TRUE, lambda: ConstExpr(True, BasicType.Boolean)
NConst |= KW_FALSE, lambda: ConstExpr(False, BasicType.Boolean)

parser = pe.Parser(NProgram)
assert parser.is_lalr_one()
# parser.print_table()

# пробельные символы
parser.add_skipped_domain('\s')
# комментарии вида { … }
parser.add_skipped_domain('\\{.*?\\}')

for filename in sys.argv[1:]:
    try:
        with open(filename) as f:
            tree = parser.parse(f.read())
            pprint(tree)
    except pe.Error as e:
        print(f'Ошибка {e.pos}: {e.message}')
    except Exception as e:
        print(e)
```



# Тестирование



# Вывод
