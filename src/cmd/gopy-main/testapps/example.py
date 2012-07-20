import time

def hello(name, counter):
    print "Hello", name, counter, "!"    

a = 5
b = 10
print a, "*", b, "=", a*b
print "Unix time?", time.time()

i = 0
while i < 5:
	hello("Go", i)
	i += 1
