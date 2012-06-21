def fibo():
	x, y = 0, 1
	c = 0
	while c < 20:
		print x
		x, y = y, x+y
		c += 1
