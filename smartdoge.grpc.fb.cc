// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: smartdoge

#include "smartdoge_generated.h"
#include "smartdoge.grpc.fb.h"

#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace SmartDoge {

static const char* IpfsStorage_method_names[] = {
  "/SmartDoge.IpfsStorage/Store",
  "/SmartDoge.IpfsStorage/Retrieve",
};

std::unique_ptr< IpfsStorage::Stub> IpfsStorage::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< IpfsStorage::Stub> stub(new IpfsStorage::Stub(channel));
  return stub;
}

IpfsStorage::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel)  , rpcmethod_Store_(IpfsStorage_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Retrieve_(IpfsStorage_method_names[1], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}
  
::grpc::Status IpfsStorage::Stub::Store(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, flatbuffers::grpc::Message<StoreResponse>* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Store_, context, request, response);
}

::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreResponse>>* IpfsStorage::Stub::AsyncStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< flatbuffers::grpc::Message<StoreResponse>>::Create(channel_.get(), cq, rpcmethod_Store_, context, request, true);
}

::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreResponse>>* IpfsStorage::Stub::PrepareAsyncStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< flatbuffers::grpc::Message<StoreResponse>>::Create(channel_.get(), cq, rpcmethod_Store_, context, request, false);
}

::grpc::Status IpfsStorage::Stub::Retrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, flatbuffers::grpc::Message<RetrieveResponse>* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Retrieve_, context, request, response);
}

::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<RetrieveResponse>>* IpfsStorage::Stub::AsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< flatbuffers::grpc::Message<RetrieveResponse>>::Create(channel_.get(), cq, rpcmethod_Retrieve_, context, request, true);
}

::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<RetrieveResponse>>* IpfsStorage::Stub::PrepareAsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< flatbuffers::grpc::Message<RetrieveResponse>>::Create(channel_.get(), cq, rpcmethod_Retrieve_, context, request, false);
}

IpfsStorage::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      IpfsStorage_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< IpfsStorage::Service, flatbuffers::grpc::Message<IpfsData>, flatbuffers::grpc::Message<StoreResponse>>(
          std::mem_fn(&IpfsStorage::Service::Store), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      IpfsStorage_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< IpfsStorage::Service, flatbuffers::grpc::Message<FileKey>, flatbuffers::grpc::Message<RetrieveResponse>>(
          std::mem_fn(&IpfsStorage::Service::Retrieve), this)));
}

IpfsStorage::Service::~Service() {
}

::grpc::Status IpfsStorage::Service::Store(::grpc::ServerContext* context, const flatbuffers::grpc::Message<IpfsData>* request, flatbuffers::grpc::Message<StoreResponse>* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status IpfsStorage::Service::Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<FileKey>* request, flatbuffers::grpc::Message<RetrieveResponse>* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


static const char* StoreContract_method_names[] = {
  "/SmartDoge.StoreContract/CompileAndStore",
  "/SmartDoge.StoreContract/Retrieve",
};

std::unique_ptr< StoreContract::Stub> StoreContract::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< StoreContract::Stub> stub(new StoreContract::Stub(channel));
  return stub;
}

StoreContract::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel)  , rpcmethod_CompileAndStore_(StoreContract_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Retrieve_(StoreContract_method_names[1], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}
  
::grpc::Status StoreContract::Stub::CompileAndStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, flatbuffers::grpc::Message<StoreContractResponse>* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_CompileAndStore_, context, request, response);
}

::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreContractResponse>>* StoreContract::Stub::AsyncCompileAndStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< flatbuffers::grpc::Message<StoreContractResponse>>::Create(channel_.get(), cq, rpcmethod_CompileAndStore_, context, request, true);
}

::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreContractResponse>>* StoreContract::Stub::PrepareAsyncCompileAndStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< flatbuffers::grpc::Message<StoreContractResponse>>::Create(channel_.get(), cq, rpcmethod_CompileAndStore_, context, request, false);
}

::grpc::Status StoreContract::Stub::Retrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, flatbuffers::grpc::Message<ContractRequestResponse>* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Retrieve_, context, request, response);
}

::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<ContractRequestResponse>>* StoreContract::Stub::AsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< flatbuffers::grpc::Message<ContractRequestResponse>>::Create(channel_.get(), cq, rpcmethod_Retrieve_, context, request, true);
}

::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<ContractRequestResponse>>* StoreContract::Stub::PrepareAsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< flatbuffers::grpc::Message<ContractRequestResponse>>::Create(channel_.get(), cq, rpcmethod_Retrieve_, context, request, false);
}

StoreContract::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      StoreContract_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< StoreContract::Service, flatbuffers::grpc::Message<ContractRawData>, flatbuffers::grpc::Message<StoreContractResponse>>(
          std::mem_fn(&StoreContract::Service::CompileAndStore), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      StoreContract_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< StoreContract::Service, flatbuffers::grpc::Message<ContractRequest>, flatbuffers::grpc::Message<ContractRequestResponse>>(
          std::mem_fn(&StoreContract::Service::Retrieve), this)));
}

StoreContract::Service::~Service() {
}

::grpc::Status StoreContract::Service::CompileAndStore(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRawData>* request, flatbuffers::grpc::Message<StoreContractResponse>* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status StoreContract::Service::Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRequest>* request, flatbuffers::grpc::Message<ContractRequestResponse>* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace SmartDoge
