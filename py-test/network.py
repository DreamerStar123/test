import os
import time
from datetime import datetime
from plyer import notification

LOGFILE = "C:\\dev\\network_status.log"
HOST = "bing.com"  # Target to ping
TIMEOUT = 60

lastStatus = ''
lastUp = 0

def log_status():
    global lastUp, lastStatus

    timestamp = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
    response = os.system(f"ping {HOST} -n 1")
    status = "UP" if response == 0 else "DOWN"
    if response == 0:
        if time.time() - lastUp > TIMEOUT:
            notification.notify(
                title="Network status",
                message="Network is UP",
                timeout=10
            )
        lastUp = time.time()
    if status != lastStatus:
        lastStatus = status
        with open(LOGFILE, "a") as file:
            file.write(f"{timestamp} - Network is {status}\n")

while True:
    log_status()
    time.sleep(1)  # Check every 60 seconds
