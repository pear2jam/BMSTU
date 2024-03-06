import numpy as np
from matplotlib import pyplot as plt


# Интеграл от -1 до 2 == 28.357
f = lambda x: -3*x*x*x+12*x*x-5*x+2 + 9*np.sin(x) - 2*np.cos(x)

# Метод средних квадратов
def mean_squares(f, a, b, n):
    h = (b-a)/n
    lin = np.linspace(a+h/2, b-h/2, n)
    return h*np.sum(f(lin))


# Метод трапеции
def trapezoid(f, a, b, n):
    h = (b-a)/n
    lin = np.array([f(a+h*i) + f(a+h*(i+1)) for i in range(n)])
    return h*np.sum(lin)/2


# Вычисление интеграла квадратичной функции
def q_int(a, b, c, l, r):
    return (-2*a*l**3+2*a*r**3-3*b*l*l+3*b*r*r-6*c*l+6*c*r)/6


# Метод Симпсона
def simpson(f, a, b, n):
    h = (b-a)/n
    s = 0
    for i in range(n):
        x = np.array([a+h*i, a+h*i+h/2, a+h*(i+1)])
        a_, b_, c_ = np.polyfit(x, f(x), 2)
        s += q_int(a_, b_, c_, a+h*i, a+h*(i+1))
    return s


n = 2
while (mean_squares(f, -1, 2, n) - mean_squares(f, -1, 2, n//2))/(2**n-1) > 0.001:
    n *= 2
print(n)

n = 2
while (trapezoid(f, -1, 2, n) - trapezoid(f, -1, 2, n//2))/(2**n-1) > 0.001:
    n *= 2
print(n)

n = 2
while (simpson(f, -1, 2, n) - simpson(f, -1, 2, n//2))/(2**n-1) > 0.001:
    n *= 2
print(n)
