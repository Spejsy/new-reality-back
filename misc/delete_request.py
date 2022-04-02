import requests

JSON = """{
    "Id": "3",
    "Title": "Newly Created Post",
    "desc": "The description for my new post",
    "content": "my articles content"
}"""

r = requests.delete("http://localhost:10000/article/1")
print(r.status_code, r.reason, r.content)
