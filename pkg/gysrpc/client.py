import socket
import json
import datetime
import requests
from bs4 import BeautifulSoup

TCP_IP = '127.0.0.1'
TCP_PORT = 9000
BUFFER_SIZE = 2048
MESSAGE = {"Url":"https://www.zoznam.sk/firma/2550207/Dajan-Daniela-Valkova-Sobrance","Selector":"div[class='col-md-8 profile middle-content']","Type":"one","Subselectors":[{"Selector":"div.row","Attribute":"text","Name":"","Split":":","Default":""},{"Selector":"div.row","Attribute":"text","Name":"CompanyName","Split":"","Default":""}]}
MESSAGE = {"Url":"https://www.profesia.sk/sitemap.php","Selector":"url","Type":"many","Subselectors":[{"Selector":"loc","Attribute":"text","Name":"url","Split":"","Default":""}]}
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((TCP_IP, TCP_PORT))
msg = json.dumps({"method":"RPCHandler.Execute","params":[MESSAGE]})
print(datetime.datetime.now().isoformat())
s.send(msg.encode())
buffer = b''
for i in range(1_000_000_000):
    data = s.recv(BUFFER_SIZE)
    #print(data)
    if data:
        buffer += data
    if len(data) <= BUFFER_SIZE and data[-5:].decode().endswith("}\n"):
        break
r = json.loads(buffer.decode())
s.close()
#print ("received data:", buffer)
print(datetime.datetime.now().isoformat())

#comparison with python
print("Pure Python")
print(datetime.datetime.now().isoformat())
r = requests.get("https://www.profesia.sk/sitemap.php")
soup = BeautifulSoup(r.text, "lxml")
locs = soup.findAll("loc")
result = [item.text for item in locs]
print(datetime.datetime.now().isoformat())
