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



