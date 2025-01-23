import json
from lxml import etree
from collections import Counter

# Load the HTML file
with open('monday-search.html', 'r', encoding='utf-8') as file:
  html_content = file.read()

# Parse the HTML content
tree = etree.HTML(html_content)

# Example XPath queries
jobs = tree.xpath('//*[@id="main"]/div[3]/div[4]/div/div[2]/div[3]/section/article')
posted = [job.xpath('.//div[2]/div[1]/small/span[2]')[0].text for job in jobs]
pm = Counter(posted)

# Print the results
print("jobs count:", len(posted))

# Save to a JSON file
with open('counter.json', 'w') as json_file:
    json.dump(dict(pm), json_file)
