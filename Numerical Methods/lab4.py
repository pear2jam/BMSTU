import numpy as np

# Решение системы методом прогонки
def solve_tridiagonal(a, b, c, d):
    """
    a, b, c, d: array-like
    a, b, c: диагонали матрицы (нижняя, главная, верхняя)
    d: свободные коэффициенты
    return: list
    """
    n = len(d)
    c_temp = [0] * n
    d_temp = [0] * n
    x = [0] * n

    c_temp[0] = c[0] / b[0]
    d_temp[0] = d[0] / b[0]

    for i in range(1, n):
        temp = b[i] - a[i] * c_temp[i - 1]
        c_temp[i] = c[i] / temp
        d_temp[i] = (d[i] - a[i] * d_temp[i - 1]) / temp

    x[n - 1] = d_temp[n - 1]
    for i in range(n - 2, -1, -1):
        x[i] = d_temp[i] - c_temp[i] * x[i + 1]

    return x


"""
Решая аналитическую задачу Коши
y''+6y=8, y(0)=2, y'(0)=6
y(x)=1/9*(12*x-7*e^(-6*x)+25)
y(1) ~ 4.10 = b
"""

n = 10

p, q, f, h = 6, 0, 8, 1/n  # Задаем параметры из условия и индивидуального варианта
low_diag, main_diag, up_diag = [0 for i in range(n)], [0 for i in range(n)], [0 for i in range(n)]
y = [0 for i in range(n)]
main_diag[0], main_diag[-1] = 1, 1
y[0], y[-1] = 2, 4.1

# Заполняем матрицу согласно методу
for i in range(1, n-1):
    low_diag[i] = 1-h/2*p
    main_diag[i] = h*h*q-2
    up_diag[i] = 1+h/2*p
    y[i] = h*h*f


# Истинная приближаемая функция
y_func = lambda x: 1/9*(12*x-7*np.exp(-6*x)+25)


# Выводим результаты
y_num, y_true = np.array(solve_tridiagonal(low_diag, main_diag, up_diag, y)), np.array([y_func(i/10) for i in range(n)])
print(max(abs(y_num - y_true)))
print(y_num, y_true)
