#define _GNU_SOURCE
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>

// Define the size of the stack for the child process as 1 megabyte
#define STACK_SIZE (1024 * 1024)

/*
    CLONE_NEWNET is a flag used with the clone system call in Linux to indicate that the cloned process should be created in a new network namespace. This is part of the Linux namespaces functionality, which is a feature of the Linux kernel that isolates and virtualizes system resources for a collection of processes.

    The value 0x40000000 is the specific bit in the flags argument of the clone system call that corresponds to CLONE_NEWNET. This value is defined in the Linux kernel headers, but it's not defined in the standard C library headers, so we need to define it manually in our program.

    When you pass CLONE_NEWNET as part of the flags argument to clone, the kernel will create the new process in a new network namespace, which has its own network devices, IP routing tables, port numbers, etc. This is used to provide network isolation for the new process.
*/
#define CLONE_NEWNET 0x40000000

static char child_stack[STACK_SIZE];

char* const child_args[] = {
    "/bin/bash",
    NULL
};

int child_main(void* arg) {
    printf("###\n");
    printf("Network interfaces in new network namespace:\n");
    system("ip addr");
    return 0;
}

int main() {

    printf("1764Snippets\n");
    printf("C netNamespace\n");

    /*
    This 1764Snippet prints the network interfaces in the initial network namespace before creating a new network namespace. Then, in the child_main function, it prints the network interfaces in the new network namespace.

    When you run this program, you should see that the initial network namespace has several network interfaces (including lo and possibly eth0, wlan0, etc.), while the new network namespace has only the loopback interface lo. This shows that the new network namespace is isolated from the host's network.

    This is a slightly modified version of a 42Snippets snippet.
    */

    system("ip addr");
    int flags = CLONE_NEWNET;
    // Create a new process
    pid_t pid = clone(
        child_main,  // The function to be executed in the new child process
        child_stack + STACK_SIZE,  // The location of the stack for the new child process
        flags | SIGCHLD,  // Flags controlling how the clone function works. CLONE_NEWNET creates a new network namespace, and SIGCHLD sends a SIGCHLD signal when the child process terminates
        NULL  // The argument to be passed to the child_main function. In this case, no argument is passed
    );
    if (pid < 0) {
        perror("clone failed");
        exit(1);
    }
    waitpid(pid, NULL, 0);
    return 0;
}
