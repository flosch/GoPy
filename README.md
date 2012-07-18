GoPy is an experimental Python interpreter implementation in Go. It parses produced Python bytecode and executes it. 

The following python code...

	import time
	a = 5
	b = 10
	print a, "*", b, "=", a*b
	print "Unix time?", time.time()

... would be executed as ...

	$ go run main.go -filename mytest.pyc 
	GoPy 0.1 - (C) 2012 Florian Schlachter, Berlin
	2012/07/16 20:44:34 [Settings filename=mytest.pyc]
	2012/07/16 20:44:34 Parsing...
	2012/07/16 20:44:34 File created: 2012-07-16 20:44:12 +0200 CEST (timestamp: 1342464252)
	2012/07/16 20:44:34 Parsing finished
	2012/07/16 20:44:34 VM created [filename=mytest.py,name=<module>]
	2012/07/16 20:44:34 Running...
	2012/07/16 20:44:34 [VM] Stacksize = 2
	5 * 10 = 50 
	Unix time? 1342634208 
	2012/07/16 20:44:34 [VM] Returning value: None (*vm.PyNone)
	2012/07/16 20:44:34 [VM] Execution of program took 135us.
	2012/07/16 20:44:34 Running finished (19 instructions ran).

I created GoPy to get in touch with the Go language and learn its basics. It's currently not more than an experimental approach and has some bugs which must be fixed in order to get more complicated programs running.
