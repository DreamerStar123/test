const { GoogleAuth } = require('google-auth-library');

// Path to your service account key file
const SERVICE_ACCOUNT_FILE = 'key.json'; // Update with your path

async function generateAccessToken() {
    const auth = new GoogleAuth({
        keyFile: SERVICE_ACCOUNT_FILE,
        scopes: ['https://www.googleapis.com/auth/cloud-platform'], // Adjust scopes as needed
    });

    const client = await auth.getClient();
    const accessToken = await client.getAccessToken();

    console.log('Access Token:', accessToken.token);
}

generateAccessToken().catch(console.error);