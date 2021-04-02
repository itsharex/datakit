#!/usr/bin/env python

# This will create golden files in a directory passed to it.
# A Test calls this internally to create the golden files
# So it can process them (so we don't have to checkin the files).

# Ensure msgpack-python and cbor are installed first, using:
#   sudo apt install python-dev (may not be necessary)
#   sudo apt install python-pip # or python3-pip
#   pip install --user msgpack-python msgpack-rpc-python cbor

# Ensure all "string" keys are utf strings (else encoded as bytes)

from __future__ import print_function
import cbor, msgpack, msgpackrpc, sys, os, threading

mylocaladdr="127.0.0.1" # localhost.localdomain localhost 127.0.0.1

def get_test_data_list():
    # get list with all primitive types, and a combo type
    l0 = [ 
        -8,
         -1616,
         -32323232,
         -6464646464646464,
         192,
         1616,
         32323232,
         6464646464646464,
         192,
         -3232.0,
         -6464646464.0,
         3232.0,
         6464.0,
         6464646464.0,
         160.0,
         1616.0,
         False,
         True,
         u"null",
         None,
         u"some&day>some<day",
         1328176922000002000,
         u"",
         -2206187877999998000,
         u"bytestring",
         270,
         u"none",
        -2013855847999995777,
         #-6795364578871345152,
         ]
    l1 = [
        { "true": True,
          "false": False },
        { "true": u"True",
          "false": False,
          "uint16(1616)": 1616 },
        { "list": [1616, 32323232, True, -3232.0, {"TRUE":True, "FALSE":False}, [True, False] ],
          "int32":32323232, "bool": True, 
          "LONG STRING": u"123456789012345678901234567890123456789012345678901234567890",
          "SHORT STRING": u"1234567890" },
        { True: "true", 138: False, "false": 200 }
        ]
    
    l = []
    l.extend(l0)
    l.append(l0)
    l.append(1)
    l.extend(l1)
    return l

def build_test_data(destdir):
    l = get_test_data_list()
    for i in range(len(l)):
        # packer = msgpack.Packer()
        serialized = msgpack.dumps(l[i])
        with open(os.path.join(destdir, str(i) + '.msgpack.golden'), 'wb') as f:
            f.write(serialized)
        serialized = cbor.dumps(l[i])
        with open(os.path.join(destdir, str(i) + '.cbor.golden'), 'wb') as f:
            f.write(serialized)

def doRpcServer(port, stopTimeSec):
    class EchoHandler(object):
        def Echo123(self, msg1, msg2, msg3):
            return ("1:%s 2:%s 3:%s" % (msg1.decode("utf-8"), msg2.decode("utf-8"), msg3.decode("utf-8")))
        def EchoStruct(self, msg):
            return ("%s" % msg)
    
    addr = msgpackrpc.Address(mylocaladdr, port)
    server = msgpackrpc.Server(EchoHandler())
    server.listen(addr)
    # run thread to stop it after stopTimeSec seconds if > 0
    if stopTimeSec > 0:
        def myStopRpcServer():
            server.stop()
        t = threading.Timer(stopTimeSec, myStopRpcServer)
        t.start()
    server.start()

def doRpcClientToPythonSvc(port):
    address = msgpackrpc.Address(mylocaladdr, port)
    client = msgpackrpc.Client(address, unpack_encoding='utf-8')
    print(client.call("Echo123", "A1", "B2", "C3"))
    print(client.call("EchoStruct", {"A" :"Aa", "B":"Bb", "C":"Cc"}))

# def doCheckSocket(port):
#     print(">>>> port: ", port, " <<<<<")
#     sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
#     result = sock.connect_ex(('127.0.0.1', port))
#     if result == 0:
#         print("\t>>>> Port is open")
#     else:
#         print("\t>>>> Port is not open")
#     sock.close()
    
def doRpcClientToGoSvc(port):
    # doCheckSocket(port)
    address = msgpackrpc.Address(mylocaladdr, port)
    client = msgpackrpc.Client(address, unpack_encoding='utf-8')
    print(client.call("TestRpcInt.Echo123", ["A1", "B2", "C3"]))
    print(client.call("TestRpcInt.EchoStruct", {"A" :"Aa", "B":"Bb", "C":"Cc"}))

def doMain(args):
    if len(args) == 2 and args[0] == "testdata":
        build_test_data(args[1])
    elif len(args) == 3 and args[0] == "rpc-server":
        doRpcServer(int(args[1]), int(args[2]))
    elif len(args) == 2 and args[0] == "rpc-client-python-service":
        doRpcClientToPythonSvc(int(args[1]))
    elif len(args) == 2 and args[0] == "rpc-client-go-service":
        doRpcClientToGoSvc(int(args[1]))
    else:
        print("Usage: test.py " + 
              "[testdata|rpc-server|rpc-client-python-service|rpc-client-go-service] ...")
    
if __name__ == "__main__":
    doMain(sys.argv[1:])
