from google_auth_oauthlib.flow import InstalledAppFlow
from googleapiclient.discovery import build

# Specify the required scopes
SCOPES = ['https://www.googleapis.com/auth/calendar']

# Run the OAuth flow
flow = InstalledAppFlow.from_client_secrets_file(
    'credentials.json', SCOPES)
creds = flow.run_local_server(port=0)

# Build the Calendar API service
service = build('calendar', 'v3', credentials=creds)

# List the user's calendars
calendar_list = service.calendarList().list().execute()
for calendar in calendar_list['items']:
    print(calendar)
