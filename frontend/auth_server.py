import grpc
from concurrent import futures
from proto import auth_pb2_grpc, auth_pb2
from google.protobuf import empty_pb2

class AuthService(auth_pb2_grpc.AuthServiceServicer):

    def Register(self, request, context):
        print(f"Register user: {request.username}")
        return empty_pb2.Empty()

    def Login(self, request, context):
        print(f"Login user: {request.username}")
        return auth_pb2.AuthResponse(token="fake_token", expires_at=9999999999)

    def Validate(self, request, context):
        print(f"Validate token: {request.token}")
        return auth_pb2.ValidateResponse(valid=True, user="test_user")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    auth_pb2_grpc.add_AuthServiceServicer_to_server(AuthService(), server)
    server.add_insecure_port('[::]:50054')
    server.start()
    print("gRPC server running on port 50054")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
