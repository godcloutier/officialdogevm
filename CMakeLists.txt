cmake_minimum_required(VERSION 3.17)
project(smartdoge)

set(CMAKE_CXX_STANDARD 11)

add_library(smartdoge SHARED smartdoge.cpp smartdoge.h WireMessage.cpp WireMessage.h)