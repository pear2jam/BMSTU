import numpy as np

def f(x):
    return np.log(np.abs(x - 0.5) + 1)

def fft(x):
    N = len(x)
    if N <= 1: return x
    even = fft(x[0::2])
    odd =  fft(x[1::2])
    T = [np.exp(-2j*np.pi*k/N)*odd[k] for k in range(N//2)]
    return [even[k] + T[k] for k in range(N//2)] + \
           [even[k] - T[k] for k in range(N//2)]


def ifft(x):
    N = len(x)
    if N <= 1: return x
    even = ifft(x[0::2])
    odd =  ifft(x[1::2])
    T= [np.exp(2j*np.pi*k/N)*odd[k] for k in range(N//2)]
    return [(even[k] + T[k])/2 for k in range(N//2)] + \
           [(even[k] - T[k])/2 for k in range(N//2)]

def fourier_series_approximation(x, yf):
    N = len(yf)
    n = np.arange(N)
    terms = np.array([2 * yf[k].real * np.cos(2 * np.pi * k * x / N) +
                      2 * yf[k].imag * np.sin(2 * np.pi * k * x / N) for k in n])
    return terms.sum(axis=0) / N

N = 128
x = np.linspace(0.0, N, N)
y = f(x)
yf = fft(y)
yi = ifft(yf)
print(yi)

y_interpolated = np.array([yi[int((0.5 + j) * N / N)].real for j in range(N)])

print("Изначальная функция:", y)
print("Интерполированные значения:", y_interpolated)
