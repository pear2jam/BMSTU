# Язык L4

## Абстрактный синтаксис  
Сформулируем абстрактный синтаксис.

Программа представляет собой набор определений функций:
```
Program → Funcs*
```
Функция состоит из заголовка и тела:
```
Funcs → FuncHeader FuncBody
```
Заголовок функции бывает 2х типов, для функции, которая
возвращает значение, и для функции, которая не вовзращает значение:
```
FuncHeader → (Type [FUNCNAME FuncParams]) | [FUNCNAME FuncParams]
```
Параметры функции — ноль и более объявлений параметров.
```
FuncParams → Var*
```
Тело функции состоит из последовательности операторов и завершается %%:
```
FuncBody → Statements %%
```
Тип — целый, символьный, логический, массив, двумерный массив:
```
Type → INT | CHAR | BOOL | ARRAY | DOUBLE_ARRAY
```
Последовательность операторов — ноль или более операторов, разделённых запятой:
```
Statements → Statement , … , Statement | ε
```
Оператор — присваивание новой переменной, вызов функции, присваивание существующей
переменной, объяавление перменной, условие с 2 ветками, условие с 1 веткой,
цикл с условием, цикл for, выход из функции, оператор-предупреждение

```
Statement → Var := Expr
          | [VARNAME VARNAME*]
          | VARNAME := Expr
          | Var
          | (? Expr) Statements +++ Statements %
          | (? Expr) Statements %
          | (& Expr) Statements %
          | (Var : Expr, Expr, Expr) Statements %
          | ^ Expr
          | \ Expr
          | ε
```
Переменные представляются как (тип имя):
```
Var → (Type VARNAME)
```
Выражение — переменная, константа, двуместная операция, одноместная операция:
```
Expr → VARNAME
     | SpecOp
     | Const
     | Expr BinOp Expr
     | UnOp Expr
```
```
SpecOp → ARRAY_ACCESS | FUNC_CALL | NEW
Const → INT_CONST | CHAR_CONST | STRING_CONST | REF_CONST | TRUE | FALSE 
BinOp → + | - | * | / | POW | AND | OR | XOR | EQ | NE | LT | GT | LE | GE | MOD
UnOp →  - | NOT
```

## Лексическая структура и синтаксис
Итоговая грамматика
```
Program → Program Func | Func
Func →  FuncHeader FuncBody
FuncHeader → (Type [VARNAME FuncParams]) | [VARNAME FuncParams]
FuncParams → FuncParams BasicVar | BasicVar
BasicVar → (Type VARNAME)
Type → BasicType | <Type>
BasicType → INTEGER | CHAR | BOOL
PointerType → BasicType | <BasicType>
FuncBody → Statements %%

Statements → Statement | Statements , Statement

Statement → ExtendedVar := Expr
          | [VARNAME Args]
          | BasicVar
          | (? Expr) Statements +++ Statements %
          | (? Expr) Statements %
          | (& Expr) Statements %
          | Cycle Statements %
          | ^ Expr
          | \ Expr  
          
Cycle → (CycleVar : Expr, Expr, INT_CONST) | (CycleVar : Expr, Expr)
CycleVar → Type VARNAME                 
Args → VARNAME | Args VARNAME

Expr → LogicalExpr
     | Expr OR LogicalExpr
     | Expr XOR LogicalExpr
     
LogicalExpr → CompareExpr | LogicalExpr AND ComareExpr
      
CompareExpr → ArithmExpr | ArithmExpr CmpOp ArithmExpr      
CmpOp → _eq_ | _ne_ | _lt_ | _gt_ | _le_ | _ge_

ArithmExpr → PowExpr | ArithmExpr AddOp PowExpr
AddOp → + | - 

PowExpr → Term | Term _pow_ PowExpr

Term → Factor | Term MulOp Factor
MulOp → * | / | MOD

Factor → NOT Spec | - Spec | Spec

Spec → [VARNAME Args] | new_ Type VARNAME | new_ Type INT_CONST | Const | Var

ExtendedVar → BasicVar | Var
Var → <Spec Expr> | VARNAME

Const → INT_CONST | CHAR_CONST | STRING_CONST | REF_CONST | TRUE | FALSE 
```
Лексическая струтктура:
```
VARNAME = [_|!|@|.|#][\p{L}]*
FUNCNAME = [\p{L}]*
REF_CONST = nothing
INT_CONST  = (([A-Za-z0-9]+{\d+})|\d+)
CHAR_CONST  = \"\p{L}?\"
STRING_CONST  = \'\.*\'
```
