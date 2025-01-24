import requests
from lxml import html

# URL of the Upwork job listing
url = 'https://www.upwork.com/jobs/~021881745060515567376'
# url = 'https://www.w3schools.com/tags/ref_httpmessages.asp'

# Send a request to the webpage
headers = {
    'Host': 'www.upwork.com',
    'User-Agent': 'PostmanRuntime/7.43.0',
    'Accept': '*/*',
    'Accept-Encoding': 'gzip, deflate, br',
    'Connection': 'keep-alive',
}

response = requests.get(url, headers=headers)

# Check if the request was successful
if response.status_code == 200:
    # Parse the HTML content
    tree = html.fromstring(response.content)

    # Example XPath queries to extract data
    title = tree.xpath('//h1/text()')  # Job title
    description = tree.xpath('//div[@class="description"]/text()')  # Job description
    posted_date = tree.xpath('//time/@datetime')  # Posted date

    # Print the results
    print("Job Title:", title)
    print("Job Description:", description)
    print("Posted Date:", posted_date)
else:
    print(f"Failed to retrieve the page. Status code: {response.status_code}")
