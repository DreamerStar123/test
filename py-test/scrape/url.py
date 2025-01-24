from selenium import webdriver

driver = webdriver.Chrome()  # Or use another driver like Firefox
driver.get("https://www.upwork.com/")
print(driver.page_source[:500])  # Print the first 500 characters
driver.quit()
