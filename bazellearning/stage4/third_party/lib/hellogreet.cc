#include "third_party/lib/include/hellogreet.hpp"
#include <string>

std::string get_greet(const std::string& who) {
  return "Hello " + who;
}