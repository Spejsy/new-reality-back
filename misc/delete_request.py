import requests

r = requests.delete("http://localhost:10000/room/1000")
print(r.status_code, r.reason, r.content)
