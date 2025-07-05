import grpc
import proto.kvstore_pb2 as pb2
import proto.kvstore_pb2_grpc as grpc2


def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = grpc2.KVStoreStub(channel)

    print("✅ Conectado ao servidor gRPC...")

    key = "user:123"
    value = "Marcus"
    set_response = stub.Set(pb2.SetRequest(key=key, value=value))
    print("📝 Set:", set_response.success)

    get_response = stub.Get(pb2.GetRequest(key=key))
    print("📥 Get:", get_response.value)

    del_response = stub.Del(pb2.DelRequest(key=key))
    print("🗑️ Del:", del_response.value)


if __name__ == "__main__":
    run()
