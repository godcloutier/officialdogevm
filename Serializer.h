//
// Created by Nicolas Cloutier on 2021-06-17.
//

#ifndef SMARTDOGE_SERIALIZER_H
#define SMARTDOGE_SERIALIZER_H

#include <fstream>
#include <boost/archive/text_oarchive.hpp>
#include <boost/archive/text_iarchive.hpp>
#include "smartdoge.h"
#include <boost/archive/binary_iarchive.hpp>
#include <boost/archive/binary_oarchive.hpp>
#include <boost/iostreams/device/array.hpp>
#include <boost/iostreams/stream.hpp>
#include <iostream>
#include <ostream>
#include <array>
#include "Bytes.h"
using namespace boost;
using namespace std;



class Serializer {

public:


    char TypeCheckSum[1]{}; // 00 = default , 01 = segwit or 0a = smartdoge
    char ProtocolVersion[1]{};
    char OptionalFourBytesRouting[4]{};
    char DataBuffer[506]{};
    char OpScript[512]{};

    void ReadIncomingStream(tcp::socket & socket);
    void Serialize(int32_t Version_Protocol, int32_t OptionalVMRouting, char nData, char nChksum);

    Serializer() {
        InitRawDataDefaults();
    }


    void InitRawDataDefaults()
    {
        memcpy(this->TypeCheckSum, "0", sizeof(this->TypeCheckSum));
        memcpy(this->ProtocolVersion, "0", sizeof(this->ProtocolVersion));
        uint32_t default_array[4] = {0x00,0x00,0x00,0x00};
        memcpy(this->OptionalFourBytesRouting, "0", sizeof(this->OptionalFourBytesRouting));
        memcpy(this->DataBuffer, "0", sizeof(this->DataBuffer));
        memcpy(this->OpScript, "0", sizeof(this->OpScript));

    }

};


#endif //SMARTDOGE_SERIALIZER_H
