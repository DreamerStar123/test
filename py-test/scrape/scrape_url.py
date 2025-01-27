import subprocess
from lxml import html

# Define the curl command
curl_command = [
    'curl',
    '--location', 'https://www.upwork.com/jobs/~021881745060515567376',
    '--header', 'Host: www.upwork.com',
    '--header', 'User-Agent: PostmanRuntime/7.43.0',
]

# Execute the curl command and capture the output
try:
    result = subprocess.run(curl_command, capture_output=True, text=True, check=True)
    output = result.stdout  # The output from the command
    # Parse the HTML content
    tree = html.fromstring(output)

    # Example XPath queries
    activity = tree.xpath('//*[@id="main"]/div/div/div/div/div[1]/section[4]/ul/li')
    proposals = activity[0].xpath('.//span[3]')[0].text.strip()
    last_viewed = activity[1].xpath('.//span[3]')[0].text.strip()
    interviewing = activity[2].xpath('.//div[1]')[0].text.strip()
    invites_sent = activity[3].xpath('.//div[1]')[0].text.strip()
    unanswered_invites = activity[4].xpath('.//div[1]')[0].text.strip()

    # Print the results
    print("proposals:           ", proposals)
    print("last_viewed:         ", last_viewed)
    print("interviewing:        ", interviewing)
    print("invites_sent:        ", invites_sent)
    print("unanswered_invites:  ", unanswered_invites)
except subprocess.CalledProcessError as e:
    print("An error occurred while executing the curl command:")
    print(e.stderr)

