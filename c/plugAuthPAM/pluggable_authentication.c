#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <security/pam_appl.h>
#include <security/pam_misc.h>

// Define conversation function
static struct pam_conv conv = {
    misc_conv,
    NULL
};

int main(int argc, char *argv[]) {
    
    printf("1764Snippets\n");
    printf("C Pluggable Authentication Modules (PAM)\n");

    /*
        Primary aim of Pluggable Authentication Modules (PAM) is to provide a flexible and modular mechanism for authentication and related services, beyond just password-based authentication.
    
        However, this snippet is password-based as a draft for a PAM snippet.
    */
    
    pam_handle_t *pamh = NULL;
    int retval;
    const char *user = "snippet1764";

    if(argc == 2) {
        user = argv[1];
    }

    // Start PAM
    retval = pam_start("check_user", user, &conv, &pamh);

    // Authentication
    if(retval == PAM_SUCCESS)
        retval = pam_authenticate(pamh, 0);    // Authenticate the user

    // Account management
    if(retval == PAM_SUCCESS)
        retval = pam_acct_mgmt(pamh, 0);   // Check if the user's account is valid

    // Clean up
    if(pam_end(pamh,retval) != PAM_SUCCESS) {  
        pamh = NULL;
        fprintf(stderr, "check_user: failed to release authenticator\n");
        exit(1);
    }

    // Evaluate the result
    if(retval == PAM_SUCCESS) {
        fprintf(stdout, "Authentication successful\n");
    } else {
        fprintf(stderr, "Authentication failed: %s\n", pam_strerror(pamh, retval));
    }

    return (retval == PAM_SUCCESS ? 0 : 1);
}
