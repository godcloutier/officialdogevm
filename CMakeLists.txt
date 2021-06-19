cmake_minimum_required(VERSION 3.17)
project(smartdoge)

set(CMAKE_CXX_STANDARD 11)

set(BOOST_ROOT "/usr/local/include/boost")
include_directories(/usr/local/include)
link_directories(/usr/local/lib)
add_library(smartdoge SHARED smartdoge.cpp smartdoge.h Bytes.h common/uint256_t.h common/sha256.cpp common/sha256.h common/sha512.cpp common/sha512.h constants.h entrypoint.cpp entrypoint.h)