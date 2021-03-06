#ifndef SMARTDOGE_SMARTDOGE_H
#define SMARTDOGE_SMARTDOGE_H

/*
 * Including generated flatbuffers files
 */
#include "smartdoge.grpc.fb.h"
#include "smartdoge_generated.h"
#include <grpcpp/impl/grpc_library.h>

#include <iostream>
#include <memory>
#include <string>



namespace SmartDoge {
    namespace root {

        class CSmartDoge
        {
        public:



            CSmartDoge() = default;


            void EmojiParser(char emoji);
            void MapContextToEmojis(char context, char emoji);





        };
    }
}



#endif //SMARTDOGE_SMARTDOGE_H
