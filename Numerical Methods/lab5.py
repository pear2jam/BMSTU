import numpy as np
from matplotlib import pyplot as plt

# Задаем исходную функцию
f = lambda x, y: x+5*y+np.exp(x*x+y*y)
eps = 1e-3

# Считаем частные производные
dfx = lambda x, y: 2*x*np.exp(x*x+y*y)+1
dfy = lambda x, y: 2*y*np.exp(x*x+y*y)+5

# Итеративно сходимся с помощью градиентного спуска к оптимальной точке
point = np.array([0., 0.])
c = 3e-2
f_list = []
while True:
    c /= 1.022
    grad = np.array([dfx(*point), dfy(*point)])
    if np.linalg.norm(grad) < eps: 
        print('Stop')
        break
    grad /= np.linalg.norm(grad)
    point -= c*grad
    f_list.append(f(*point))

#аналитически: (-0.191911, -0.959555)

plt.plot(f_list), print(point)
