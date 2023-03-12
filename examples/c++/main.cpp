// g++ -o myapp main.cpp -lcurl -lssl -lcrypto

#include <iostream>
#include <curl/curl.h>
#include <sstream>
#include <string>

size_t write_data(void *ptr, size_t size, size_t nmemb, void *stream) {
  std::string data((const char*) ptr, (size_t) size * nmemb);
  *((std::stringstream*) stream) << data;
  return size * nmemb;
}

int main() {
  CURL *curl;
  CURLcode res;
  std::stringstream out;
  
  curl = curl_easy_init();
  if (curl) {
    curl_easy_setopt(curl, CURLOPT_URL, "http://localhost:8080/v1/cep/08226021");
    curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_data);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &out);
    res = curl_easy_perform(curl);
    if (res != CURLE_OK) {
      std::cerr << "Erro ao fazer a chamada GET: " << curl_easy_strerror(res) << std::endl;
    } else {
      std::cout << out.str() << std::endl;
    }
    curl_easy_cleanup(curl);
  }
  return 0;
}

