## Лексическая структура

```
VARNAME = [_|!|@|.|#][\p{L}]+
FUNCNAME = [\p{L}]*
REF_CONST = nothing
INT_CONST  = (([A-Za-z0-9]+{\d+})|\d+)
CHAR_CONST  = \"\p{L}?\"
STRING_CONST  = \'.*\'
BOOL_CONST = true | false
VAR_TYPE = int | char | bool
Сomments = {.*}
```

## Грамматика языка

```
Program -> (Func)+
Func ->  FuncHeader FuncBody 
FuncHeader -> "(" Type "[" FUNCNAME FuncParams "]" ")" 
            | "[" FUNCNAME FuncParams "]" 
FuncBody -> Statements "%%"
FuncParams -> (BasicVar)*
BasicVar -> "(" Type VARNAME ")"
Type -> VAR_TYPE | "<" Type ">"
Statements -> Statement ("," Statement)*
Statement -> "^" Expr 
           | "\\" Expr
           | Var ":=" Expr
           | "[" FUNCNAME Args "]"
           | "(" StatementTail
StatementTail -> Type VARNAME ((")" (":=" Expr)?) | (Cycle Statements "%"))
               | "&" Expr ")" Statements "%"
               | "?" Expr ")" Statements ("+++" Statements)? "%"
Cycle -> ":" Expr "," Expr ("," INT_CONST)? ")"
Args -> (Spec)+ 
Expr -> LogicalExpr ((_or_ | _xor_) LogicalExpr)*
LogicalExpr -> CompareExpr (_and_ CompareExpr)*
CompareExpr -> ArithmExpr (CmpOp ArithmExpr)?
CmpOp → _eq_ | _ne_ | _lt_ | _gt_ | _le_ | _ge_
ArithmExpr -> PowExpr (("+" | "-")  PowExpr)*
PowExpr -> Term (_pow_ PowExpr)?
Term -> Factor (("*" | "/" | _mod_) Factor)*
Factor -> (not_ | "-")? Spec
FuncCall -> "[" FUNCNAME Args "]"
Spec -> FuncCall 
      | new_ Type (VARNAME | INT_CONST) 
      | Const
      | Var 
      | "(" Expr ")"
Var -> VARNAME | "<" Spec Expr ">" .
Const → INT_CONST | CHAR_CONST | STRING_CONST | REF_CONST | BOOL_CONST 
```
