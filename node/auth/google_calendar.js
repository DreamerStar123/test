const { google } = require('googleapis');
const express = require('express');
const session = require('express-session');
require("dotenv").config();

const app = express();
const PORT = 5000;

// Replace these with your own credentials
const REDIRECT_URI = 'http://localhost:5000/oauth2callback';

// Create an OAuth2 client
const oAuth2Client = new google.auth.OAuth2(process.env.GCP_CLIENT_ID, process.env.GCP_CLIENT_SECRET, REDIRECT_URI);

// Set up session middleware
app.use(session({ secret: 'your-secret-key', resave: false, saveUninitialized: true }));

// Generate the URL for consent
app.get('/auth', (req, res) => {
  const authUrl = oAuth2Client.generateAuthUrl({
    access_type: 'offline',
    scope: ['https://www.googleapis.com/auth/calendar.readonly'],
  });
  res.redirect(authUrl);
});

// Handle the OAuth2 callback
app.get('/oauth2callback', async (req, res) => {
  const { code } = req.query;
  const { tokens } = await oAuth2Client.getToken(code);
  console.log(code);
  oAuth2Client.setCredentials(tokens);
  req.session.tokens = tokens; // Store tokens in session
  res.redirect('/events'); // Redirect to events page
});

// List the user's upcoming events
app.get('/events', async (req, res) => {
  oAuth2Client.setCredentials(req.session.tokens);
  console.log(req.session.tokens);
  
  const calendar = google.calendar({ version: 'v3', auth: oAuth2Client });
  calendar.events.list({
    calendarId: 'primary',
    // timeMin: (new Date()).toISOString(),
    maxResults: 10,
    singleEvents: true,
    orderBy: 'startTime',
  }, (err, response) => {
    if (err) return res.status(500).send('The API returned an error: ' + err);
    console.log(response.data);
    const events = response.data.items;
    if (events.length) {
      const eventsList = events.map(event => {
        const start = event.start.dateTime || event.start.date;
        return `${event.summary} (${start})`;
      }).join('<br>');
      res.send('Upcoming events:<br>' + eventsList);
    } else {
      res.send('No upcoming events found.');
    }
  });
});

// Start the server
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});