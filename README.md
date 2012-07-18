GoPy is an experimental Python interpreter implementation in Go. It parses produced Python bytecode and executes it. 

The following python code (see src/main/testapps/example.py)...

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

... would be executed as ...

	GoPy 0.1 - (C) 2012 Florian Schlachter, Berlin
	2012/07/18 22:03:20 [Settings filename=testapps/example.pyc]
	2012/07/18 22:03:20 Parsing...
	2012/07/18 22:03:20 File created: 2012-07-18 22:03:05 +0200 CEST (timestamp: 1342641785)
	2012/07/18 22:03:20 Parsing finished
	2012/07/18 22:03:20 VM created [filename=example.py,name=<module>]
	2012/07/18 22:03:20 Running...
	2012/07/18 22:03:20 [VM] Stacksize = 3
	5 * 10 = 50 
	Unix time? 1342641800 
	Hello Go 0 ! 
	Hello Go 1 ! 
	Hello Go 2 ! 
	Hello Go 3 ! 
	Hello Go 4 ! 
	2012/07/18 22:03:20 [VM] Returning value: None (*vm.PyNone)
	2012/07/18 22:03:20 [VM] Execution of program took 1.252ms.
	2012/07/18 22:03:20 Running finished (166 instructions ran).

I created GoPy to get in touch with the Go language and learn its basics. It's currently not more than an experimental approach and has some bugs which must be fixed in order to get more complicated programs running.
