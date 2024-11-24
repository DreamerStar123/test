const msal = require('@azure/msal-node');
const { Client } = require('@microsoft/microsoft-graph-client');

const config = {
    auth: {
        clientId: '4ba4a1e1-6337-4b28-9f8b-06c82e76e50c',
        authority: 'https://login.microsoftonline.com/390e3543-289a-4ae8-9776-1314872f4235',
        clientSecret: '3KT8Q~dGTBYVwioLq~joVlZc0sQcrB6RdTg-cdsj',
    }
};

const cca = new msal.ConfidentialClientApplication(config);

async function getAccessToken() {
    const clientCredentialRequest = {
        scopes: ['https://graph.microsoft.com/.default'],
    };

    try {
        const response = await cca.acquireTokenByClientCredential(clientCredentialRequest);
        return response.accessToken;
    } catch (error) {
        console.error(`Error acquiring token: ${error}`);
    }
}

async function getCalendarEvents() {
    const accessToken = await getAccessToken();

    const client = Client.init({
        authProvider: (done) => {
            done(null, accessToken);
        }
    });

    try {
        const events = await client.api('/users').get();
        console.log(events);
    } catch (error) {
        console.error(`Error fetching calendar events: ${error}`);
    }
}

getCalendarEvents();
