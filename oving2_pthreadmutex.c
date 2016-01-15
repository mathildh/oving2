#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
int i = 0;
pthread_mutex_t mymutex;
void* thread_increase_i(){
	int j;
	for (j = 0; j < 999999; j++){
		pthread_mutex_lock(&mymutex);
		i++;
		pthread_mutex_unlock(&mymutex);
	}
	
}
void* thread_decrease_i(){
	int j;
	for (j = 0; j < 1000000; j++){
		pthread_mutex_lock(&mymutex);
		i--;
		pthread_mutex_unlock(&mymutex);
	}
}

int main(){
	
	printf("Hello from main!\n");
	pthread_t thread_1;
	pthread_t thread_2;
	pthread_mutex_init(&mymutex,NULL);
	
	pthread_create(&thread_1,NULL,thread_increase_i,NULL);
	pthread_create(&thread_2,NULL,thread_decrease_i,NULL);
	pthread_join(thread_1, NULL);
	pthread_join(thread_2, NULL);

	pthread_mutex_destroy(&mymutex);

	printf("Done, result: %i\n",i); 

}