import schedule
import subprocess
import time
from lxml import html
import logging
from plyer import notification

# Configure logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s', filename='scrape.log', filemode='a')

joblist = ["021896395618817438243"]

def scrape_url(job_id):
    # Define the curl command
    curl_command = [
        'curl',
        '--location', f'https://www.upwork.com/jobs/~{job_id}',
        '--header', 'Host: www.upwork.com',
        '--header', 'User-Agent: PostmanRuntime/7.43.0',
    ]

    # Execute the curl command and capture the output
    try:
        result = subprocess.run(curl_command, capture_output=True, text=True, check=True)
        output = result.stdout  # The output from the command
        # Parse the HTML content
        tree = html.fromstring(output)
        
        tree.xpath('//*[@id="main"]/div/div/div[1]/div/div/section')
        title = tree.xpath('//*[@id="main"]/div/div/div[1]/div/div/header/h4')[0].text.strip()

        # Example XPath queries
        activity = tree.xpath('//*[@id="main"]/div/div/div[1]/div/div/section[4]/ul/li')
        if (len(activity) == 0):
            activity = tree.xpath('//*[@id="main"]/div/div/div[1]/div/div/section[5]/ul/li')
        if (len(activity) == 0):
            logging.error("Could not find the activity section in the HTML content")
            return
        
        proposals = 0
        last_viewed = "not viewed"
        interviewing = 0
        invites_sent = 0
        unanswered_invites = 0
        
        if len(activity) == 5:
            # viewed
            proposals = activity[0].xpath('.//span[3]')[0].text.strip()
            last_viewed = activity[1].xpath('.//span[3]')[0].text.strip()
            interviewing = activity[2].xpath('.//div[1]')[0].text.strip()
            invites_sent = activity[3].xpath('.//div[1]')[0].text.strip()
            unanswered_invites = activity[4].xpath('.//div[1]')[0].text.strip()
        elif len(activity) == 4:
            # not viewed
            proposals = activity[0].xpath('.//span[3]')[0].text.strip()
            interviewing = activity[1].xpath('.//div[1]')[0].text.strip()
            invites_sent = activity[2].xpath('.//div[1]')[0].text.strip()
            unanswered_invites = activity[3].xpath('.//div[1]')[0].text.strip()

        # Print the results
        logging.info(f"title:               {title}", )
        logging.info(f"proposals:           {proposals}", )
        logging.info(f"last_viewed:         {last_viewed}")
        logging.info(f"interviewing:        {interviewing}")
        logging.info(f"invites_sent:        {invites_sent}")
        logging.info(f"unanswered_invites:  {unanswered_invites}")

        if last_viewed != "not viewed":
            # Send desktop notification
            notification.notify(
                title="Scrape Result",
                message=f"Title: {title}\nProposals: {proposals}\nLast Viewed: {last_viewed}\nInterviewing: {interviewing}\nInvites Sent: {invites_sent}\nUnanswered Invites: {unanswered_invites}",
                timeout=10
            )
    except subprocess.CalledProcessError as e:
        logging.info("An error occurred while executing the curl command:")
        logging.info(e.stderr)

def scrape_all():
    logging.info("Scraping all jobs...")
    start_time = time.time()
    for job in joblist:
        scrape_url(job)
    end_time = time.time()
    elapsed_time = end_time - start_time
    logging.info(f"Time taken to scrape: {elapsed_time:.2f} seconds")

scrape_all()

# Schedule the job every 10 minutes
schedule.every(5).minutes.do(scrape_all)

while True:
    schedule.run_pending()
    time.sleep(10)
