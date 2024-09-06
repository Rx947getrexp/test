import hashlib
import hmac
import json
import logging
import telnetlib
import time
import requests
url = "http://127.0.0.1:13002/machine_states_witching"
secret_key = "3f5202f0-4ed3-4456-80dd-13638c975bda"
params = {"Ip": "213.159.68.106", "status": "1"}
signature = hmac.new(
    secret_key.encode(), secret_key.encode(), hashlib.sha256
).hexdigest()
headers = {"X-Signature": signature}
response = requests.post(url, headers=headers, params=params)
print(response.text)
