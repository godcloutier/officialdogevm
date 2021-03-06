// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: smartdoge
#ifndef GRPC_smartdoge__INCLUDED
#define GRPC_smartdoge__INCLUDED

#include "smartdoge_generated.h"
#include "flatbuffers/grpc.h"

#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/proto_utils.h>
#include <grpcpp/impl/codegen/rpc_method.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/status.h>
#include <grpcpp/impl/codegen/stub_options.h>
#include <grpcpp/impl/codegen/sync_stream.h>

namespace grpc {
class CompletionQueue;
class Channel;
class ServerCompletionQueue;
class ServerContext;
}  // namespace grpc

namespace SmartDoge {

class IpfsStorage final {
 public:
  static constexpr char const* service_full_name() {
    return "SmartDoge.IpfsStorage";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status Store(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, flatbuffers::grpc::Message<StoreResponse>* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreResponse>>> AsyncStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreResponse>>>(AsyncStoreRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreResponse>>> PrepareAsyncStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreResponse>>>(PrepareAsyncStoreRaw(context, request, cq));
    }
    virtual ::grpc::Status Retrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, flatbuffers::grpc::Message<RetrieveResponse>* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<RetrieveResponse>>> AsyncRetrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<RetrieveResponse>>>(AsyncRetrieveRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<RetrieveResponse>>> PrepareAsyncRetrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<RetrieveResponse>>>(PrepareAsyncRetrieveRaw(context, request, cq));
    }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreResponse>>* AsyncStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreResponse>>* PrepareAsyncStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<RetrieveResponse>>* AsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<RetrieveResponse>>* PrepareAsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status Store(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, flatbuffers::grpc::Message<StoreResponse>* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreResponse>>> AsyncStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreResponse>>>(AsyncStoreRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreResponse>>> PrepareAsyncStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreResponse>>>(PrepareAsyncStoreRaw(context, request, cq));
    }
    ::grpc::Status Retrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, flatbuffers::grpc::Message<RetrieveResponse>* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<RetrieveResponse>>> AsyncRetrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<RetrieveResponse>>>(AsyncRetrieveRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<RetrieveResponse>>> PrepareAsyncRetrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<RetrieveResponse>>>(PrepareAsyncRetrieveRaw(context, request, cq));
    }
  
   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreResponse>>* AsyncStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreResponse>>* PrepareAsyncStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<IpfsData>& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<RetrieveResponse>>* AsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<RetrieveResponse>>* PrepareAsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<FileKey>& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_Store_;
    const ::grpc::internal::RpcMethod rpcmethod_Retrieve_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
  
  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status Store(::grpc::ServerContext* context, const flatbuffers::grpc::Message<IpfsData>* request, flatbuffers::grpc::Message<StoreResponse>* response);
    virtual ::grpc::Status Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<FileKey>* request, flatbuffers::grpc::Message<RetrieveResponse>* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_Store : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_Store() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_Store() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Store(::grpc::ServerContext* context, const flatbuffers::grpc::Message<IpfsData>* request, flatbuffers::grpc::Message<StoreResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestStore(::grpc::ServerContext* context, flatbuffers::grpc::Message<IpfsData>* request, ::grpc::ServerAsyncResponseWriter< flatbuffers::grpc::Message<StoreResponse>>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_Retrieve : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_Retrieve() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_Retrieve() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<FileKey>* request, flatbuffers::grpc::Message<RetrieveResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestRetrieve(::grpc::ServerContext* context, flatbuffers::grpc::Message<FileKey>* request, ::grpc::ServerAsyncResponseWriter< flatbuffers::grpc::Message<RetrieveResponse>>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef   WithAsyncMethod_Store<  WithAsyncMethod_Retrieve<  Service   >   >   AsyncService;
  template <class BaseClass>
  class WithGenericMethod_Store : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_Store() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_Store() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Store(::grpc::ServerContext* context, const flatbuffers::grpc::Message<IpfsData>* request, flatbuffers::grpc::Message<StoreResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_Retrieve : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_Retrieve() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_Retrieve() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<FileKey>* request, flatbuffers::grpc::Message<RetrieveResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Store : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_Store() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler< flatbuffers::grpc::Message<IpfsData>, flatbuffers::grpc::Message<StoreResponse>>(std::bind(&WithStreamedUnaryMethod_Store<BaseClass>::StreamedStore, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_Store() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Store(::grpc::ServerContext* context, const flatbuffers::grpc::Message<IpfsData>* request, flatbuffers::grpc::Message<StoreResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedStore(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< flatbuffers::grpc::Message<IpfsData>,flatbuffers::grpc::Message<StoreResponse>>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Retrieve : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_Retrieve() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler< flatbuffers::grpc::Message<FileKey>, flatbuffers::grpc::Message<RetrieveResponse>>(std::bind(&WithStreamedUnaryMethod_Retrieve<BaseClass>::StreamedRetrieve, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_Retrieve() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<FileKey>* request, flatbuffers::grpc::Message<RetrieveResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedRetrieve(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< flatbuffers::grpc::Message<FileKey>,flatbuffers::grpc::Message<RetrieveResponse>>* server_unary_streamer) = 0;
  };
  typedef   WithStreamedUnaryMethod_Store<  WithStreamedUnaryMethod_Retrieve<  Service   >   >   StreamedUnaryService;
  typedef   Service   SplitStreamedService;
  typedef   WithStreamedUnaryMethod_Store<  WithStreamedUnaryMethod_Retrieve<  Service   >   >   StreamedService;
};

class StoreContract final {
 public:
  static constexpr char const* service_full_name() {
    return "SmartDoge.StoreContract";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status CompileAndStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, flatbuffers::grpc::Message<StoreContractResponse>* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreContractResponse>>> AsyncCompileAndStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreContractResponse>>>(AsyncCompileAndStoreRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreContractResponse>>> PrepareAsyncCompileAndStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreContractResponse>>>(PrepareAsyncCompileAndStoreRaw(context, request, cq));
    }
    virtual ::grpc::Status Retrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, flatbuffers::grpc::Message<ContractRequestResponse>* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<ContractRequestResponse>>> AsyncRetrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<ContractRequestResponse>>>(AsyncRetrieveRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<ContractRequestResponse>>> PrepareAsyncRetrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<ContractRequestResponse>>>(PrepareAsyncRetrieveRaw(context, request, cq));
    }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreContractResponse>>* AsyncCompileAndStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<StoreContractResponse>>* PrepareAsyncCompileAndStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<ContractRequestResponse>>* AsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< flatbuffers::grpc::Message<ContractRequestResponse>>* PrepareAsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status CompileAndStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, flatbuffers::grpc::Message<StoreContractResponse>* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreContractResponse>>> AsyncCompileAndStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreContractResponse>>>(AsyncCompileAndStoreRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreContractResponse>>> PrepareAsyncCompileAndStore(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreContractResponse>>>(PrepareAsyncCompileAndStoreRaw(context, request, cq));
    }
    ::grpc::Status Retrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, flatbuffers::grpc::Message<ContractRequestResponse>* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<ContractRequestResponse>>> AsyncRetrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<ContractRequestResponse>>>(AsyncRetrieveRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<ContractRequestResponse>>> PrepareAsyncRetrieve(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<ContractRequestResponse>>>(PrepareAsyncRetrieveRaw(context, request, cq));
    }
  
   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreContractResponse>>* AsyncCompileAndStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<StoreContractResponse>>* PrepareAsyncCompileAndStoreRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRawData>& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<ContractRequestResponse>>* AsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< flatbuffers::grpc::Message<ContractRequestResponse>>* PrepareAsyncRetrieveRaw(::grpc::ClientContext* context, const flatbuffers::grpc::Message<ContractRequest>& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_CompileAndStore_;
    const ::grpc::internal::RpcMethod rpcmethod_Retrieve_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
  
  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status CompileAndStore(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRawData>* request, flatbuffers::grpc::Message<StoreContractResponse>* response);
    virtual ::grpc::Status Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRequest>* request, flatbuffers::grpc::Message<ContractRequestResponse>* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_CompileAndStore : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_CompileAndStore() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_CompileAndStore() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status CompileAndStore(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRawData>* request, flatbuffers::grpc::Message<StoreContractResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestCompileAndStore(::grpc::ServerContext* context, flatbuffers::grpc::Message<ContractRawData>* request, ::grpc::ServerAsyncResponseWriter< flatbuffers::grpc::Message<StoreContractResponse>>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_Retrieve : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_Retrieve() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_Retrieve() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRequest>* request, flatbuffers::grpc::Message<ContractRequestResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestRetrieve(::grpc::ServerContext* context, flatbuffers::grpc::Message<ContractRequest>* request, ::grpc::ServerAsyncResponseWriter< flatbuffers::grpc::Message<ContractRequestResponse>>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef   WithAsyncMethod_CompileAndStore<  WithAsyncMethod_Retrieve<  Service   >   >   AsyncService;
  template <class BaseClass>
  class WithGenericMethod_CompileAndStore : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_CompileAndStore() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_CompileAndStore() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status CompileAndStore(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRawData>* request, flatbuffers::grpc::Message<StoreContractResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_Retrieve : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_Retrieve() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_Retrieve() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRequest>* request, flatbuffers::grpc::Message<ContractRequestResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_CompileAndStore : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_CompileAndStore() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler< flatbuffers::grpc::Message<ContractRawData>, flatbuffers::grpc::Message<StoreContractResponse>>(std::bind(&WithStreamedUnaryMethod_CompileAndStore<BaseClass>::StreamedCompileAndStore, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_CompileAndStore() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status CompileAndStore(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRawData>* request, flatbuffers::grpc::Message<StoreContractResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedCompileAndStore(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< flatbuffers::grpc::Message<ContractRawData>,flatbuffers::grpc::Message<StoreContractResponse>>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Retrieve : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_Retrieve() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler< flatbuffers::grpc::Message<ContractRequest>, flatbuffers::grpc::Message<ContractRequestResponse>>(std::bind(&WithStreamedUnaryMethod_Retrieve<BaseClass>::StreamedRetrieve, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_Retrieve() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Retrieve(::grpc::ServerContext* context, const flatbuffers::grpc::Message<ContractRequest>* request, flatbuffers::grpc::Message<ContractRequestResponse>* response) final override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedRetrieve(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< flatbuffers::grpc::Message<ContractRequest>,flatbuffers::grpc::Message<ContractRequestResponse>>* server_unary_streamer) = 0;
  };
  typedef   WithStreamedUnaryMethod_CompileAndStore<  WithStreamedUnaryMethod_Retrieve<  Service   >   >   StreamedUnaryService;
  typedef   Service   SplitStreamedService;
  typedef   WithStreamedUnaryMethod_CompileAndStore<  WithStreamedUnaryMethod_Retrieve<  Service   >   >   StreamedService;
};

}  // namespace SmartDoge


#endif  // GRPC_smartdoge__INCLUDED
