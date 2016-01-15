package main 
import (
	"fmt"
	"runtime"
	//"time"
	)

var i int = 0

func thread_increase_i(channel chan string){
	
	for j := 0; j < 1000000; j++ {
		copy_i := <-channel
		i++
		//copy_i++
		//i = copy_i
		channel <- copy_i
	}
	

}

func thread_decrease_i(channel chan string){
	dummy := <- quit
	for j := 0; j < 999999; j++ {
		copy_i := <-channel
		i--
		//copy_i--
		//i = copy_i
		channel <- copy_i
	}
	quit <- dummy
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	channel := make(chan string, 1)
	channel <- "no"
	
	go thread_decrease_i(channel)
	go thread_increase_i(channel)

	time.Sleep(1000*time.Millisecond)
	fmt.Println(i)
}