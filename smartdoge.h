#ifndef SMARTDOGE_SMARTDOGE_H
#define SMARTDOGE_SMARTDOGE_H

#include <iostream>

class CSmartDoge
{
public:

    static const int32_t Version_Protocol_No_Routing = 0x0a;
    static const int32_t Version_Protocol_RouteNext4Bytes = 0x0b;
    static const int32_t Version_Protocol_EmojisVM = 0x0c;

    CSmartDoge() = default;

    void Serialize(int32_t Version_Protocol, int32_t OptionalVMRouting, char nData, char nChksum);
    void Unserialize(char serializedData);
    void EmojiParser(char emoji);
    void MapContextToEmojis(char context, char emoji);
    void ContractSigner();
    void VerifyContractSignature();




};


#endif //SMARTDOGE_SMARTDOGE_H
