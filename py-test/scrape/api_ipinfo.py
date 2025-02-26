import ipinfo
import logging

prev_ip = ''

def scrape_ip():
    global prev_ip
    
    access_token = 'b5929ace43cdca'  # Replace with your IPinfo access token
    handler = ipinfo.getHandler(access_token)

    try:
        details = handler.getDetails()
        ip = details.ip
        city = details.city
        region = details.region
        country = details.country
        org = details.org
        
        # Check if the IP address has changed
        if prev_ip != ip:
            print(f'ip changed: {prev_ip}, {ip}')
            prev_ip = ip

            # Print the results
            print(f"IP: {ip}")
            print(f"City: {city}")
            print(f"Region: {region}")
            print(f"Country: {country}")
            print(f"Organization: {org}")

            # Log the results
            logging.info(f"IP: {ip}")
            logging.info(f"City: {city}")
            logging.info(f"Region: {region}")
            logging.info(f"Country: {country}")
            logging.info(f"Organization: {org}")

    except Exception as e:
        logging.error("An error occurred while fetching IP details:")
        logging.error(e)
