import hashlib

# Required authentication and transaction data
apiUser = "COI3946"
apiPassword = "5149da0bf59f"
apiCmd = "700"
merchantTransactionId = "6188a25b-fb7c-4c31-98ec-a2c6302d0718"
amount = "10000"
currencyCode = "EUR"
token = "fRnYLjIkSxS4Z0bqumlv-g"
apiKey = "AEC35C1A-4BE2-CB58-A544-5C8E1141B3A9"  # replace with your actual API key

# Step 1: Compute SHA1 checksum
checksum_string = apiUser + apiPassword + apiCmd + merchantTransactionId + amount + currencyCode + token + apiKey
checksum = hashlib.sha1(checksum_string.encode('utf-8')).hexdigest()

print(checksum)