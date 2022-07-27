from time import sleep
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait as Wait

print("Start")

# Initialise web driver for Firefox
options = webdriver.FirefoxOptions()
driver = webdriver.Remote(
  command_executor="http://127.0.0.1:4444",
  options=options,
)

print("Send HTTP GET request")
driver.get("https://recaptcha-demo.appspot.com/recaptcha-v3-request-scores.php")

print("Wait for 2 sec")
sleep(2)

# Get button to trigger reCAPTCHA action
print("Click button")
buttonElem = driver.find_element(By.CSS_SELECTOR, "button.go")
buttonElem.click()

print("Wait for 5 sec")
sleep(5)
responseElem = driver.find_element(By.CSS_SELECTOR, "pre.response")
print("Response element is displayed?: " + str(responseElem.is_displayed()))
print("Response text: " + responseElem.text)

driver.quit()
print("End")

"""
# REFERENCES
- https://www.selenium.dev/documentation/webdriver/remote_webdriver/
- https://selenium-python.readthedocs.io/getting-started.html
- https://stackoverflow.com/questions/15937966/in-python-selenium-how-does-one-find-the-visibility-of-an-element

# LINE COUNT
Total:
Reused:
Written:
Referenced: 3
"""
