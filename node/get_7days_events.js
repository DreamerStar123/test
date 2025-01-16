const admin = require('firebase-admin');
const { google } = require('googleapis');
const { Client } = require('@microsoft/microsoft-graph-client');
const serviceAccount = require('./turbo-tabs-firebase-adminsdk-vk1px-e1798541d6.json'); // Path to your service account JSON file
const { refreshGoogleAccessToken, refreshMicrosoftAccessToken } = require('./refreshAccessToken');
require("dotenv").config();

admin.initializeApp({
    credential: admin.credential.cert(serviceAccount),
});

const db = admin.firestore();

const getUsers = async () => {
    try {
        let nextPageToken;
        const googleUsers = [];
        const microsoftUsers = [];

        do {
            const listUsersResult = await admin.auth().listUsers(1000, nextPageToken); // Batch of 1000 users
            listUsersResult.users.forEach((userRecord) => {
                // Check if user signed in with Google
                const isGoogleUser = userRecord.providerData.some(
                    (provider) => provider.providerId === "google.com"
                );
                const isMicrosoftUser = userRecord.providerData.some(
                    (provider) => provider.providerId === "microsoft.com"
                );
                if (isGoogleUser)
                    googleUsers.push(userRecord);
                if (isMicrosoftUser)
                    microsoftUsers.push(userRecord);
            });
            nextPageToken = listUsersResult.pageToken;
        } while (nextPageToken);

        console.log("Google Users:", googleUsers.length);
        console.log("Microsoft Users:", microsoftUsers.length);
        return { googleUsers, microsoftUsers };
    } catch (error) {
        console.error("Error listing users:", error);
        return null;
    }
}

const getGoogleCalendars = async (accessToken) => {
    const oauth2Client = new google.auth.OAuth2();
    oauth2Client.setCredentials({ access_token: accessToken });

    const calendar = google.calendar({ version: 'v3', auth: oauth2Client });
    const res = await calendar.calendarList.list();
    const calendars = res.data;

    return calendars;
}

const getGoogleEvents = async (accessToken) => {
    const oauth2Client = new google.auth.OAuth2();
    oauth2Client.setCredentials({ access_token: accessToken });

    const calendar = google.calendar({ version: 'v3', auth: oauth2Client });

    const now = new Date();
    const timeMin = now.toISOString();
    const timeMax = new Date(now.getTime() + 7 * 24 * 60 * 60 * 1000).toISOString(); // 7 days from now
    const res = await calendar.events.list({
        calendarId: 'primary',
        timeMin,
        timeMax,
        singleEvents: true,
        orderBy: 'startTime',
    });
    const events = res.data;

    return events;
}

const getMicrosoftEvents = async (accessToken) => {
    let events = [];
    const client = Client.init({
        authProvider: (done) => {
            done(null, accessToken);
        }
    });

    try {
        const now = new Date();
        const timeMin = now.toISOString();
        const timeMax = new Date(now.getTime() + 7 * 24 * 60 * 60 * 1000).toISOString(); // 7 days from now
        events = await client.api(`/me/calendarview?startDateTime=${timeMin}&endDateTime=${timeMax}`).get();
    } catch (error) {
        console.error(`Error fetching calendar events: ${error}`);
    }
    return events;
}

const main = async () => {
    try {
        const { googleUsers, microsoftUsers } = await getUsers();
        if (googleUsers) {
            googleUsers.forEach(async (user) => {
                // Reference to a document in the "orders" collection
                const docRef = db.collection("users").doc(user.uid);

                // Data to store
                let userData = (await docRef.get()).data();
                const refreshToken = userData.refreshToken;
                if (refreshToken) {
                    const accessToken = await refreshGoogleAccessToken(refreshToken);
                    const events = await getGoogleEvents(accessToken);

                    userData = {
                        ...userData,
                        calendars: events.items
                    }

                    // Save the document
                    await docRef.set(userData);
                    console.log(`Document written for ${userData.email}`);
                }
            });
        }
        if (microsoftUsers) {
            microsoftUsers.forEach(async (user) => {
                // Reference to a document in the "orders" collection
                const docRef = db.collection("users").doc(user.uid);

                // Data to store
                let userData = (await docRef.get()).data();
                const refreshToken = userData.refreshToken;
                if (refreshToken) {
                    const accessToken = await refreshMicrosoftAccessToken(refreshToken);
                    const events = await getMicrosoftEvents(accessToken);

                    userData = {
                        ...userData,
                        calendars: events.value
                    }

                    // Save the document
                    await docRef.set(userData);
                    console.log(`Document written for ${userData.email}`);
                }
            })
        }
    } catch (e) {
        console.error("Error adding document: ", e.status);
    }
}

main();
