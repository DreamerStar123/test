import schedule
import subprocess
import time
from lxml import html
from datetime import datetime
from plyer import notification
import pandas as pd
from openpyxl import load_workbook

file_path = 'E:\\note\\task\\upwork\\upwork_stats.xlsx'

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
        
        entry = search_pane.xpath('.//div[2]/div/div/div/label[1]/span[2]/small')
        intermediate = search_pane.xpath('.//div[2]/div/div/div/label[2]/span[2]/small')
        expert = search_pane.xpath('.//div[2]/div/div/div/label[3]/span[2]/small')
        hourly = search_pane.xpath('.//div[3]/div/div/div[1]/label/span[2]/small')
        fixed = search_pane.xpath('.//div[3]/div/div/div[2]/label/span[2]/small')
        fixed_less_100 = search_pane.xpath('.//div[3]/div/div/div[2]/div/div[1]/label[1]/span[2]/small')
        fixed_100_500 = search_pane.xpath('.//div[3]/div/div/div[2]/div/div[1]/label[2]/span[2]/small')
        fixed_500_1K = search_pane.xpath('.//div[3]/div/div/div[2]/div/div[1]/label[3]/span[2]/small')
        fixed_1K_5K = search_pane.xpath('.//div[3]/div/div/div[2]/div/div[1]/label[4]/span[2]/small')
        fixed_5K_plus = search_pane.xpath('.//div[3]/div/div/div[2]/div/div[1]/label[5]/span[2]/small')
        no_hire = search_pane.xpath('.//div[4]/div/div/div/label[1]/span[2]/small')
        hire_1_9 = search_pane.xpath('.//div[4]/div/div/div/label[2]/span[2]/small')
        hire_10_plus = search_pane.xpath('.//div[4]/div/div/div/label[3]/span[2]/small')
        len_less_one_month = search_pane.xpath('.//div[7]/div/div/div/label[1]/span[2]/small')
        len_1_3_month = search_pane.xpath('.//div[7]/div/div/div/label[2]/span[2]/small')
        len_3_6_month = search_pane.xpath('.//div[7]/div/div/div/label[3]/span[2]/small')
        len_6_month_plus = search_pane.xpath('.//div[7]/div/div/div/label[4]/span[2]/small')
        hrs_less_30 = search_pane.xpath('.//div[8]/div/div/div/label[1]/span[2]/small')
        hrs_30_plus = search_pane.xpath('.//div[8]/div/div/div/label[2]/span[2]/small')
        contract_to_hire = search_pane.xpath('.//div[9]/div/div/div/label/span[2]/small')

        # Remove specified characters and convert to integer
        new_row = [
            datetime.now().isoformat(), 
            convert_int(entry[0].text.strip()), 
            convert_int(intermediate[0].text.strip()), 
            convert_int(expert[0].text.strip()), 
            convert_int(hourly[0].text.strip()), 
            convert_int(fixed[0].text.strip()), 
            convert_int(fixed_less_100[0].text.strip()), 
            convert_int(fixed_100_500[0].text.strip()), 
            convert_int(fixed_500_1K[0].text.strip()), 
            convert_int(fixed_1K_5K[0].text.strip()), 
            convert_int(fixed_5K_plus[0].text.strip()), 
            convert_int(no_hire[0].text.strip()), 
            convert_int(hire_1_9[0].text.strip()), 
            convert_int(hire_10_plus[0].text.strip()), 
            convert_int(len_less_one_month[0].text.strip()), 
            convert_int(len_1_3_month[0].text.strip()), 
            convert_int(len_3_6_month[0].text.strip()), 
            convert_int(len_6_month_plus[0].text.strip()), 
            convert_int(hrs_less_30[0].text.strip()), 
            convert_int(hrs_30_plus[0].text.strip()), 
            convert_int(contract_to_hire[0].text.strip()), 
        ]
        
        # Print the results
        print(f"hourly: {new_row}")
        
        # Load the workbook and select the active worksheet
        wb = load_workbook(file_path)
        ws = wb.active
        # Append a new row
        ws.append(new_row)
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
schedule.every(20).minutes.do(scrape_all)

while True:
    schedule.run_pending()
    time.sleep(10)
