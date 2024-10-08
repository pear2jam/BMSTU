% Лабораторная работа № 1.1. Раскрутка самоприменимого компилятора 
% 30 марта 2024 г.
% Наумов Сергей, ИУ9-62Б

# Цель работы
Целью данной работы является ознакомление с раскруткой самоприменимых компиляторов на примере модельного компилятора.

# Индивидуальный вариант
Компилятор P5. Обеспечить возможность использования в числовых литералах незначимый знак _ 
(например, число 10 000 можно записать и как 10000, и как 10_000, и как 100__00).

# Реализация

В рамках лабораторной работы будем работать с компилятором языка Pascal P5. Нужно изменить исходный код компилятора так,
чтобы при формировании числовых литералов можно было использовать незначимый знак '_' \
То есть литерал 100 может быть записан как 10_0 и как 1_____0

Для этого обратимся к книге документации компилятора P4 : P4 Compiler and Interpreter 
(https://homepages.cwi.nl/~steven/pascal/book/pascalimplementation.html) \
Несмотря на то, что мы работаем с компилятором P5, код для работы с численными литералами в них идентичен. \
Перейдем к графе 'Real and Integer Numbers' в разделе Input and Lexical Analysis : 
здесь описаны блоки кода которые отвечают за 'сбор' разных частей численного литерала, обработки точки и символа экспоненты,
формирования литерала числа с плавающей точкой и целочисленного литерала.

```pascal
repeat i := i+1; if i<= digmax then digit[i] := ch; nextch
until chartp[ch] <> number;
```

Здесь мы проходимся по литералу и записываем в массив хранения литерала символы пока мы не встретим не цифру.
Далее также мы можем считать точку или экспоненту, и досчитывать правую часть если такая имеется

Добавим в условие продолжение считывания незначимый символ '_'

```pascal
until (chartp[ch] <> number) or (ch <> '_');
```

Далее изменим часть формирования численного значения:

```pascal
ival := ival*10+ordint[digit[k]]
```

Сделаем так, что если мы встречаем незначимый символ, мы не меняем результирующее значение

```pascal
if digit[k] <> '_' then begin
ival := ival*10+ordint[digit[k]]
end
```

Таким образом мы добавили незначмый символ '_'

Следущим этапом надо собрать новый компилятор в котором есть соответсвующая новая возможность \
Теперь возьмем измененный код комплятора и добавим в сам код новую возможность

Изменим строчку с формированием результата численного литерала и добавим символы '_'
```
ival := ival*1__0+ordint[digit[k]]
```
После чего снова соберем уже третюю версию компилятора

В качестве результата выведем разницу между парами файлов (pcom.pas, pcom2.pas) и (pcom2.pas, pcom3.pas)

pcom.pas - исходный код, pcom2.pas - добавление новых возможностей
```
naumov@naumov:~/Downloads/Comp/lab1.1$ diff -uw  pcom.pas pcom2.pas
--- pcom.pas	2024-03-23 15:58:47.367168964 +0300
+++ pcom2.pas	2024-03-31 14:58:12.549237785 +0300
@@ -1363,7 +1363,7 @@
       number:
         begin op := noop; i := 0;
           repeat i := i+1; if i<= digmax then digit[i] := ch; nextch
-          until chartp[ch] <> number;
+          until (chartp[ch] <> number) and (ch <> '_');
           if ((ch = '.') and (input^ <> '.') and (input^ <> ')')) or 
              (lcase(ch) = 'e') then
             begin
@@ -1412,8 +1412,11 @@
                   begin ival := 0;
                     for k := 1 to i do
                       begin
-                        if ival <= mxint10 then
+                        if ival <= mxint10 then begin
+                          if digit[k] <> '_' then begin
                           ival := ival*10+ordint[digit[k]]
+                          end
+                        end
                         else begin error(203); ival := 0 end
                       end;
                     sy := intconst
```

pcom2.pas и pcom3.pas - добавили новую возможность в сам компилятор

```
naumov@naumov:~/Downloads/Comp/lab1.1$ diff -uw  pcom2.pas pcom3.pas
--- pcom2.pas	2024-03-31 15:11:28.010108090 +0300
+++ pcom3.pas	2024-03-25 15:55:35.048965024 +0300
@@ -1362,7 +1362,9 @@
       end;
       number:
         begin op := noop; i := 0;
-          repeat i := i+1; if i<= digmax then digit[i] := ch; nextch
+          repeat i := i+1; 
+          if ((i<= digmax) and (ch <> '_')) then digit[i] := ch;
+          if ((i<= digmax) and (ch = '_')) then digit[i] := ch; nextch
           until (chartp[ch] <> number) and (ch <> '_');
           if ((ch = '.') and (input^ <> '.') and (input^ <> ')')) or 
              (lcase(ch) = 'e') then
@@ -1414,7 +1416,7 @@
                       begin
                         if ival <= mxint10 then begin
                           if digit[k] <> '_' then begin
-                            ival := ival*10+ordint[digit[k]]
+                            ival := ival*1__0+ordint[digit[k]]
                           end
                         end
                         else begin error(203); ival := 0 end
```
# Тестирование

Составим программу на языке Pascal в котором зададим численную константу и выведем ее в консоль

```pascal
program hello(output);

var

n : integer;

begin
   n := 1___45;
   writeln(n);
end.
```

Результат компиляции и ее исполнения:

```
P5 Pascal compiler vs. 1.0


     1       40 program hello(output); 
     2       40  
     3       40 var 
     4       40  
     5       40 n : integer; 
     6       44  
     7       44 begin 
     8        3     n := 1___45; 
     9        7     write(n); 
    10       13 end. 

Errors in program: 0

program complete
P5 Pascal interpreter vs. 1.0

Assembling/loading program
Running program

        145
program complete
```
Как видно, компилятор не учел символы '_' при формировании численного литерала

# Вывод
Работа с исходным кодом: \
Я познакомился с исходным кодом компилятора Pascal P5 и проанализировали его структуру, особенно в части работы с 
числовыми литералами.
На основе изученного кода я внес изменения, позволяющие компилятору распознавать и обрабатывать незначимый символ '_' 
\в числовых литералах.
После внесения изменений я собрал новую версию компилятора и протестировали его, убедившись, что добавленная функциональность
работает корректно.

Работа с документацией: \
Я получил опыт работы с документацией к коду. Этот навык в будущем пригодится для написания докуменации как к своим проектам, 
так и к упрощению понимая чужого кода.
Чтение документации дало понимание того, как реализована обработка числовых литералов внутри компилятора и какие изменения нужно
внести для решения задачи.

Работа с лексическим анализом: \
Изучение кода, отвечающего за лексический анализ числовых литералов, позволило лучше понять принципы работы компилятора на 
этапе разбора исходного кода.

Процесс итеративной разработки: \
В процессе выполнения лабораторной использовалась методика итеративной разработки, последовательно вносились изменения в исходный код, 
собирались новые версии компилятора и проверялась их работоспособность. 

Эта лабораторная работа позволила не только ознакомиться с процессом раскрутки самоприменимого компилятора, но и приобрести 
практические навыки работы с исходным кодом компилятора, его изменениями и тестированием. 
Такой опыт будет полезен при дальнейшей разработке программного обеспечения и работы с большими проектами.

