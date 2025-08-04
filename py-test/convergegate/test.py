import hashlib

# Required authentication and transaction data
apiUser = "3987skk"
apiPassword = "27f746d718c5"
apiCmd = "700"
merchantTransactionId = "e0ffd82c-ad0b-47f4-83ad-7749ff7850c4"
amount = "10100"
currencyCode = "USD"
token = "b8qqhGZBTsOvNjsWk6CgiD"
ccNumber = "4000000000000002"
ccv = "123"
nameOnCard = "Demo Client"
apiKey = "32F936C6-0291-52C9-745A-F7E90741CA3C"  # replace with your actual API key

# Step 1: Compute SHA1 checksum
# checksum_string = apiUser + apiPassword + apiCmd + merchantTransactionId + amount + currencyCode + token + ccv + apiKey
# checksum_string = apiUser + apiPassword + apiCmd + merchantTransactionId + amount + currencyCode + ccNumber + ccv + nameOnCard + apiKey
checksum_string = apiUser + apiPassword + apiCmd + apiKey
checksum = hashlib.sha1(checksum_string.encode('utf-8')).hexdigest()

print(checksum)