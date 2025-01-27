import subprocess

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
    print("Curl Output:")
    print(output)
except subprocess.CalledProcessError as e:
    print("An error occurred while executing the curl command:")
    print(e.stderr)