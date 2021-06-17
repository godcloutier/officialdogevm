//
// Created by Nicolas Cloutier on 2021-06-17.
//

#ifndef SMARTDOGE_WIREMESSAGE_H
#define SMARTDOGE_WIREMESSAGE_H

#include <cstdint>

unsigned char ProtocolVersions[3] = { 0x0a, 0x0b, 0x0c  };
static const int32_t Version_Protocol_No_Routing = ProtocolVersions[0];
    static const int32_t Version_Protocol_RouteNext4Bytes = ProtocolVersions[1];
    static const int32_t Version_Protocol_EmojisVM = ProtocolVersions[2];


class WireMessage {

};


#endif //SMARTDOGE_WIREMESSAGE_H
