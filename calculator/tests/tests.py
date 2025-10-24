import requests
import time

BASE_URL = "http://localhost:8081"

def test_add():
    r = requests.get(f"{BASE_URL}/add")
    assert r.status_code == 200
    print("Get user OK")

if __name__ == "__main__":
    test_add()
