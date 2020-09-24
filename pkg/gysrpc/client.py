import socket
import json

TCP_IP = '127.0.0.1'
TCP_PORT = 9000
BUFFER_SIZE = 4096
MESSAGE = {"Url":"https://www.zoznam.sk/firma/2550207/Dajan-Daniela-Valkova-Sobrance","Selector":"div[class='col-md-8 profile middle-content']","Type":"one","Subselectors":[{"Selector":"div.row","Attribute":"text","Name":"","Split":":","Default":""},{"Selector":"div.row","Attribute":"text","Name":"CompanyName","Split":"","Default":""}]}
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((TCP_IP, TCP_PORT))
msg = json.dumps({"method":"RPCHandler.Execute","params":[MESSAGE]})
s.send(msg.encode())
data = s.recv(BUFFER_SIZE)
s.close()
print ("received data:", data)