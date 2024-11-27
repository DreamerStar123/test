require("dotenv").config();
const express = require("express");
const bodyParser = require("body-parser");
const axios = require("axios");

const app = express();
app.use(bodyParser.urlencoded({ extended: true }));

// OAuth URLs
const AUTHORIZE_URL = `https://login.microsoftonline.com/${process.env.TENANT_ID}/oauth2/v2.0/authorize`;
const TOKEN_URL = `https://login.microsoftonline.com/${process.env.TENANT_ID}/oauth2/v2.0/token`;

// Step 1: Redirect user to Microsoft's authorization page
app.get("/auth", (req, res) => {
  const params = new URLSearchParams({
    client_id: process.env.CLIENT_ID,
    response_type: "code",
    redirect_uri: process.env.REDIRECT_URI,
    scope: "Calendars.ReadWrite offline_access",
    response_mode: "query",
    // state: "12345", // Optional for CSRF protection
  });

  res.redirect(`${AUTHORIZE_URL}?${params.toString()}`);
});

// Step 2: Handle callback and exchange authorization code for tokens
app.get("/auth/callback", async (req, res) => {
  const { code, state, error } = req.query;

  if (error) {
    return res.send(`Error: ${error}`);
  }

  try {
    // Exchange authorization code for tokens
    const tokenResponse = await axios.post(
      TOKEN_URL,
      new URLSearchParams({
        client_id: process.env.CLIENT_ID,
        client_secret: process.env.CLIENT_SECRET,
        grant_type: "authorization_code",
        code: code,
        redirect_uri: process.env.REDIRECT_URI,
      }),
      {
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
      }
    );

    const tokens = tokenResponse.data;

    // Display the tokens
    res.json({
      access_token: tokens.access_token,
      refresh_token: tokens.refresh_token,
      expires_in: tokens.expires_in,
      scope: tokens.scope,
    });
  } catch (err) {
    console.error(err.response ? err.response.data : err.message);
    res.status(500).send("Error exchanging code for tokens.");
  }
});

// Step 3: Use refresh token to get a new access token
app.post("/auth/refresh", async (req, res) => {
  const { refresh_token } = req.body;

  try {
    const tokenResponse = await axios.post(
      TOKEN_URL,
      new URLSearchParams({
        client_id: process.env.CLIENT_ID,
        client_secret: process.env.CLIENT_SECRET,
        grant_type: "refresh_token",
        refresh_token: refresh_token,
      }),
      {
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
      }
    );

    const tokens = tokenResponse.data;

    // Display the refreshed tokens
    res.json({
      access_token: tokens.access_token,
      refresh_token: tokens.refresh_token,
      expires_in: tokens.expires_in,
      scope: tokens.scope,
    });
  } catch (err) {
    console.error(err.response ? err.response.data : err.message);
    res.status(500).send("Error refreshing token.");
  }
});

// Start server
app.listen(3000, () => {
  console.log("Server running on http://localhost:3000");
});
