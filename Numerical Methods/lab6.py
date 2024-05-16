import numpy as np
import matplotlib.pyplot as plt

# Задаем аналитическое точное решение
exact_solution = lambda x: (12*x - 7*np.exp(-6*x) + 25) / 9

# Метрика сравнения результатов
metric = lambda x, y: np.mean(np.abs(x-y))

# Реализация метода Рунге-Кнутта
def runge_kutta_second_order(f, y0, x0, xf, h, exact):
    x_values = np.arange(x0, xf + h, h)
    n = len(x_values)
    y_values = np.zeros((n, len(y0)))
    y_values[0] = y0
    
    for i in range(1, n):
        x = x_values[i]
        y_prev = y_values[i - 1]
        
        k1 = h * f(x, y_prev)
        k2 = h * f(x + h/2, y_prev + k1/2)
        k3 = h * f(x + h/2, y_prev + k2/2)
        k4 = h * f(x + h, y_prev + k3)
        
        y_values[i] = y_prev + (k1 + 2*k2 + 2*k3 + k4) / 6
        print(f'step {i}: metric: {metric(exact(x_values), y_values[:,0])}')
    return x_values, y_values

# Задаем исходную функцию
def f(x, y):
    dydx = np.zeros_like(y)
    dydx[0] = y[1]  #y' = v
    dydx[1] = 8 - 6*y[1]  #v' = 8 - 6*v
    return dydx

# начальные значения
y0 = np.array([2, 6])
x0, xf = 0, 1
h = 0.01

x_values, y_values = runge_kutta_second_order(f, y0, x0, xf, h, exact_solution)

exact_values = exact_solution(t_values)

# Вывод результатов
plt.plot(x_values, y_values[:, 0], label='Приблизительное решение')
plt.plot(x_values, exact_values, label='Точное решение')
plt.xlabel('x')
plt.ylabel('y')
plt.legend()
plt.grid(True)
plt.show()
