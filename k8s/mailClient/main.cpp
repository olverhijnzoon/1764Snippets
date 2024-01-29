#include <curl/curl.h>
#include <string>
#include <unistd.h>
#include <iostream>

/*
    This function will be used as a callback for libcurl's CURLOPT_READFUNCTION option. The parameters are as follows:

    void* ptr       ..  A buffer to which the data should be written.
    size_t size     ..  Size of each data unit.
    size_t nmemb    ..  Number of units to transfer. The total size to be written is size * nmemb.
    void* userp     ..  Custom pointer set by CURLOPT_READDATA to pass data to the function.
*/
size_t ReadCallback(void* ptr, size_t size, size_t nmemb, void* userp) {

    // This line takes the userp pointer, which is a void pointer, and casts it to a std::string* pointer, which means it's now pointing to a std::string object. We can use s to access and manipulate the string object in this function.
    std::string* s = static_cast<std::string*>(userp);

    // This checks if the buffer size (size * nmemb) is less than 1, which means there's no room to write data, or if the string s is empty. If either condition is true, the function returns 0 to indicate that no data was written.
    if (size * nmemb < 1 || s->empty())
        return 0;

    // This line takes the first character from the string s (retrieved by s->front()) and puts it in the buffer ptr. The ptr pointer is first cast to a char* pointer, then dereferenced to store the character. Since the size of a char is 1 byte, this operation is safe as long as size * nmemb >= 1.
    *(static_cast<char*>(ptr)) = s->front();

    // This line removes the first character from the string s that we've just written to the buffer as it is supposed to "consume" the data it reads.
    s->erase(0, 1);

    // Finally, the function returns 1 to indicate that 1 unit of data was written. In this case, 1 unit corresponds to 1 character. This return value is important, as it tells libcurl how much data was transferred so it can update its counters and continue with the remaining data.
    return 1;
}

int main() {
    std::cout << "# 1764Snippets" << std::endl;
    std::cout << "## Mail & Curl" << std::endl;

    /*
        The callback function is used as a custom data provider for libcurl's HTTP POST or FTP upload operations (in this specific case, for sending an email using SMTP protocol). It reads the contents of a std::string one character at a time, placing each character in the buffer provided by libcurl. This allows you to send a std::string as the body of the email using SMTP. The function also returns the number of units actually written (in this case, 1 for each character or 0 if there are no more characters), so libcurl knows when all data has been sent.
    */

    CURL *curl;
    CURLcode res;

    curl_global_init(CURL_GLOBAL_ALL);

    while(true) {
        curl = curl_easy_init();
        if(curl) {
            struct curl_slist *recipients = NULL;
            std::string from = "1764Snippets Sender <from@mailhog.example>";
            std::string to = "to@mailhog.example";
            std::string mailbody = "Subject: 1764Snippets\r\n"
                                   "\r\n"
                                   "This is a test email.";

            // Set up connection parameters
            curl_easy_setopt(curl, CURLOPT_URL, "smtp://mailhog-smtp-service.default.svc.cluster.local:1025");
            //curl_easy_setopt(curl, CURLOPT_URL, "smtp://localhost:31025");
            curl_easy_setopt(curl, CURLOPT_MAIL_FROM, from.c_str());
            recipients = curl_slist_append(recipients, to.c_str());
            curl_easy_setopt(curl, CURLOPT_MAIL_RCPT, recipients);

            // Set up email body
            curl_easy_setopt(curl, CURLOPT_READDATA, &mailbody);
            curl_easy_setopt(curl, CURLOPT_READFUNCTION, ReadCallback);
            curl_easy_setopt(curl, CURLOPT_UPLOAD, 1L);

            // Send the email
            res = curl_easy_perform(curl);

            // Check for errors
            if(res != CURLE_OK) {
                fprintf(stderr, "curl_easy_perform() failed: %s\n", curl_easy_strerror(res));
            }

            // Clean up
            curl_slist_free_all(recipients);
            curl_easy_cleanup(curl);
        }

        sleep(42);
    }

    curl_global_cleanup();

    return (int)res;
}
