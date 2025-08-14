const { TranslationServiceClient } = require("@google-cloud/translate").v3;
require("dotenv/config");

const client = new TranslationServiceClient();

async function translateText(text, targetLanguage = "es") {
  const projectId = process.env.PROJECT_ID;
  const location = "global"; // or 'us-central1'

  const request = {
    parent: `projects/${projectId}/locations/${location}`,
    contents: [text],
    mimeType: "text/plain",
    targetLanguageCode: targetLanguage,
  };

  const [response] = await client.translateText(request);
  return response.translations[0].translatedText;
}

// Example usage
translateText("Hello world", "fr").then(console.log); // Bonjour le monde
