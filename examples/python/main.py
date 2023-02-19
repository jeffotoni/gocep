import requests
resp = requests.get('http://localhost:8080/api/v1/08226021')
print("Status:" + str(resp.status_code))
print(resp.json())