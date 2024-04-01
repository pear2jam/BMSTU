% Лабораторная работа № 2.1. Синтаксические деревья
% 30 марта 2024 г.
% Наумов Сергей, ИУ9-62Б

# Цель работы

Целью данной работы является изучение представления синтаксических деревьев в 
памяти компилятора и приобретение навыков преобразования синтаксических деревьев.

# Индивидуальный вариант

Подсчёт общего количества итераций всех циклов в процессе выполнения программы.

# Реализация

Выполнение лабораторной работы состоит из нескольких этапов:

* подготовка исходного текста демонстрационной программы, которая в дальнейшем
 будет выступать в роли объекта преобразования (демонстрационная программа
 должна размещаться в одном файле и содержать функцию main)
* компиляция и запуск программы astprint для изучения структуры синтаксического
 дерева демонстрационной программы
* разработка программы, осуществляющей преобразование синтаксического дерева
 и порождение по нему новой программы;
* тестирование работоспособности разработанной программы на исходном тексте
 демонстрационной программы.

В качестве демонстрационной программы напишем программу в которой будут 
рассматриваться различные виды циклов:
* Цикл с фиксированным количеством итераций
  ```go
  for i := 0; i < 5; i++ {
		fmt.Printf("pre-known loop: %d\n", i)
	}
  ```
* Цикл for с количеством итераций заданным пользователем
  Число n считывается с ввода
  ```go
  for i := 0; i < n; i++ {
		fmt.Printf("user-defined loop: %d\n", i)
	}
  ```
* Цикл for без параметров с break при определенном условии
  ```go
  cnt := 17
	for {
		fmt.Printf("Break-exit loop: cnt = %d\n", cnt)
		cnt -= 2
		if cnt < 0 {
			break
		}
	}
  ```
* Итерация по массиву
  ```
  a := []string{"1", "2", "3", "4", "apple"}
	for _, s := range a {
		fmt.Printf("Array-iter loop: value = %s\n", s)
	}
  ```

* Цикл в вызываемой функции
```go
func iter() {
	// Цикл for с заданным количеством итераций в функции
	for i := 0; i < 8; i++ {
		fmt.Printf("pre-known loop in function: %d\n", i)
	}
}

iter()
```

Далее рассмотрим код astprint.go

```go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: astprint <filename.go>\n")
		return
	}

	// Создаём хранилище данных об исходных файлах
	fset := token.NewFileSet()

	// Вызываем парсер
	if file, err := parser.ParseFile(
		fset,                 // данные об исходниках
		os.Args[1],           // имя файла с исходником программы
		nil,                  // пусть парсер сам загрузит исходник
		parser.ParseComments, // приказываем сохранять комментарии
	); err == nil {
		// Если парсер отработал без ошибок, печатаем дерево
		ast.Fprint(os.Stdout, fset, file, nil)
	} else {
		// в противном случае, выводим сообщение об ошибке
		fmt.Printf("Error: %v", err)
	}
}
```

В этой функции реализован обход синтаксического дерева в глубину. 
Она вызывает переданную ей в качестве параметра функцию для каждого 
посещённого узла дерева. С помощью этой функции удобно осуществлять 
поиск узлов определённого типа в дереве.

Соберем этот код и применим к тестовому примеру:

Вывод довольно большой, но из него можно вычленить важные фрагменты
для нашей лабораторной работы

```
0: *ast.ForStmt {
    57  .  .  .  .  .  .  For: test.go:9:2
    58  .  .  .  .  .  .  Init: *ast.AssignStmt {
    59  .  .  .  .  .  .  .  Lhs: []ast.Expr (len = 1) {
    60  .  .  .  .  .  .  .  .  0: *ast.Ident {
    61  .  .  .  .  .  .  .  .  .  NamePos: test.go:9:6
    62  .  .  .  .  .  .  .  .  .  Name: "i"
...
...
...
```

Здесь можно увидеть первый цикл for в программе

```
9: *ast.ExprStmt {
   769  .  .  .  .  .  .  X: *ast.CallExpr {
   770  .  .  .  .  .  .  .  Fun: *ast.Ident {
   771  .  .  .  .  .  .  .  .  NamePos: test.go:49:2
   772  .  .  .  .  .  .  .  .  Name: "iter"
...
...
...
```

А здесь вызов функции с циклом внутри. Эти блоки также важны в лабороторной,
так как итерации в вызываемых функциях мы также подсчитываем


Теперь рассмотрим функцию преобразования кода, так чтобы выполнялся 
функционал лабороторной

Для того чтобы подсчитать общее количество итераций, объявим счетчик
и будем его инкрементировать в телах всех циклов

```go
package main

import (
	"fmt"
	"os"

	"go/ast"
	"go/format"
	"go/token"
	"go/parser"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error in parsing")
		return
	}
	addLoopIterationCount(file)
	format.Node(os.Stdout, fset, file)
}


func declareIntegerVariable(file *ast.File, value int) {
	var firstDecl, secDecl []ast.Decl
	if len(file.Decls) > 0 {
		if a, ok := file.Decls[0].(*ast.GenDecl); ok {
			if a.Tok == token.IMPORT {
				firstDecl, secDecl = []ast.Decl{file.Decls[0]}, file.Decls[1:]
			} else {secDecl = file.Decls}
		}
	}

	file.Decls = append(firstDecl,
		&ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{ast.NewIdent("__COUNTER__")},
					Type:  ast.NewIdent("int"),
					Values: []ast.Expr{
						&ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%d", value)},
					},
				},
			},
		},
	)
	file.Decls = append(file.Decls, secDecl...)
}

func findFuncDeclaration(file *ast.File, name string) *ast.FuncDecl {
	for _, declaration := range file.Decls {
		if funcDecl, ok := declaration.(*ast.FuncDecl); ok {
			if funcDecl.Name.Name == name {return funcDecl}
		}
	}

	return nil
}

func printLoopCounter(file *ast.File) {
	mainFunc := findFuncDeclaration(file, "main")

	mainFunc.Body.List = append(
		mainFunc.Body.List,
		&ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent("fmt"), Sel: ast.NewIdent("Printf")},
				Args: []ast.Expr{
					&ast.BasicLit{Kind:  token.STRING, Value: "\"Total: %d\\n\""},
					&ast.Ident{Name: "__COUNTER__"},
				},
			},
		},
	)
}

func addLoopIterationCount(file *ast.File) {
	declareIntegerVariable(file, 0)
	printLoopCounter(file)
	incStatemet := &ast.IncDecStmt{
		X: &ast.Ident{
			Name: "__COUNTER__",
		},
		Tok: token.INC,
	}
	ast.Inspect(file, func(node ast.Node) bool {
		if x, ok := node.(*ast.ForStmt); ok {
			x.Body.List = append(x.Body.List, incStatemet)
		} else if x, ok := node.(*ast.RangeStmt); ok {
			x.Body.List = append(x.Body.List, incStatemet)
		}
		return true
	})
	
}
```

* В функции main создается новый набор токенов с помощью token.NewFileSet().
* Файл программы парсится с использованием пакета parser.ParseFile(),
  который возвращает синтаксическое дерево программы и ошибку (если есть).
* Если происходит ошибка в парсинге, программа выводит сообщение об ошибке
  и завершается.
* Вызывается функция addLoopIterationCount, которая изменяет синтаксическое
  дерево программы, добавляя переменную __COUNTER__, итерацию счетчика в
  циклах и вывод общего количества итераций в функции main.
* Затем собранное дерево форматируется обратно в код с помощью format.Node()
  и выводится на стандартный вывод.


* Функция declareIntegerVariable добавляет объявление целочисленной переменной
  __COUNTER__ в начало файла. Она проверяет, есть ли уже объявления в файле, 
  и если есть, разделяет их на импорты и остальные объявления.
* Функция findFuncDeclaration находит объявление функции по ее имени в
  синтаксическом дереве файла. Она перебирает все объявления в файле и
  возвращает первое совпадение с именем функции.
* Функция printLoopCounter добавляет вызов fmt.Printf в функцию main, который 
  выводит общее количество итераций счетчика.
* Функция addLoopIterationCount выполняет все необходимые шаги для добавления
  счетчика итераций в программу. Она вызывает функции declareIntegerVariable 
  и printLoopCounter, затем создает инструкцию инкремента __COUNTER__ и
  добавляет ее в тело всех циклов в программе с помощью ast.Inspect,
  который перебирает все узлы в синтаксическом дереве.

# Тестирование

Применим функцию transform к коду программы test.go
и получим на выходе новую программу, в рамках которой считается количетсво
итераций

```go
package main

import (
        "fmt"
)

var __COUNTER__ int = 0

func iter() {
        // Цикл for с заданным количеством итераций в функции
        for i := 0; i < 8; i++ {
                fmt.Printf("pre-known loop in function: %d\n", i)
                __COUNTER__++
        }
}

func main() {
        var n int
        fmt.Print("Enter the number of iterations for the loop: ")
        if _, err := fmt.Scan(&n); err != nil {
                panic(err)
        }

        // Примеры различных циклов

        // Цикл for с заданным количеством итераций
        for i := 0; i < 5; i++ {
                fmt.Printf("pre-known loop: %d\n", i)
                __COUNTER__++
        }

        // Цикл for с количеством итераций заданным пользователем
        for i := 0; i < n; i++ {
                fmt.Printf("user-defined loop: %d\n", i)
                __COUNTER__++
        }

        // Цикл без условия с выходом при определенном условии
        cnt := 17
        for {
                fmt.Printf("Break-exit loop: cnt = %d\n", cnt)
                cnt -= 2
                if cnt < 0 {
                        break
                }
                __COUNTER__++
        }

        // Пробежаться по массиву
        a := []string{"1", "2", "3", "4", "apple"}
        for _, s := range a {
                fmt.Printf("Array-iter loop: value = %s\n", s)
                __COUNTER__++
        }

        iter()
        fmt.Printf("Total: %d\n", __COUNTER__)
}
```

Исполним полученную программу:

```
Enter the number of iterations for the loop: 6
pre-known loop: 0   
pre-known loop: 1   
pre-known loop: 2   
pre-known loop: 3   
pre-known loop: 4   
user-defined loop: 0
user-defined loop: 1
user-defined loop: 2
user-defined loop: 3
user-defined loop: 4
user-defined loop: 5
Break-exit loop: cnt = 17
Break-exit loop: cnt = 15
Break-exit loop: cnt = 13
Break-exit loop: cnt = 11
Break-exit loop: cnt = 9
Break-exit loop: cnt = 7
Break-exit loop: cnt = 5
Break-exit loop: cnt = 3
Break-exit loop: cnt = 1
Array-iter loop: value = 1
Array-iter loop: value = 2
Array-iter loop: value = 3
Array-iter loop: value = 4
Array-iter loop: value = apple
pre-known loop in function: 0
pre-known loop in function: 1
pre-known loop in function: 2
pre-known loop in function: 3
pre-known loop in function: 4
pre-known loop in function: 5
pre-known loop in function: 6
pre-known loop in function: 7
Total: 32
```

Ответ соответствует действительности

# Вывод

В ходе выполнения лабораторной работы я изучил представление синтаксических
деревьев в памяти компилятора и приобрел навыки преобразования синтаксических
деревьев.

Для достижения цели работы я выполнил следующие шаги:

1. Подготовил исходный текст демонстрационной программы,
 в которой рассматривались различные виды циклов.
2. Изучил структуру синтаксического дерева демонстрационной программы
 с помощью программы `astprint.go`, которая осуществляет обход дерева
   в глубину и выводит его структуру.
3. Разработал программу `transform.go`, которая преобразует синтаксическое
    дерево программы, подсчитывая общее количество итераций в циклах.
4. Провел тестирование разработанной программы на тестовом примере,
   убедившись в правильности работы.

В результате выполнения работы я научился работать с синтаксическими деревьями
в рамках компилятора Go. Я понял, что синтаксические деревья представляют 
собой структуру данных, которая отражает синтаксическую структуру программы,
разбивая её на составные элементы (узлы) и определяя отношения между ними. 

В ходе работы я понял, как строятся синтаксические деревья: при компиляции
исходного кода программа разбирается на лексемы, затем лексемы группируются
в токены, а токены в свою очередь формируют узлы синтаксического дерева, 
представляющие собой абстрактное синтаксическое представление программы.

Также я осознал важность работы с синтаксическими деревьями для множества задач,
включая анализ кода, автоматическое преобразование программ, 
инструменты статического анализа и другие.

