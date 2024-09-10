using Pkg
using Plots
gr()

t = range(0, 20π, length=2000)


a = 2
b = 2
c = 1

X = a * cos.(t)
Y = a * sin.(t)
Z = b * t + c * t .* cos.(t)

plot(X, Y, Z)


gui()

readline()
