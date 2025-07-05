import grpc
import proto.kvstore_pb2 as pb2
import proto.kvstore_pb2_grpc as grpc2



def run():
    """
    Feito para envio em lotes ou consumo em lotes.
    """
    channel = grpc.insecure_channel('localhost:50051')
    stub = grpc2.KVStoreStub(channel)

    print("✅ Conectado ao servidor gRPC...")

    keys_values = [
        ("user:1", "Marcus"),
        ("user:2", "Anna"),
        ("user:3", "Lucas"),
    ]

    def request_generator():
        for key, val in keys_values:
            print(f"Enviando Set key={key} value={val}")
            yield pb2.SetRequest(key=key, value=val)

    responses = stub.StreamSet(request_generator())

    # é meio q obrigatorio consumir as respostas, msm q n faça nada pra n travar o stream.
    for response in responses:
        print(f"Resposta: key={response.key} success={response.success}")


if __name__ == "__main__":
    run()

