cmake_minimum_required(VERSION 3.17)
project(smartdoge)

set(CMAKE_CXX_STANDARD 11)

set(BOOST_ROOT "/usr/local/include/boost")
include_directories(/usr/local/include)
link_directories(/usr/local/lib)
add_library(smartdoge SHARED smartdoge.cpp smartdoge.h WireMessage.cpp WireMessage.h Bytes.h)