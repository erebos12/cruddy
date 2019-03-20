import requests
import time 
import json

sum_roundtrip = 0
for i in range(1000):
    user = "user" + str(i)
    passwd = "1234" + str(i)
    payload = {'username': user, 'password': passwd, 'isAdmin': 'true'}
    start = time.time()
    r = requests.post("http://localhost:8080/users", data=json.dumps(payload))
    roundtrip = time.time() - start
    #time.sleep(1) 
    #print(r.status_code, r.reason, roundtrip)
    sum_roundtrip += roundtrip

print("Average roundtrip: " + str((sum_roundtrip/i)))