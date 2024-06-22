% Лабораторная работа № 3.2. «Форматтер исходных текстов»
% 21 июня 2024 г.
% Наумов Сергей ИУ9-62Б

# Цель работы
Целью данной работы является приобретение навыков использования генератора синтаксических анализаторов bison.

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



Программа представляет из себя два файла: лексер и парсер

`Lexer.l`

Этот файл Flex представляет собой лексический анализатор, 
который разбивает входной текст на токены,
используемые парсером Bison.

**Опции flex:**

`%option reentrant noyywrap bison-bridge bison-locations`

Здесь задаются опции для генератора Flex:

-reentrant - создание потокобезопасного анализатора.
-noyywrap - отключение функции yywrap.
-bison-bridge - включение совместимости с Bison.
-bison-locations - использование функций Bison для 
 отслеживания расположения токенов.

**Правила распознавания токенов**

Правила определяют, как текст преобразуется в токены

```
[\n]+ {
printf("\n");
return NL_TOKEN;
}
(( [a-zA-Z0-9]+\{[0-9]+\} )|[0-9]+) {
    yylval->NUM_TOKEN = yytext;
    return NUM_TOKEN;
}

([@|.|#|_|!][a-zA-Z]+)  {
    yylval->VNAME_TOKEN = yytext;
    return VNAME_TOKEN;
}

\"[a-zA-Z]?\"   {
    yylval->CHAR_CONST_TOKEN = yytext;
    return CHAR_CONST_TOKEN;
}

[a-zA-Z]*  {
    yylval->FUNCNAME_TOKEN = yytext;
    return FUNCNAME_TOKEN;
}

(\'[a-zA-Z]*\')   {
    yylval->STR_TOKEN = yytext;
    return STR_CONST_TOKEN;
}
```

**Определение токенов**

Далее определеям токены языка L4

Базовые переменные
```
int return INT_TOKEN;
char return CHAR_TOKEN;
bool return BOOL_TOKEN;
```

Логические
```
true return TRUE_TOKEN;
false return FALSE_TOKEN;

_and_ return AND_TOKEN;
_or_ return OR_TOKEN;
_xor_ return XOR_TOKEN;
_eq_ return EQ_TOKEN;
_lt_ return LT_TOKEN;
_le_ return LE_TOKEN;
_mod_ return MOD_TOKEN;
_ne_ return NE_TOKEN;
```

Арифметические
```
\+  return PLUS_TOKEN;
\-  return SUB_TOKEN;
\*  return MUL_TOKEN;
\/  return DIV_TOKEN;
_pow_ return POW_TOKEN;

```

Скобки
```
\(  return BREAK1_TOKEN;
\)  return BREAK2_TOKEN;
\[  return BREAK3_TOKEN;
\]  return BREAK4_TOKEN;
\<  return BREAK5_TOKEN;
\>  return BREAK6_TOKEN;
```

Управляющие токены и спец символы
```
\+\+\+ return ELSE_TOKEN;
\& return WHILE_TOKEN;
:=  return ASSIGN_TOKEN;
(%%) return STMT_END_TOKEN;
(%) return STMT_EXPR_END_TOKEN;
\? return IF_TOKEN;
\: return DOUBLEDOT_TOKEN;
\^  return RETURN_TOKEN;
```

**Пользовательские действия**

Здесь мы используем макрос YY_USER_ACTION
Если флаг continued в структуре inf не установлен, то текущие
значения столбца и строки  копируются в поля first_column и 
first_line структуры yylloc. Это указывает на начало текущей лексемы.

Флаг continued устанавливается в false,
что означает, что текущая лексема не является
продолжением предыдущей.

Затем выполняется цикл for, который проходит
по каждому символу в текущей лексеме (yytext),
длина которой определяется переменной yyleng

Проверяем на переход на новую строку
После завершения цикла значения cur_line и cur_column
копируются в поля last_line и last_column 
структуры yylloc - указывает на конец текущего лексемы.


```
#define YY_USER_ACTION \
  { \
    struct Extra *inf = yyextra; \
    if (!inf->continued) { \
      yylloc->first_column = inf->cur_column; \
      yylloc->first_line = inf->cur_line; \
    } \
    inf->continued = false; \
    for (int i = 0; i < yyleng; i++) { \
      if (yytext[i] == '\n') { \
        inf->cur_column = 1; \
        ++inf->cur_line; \
      } else { \
        ++inf->cur_column; \
      } \
    }
    yylloc->last_line = inf->cur_line; \
    yylloc->last_column = inf->cur_column; \
  }
```

Теперь перейдем к `parser.y`

Этот файл реализует парсинг входного файла

**Настройка Bison и среды**

Включение заголовочных файлов и определение макросов 
для использования чистого API Bison и местоположений токенов.
Определение параметров для лексера и парсера, таких как yyscan_t 
lexer, буфер для результатов парсинга, счетчик отступов и флаг 
для добавления отступов.
Определение объединения %union, которое используется для 
хранения различных типов токенов.

```
%{
#include <stdio.h>
#include "lexer.h"
%}

%define api.pure
%locations

// Параметры для лексера и парсера
%lex-param {yyscan_t lexer} // Параметр лексера
%parse-param {yyscan_t lexer} // Параметр парсера
%parse-param {long buff[50]} // Буфер для хранения результатов парсинга
%parse-param {int tabulates} // Счетчик отступов
%parse-param {bool to_tabulate} // Флаг для добавления отступов

// Объединение для хранения различных типов токенов
%union {
    char* NUM_TOKEN;
    char* CHAR_CONST_TOKEN;
    char* STR_TOKEN;
    char* VNAME_TOKEN;
    char* FUNCNAME_TOKEN;
    char* COMMENT_TOKEN;
}
```

**Приоритет операций**

Здесь мы используем директивы %left, %right, 
которые определяют ассоциативность и приоритет операторов

%left OR_TOKEN XOR_TOKEN: Операторы "или"  и "исключающее или" 
имеют одинаковый приоритет и левую ассоциативность. 
Это означает, что если в выражении встречаются несколько
таких операторов, они будут группироваться слева направо.

%left AND_TOKEN: Оператор "и" имеет более высокий приоритет,
чем операторы "или" и "исключающее или", и 
также имеет левую ассоциативность.

И так далее

```
%left OR_TOKEN XOR_TOKEN
%left AND_TOKEN
%left PLUS_TOKEN SUB_TOKEN
%right POW_TOKEN
%left MUL_TOKEN DIV_TOKEN MOD_TOKEN
%left NOT_TOKEN UNARY_SUB_TOKEN
%left NEW_TOKEN
%left EQ_TOKEN NE_TOKEN LT_TOKEN GT_TOKEN LE_TOKEN GE_TOKEN
```

**Определение токенов**

```
%token INT_TOKEN CHAR_TOKEN BOOL_TOKEN 
%token WHILE_TOKEN WARNING_TOKEN RETURN_TOKEN 
%token BREAK1_TOKEN BREAK2_TOKEN BREAK3_TOKEN
%token BREAK4_TOKEN BREAK5_TOKEN BREAK6_TOKEN

%token STMT_END_TOKEN COMMA_TOKEN ASSIGN_TOKEN
%token IF_TOKEN ELSE_TOKEN STMT_EXPR_END_TOKEN

%token DOUBLEDOT_TOKEN TRUE_TOKEN FALSE_TOKEN

%token <NUM_TOKEN> NUM_TOKEN
%token <CHAR_CONST_TOKEN> CHAR_CONST_TOKEN
%token <STR_TOKEN> STR_CONST_TOKEN
%token <FUNCNAME_TOKEN> FUNCNAME_TOKEN
%token <VNAME_TOKEN> VNAME_TOKEN
%token <COMMENT_TOKEN> COMMENT_TOKEN
```

**Функции лексера и вспомогательная**

`yylex`: Это функция лексического анализатора,
которая вызывается парсером Bison для получения
следующего токена из входного потока. Она возвращает
целочисленный код токена, который соответствует одному 
из определенных в грамматике токенов.

`yyerror`: Это функция обработки ошибок, 
которая вызывается парсером Bison, когда 
обнаруживается синтаксическая ошибка. 

`tabulate`: Вспомогательная функция для расстановки
отступов

```
%{
int yylex(YYSTYPE *yylval_param, YYLTYPE *yylloc_param, yyscan_t lexer);
void yyerror(YYLTYPE *loc, yyscan_t lexer, long buff[50], int tabulates, bool to_tabulate, const char *message);

void tabulate(int tabulates) {
    for(int i = 0; i < tabulates; i++) {
        printf("    ");
    }
}
%}
```

**Правила парсера**

Правила для парсинга следуют из описанной грамматики языка
В процессе обработки этих правил мы также занимаемся
расстановкой отсупов, основывясь на грамматику

Например:

`Prog` - рекурсивное правило, которое позволяет определить программу,
состоящую из одной или нескольких функций. Каждая функция может быть
предварена комментариями `CmtTab`. Если встречается символ новой 
строки `NL_add`, выводится новая строка, и счетчик табуляции сбрасывается.



```
Prog:
Prog NL_add {printf("\n"); tabulates=0;} CmtTab Function
| Function
;
```
`Function` -  Описывает структуру функции, начиная с заголовка
`FunctionHeader` и заканчивая телом функции `FunctionBody`. После заголовка
функции, если флаг to_tabulate не установлен, выводится пробел, и 
счетчик табуляции увеличивается, что позволяет поддерживать правильное
форматирование кода.

```
Function:
FunctionHeader NL_add {if (!to_tabulate) printf(" "); tabulates++;} CmtTab FunctionBody
;
```

`CmtTab:` - управляет выводом комментариев. Если флаг `to_tabulate` установлен,
перед выводом комментария выполняется табуляция. После вывода 
комментария флаг сбрасывается, и следует новая строка.

```
CmtTab:
| COMMENT_TOKEN {if (to_tabulate) {tabulate(tabulates); to_tabulate = false;} printf("%s", $COMMENT_TOKEN);} NL_add CmtTab
;
```

`FunctionParams` - Управляет форматированием параметров функции.
Если флаг `to_tabulate` установлен, перед параметром выполняется 
табуляция. В противном случае выводится пробел. 
Если параметр является началом нового блока параметров, 
счетчик табуляции увеличивается.

```
FunctionParams:
FunctionParams {if (to_tabulate) tabulate(tabulates); else printf(" ");} BasicVar
| {if (to_tabulate) {tabulates++; tabulate(tabulates);}} BasicVar
;
```

**Функция main**

```
int main(int argc, char *argv[]) {
    long buff[50] = {0};
    int tabulates = 0;
    yyscan_t lexer;
    struct Extra extra;

    bool to_tabulate = true;

    // Считаем файл
    FILE *input_file = fopen(argv[1], "r");
    

    // Иннициализируем лексер

    extra.continued = false;
    extra.cur_line = 1;
    extra.cur_column = 1;

    yylex_init(&lexer);
    yylex_init_extra(&extra, &lexer);
    yyset_in(input_file, lexer);
    
    // Выполняем парсинг

    yyparse(lexer, buff, tabulates, to_tabulate);

    yylex_destroy(lexer);

    return 0;
}
```

# Тестирование

Напишем тестовую программу, содеражащию как можно
больше сущностей языка

```
(<bool>
[Double(int #A)] 
){Text is text} 
(int #res) := new_ int #size, 
#old := _a * (3 + 3),
<!A 5> := <#B 6> * 10, 
<[Func #old #y] !i + 1> := 0, 
(<int>#ffff) :=  new_ <int> 10,
                     (&#at _eq_ 0)#at := -100%, 
(<int> #j : 0, #size - 1) 
<#res #j> := <!A#j> * 2
{ <#res 5> := <!A #j> * 2 } 
%, 
(? #a _eq_ 0) 
#peremennaya := 3000
+++ 
(? #a _eq_ 0) 
#peremennaya := 2000
+++ 
#peremennaya := 1000
% 
%, 
^ _return
%% 
```

Вывод программы:
```
(<bool>
    [Double (int #A)]
) {Text is text}
    (int #res) := new_ int #size,
    #old := _a * (3 + 3),
    <!A 5> := <#B 6> * 10,
    <[Func #old #y] !i + 1> := 0,
    (<int> #ffff) := new_ <int> 10,
    (& #at _eq_ 0) #at := -100 %,
    (<int> #j : 0, #size - 1)
        <#res #j> := <!A #j> * 2
        { <#res 5> := <!A #j> * 2 }
    %,
    (? #a _eq_ 0)
        #peremennaya := 3000
    +++
        (? #a _eq_ 0)
            #peremennaya := 2000
        +++
            #peremennaya := 1000
        %
    %,
    ^ _return
```
# Вывод

В ходе выполнения лабораторной работы я приобрел ряд важных навыков и знаний:

1. **Изучение Bison и Flex**: Я научился использовать генератор синтаксических
   анализаторов Bison в сочетании с лексическим анализатором Flex.
   Это позволило мне глубже понять процесс компиляции и трансляции
   языков программирования.

2. **Реализация парсера**: Я освоил создание парсера с помощью Bison,
   что включало написание правил парсинга и функций для обработки
   синтаксических конструкций.

3. **Управление отступами**: Я научился управлять отступами
   в сгенерированном коде, освоив слабое форматирование

4. **Тестирование**: Я провел тестирование парсера, создав тестовую программу,
   содержащую различные сущности языка L4, что позволило мне проверить
   корректность работы парсера.

5. **Прикладная задача**: В рамках лабороторной работы я разработал функционал,
   который используются во многих реальных языках программирования
