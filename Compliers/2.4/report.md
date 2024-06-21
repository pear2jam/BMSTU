% Лабораторная работа № 2.4 «Рекурсивный спуск»
% 20 июня 2024 г.
% Наумов Сергей ИУ9-62Б

# Цель работы
Целью данной работы является изучение алгоритмов
построения парсеров методом рекурсивного спуска.

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

И на основании этой грамматики напишешем парсер, работающий по принципу
рекрсивного спуска

Реализация парсера содержится в следующих файлах:


`main.py`
Здесь происходит подключение лексера (scanner) и непосредственно
парсера.
После чего происходит считывание программы и построение дерева

```
from scanner import *
from parser import *

import sys
import os

def new_scanner(reader):
    return Scanner(reader)

def new_parser(scanner):
    return Parser(scanner)

def main():
    try:
        if len(sys.argv) < 2:
            raise Exception("usage must be: python main.py\
 <fileTag.txt>\n")
        
        file_path = sys.argv[1]

        with open(file_path, 'r') as file:
            reader = file

            scanner = new_scanner(reader)

            parser = new_parser(scanner)

            program = parser.parse_program()
            program.print_program("")

            print("COMMENTS:")
            scanner.print_comments()

    except Exception as e:
        print(e)
        os._exit(1)

if __name__ == "__main__":
    main()
```

`content.py`

В этом файле реализован класс представления фрагемента программы
с которым мы работаем на этапе лексического анализа и последующего
синтаксического анализа.

```
from position import *

class Content:
    def __init__(self, starting, following):
        self.starting = starting
        self.following = following

    def __str__(self):
        return f"{self.starting.__str__()} - \
{self.following.__str__()}"

def new_content(starting, following):
    return Content(starting, following)
```

`node.py`

Файл реализует функционал узла дерева для разных типов данных
которые учавствуют в синтаксическом разборе

```
class NodePrinter:
    def print(self, indent):
        pass

class Program:
    def __init__(self, funcs):
        self.funcs = funcs

    def print(self, indent):
        print(indent + "Program:")
        for func in self.funcs:
            func.print(indent + "\t")

class Func:
    def __init__(self, header, body):
        self.header = header
        self.body = body

    def print(self, indent):
        print(indent + "Func:")
        self.header.print(indent + "\t")
        self.body.print(indent + "\t")

class FuncHeader:
    def __init__(self, type, func_name, params):
        self.type = type
        self.func_name = func_name
        self.params = params

    def print(self, indent):
        print(indent + "FuncHeader:")
        if self.type is not None:
            print(indent + "\t" + "FuncType:")
            self.type.print(indent + "\t" + "  ")
        print(indent + "\t" + "FuncName: " + \
self.func_name.val)
        self.params.print(indent + "\t")

class FuncBody:
    def __init__(self, statements):
        self.statements = statements

    def print(self, indent):
        print(indent + "FuncBody:")
        self.statements.print(indent + "\t")

class FuncParams:
    def __init__(self, basic_vars):
        self.basic_vars = basic_vars

    def print(self, indent):
        print(indent + "FuncParams:")
        for basic_var in self.basic_vars:
            basic_var.print(indent + "\t")

class BasicVar:
    def __init__(self, type, var_name):
        self.type = type
        self.var_name = var_name

    def print(self, indent):
        print(indent + "BasicVar:")
        print(indent + "\t" + "VarType:")
        self.type.print(indent + "\t" + "  ")
        print(indent + "\t" + "VarName: " + \
self.var_name.val)

class Type:
    def print(self, indent):
        pass

    def type_func(self):
        pass

class BasicType(Type):
    def __init__(self, var_type):
        self.var_type = var_type

    def print(self, indent):
        print(indent + "BasicType: " + self.var_type.val)


class Statements:
    def __init__(self, statements):
        self.statements = statements

    def print(self, indent):
        print(indent + "Statements: ")
        for statement in self.statements:
            statement.print(indent + "\t")

class Statement:
    def print(self, indent):
        pass

    def statement_func(self):
        pass

class ReturnStatement(Statement):
    def __init__(self, expr):
        self.expr = expr

    def print(self, indent):
        print(indent + "ReturnStatement: ")
        self.expr.print(indent + "\t")

    def statement_func(self):
        pass

class WarningStatement(Statement):
    def __init__(self, expr):
        self.expr = expr

    def print(self, indent):
        print(indent + "WarningStatement: ")
        self.expr.print(indent + "\t")

    def statement_func(self):
        pass

class AssignmentStatement(Statement):
    def __init__(self, var, expr):
        self.var = var
        self.expr = expr

    def print(self, indent):
        print(indent + "AssignmentStatement: ")
        self.var.print(indent + "\t")
        self.expr.print(indent + "\t")

    def statement_func(self):
        pass

class WhileStatement(Statement):
    def __init__(self, expr, statements):
        self.expr = expr
        self.statements = statements

    def print(self, indent):
        print(indent + "WhileStatement: ")
        indent += "  "
        print(indent + "WhileCondition: ")
        self.expr.print(indent + "\t")
        print(indent + "WhileBody: ")
        self.statements.print(indent + "\t")

    def statement_func(self):
        pass

class IfStatement(Statement):
    def __init__(self, expr, then_branch, else_branch=None):
        self.expr = expr
        self.then_branch = then_branch
        self.else_branch = else_branch

    def print(self, indent):
        print(indent + "IfStatement: ")
        indent += "  "
        print(indent + "IfCondition: ")
        self.expr.print(indent + "\t")
        print(indent + "ThenBranchBody: ")
        self.then_branch.print(indent + "\t")
        if self.else_branch is not None:
            print(indent + "ElseBranchBody: ")
            self.else_branch.print(indent + "\t")

    def statement_func(self):
        pass


class Expression:
    def print(self, indent):
        pass

    def expr_func(self):
        pass


class BinaryExpression(Expression):
    def __init__(self, left, operator, right):
        self.left = left
        self.operator = operator
        self.right = right

    def print(self, indent):
        print(indent + "BinaryExpression: ")
        self.left.print(indent + "\t")
        print(indent + "Operator: " + self.operator)
        self.right.print(indent + "\t")

    def expr_func(self):
        pass


class Literal(Expression):
    def __init__(self, value):
        self.value = value

    def print(self, indent):
        print(indent + "Literal: " + str(self.value))

    def expr_func(self):
        pass


class Variable(Expression):
    def __init__(self, name):
        self.name = name

    def print(self, indent):
        print(indent + "Variable: " + self.name)

    def expr_func(self):
        pass

class UnOpExpr:
    def __init__(self, op, expr):
        self.op = op
        self.expr = expr

    def print(self, indent):
        print(indent + "UnOpExpr: ")
        print(indent + "  " + "Op: " + self.op.val)
        self.expr.print(indent + "  ")

    def expr_func(self):
        pass

class AllocExpr:
    def __init__(self, type_, size):
        self.type_ = type_
        self.size = size

    def print(self, indent):
        print(indent + "AllocExpr: ")
        self.type_.print(indent + "  ")
        print(indent + "  " + "Size: " + self.size.val)

    def expr_func(self):
        pass
        

class FuncCall:
    def __init__(self, func_name, args):
        self.func_name = func_name
        self.args = args

    def print(self, indent):
        print(indent + "FuncCall: ")
        print(indent + "  " + "FuncName: " + \
self.func_name.val)
        print(indent + "  " + "Args: ")
        for arg in self.args:
            arg.print(indent + "    ")

    def expr_func(self):
        pass

    def statement_func(self):
        pass

class Cycle:
    def __init__(self, start, end, step=None):
        self.start = start
        self.end = end
        self.step = step
```

`parser.py`

Непосредственно сам рекурсивный парсер

```
from scanner import *
from content import *

class Parser:
    def __init__(self, scanner):
        self.sym = scanner.next_token()
        self.scanner = scanner

    def expect_str(self, *vals):
        for val in vals:
            if self.sym.val == val:
                sym = self.sym
                self.sym = self.scanner.next_token()
                if self.sym.tag == ERR:
                    raise Exception("parse error:\
 unexpected token")
                return sym
        raise Exception(f"parse error: expected {vals}, but got\
 {self.sym.val}")

    def expect_tags(self, *tags):
        for tag in tags:
            if self.sym.tag == tag:
                sym = self.sym
                self.sym = self.scanner.next_token()
                if self.sym.tag == ERR:
                    raise Exception("parse\
 error: unexpectedtoken")
                return sym
        raise Exception(f"parse error: expected \
{tags}, but got {self.sym.tag}")

    def program(self):
        funcs = []
        funcs.append(self.func())
        while self.sym.val in ["[", "("]:
            funcs.append(self.func())
        return Program(funcs)

    def func(self):
        header = self.func_header()
        body = self.func_body()
        return Func(header, body)

    def func_header(self):
        if self.sym.val == "(":
            self.expect_str("(")
            type_ = self.type()
            self.expect_str("[")
            func_name = self.expect_tags(FUNCNAME)
            params = self.func_params()
            self.expect_str("]")
            self.expect_str(")")
            return FuncHeader(type_, func_name, params)
        self.expect_str("[")
        func_name = self.expect_tags(FUNCNAME)
        params = self.func_params()
        self.expect_str("]")
        return FuncHeader(None, func_name, params)

    def func_body(self):
        statements = self.statements()
        self.expect_str("%%")
        return FuncBody(statements)

    def func_params(self):
        basic_vars = []
        while self.sym.val == "(":
            basic_vars.append(self.basic_var())
        return FuncParams(basic_vars)

    def basic_var(self):
        self.expect_str("(")
        type_ = self.type()
        var_name = self.expect_tags(VARNAME)
        self.expect_str(")")
        return BasicVar(type_, var_name)

    def type(self):
        if self.sym.val == "<":
            self.expect_str("<")
            type_ = self.type()
            self.expect_str(">")
            return ArrType(type_)
        var_type = self.expect_str("int", "bool", "char")
        return VarType(var_type)

    def statements(self):
        statements = []
        statements.append(self.statement())
        while self.sym.val == ",":
            self.expect_str(",")
            statements.append(self.statement())
        return Statements(statements)

    def statement(self):
        if self.sym.val == "^":
            self.expect_str("^")
            expr = self.expr()
            return ReturnStatement(expr)
        elif self.sym.val == "\\":
            self.expect_str("\\")
            expr = self.expr()
            return WarningStatement(expr)
        elif self.sym.val == "[":
            self.expect_str("[")
            func_name = self.expect_tags(FUNCNAME)
            args = self.args()
            self.expect_str("]")
            return FuncCall(func_name, args)
        elif self.sym.tag == VARNAME or self.sym.val \
== "<":
            var = self.var()
            self.expect_str(":=")
            expr = self.expr()
            return AssignmentStatement(var, expr)
        else:
            self.expect_str("(")
            return self.statement_tail()

    def statement_tail(self):
        if self.sym.val in ["<", "int", "bool", "char"]:
            type_ = self.type()
            var_name = self.expect_tags(VARNAME)
            if self.sym.val == ")":
                self.expect_str(")")
                if self.sym.val == ":=":
                    self.expect_str(":=")
                    expr = self.expr()
                    return AssignmentStatement(\
VarDeclaration\Statement(type_, var_name), expr)
                return VarDeclarationStatement(type_, \
var_name)
            elif self.sym.val == ":":
                cycle = self.cycle()
                statements = self.statements()
                self.expect_str("%")
                return ForStatement(type_, var_name, \
cycle.start, cycle.end, cycle.step, statements)
            else:
                raise Exception("parse error")
        elif self.sym.val == "?":
            self.expect_str("?")
            expr = self.expr()
            self.expect_str(")")
            then_branch = self.statements()
            else_branch = None
            if self.sym.val == "+++":
                self.expect_str("+++")
                else_branch_ = self.statements()
                else_branch = else_branch_
            self.expect_str("%")
            return IfStatement(expr, then_branch, \
else_branch)
        else:
            self.expect_str("&")
            expr = self.expr()
            self.expect_str(")")
            statements = self.statements()
            self.expect_str("%")
            return WhileStatement(expr, statements)

    def cycle(self):
        self.expect_str(":")
        start_expr = self.expr()
        self.expect_str(",")
        end_expr = self.expr()
        step = None
        if self.sym.val == ",":
            self.expect_str(",")
            token_ = self.expect_tags(INT_CONST)
            step = token_
        self.expect_str(")")
        return Cycle(start_expr, end_expr, step)

    def args(self):
        specs = []
        specs.append(self.spec())
        while self.sym.val in ["(", "[", "<", "new_",\
 "true", "false"] or self.sym.tag in [INT_CONST,
CHAR_CONST, STRING_CONST, VARNAME]:
            specs.append(self.spec())
        return specs

    def expr(self):
        expr = self.logical_expr()
        while self.sym.val in ["_or_", "_xor_"]:
            op = self.expect_str("_or_", "_xor_")
            right_expr = self.logical_expr()
            expr = BinOpExpr(expr, op, right_expr)
        return expr

    def logical_expr(self):
        expr = self.compare_expr()
        while self.sym.val == "_and_":
            op = self.expect_str("_and_")
            right_expr = self.compare_expr()
            expr = BinOpExpr(expr, op, right_expr)
        return expr

    def compare_expr(self):
        expr = self.arithm_expr()
        if self.sym.val in ["_eq_", "_ne_", \
"_lt_", "_gt_", "_le_", "_ge_"]:
            op = self.cmp_op()
            right_expr = self.arithm_expr()
            expr = BinOpExpr(expr, op, right_expr)
        return expr

    def cmp_op(self):
        return self.expect_str("_eq_", "_ne_", \
"_lt_", "_gt_", "_le_", "_ge_")

    def arithm_expr(self):
        expr = self.pow_expr()
        while self.sym.val in ["+", "-"]:
            op = self.expect_str("+", "-")
            right_expr = self.pow_expr()
            expr = BinOpExpr(expr, op, right_expr)
        return expr

    def pow_expr(self):
        expr = self.term()
        if self.sym.val == "_pow_":
            op = self.expect_str("_pow_")
            right_expr = self.pow_expr()
            expr = BinOpExpr(expr, op, right_expr)
        return expr

    def term(self):
        expr = self.factor()
        while self.sym.val in ["*", "/", "_mod_"]:
            op = self.expect_str("*", "/", "_mod_")
            right_expr = self.factor()
            expr = BinOpExpr(expr, op, right_expr)
        return expr

    def factor(self):
        if self.sym.val == "not_":
            self.expect_str("not_")
        elif self.sym.val == "-":
            self.expect_str("-")
        return self.spec()

    def func_call(self):
        self.expect_str("[")
        func_name = self.expect_tags(FUNCNAME)
        args = self.args()
        self.expect_str("]")
        return FuncCall(func_name, args)

    def spec(self):
        if self.sym.val == "(":
            self.expect_str("(")
            expr = self.expr()
            self.expect_str(")")
            return expr
        elif self.sym.val == "<" or self.sym.tag == VARNAME:
            var_expr = self.var()
            return var_expr
        elif self.sym.val in ["true", "false", "nothing"]\
 or self.sym.tag in [INT_CONST, CHAR_CONST, STRING_CONST]:
            return self.const()
        elif self.sym.val == "[":
            return self.func_call()
        else:
            self.expect_str("new_")
            type_ = self.type()
            size = self.expect_tags(VARNAME, INT_CONST)
            return AllocExpr(type_, size)

    def var(self):
        if self.sym.val == "<":
            self.expect_str("<")
            array_name = self.spec()
            expr = self.expr()
            self.expect_str(">")
            return Var(Token(), array_name, expr)
        var_name = self.expect_tags(VARNAME)
        return Var(var_name, None, None)

    def const(self):
        if self.sym.tag in [INT_CONST, CHAR_CONST, \
STRING_CONST]:
            val = self.expect_tags(INT_CONST, \
CHAR_CONST, STRING_CONST)
            return ConstExpr(None, val)
        str_ = self.expect_str("true", "false", \
"nothing")
        return ConstExpr(None, str_)
```

`position.py`
Файл который реализует функционал работы с позициями в тексте

```
import io
import unicodedata

class Position:
    def __init__(self, reader):
        self.symb = None
        self.line = 1
        self.pos = 1
        self.reader = reader
        self._initialize()

    def _initialize(self):
        try:
            self.symb = self.reader.read(1)
        except IOError:
            self.symb = None

    def __str__(self):
        return f"({self.line},{self.pos})"

    def cp(self):
        return ord(self.symb) if self.symb else -1

    def is_white_space(self):
        return self.symb == ' '

    def is_letter(self):
        return unicodedata.category(self.symb).startswith('L')\
if self.symb else False

    def is_underlining(self):
        return self.symb == '_'

    def is_digit(self):
        return self.symb.isdigit() if self.symb else False

    def is_letter_or_digit(self):
        return self.is_digit() or self.is_letter()

    def is_new_line(self):
        return self.symb == '\n'

    def is_close_bracket(self):
        return self.symb in [')', ']', '>', ',']

    def next(self):
        try:
            next_symb = self.reader.read(1)
            if next_symb:
                if self.is_new_line():
                    self.line += 1
                    self.pos = 1
                else:
                    self.pos += 1
                self.symb = next_symb
            else:
                self.symb = None
        except IOError:
            self.symb = None
        return self

    def skip_errors(self):
        while not self.is_white_space() and self.symb \
is not None:
            if self.next() == self:
                break

    def get_symbol(self):
        return self.symb
```

`scanner.py`
Реализация лексического анализа

```
import re
from token import *
from tag import *
from position import *
from content import *

class Comment:
    def __init__(self, starting, following, value):
        self.fragment = self.new_fragment(starting, \
following)
        self.value = value

    def new_fragment(self, starting, following):
        return (starting, following)

    def __str__(self):
        return f"COMMENT {self.fragment[0]}-\
{self.fragment[1]}:{self.value}"

class Scanner:
    def __init__(self, program_file):
        self.program_reader = program_file
        self.cur_pos = self.new_position(program_file)
        self.comments = []

    def new_position(self, program_file):
        return Position(program_file)

    def print_comments(self):
        for comm in self.comments:
            print(comm)

    def next_token(self):
        cur_word = ""
        while self.cur_pos.cp() != -1:
            while self.cur_pos.is_white_space():
                self.cur_pos.next()
            start = self.cur_pos

            if self.cur_pos.cp() == '\n':
                self.cur_pos.next()
            elif self.cur_pos.cp() == ':':
                cur_word += chr(self.cur_pos.cp())
                self.cur_pos.next()
                if self.cur_pos.get_symbol() == '=':
                    cur_word += chr(self.cur_pos.cp())
                    pos = self.cur_pos
                    self.cur_pos.next()
                    return self.new_token("SPEC_SYMB",\
 start, pos,cur_word)

                return self.new_token("SPEC_SYMB", \
start, start, cur_word)
            elif self.cur_pos.cp() == '%':
                cur_word += chr(self.cur_pos.cp())
                self.cur_pos.next()
                if self.cur_pos.get_symbol() == '%':
                    cur_word += chr(self.cur_pos.cp())
                    pos = self.cur_pos
                    self.cur_pos.next()
                    return self.new_token("SPEC_SYMB", \
start, pos, cur_word)

                return self.new_token("SPEC_SYMB", start, start,\
 cur_word)
            elif self.cur_pos.cp() == '"':
                self.cur_pos.next()
                pos = None
                if not self.cur_pos.is_new_line() and \
self.cur_pos.cp() != -1 and self.cur_pos.get_symbol() != '"':
                    cur_word += chr(self.cur_pos.cp())
                    pos = self.cur_pos
                    self.cur_pos.next()

                    if self.cur_pos.get_symbol() == '"':
                        pos = self.cur_pos
                        self.cur_pos.next()
                        return self.new_token("CHAR_CONST", start,\
 pos, cur_word)

                self.compiler.add_message(True, start,\
 "invalid syntax")
                self.cur_pos.skip_errors()

                return self.new_token("ERR", self.cur_pos, \
self.cur_pos, "")

            elif self.cur_pos.cp() == '\'':
                self.cur_pos.next()
                pos = None
                while self.cur_pos.cp() != -1 and \
self.cur_pos.get_symbol()\ != '\'' and \
not self.cur_pos.is_new_line():
                    cur_word += chr(self.cur_pos.cp())
                    pos = self.cur_pos
                    self.cur_pos.next()
                if self.cur_pos.get_symbol() == '\'':
                    pos = self.cur_pos
                    self.cur_pos.next()
                    return self.new_token("STRING_CONST",\
 start, pos, cur_word)

                self.compiler.add_message(True, start,\
 "invalid syntax")
                self.cur_pos.skip_errors()

                return self.new_token("ERR", self.cur_pos, \
self.cur_pos, "")

            elif self.cur_pos.cp() == '+':
                cur_word += chr(self.cur_pos.cp())
                self.cur_pos.next()
                pos = self.cur_pos
                while self.cur_pos.get_symbol() == '+':
                    cur_word += chr(self.cur_pos.cp())
                    pos = self.cur_pos
                    self.cur_pos.next()

                if not self.cur_pos.is_new_line() and not \
self.cur_pos.is_white_space() and self.cur_pos.cp() != -1:
                    self.compiler.add_message(True, start, \
"invalid syntax")
                    self.cur_pos.skip_errors()

                    return self.new_token("ERR", self.cur_pos, \
self.cur_pos, "")

                if cur_word == "+++" or cur_word == "+":
                    return self.new_token("SPEC_SYMB", start,\
 pos, cur_word)

                self.compiler.add_message(True, start, "invalid\
 syntax")
                self.cur_pos.skip_errors()

                return self.new_token("ERR", self.cur_pos, \
self.cur_pos, "")
            elif self.cur_pos.cp() == '{':
                self.cur_pos.next()
                while not self.cur_pos.is_new_line() and \
self.cur_pos.cp() != -1 and self.cur_pos.get_symbol() != '}':
                    cur_word += chr(self.cur_pos.cp())
                    self.cur_pos.next()

                self.comments.append(Comment(start, \
self.cur_pos, cur_word))
                self.cur_pos.next()
                cur_word = ""
            elif self.cur_pos.cp() in ['_', '!', '@', '.', '#']:
                cur_word += chr(self.cur_pos.cp())
                self.cur_pos.next()
                pos = None
                while self.cur_pos.is_letter():
                    cur_word += chr(self.cur_pos.cp())
                    pos = self.cur_pos
                    self.cur_pos.next()

                if self.cur_pos.is_underlining():
                    cur_word += chr(self.cur_pos.cp())
                    pos = self.cur_pos
                    self.cur_pos.next()
                    if cur_word in \
keywords_with_underlining_start:
                        return self.new_token("KEYWORD", start,\
 pos, cur_word)

                    self.compiler.add_message(True, start, \
"invalid syntax")
                    self.cur_pos.skip_errors()

                    return self.new_token("ERR", self.cur_pos, \
self.cur_pos, "")
                if not self.cur_pos.is_close_bracket() and not\
 self.cur_pos.is_new_line() and not \
self.cur_pos.is_white_space() and self.cur_pos.cp() != -1:
                    self.compiler.add_message(True, start, \
"invalid syntax")
                    self.cur_pos.skip_errors()

                    return self.new_token("ERR", self.cur_pos,\
 self.cur_pos, "")

                return self.new_token("VARNAME", start, pos,\
 cur_word)
            else:
                if chr(self.cur_pos.get_symbol()) in \
spec_symbs_in_one_rune:
                    cur_word += chr(self.cur_pos.cp())
                    self.cur_pos.next()
                    return self.new_token("SPEC_SYMB", \
start, start, cur_word)

                if self.cur_pos.is_digit():
                    cur_word += chr(self.cur_pos.cp())
                    self.cur_pos.next()
                    pos = None
                    while self.cur_pos.is_letter_or_digit():
                        cur_word += chr(self.cur_pos.cp())
                        pos = self.cur_pos
                        self.cur_pos.next()

                    if self.cur_pos.get_symbol() == '{':
                        cur_word += chr(self.cur_pos.cp())
                        self.cur_pos.next()
                        while self.cur_pos.is_digit():
                            cur_word += chr(self.cur_pos.cp())
                            pos = self.cur_pos
                            self.cur_pos.next()

                        if self.cur_pos.get_symbol() != '}' \
and not self.cur_pos.is_close_bracket() and not \
self.cur_pos.is_new_line() and not self.cur_pos.\
is_white_space()and self.cur_pos.cp() != -1:
                            self.compiler.add_message(True, \
start, "invalid syntax")
                            self.cur_pos.skip_errors()

                            return self.new_token("ERR", \
self.cur_pos, self.cur_pos, "")

                        cur_word += chr(self.cur_pos.cp())
                        pos = self.cur_pos
                        self.cur_pos.next()
                        return self.new_token("INT_CONST", \
start, pos, cur_word)

                    if not self.cur_pos.is_close_bracket() \
and not self.cur_pos.is_new_line() and not \
self.cur_pos.is_white_space() and self.cur_pos.cp() != -1:
                        self.compiler.add_message(True, \
start, "invalid syntax")
                        self.cur_pos.skip_errors()

                        return self.new_token("ERR", \
self.cur_pos, self.cur_pos, "")

                    return self.new_token("INT_CONST", \
start, pos, cur_word)

                if self.cur_pos.is_letter():
                    cur_word += chr(self.cur_pos.cp())
                    self.cur_pos.next()
                    pos = None
                    while self.cur_pos.is_letter():
                        cur_word += chr(self.cur_pos.cp())
                        pos = self.cur_pos
                        self.cur_pos.next()

                    if self.cur_pos.is_digit():
                        continue

                    if self.cur_pos.is_underlining():
                        cur_word += chr(self.cur_pos.cp())
                        pos = self.cur_pos
                        self.cur_pos.next()
                        if cur_word in keywords:
                            return self.new_token("KEYWORD",\
 start, pos, cur_word)

                        self.compiler.add_message(True, \
start, "invalid syntax")
                        self.cur_pos.skip_errors()

                        return self.new_token("ERR",\
 self.cur_pos, self.cur_pos, "")

                    if not self.cur_pos.is_close_bracket()\
 and not self.cur_pos.is_new_line() and not \
self.cur_pos.is_white_space() and self.cur_pos.cp() != -1:
                        self.compiler.add_message(True, \
start, "invalid syntax")
                        self.cur_pos.skip_errors()

                        return self.new_token("ERR", \
self.cur_pos, self.cur_pos, "")

                    if cur_word in keywords:
                        return self.new_token("KEYWORD", \
start, pos, cur_word)

                    return self.new_token("FUNCNAME",\
 start, pos, cur_word)

                self.compiler.add_message(True, start,\
 "invalid syntax")
                self.cur_pos.skip_errors()

                return self.new_token("ERR", \
self.cur_pos, self.cur_pos, "")

        return self.new_token("EOP", self.cur_pos, \
self.cur_pos, "")

    def new_token(self, token_type, start, pos, cur_word):
        return Token(token_type, start, pos, cur_word)
```

`tag.py`

Файл содержащий теги и информацию о них для разных лексем:
для лексического анализа

```
from enum import Enum

class DomainTag(Enum):
    VARNAME = 0
    FUNCNAME = 1
    INT_CONST = 2
    CHAR_CONST = 3
    STRING_CONST = 4
    KEYWORD = 5
    SPEC_SYMB = 6
    ERR = 7
    EOP = 8

tag_to_string = {
    DomainTag.VARNAME: "VARNAME",
    DomainTag.FUNCNAME: "FUNCNAME",
    DomainTag.INT_CONST: "INT_CONST",
    DomainTag.CHAR_CONST: "CHAR_CONST",
    DomainTag.STRING_CONST: "STRING_CONST",
    DomainTag.KEYWORD: "KEYWORD",
    DomainTag.SPEC_SYMB: "SPEC_SYMB",
    DomainTag.EOP: "EOP",
    DomainTag.ERR: "ERR",
}

keywords = {
    "bool": {},
    "char": {},
    "int": {},
    "false": {},
    "true": {},
    "nothing": {},
    "new_": {},
    "not_": {},
}

keywords_with_underlining_start = {
    "_and_": {},
    "_eq_": {},
    "_ge_": {},
    "_gt_": {},
    "_le_": {},
    "_lt_": {},
    "_mod_": {},
    "_ne_": {},
    "_or_": {},
    "_pow_": {},
    "_xor_": {},
}

spec_symbs_in_one_rune = {
    "<": {},
    ">": {},
    "(": {},
    ")": {},
    "[": {},
    "]": {},
    ",": {},
    "?": {},
    "&": {},
    "\\\\": {},
    "^": {},
    "-": {},
    "*": {},
    "/": {},
}
```

`token.py`

Файл реализующий представление токена и функционал для работы
с ним

```
class Token:
    def __init__(self, tag, starting, following, val):
        self.tag = tag
        self.coords = self.new_fragment(starting, following)
        self.val = val

    @staticmethod
    def new_fragment(starting, following):
        return {"starting": starting, "following": following}

    def __str__(self):
        return f"{tag_to_string[self.tag]} \
{self.coords}: {self.val}"

    def get_tag(self):
        return self.tag
```

# Тестирование

Для тестирования возьмем программу из
лабораторной 2.2

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

Применим функционал лабораторной к коду
и получим дерево разбора

```
Program:
Func:
FuncHeader:
FuncType:
RefType:
BasicType: int
FuncName: DoubleArrayElements
FuncParams:
BasicVar:
    VarType:
      RefType:
        RefType:
          BasicType: int
    VarName: !array
FuncBody:
Statements:
AssignmentStatement:
    VarDeclarationStatement:
            VarType:
              BasicType: int
            VarName: #length
    FuncCall:
            FuncName: length
            Args:
              VarExpr:
                    VarExprName: !array
ForStatement:
    ForHeader:
      VarType:
        RefType:
          BasicType: int
      VarName: #i
      ForStart:
        ConstExpr:
            ConstType: INT_CONST
            Val: 0
      ForEnd:
        BinOpExpr:
            VarExpr:
                    VarExprName: #length
            Op: -
            ConstExpr:
                    ConstType: INT_CONST
                    Val: 1
ReturnStatement:
    VarExpr:
            VarExprName: !array
Func:
FuncHeader:
FuncName: PrintArray
FuncParams:
BasicVar:
    VarType:
      RefType:
        RefType:
          BasicType: int
    VarName: !array
FuncBody:
Statements:
AssignmentStatement:
    VarDeclarationStatement:
            VarType:
              BasicType: int
            VarName: #length
    FuncCall:
            FuncName: length
            Args:
              VarExpr:
                    VarExprName: !array
ForStatement:
    ForHeader:
      VarType:
        RefType:
          BasicType: int
      VarName: #i
      ForStart:
        ConstExpr:
            ConstType: INT_CONST
            Val: 0
      ForEnd:
        BinOpExpr:
            VarExpr:
                    VarExprName: #length
            Op: -
            ConstExpr:
                    ConstType: INT_CONST
                    Val: 1

```

# Вывод

## Чему я научился:

- Изучение алгоритмов построения парсеров методом рекурсивного спуска.
- Понимание принципов работы рекурсивного парсера и его структуры.
- Преобразование формальной грамматики в набор функций для
  синтаксического анализа.
- Написание кода для лексического анализатора и парсера на Python.

## Что было сделано:

- Разработка грамматики языка L4, включающей описание типов данных, выражений,
  операторов и структур управления.
- Создание основных компонентов парсера: лексический анализатор
  (`scanner.py`), синтаксический анализатор (`parser.py`),
  узлы дерева разбора (`node.py`) и вспомогательные классы
  (`content.py`, `position.py`, `token.py`, `tag.py`).
- Успешноеестирование работы парсера на примере программы из
  лабораторной работы 2.2
