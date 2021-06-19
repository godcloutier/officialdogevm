//
// Created by Nicolas Cloutier on 2021-06-19.
//

#ifndef SMARTDOGE_ENTRYPOINT_H
#define SMARTDOGE_ENTRYPOINT_H
#include "flatbuffers/flatbuffers.h"
#include "smartdoge.grpc.fb.h"
#include "smartdoge_generated.h"
#include <iostream> // C++ header file for printing
#include <fstream> // C++ header file for file access

namespace SmartDoge {
    namespace root {

        class entrypoint {

        };

    class StoreContractServiceImpl final : public StoreContract::Service {

        grpc::Status CompileAndStore(grpc::ServerContext *context,
                                             const flatbuffers::grpc::Message<ContractRawData> *request_msg,
                                             flatbuffers::grpc::Message<StoreContractResponse> *response_msg) override {
            flatbuffers::grpc::MessageBuilder mb_;

            const ContractRawData *request = request_msg->GetRoot();
            const uint8_t *ctx = request->context()->Data();
            const uint8_t *e = request->emoji_set()->Data();
            const uint8_t *owner = request->owner_address()->Data();
            const uint8_t *raw = request->raw_data()->Data();

            /*
             *  Compilation and storage logic below
             *
             */



            // Create Reply

            auto resp = CreateStoreContractResponse(
                    {0},
                    200,
                    {0},
                    ":):(:P",
                    {0}
                    );

            mb_.Finish(resp);

            *response_msg = mb_.ReleaseMessage<StoreContractResponse>();
            assert(response_msg->Verify());

            // return ok status code
            return grpc::Status::OK;


        }
    };

    }
}


#endif //SMARTDOGE_ENTRYPOINT_H
