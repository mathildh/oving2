package main 
import (
	"fmt"
	"runtime"
	)


func thread_increase_i(channel chan int, finished chan int){
	for j := 0; j < 1000000; j++ {
		channel <- 1
	}
	finished <- 1
	
}

func thread_decrease_i(channel chan int, finished chan int){
	for j := 0; j < 999999; j++ {
		channel <- 1
	}
	finished <- 1
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	i := 0
	channel_increase := make(chan int)
	channel_decrease := make(chan int)
	channel_finished1 := make(chan int)
	channel_finished2 := make(chan int)

	channel1Finished := bool(false)
	channel2Finished := bool(false)

	go thread_decrease_i(channel_decrease, channel_finished1)
	go thread_increase_i(channel_increase, channel_finished2)

	for !(channel1Finished && channel2Finished) {
		select{
			case <-channel_increase:
				i++
			case <-channel_decrease:
				i--
			case <-channel_finished1:
				channel1Finished = true
			case <-channel_finished2:
				channel2Finished = true
		}
		//if channel1Finished && channel2Finished{
		//	break
		//}
	}
	
	fmt.Println(i)
}