import os
import time
from datetime import datetime

LOGFILE = "network_status.log"
HOST = "google.com"  # Target to ping

def log_status():
    timestamp = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
    response = os.system(f"ping {HOST} -n 1")
    status = "UP" if response == 0 else "DOWN"
    with open(LOGFILE, "a") as file:
        file.write(f"{timestamp} - Network is {status}\n")

while True:
    log_status()
    time.sleep(1)  # Check every 60 seconds
