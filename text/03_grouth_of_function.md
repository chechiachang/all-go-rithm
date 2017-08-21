Chpt3. Growth of Function
===

# Theta-notation - asymptotical tight bound

Sandswitched by c1,c2
require f(n) asymptotically nonnegtive
T(n) = theta(n^2)

theta(g(n)) = {
	f(n), exists c1,c2,n0 such that 0 <= c1g(n) <= f(n) <= c2g(n) for n >=n0
}

example.1	1/2n^2 -3n = theta(n^2)

find a set of (c1, c2) such that c1n^2 <= 1/2n^2 - 3n <= c2n^2
	-> c1 <= 1/2 - 3/n <= c2
	-> take n >= 1, c1 = 1/14 <= 1/2 -3/n <= 1/2 = c2 
	-> for n >= 1, there exists (c1, c2) = (1/2, 1/14) such that ...

example.2	6n^3 =/= theta(n^2)
	-> there exists no c2 such that 6n^3 <= c2n^2 for all n larger enough

# O notation - asymptotical upper bound
T(n) = O(n)

O(g(n)) = {
	f(n), there exists c, n0 such that 0 <= f(n) <= cg(n) for all n >= n0
}


# Monotonicity

f(n) is monotonically increasing if m <= n implies	f(m) <= f(n)
		monotonically decreasing					f(m) >= f(n)
		strictly increasing		 if m < n implies	f(m) < f(n)
				 decreasing							f(m) > f(n)

# Floors & ceilings

x - 1 < floor(x) <= x <= ceiling(x) <= x + 1

for n <- integer
	f(n/2) + c(n/2) = n

for n <- real number
	c( c(n/a) / b ) = c( n/ab )
	f( f(n/a) / b ) = f( n/ab )

	c(a/b) <= (a + (b - 1))/b
	f(a/b) <= (a - (b = 1))/b

# Modular arithmetic

a mod n = a - f(a/n)n

floor(a/n) = a/n - c, where c is the remainer

n is remainder or residue

a % n = b % n 
-> a -= b ( mod n ), a is equivilent to b, modulo to n

# Polynominal

a polynominal in n of degree d
i <- range 0 to d {
	return += ai * n^i
}

coefficients a0, ... ai,... ad, where ad =/= 0

f(n) is polynominally bound if f(n) = O(n^k)

# Exponentials

-> lim (n -> infinity) n^b / a^n
exponential beats polynominal -> n^b = o(a^n)

e^x = lim(0, inf) x^i/i! = 1 + x + x^2/2! + ...
e^x >= 1 + x

1 + x <= e^x <= 1 + x + x ^2

lim(n-> inf) (1 + x/n)^n = e^x

# Logarithms


a = b^logb\a

	log a = log (b^logb\a) = logb * logb\a = logb * loga/logb = loga

logc\ab = logc\a + logc\b

	logc\ab = log(ab) / logc = loga/logc + logb/logc = logc\a + logc\b

logb\a^n = nlogb\a

	logb\a^n = logb(a * a * a ...* a) = nlogb\a

# Factorial

Stirling approximation:
	
		n! -> (2Pin)^-2 * (n/e)^n





