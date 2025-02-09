import requests
from lxml import html
import logging

def scrape_ip():
    # Execute the curl command and capture the output
    try:
        response = requests.get(
          'https://www.ipaddress.my', 
          headers={
            'Host': 'www.ipaddress.my',
            'User-Agent': 'PostmanRuntime/7.43.0'
          })
        output = response.text

        # Parse the HTML content
        tree = html.fromstring(output)
        
        ip = tree.xpath('/html/body/div[1]/div[2]/div[2]/div/div/ul/li[1]/span')[0].text.strip()

        # Print the results
        print(ip)
        # logging.info(f"title:               {title}", )
    except requests.exceptions.RequestException as e:
        logging.error("An error occurred while making the GET request:")
        logging.error(e)
