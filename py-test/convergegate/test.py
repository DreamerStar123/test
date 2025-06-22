import hashlib

# Required authentication and transaction data
apiUser = "COI3946"
apiPassword = "5149da0bf59f"
apiCmd = "850"
merchantTransactionId = "6188a25b-fb7c-4c31-98ec-a2c6302d0718"
amount = "10000"
currencyCode = "EUR"
token = "0SMUCnuiS0GGJX49WVmqsg"
ccNumber = "4242424242424242"
ccv = "123"
nameOnCard = "Demo Client"
apiKey = "AEC35C1A-4BE2-CB58-A544-5C8E1141B3A9"  # replace with your actual API key

# Step 1: Compute SHA1 checksum
# checksum_string = apiUser + apiPassword + apiCmd + merchantTransactionId + amount + currencyCode + token + apiKey
# checksum_string = apiUser + apiPassword + apiCmd + merchantTransactionId + amount + currencyCode + ccNumber + ccv + nameOnCard + apiKey
checksum_string = apiUser + apiPassword + apiCmd + merchantTransactionId + amount + currencyCode + apiKey
checksum = hashlib.sha1(checksum_string.encode('utf-8')).hexdigest()

print(checksum)