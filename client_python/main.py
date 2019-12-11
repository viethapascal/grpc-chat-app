import hashlib
import logging
import argparse
import sys
import time
from multiprocessing import Process

import grpc

import service_pb2
import service_pb2_grpc


def gen_encrypt_string(name):
    ts = int(time.time())
    id = str(ts) + name

    sha_signature = hashlib.sha256(id.encode()).hexdigest()
    return sha_signature

class Client:
    def __init__(self):

        parser = argparse.ArgumentParser("Simple argument parser")
        parser.add_argument(
            "-N",
            help="User name",
            default="Python Client"
        )
        self.args = parser.parse_args()

        self.user = self.create_user()
        self.connect = self.create_connect()


    def create_user(self):
        name = self.args.N
        id = gen_encrypt_string(name)

        return service_pb2.User(
            id=id,
            display_name=name
        )

    def create_connect(self):
        return service_pb2.Connect(
            user=self.user,
            active=True
        )

    def receive_message(self, stream):
        res = stream.CreateStream(self.connect)
        print self.connect
        while True:
            for r in res:
                print r.user.display_name + ": "+r.message

    def broadcast_message(self, stub):
        while True:
            content = sys.stdin.readline()
            ts = str(int(time.time()))
            user = self.user
            msg = service_pb2.Message(
                id=gen_encrypt_string(user.display_name),
                user=user,
                message=content,
                timestamp=ts
            )
            stub.BroadcastMessage(msg)

    def run(self):
        with grpc.insecure_channel("localhost:17100") as channel:
            stub = service_pb2_grpc.BroadcastStub(channel)
            p2 = Process(target=self.receive_message, args=(stub,))
            p2.start()
            self.broadcast_message(stub)
            p2.join()


if __name__ == "__main__":
    logging.basicConfig()
    client = Client()
    client.run()

