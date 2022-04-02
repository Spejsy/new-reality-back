import requests

JSON = """{
    "Id": "999",
    "Test": "Test"
}"""

r = requests.post("http://localhost:10000/room", data=JSON)
print(r.status_code, r.reason, r.content)
