import requests
resp = requests.get('http://localhost:8080/v1/cep/08226021')
print("Status:" + str(resp.status_code))
print(resp.json())