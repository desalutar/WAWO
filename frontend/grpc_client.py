import grpc
from proto import auth_pb2_grpc

channel = grpc.insecure_channel("auth:50054")
stub = auth_pb2_grpc.AuthServiceStub(channel)
# chats = {}
# messages = []