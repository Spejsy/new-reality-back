import requests

JSON = """{
    "id": "999",
    "users": ["szymon", "igor"],
    "smallTasks": {"task0": [true, false], "task1": [true, false]},
    "complexTasks": {"task0": [0.7, 0.1], "task1": [0.1, 0.99]},
    "comments": [{"user": "szymon", "text": "dobra robota"}, {"user": "igor", "text": "super"}]
}"""

r = requests.put("http://localhost:10000/room/1001", data=JSON)
print(r.status_code, r.reason, r.content)
