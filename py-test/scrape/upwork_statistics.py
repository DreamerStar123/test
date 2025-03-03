import schedule
import subprocess
import time
from lxml import html
from datetime import datetime
from plyer import notification
import pandas as pd
from openpyxl import load_workbook

file_path = 'upwork_stats.xlsx'

def convert_int(val):
    return int(val.translate(str.maketrans('', '', '(),')))

def scrape_url():
    # Define the curl command
    curl_command = [
        'curl',
        '--location', f'https://www.upwork.com/nx/search/jobs',
        '--header', 'Host: www.upwork.com',
        '--header', 'User-Agent: PostmanRuntime/7.43.0',
    ]
    
    # Execute the curl command and capture the output
    try:
        result = subprocess.run(curl_command, capture_output=True, text=True, check=True, encoding='utf-8')
        output = result.stdout  # The output from the command
        
        if output is None:
            print('output is None')
            return
        
        # Parse the HTML content
        tree = html.fromstring(output)
        # Example XPath queries
        main_pane = tree.xpath('//*[@id="main"]/div/div/div/div[2]/div')
        if len(main_pane) == 0:
            print('len(main_pane) == 0')
            return
        if len(main_pane) == 2:
            search_pane = main_pane[0]
        elif len(main_pane) == 3:
            search_pane = main_pane[1]
        else:
            print('invalid len(main_pane)')
            return
        hourly = search_pane.xpath('.//div[3]/div/div/div[1]/label/span[2]/small')
        fixed = search_pane.xpath('.//div[3]/div/div/div[2]/label/span[2]/small')
        # Remove specified characters and convert to integer
        hourly = convert_int(hourly[0].text.strip())
        fixed = convert_int(fixed[0].text.strip())
        # Print the results
        print(f"hourly: {hourly}")
        print(f"fixed: {fixed}")
        
        # Load the workbook and select the active worksheet
        wb = load_workbook(file_path)
        ws = wb.active
        # Append a new row
        ws.append([datetime.now().isoformat(), hourly, fixed])
        # Save the workbook
        wb.save(file_path)
    except subprocess.CalledProcessError as e:
        print("An error occurred while executing the curl command:")
        print(e.stderr)

def scrape_all():
    print("Scraping all jobs...")
    start_time = time.time()
    scrape_url()
    end_time = time.time()
    elapsed_time = end_time - start_time
    print(f"Time taken to scrape: {elapsed_time:.2f} seconds")

scrape_all()

# Schedule the job every 10 minutes
schedule.every(5).minutes.do(scrape_all)

while True:
    schedule.run_pending()
    time.sleep(10)
