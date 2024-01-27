#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/stat.h>
#include <seccomp.h>

int main() {
    
    printf("1764Snippets\n");
    printf("C seccomp\n");
    
    // Initialize the seccomp context
    scmp_filter_ctx ctx;
    ctx = seccomp_init(SCMP_ACT_KILL); // default action: kill the process
    if (ctx == NULL) {
        perror("seccomp_init failed");
        return -1;
    }

    // Allow read(), write(), and exit() syscalls
    if (seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(read), 0) < 0 ||
        seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(write), 0) < 0 ||
        seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(exit), 0) < 0) {
        perror("seccomp_rule_add failed");
        seccomp_release(ctx);
        return -1;
    }

    // Load the filter into the kernel
    if (seccomp_load(ctx) < 0) {
        perror("seccomp_load failed");
        seccomp_release(ctx);
        return -1;
    }

    // Now the process is restricted to the allowed syscalls
    printf("Seccomp filter loaded. Process is now restricted.\n");

    // Example of allowed syscall (write)
    write(STDOUT_FILENO, "1764Snippets write\n", 14);
    printf("\n");

    // Trying a "Bad system call"
    printf("Attempting to create a directory...\n");
    if (mkdir("test_directory", 0777) == -1) {
        perror("mkdir failed");
    } else {
        printf("Directory created successfully.\n");
    }

    // Clean up
    seccomp_release(ctx);

    return 0;
}
