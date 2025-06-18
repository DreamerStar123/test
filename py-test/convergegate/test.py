import hashlib

# Required authentication and transaction data
apiUser = "UsrGPNSub1"
apiPassword = "PwdGPNSub1"
apiCmd = "700"
merchantTransactionId = "8825-5236"
amount = "1.71"
currencyCode = "USD"
ccNumber = "4420151544181238"
ccv = "100"
nameOnCard = "John Doe"
apiKey = "YOUR_API_KEY"  # replace with your actual API key

# Step 1: Compute SHA1 checksum
checksum_string = (
    apiUser + apiPassword + apiCmd + merchantTransactionId + amount +
    currencyCode + ccNumber + ccv + nameOnCard + apiKey
)
checksum = hashlib.sha1(checksum_string.encode('utf-8')).hexdigest()

print(checksum)