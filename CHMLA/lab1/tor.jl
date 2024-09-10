using Pkg
using Plots
gr()


u = range(0, 2π, length=100)
v = range(0, 2π, length=100)

U = repeat(u, 1, length(v))
V = repeat(v', length(u), 1)

R = 3 
r = 1

X = (R .+ r .* cos.(V)) .* cos.(U)
Y = (R .+ r .* cos.(V)) .* sin.(U)
Z = r .* sin.(V)

plot(X, Y, Z, st=:surface)

gui()

readline()
