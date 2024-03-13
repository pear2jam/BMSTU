import numpy as np

def calc_a(x, y):
    return (x+y)/2

def calc_g(x, y):
    return np.sqrt(x*y)

def calc_h(x, y):
    return 2/(1/x+1/y)

# Вариант 20
x = np.linspace(1, 5, 9)
y = np.array([3, 2.1, 2.21, 2.04, 1.13, 1.18, 1.27, 1.10, 0.86])

# Аппроксимация полиномом
z = np.poly1d(np.polyfit(x, y, len(x)-1))

# Вычисляем средние
x_a, x_g, x_h = calc_a(x[0], x[-1]), calc_g(x[0], x[-1]), calc_h(x[0], x[-1])
y_a, y_g, y_h = calc_a(y[0], y[-1]), calc_g(y[0], y[-1]), calc_h(y[0], y[-1])
z_a, z_g, z_h = calc_a(z(x[0]), z(x[-1])), calc_g(z(x[0]), z(x[-1])), calc_h(z(x[0]), z(x[-1]))

print('X:', x_a, x_g, x_h)
print('Y:', y_a, y_g, y_h)
print('Z:', z_a, z_g, z_h)


print('1: ', np.abs(z_a-y_a))
print('2: ', np.abs(z_g-y_g))
print('3: ', np.abs(z_a-y_g))
print('4: ', np.abs(z_g-y_a))
print('5: ', np.abs(z_h-y_a))
print('6: ', np.abs(z_a-y_h))
print('7: ', np.abs(z_h-y_h))
print('8: ', np.abs(z_h-y_g))
print('9: ', np.abs(z_g-y_h))

# Argmin == 1 --> f(x) = ax+b


# Считаем коэффициенты системы для вычисления коэфициентов a, b функции аппроксимации f
a1, b1, c1 = sum(x*x), sum(x), sum(x*y)
a2, b2, c2 = sum(x), len(x)+1, sum(y)


a, b = np.linalg.solve(np.array([[a1, b1], [a2, b2]]), np.array([c1, c2]))

# Итоговая функция
f = lambda x: a*x+b


# Среднеквадратичное отклонение
np.mean((f(x)-y)*(f(x)-y))
