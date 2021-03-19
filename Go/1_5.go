package main

import "fmt"
import "math/big"

func main(){

	a11, a12, a21, a22 := big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)
	b11, b12, b21, b22 := big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(0)
	var n int32
	fmt.Scanln(&n)
	n -= 1
	ae, af, bg, bh, ce, cf, dg, dh := big.NewInt(0),big.NewInt(0),big.NewInt(0),big.NewInt(0),big.NewInt(0),big.NewInt(0),big.NewInt(0),big.NewInt(0)
	for ;n > 0; {
		if n % 2 == 1{
			ae.Set(a11)
			af.Set(a11)
			bg.Set(a12)
			bh.Set(a12)
			ce.Set(a21)
			cf.Set(a21)
			dg.Set(a22)
			dh.Set(a22)

			a11.Set(ae.Add(ae.Mul(ae, b11), bg.Mul(bg, b21)))
			a12.Set(af.Add(af.Mul(af, b12), bh.Mul(bh, b22)))
			a21.Set(ce.Add(ce.Mul(ce, b11), dg.Mul(dg, b21)))
			a22.Set(cf.Add(cf.Mul(cf, b12), dh.Mul(dh, b22)))
		}

		_b11, _b12, _b21, _b22 := big.NewInt(0),big.NewInt(0),big.NewInt(0),big.NewInt(0)
		ae.Set(b11)
		af.Set(b11)
		bg.Set(b12)
		bh.Set(b12)
		ce.Set(b21)
		cf.Set(b21)
		dg.Set(b22)
		dh.Set(b22)
		
		_b11.Set(b11)
		_b12.Set(b12)
		_b21.Set(b21)
		_b22.Set(b22)
		
		b11.Set(ae.Add(ae.Mul(ae, _b11), bg.Mul(bg, _b21)))
		b12.Set(af.Add(af.Mul(af, _b12), bh.Mul(bh, _b22)))
		b21.Set(ce.Add(ce.Mul(ce, _b11), dg.Mul(dg, _b21)))
		b22.Set(cf.Add(cf.Mul(cf, _b12), dh.Mul(dh, _b22)))
	
		n /= 2
	}
	fmt.Println(a11)
}