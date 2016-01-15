from threading import Thread 
from threading import Lock

i = 0
mutex = Lock()

def thread_increase_i():
	global i
	for j in range(0,1000000):
		mutex.acquire()
		i += 1 
		mutex.release()

def thread_decrease_i():
	global i
	for j in range(0,999999):
		mutex.acquire()
		i -= 1
		mutex.release()

def main():
	thread_add = Thread(target = thread_increase_i,args = (),)
	thread_sub = Thread(target = thread_decrease_i,args = (),)
	thread_add.start()
	thread_sub.start()
	thread_add.join()
	thread_sub.join()

	print i

main()
