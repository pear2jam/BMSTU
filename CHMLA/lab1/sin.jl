using Pkg
using Plots
gr()

x = range(-2π, 2π, length=100)
y = sin.(x)

plot(x, y, title="Это график")

display(plot)

gui()
readline()
