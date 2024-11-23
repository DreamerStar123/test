const { google } = require('googleapis');
const express = require('express');
const app = express();
const PORT = process.env.PORT || 3000;

const CLIENT_ID = '1236426385-a7gluu9h9fh61tiinlginq7mgq6v1la9.apps.googleusercontent.com';
const CLIENT_SECRET = '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';

const fetchCalendarEvents = async (req, res) => {
    const oAuth2Client = new google.auth.OAuth2(
        CLIENT_ID, CLIENT_SECRET
    );
    oAuth2Client.setCredentials({ access_token: req.query.token });

    const calendar = google.calendar({ version: 'v3', auth: oAuth2Client });
    console.log(calendar.events);
    // const events = await calendar.events.list({
    //     calendarId: 'primary',
    //     maxResults: 10,
    //     singleEvents: true,
    //     orderBy: 'startTime',
    // });

    res.send('ok');
};


// Middleware to parse JSON bodies
app.use(express.json());

// Define a simple route
app.get('/', fetchCalendarEvents);

// Start the server
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
