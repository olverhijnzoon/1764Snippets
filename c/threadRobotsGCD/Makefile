BINARY=threadRobotsGCD

run_c: build_c
	./$(BINARY)X

build_c:
	gcc -o $(BINARY)X $(BINARY).c

clean_c:
	rm -f $(BINARY)

debug_c:
	brew install llvm
	gcc -g -o $(BINARY)X $(BINARY).c
	lldb ./$(BINARY)X

	# Example
	#
	#	(lldb) breakpoint set --file threadRobotsGCD.c --line 42
	#	Breakpoint 1: where = threadRobotsGCDX`main + 272 at threadRobotsGCD.c:43:10, address = 0x0000000100003a60
	# 	(lldb) r
	#	* thread #1, queue = 'com.apple.main-thread', stop reason = breakpoint 1.1
	#		frame #0: 0x0000000100003a60 threadRobotsGCDX`main at threadRobotsGCD.c:43:10
	#	40           printf("Robot %d is idle\n",i+1);
	#	41       }
	#	42  
	#	-> 43       for(i=0;i<N;i++)
	#	44           pthread_join(thread_id[i],NULL);
	#	45   }
	#	46  
	#	(lldb) thread list
	#	Process 54651 stopped
	#	* thread #1: tid = 0x3703b5, 0x0000000100003a60 threadRobotsGCDX`main at threadRobotsGCD.c:43:10, queue = 'com.apple.main-thread', stop reason = breakpoint 1.1
	#	thread #2: tid = 0x370408, 0x0000000181594e68 libsystem_kernel.dylib`__semwait_signal + 8
	#	thread #3: tid = 0x370409, 0x0000000181594e68 libsystem_kernel.dylib`__semwait_signal + 8
	#	thread #4: tid = 0x37040a, 0x0000000181594e68 libsystem_kernel.dylib`__semwait_signal + 8
	#	thread #5: tid = 0x37040b, 0x0000000181594e68 libsystem_kernel.dylib`__semwait_signal + 8
	#	thread #6: tid = 0x37040c, 0x0000000181594e68 libsystem_kernel.dylib`__semwait_signal + 8
	#	(lldb) thread select 3
	#	* thread #3
	#		frame #0: 0x0000000181594e68 libsystem_kernel.dylib`__semwait_signal + 8
	#	libsystem_kernel.dylib`:
	#	->  0x181594e68 <+8>:  b.lo   0x181594e88               ; <+40>
	#		0x181594e6c <+12>: pacibsp 
	#		0x181594e70 <+16>: stp    x29, x30, [sp, #-0x10]!
	#		0x181594e74 <+20>: mov    x29, sp
	#	(lldb) thread backtrace
	#	* thread #3
	#	* frame #0: 0x0000000181594e68 libsystem_kernel.dylib`__semwait_signal + 8
	#		frame #1: 0x00000001814742f8 libsystem_c.dylib`nanosleep + 220
	#		frame #2: 0x000000018147d0bc libsystem_c.dylib`sleep + 52
	#		frame #3: 0x0000000100003b28 threadRobotsGCDX`robot(num=0x0000000100008004) at threadRobotsGCD.c:52:9
	#		frame #4: 0x00000001815d2034 libsystem_pthread.dylib`_pthread_start + 136
	#	(lldb) frame select 3
	#	frame #3: 0x0000000100003b28 threadRobotsGCDX`robot(num=0x0000000100008004) at threadRobotsGCD.c:52:9
	#	49       while(1) {
	#	50           int* i = num;
	#	51           printf("Robot %d is working\n",*i+1);
	#	-> 52           sleep(42);
	#	53           take_charger(*i);
	#	54           put_charger(*i);
	#	55       }
	#	(lldb) frame info
	#	frame #3: 0x0000000100003b28 threadRobotsGCDX`robot(num=0x0000000100008004) at threadRobotsGCD.c:52:9
	#	(lldb) frame variable
	#	(void *) num = 0x0000000100008004
	#	(int *) i = 0x0000000100008004
	#	(lldb) p *i
	#	(int) 1