#ifndef SMARTDOGE_SMARTDOGE_H
#define SMARTDOGE_SMARTDOGE_H

#include <iostream>
#include <boost/asio.hpp>
using boost::asio::ip::tcp;

class CSmartDoge
{
public:

    static const int32_t Version_Protocol_No_Routing = 0x0a;
    static const int32_t Version_Protocol_RouteNext4Bytes = 0x0b;
    static const int32_t Version_Protocol_EmojisVM = 0x0c;

    CSmartDoge() = default;


    void EmojiParser(char emoji);
    void MapContextToEmojis(char context, char emoji);
    void ContractSigner();
    void VerifyContractSignature();




};


#endif //SMARTDOGE_SMARTDOGE_H
