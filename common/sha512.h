//
// Created by Nicolas Cloutier on 2021-06-19.
//

#ifndef SMARTDOGE_SHA512_H
#define SMARTDOGE_SHA512_H

#include <utility>

#include "SHA512/SHA512.h"

class sha512 {

   /*
    * call with

        sha512::encode(std::move(str));

*/


    static std::basic_string<char, std::char_traits<char>, std::allocator<char>> encode(std::string str) {
        return SHA512().hash(std::move(str));
    }

};


#endif //SMARTDOGE_SHA512_H
