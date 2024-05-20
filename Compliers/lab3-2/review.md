Грамматика и Лексическая структура

```
Program → Program Func | Func
Func →  FuncHeader FuncBody
FuncHeader → (Type [FUNCNAME FuncParams]) | [FUNCNAME FuncParams]
FuncParams → FuncParams BasicVar | BasicVar
BasicVar → (Type VARNAME)
Type → BasicType | <Type>
BasicType → INTEGER | CHAR | BOOL
FuncBody → Statements %%
Statements → Statement | Statements , Statement
Statement → Var := Expr
          | [FUNCNAME Args]
          | BasicVar
          | ( StatementTail
          | ^ Expr
          | \ Expr  
StatementTail → Type VARNAME ) InnerVar | & Expr) Statements % | Cycle Statements % 
          | (? Expr) Statements ElseStatement
ElseStatement → +++ Statements % | %
InnerVar → EPS | := Expr
Cycle → CycleVar : Expr, Expr CycleTail
CycleTail → , INT_CONST ) | )
CycleVar → Type VARNAME                 
Args → VARNAME | Args VARNAME
Expr → LogicalExpr
     | Expr _or_ LogicalExpr
     | Expr _xor_ LogicalExpr
LogicalExpr → CompareExpr | LogicalExpr _and_ ComareExpr
CompareExpr → ArithmExpr | ArithmExpr CmpOp ArithmExpr      
CmpOp → _eq_ | _ne_ | _lt_ | _gt_ | _le_ | _ge_
ArithmExpr → PowExpr | ArithmExpr AddOp PowExpr
AddOp → + | - 
PowExpr → Term | Term _pow_ PowExpr
Term → Factor | Term MulOp Factor
MulOp → * | / | _mod_
Factor → not_ Spec | - Spec | Spec
Spec → [VARNAME Args] | new_ Type NewTail | Const | Var
NewTail → VARNAME | INT_CONST
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
