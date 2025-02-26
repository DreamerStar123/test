import schedule
import time
from api_ipinfo import scrape_ip
import logging

# Configure logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s', filename='ip.log', filemode='a')

# Schedule the job every 10 minutes
# schedule.every(5).seconds.do(scrape_ip)
schedule.every(10).seconds.do(scrape_ip)

scrape_ip()
while True:
    schedule.run_pending()
    time.sleep(1)
