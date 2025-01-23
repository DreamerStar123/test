from lxml import etree

# Load the HTML file
with open('temp.html', 'r', encoding='utf-8') as file:
  html_content = file.read()

# Parse the HTML content
tree = etree.HTML(html_content)

# Example XPath queries
activity = tree.xpath('//*[@id="main"]/div[3]/div[4]/div/div/div[1]/div[1]/section[8]/section[2]/ul/li')
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
