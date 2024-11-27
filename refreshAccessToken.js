const axios = require('axios');
require("dotenv").config();

const CLIENT_ID = '727774901928-p95rnppd6rfhisjvj9nk7rp88qgbbeci.apps.googleusercontent.com';
const CLIENT_SECRET = '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';

async function refreshGoogleAccessToken(refreshToken) {
  try {
    const response = await axios.post('https://oauth2.googleapis.com/token', {
      client_id: CLIENT_ID,
      client_secret: CLIENT_SECRET,
      refresh_token: refreshToken,
      grant_type: 'refresh_token',
    }, {
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
    });

    const { access_token, expires_in } = response.data;
    // console.log("New Access Token:", access_token);
    console.log("Expires In:", expires_in, "seconds");
    return access_token;
  } catch (error) {
    console.error("Error refreshing access token:", error.response?.data || error.message);
    throw error;
  }
}

async function refreshMicrosoftAccessToken(refreshToken) {
  try {
    const response = await axios.post('https://login.microsoftonline.com/common/oauth2/v2.0/token', {
      client_id: process.env.CLIENT_ID,
      client_secret: process.env.CLIENT_SECRET,
      refresh_token: refreshToken,
      grant_type: 'refresh_token',
    }, {
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
    });

    const { access_token, expires_in } = response.data;
    // console.log("New Access Token:", access_token);
    console.log("Expires In:", expires_in, "seconds");
    return access_token;
  } catch (error) {
    console.error("Error refreshing access token:", error.response?.data || error.message);
    throw error;
  }
}

module.exports = {
  refreshGoogleAccessToken,
  refreshMicrosoftAccessToken
}
