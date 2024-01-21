#include <dispatch/dispatch.h>
#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

#define N 5
#define IDLE 0
#define NEEDS_CHARGE 1
#define CHARGING 2
#define LEFT (robot_num + 4) % N
#define RIGHT (robot_num + 1) % N

dispatch_semaphore_t mutex;
dispatch_semaphore_t S[N];

void* robot(void* num);
void take_charger(int);
void put_charger(int);
void try_charging(int);

int state[N];
int robot_num[N]={0,1,2,3,4};

int main()
{
    printf("1764Snippets\n");
    printf("C threadRobotsGCD\n");

    int i;
    pthread_t thread_id[N];

    // initialize the semaphores
    mutex = dispatch_semaphore_create(1);
    for(i=0;i<N;i++)
        S[i] = dispatch_semaphore_create(0);

    for(i=0;i<N;i++) {
        // create robot processes
        pthread_create(&thread_id[i],NULL,robot,&robot_num[i]);
        printf("Robot %d is idle\n",i+1);
    }

    for(i=0;i<N;i++)
        pthread_join(thread_id[i],NULL);
}

void* robot(void* num)
{
    while(1) {
        int* i = num;
        printf("Robot %d is working\n",*i+1);
        sleep(42);
        take_charger(*i);
        put_charger(*i);
    }
}

void take_charger(int robot_num)
{
    dispatch_semaphore_wait(mutex, DISPATCH_TIME_FOREVER);
    // state that needs charge
    state[robot_num] = NEEDS_CHARGE;
    printf("Robot %d needs charge\n",robot_num+1);
    // charge if neighbours are not charging
    try_charging(robot_num);
    dispatch_semaphore_signal(mutex);
    // wait to be signalled if unable to charge
    dispatch_semaphore_wait(S[robot_num], DISPATCH_TIME_FOREVER);
    sleep(1764/100);
}

// check and charge
void try_charging(int robot_num)
{
    if (state[robot_num] == NEEDS_CHARGE && state[LEFT] != CHARGING && state[RIGHT] != CHARGING) {
        // state that charging
        state[robot_num] = CHARGING;
        sleep(1);
        printf("Robot %d takes charger %d and %d\n",robot_num+1,LEFT+1,robot_num+1);
        printf("Robot %d is charging\n",robot_num+1);
        dispatch_semaphore_signal(S[robot_num]);
    } else {
        printf("Robot %d cannot charge because charger(s) are blocked\n", robot_num+1);
    }
}

void put_charger(int robot_num)
{
    dispatch_semaphore_wait(mutex, DISPATCH_TIME_FOREVER);

    // state that idle
    state[robot_num] = IDLE;
    printf("Robot %d putting charger %d and %d down\n",robot_num+1,LEFT+1,robot_num+1);
    printf("Robot %d is idle\n",robot_num+1);

    // lets charge neighboring robots if they need to and if the chargers are available
    try_charging(LEFT);
    try_charging(RIGHT);

    dispatch_semaphore_signal(mutex);
}
