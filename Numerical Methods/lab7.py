import numpy as np

f = lambda x: 5
y_true = lambda x: 1/9*(12*x-7*np.exp(-6*x)+25)

def method(n, a, b, h, p, q, x_0, o_h):
    y_0 = [0] * (n + 1)
    y_1 = [0] * (n + 1)
    y_0[0] = a
    y_0[1] = a + o_h
    y_1[1] = o_h

    for i in range(1, n):
        y_0[i + 1] = (h * h * f(x_0 + i * h) - (1 - h / 2 * p) * y_0[i - 1] - (h * h * q - 2) * y_0[i]) / (1 + h / 2 * p)
        y_1[i + 1] = ((h / 2 * p - 1) * y_1[i - 1] - (h * h * q - 2) * y_1[i]) / (1 + h / 2 * p)

    if y_1[n] == 0:
        return solve(n, a, b, h, p, q, x_0, o_h + 1)
    else:
        c1 = (b - y_0[n]) / y_1[n]
        return [y_0[i] + c1 * y_1[i] for i in range(n+1)]

def solve(mid, top, bot, b):
    n = len(b)
    x = [0] * n
    v = [0] * n
    u = [0] * n

    v[0] = -top[0] / mid[0]
    u[0] = b[0] / mid[0]
    for i in range(1, n):
        v[i] = -top[i] / (bot[i] * v[i - 1] + mid[i])
        u[i] = (b[i] - bot[i] * u[i - 1]) / (bot[i] * v[i - 1] + mid[i])

    x[n - 1] = u[n - 1]
    for i in range(n - 1, 0, -1):
        x[i - 1] = v[i - 1] * x[i] + u[i - 1]
    return x

x_0 = 0
x_n = 1
h = (x_n - x_0)/n

a, b = y_true(x_0), y_true(x_n)

p, q = 6, 0

n = 10

mid = [h * h * q - 2]
top = [1 + h / 2 * p]
result = [h * h * f(x_0 + h) - a * (1 - h / 2 * p)]
bot = [0]

bot += [1 - h / 2 * p for i in range(2, n - 1)]
mid += [h * h * q - 2 for i in range(2, n - 1)]
top += [1 + h / 2 * p for i in range(2, n - 1)]
result += [h * h * f(x_0 + h * i) for i in range(2, n - 1)]

bot += [1 - h / 2 * p]
mid += [h * h * q - 2]
result += [h * h * f(x_0 + (n - 1) * h) - b * (1 + h / 2 * p)]
top += [0]


x_arr = [x_0 + i * h for i in range(n + 1)]
y_true_arr = [y_true(x_0 + i * h) for i in range(n + 1)]
y_ev_arr = [a] + solve(mid, top, bot, result) + [b]
y_solve = method(n, a, b, h, p, q, x_0, h)

print("x\t\ty_true\t\ty_solve")
for i in range(n + 1):
    print(f"{x_arr[i]:.6f}\t{y_true_arr[i]:.6f}\t{y_solve[i]:.6f}")

diff = [y_true_arr[i] - y_ev_arr[i] for i in range(n + 1)]
print(np.mean(np.array(diff)/np.array(y_solve)))
