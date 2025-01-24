const axios = require('axios');

const apiKey = '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'; // Replace with your API key
const apiUrl = 'https://api.openai.com/v1/chat/completions';

async function sendMessageToChatGPT(message) {
    try {
        const response = await axios.post(
            apiUrl,
            {
                model: 'gpt-4o-mini', // Replace with 'gpt-4' if needed
                messages: [
                    { role: 'system', content: 'You are a helpful assistant.' },
                    { role: 'user', content: message },
                ],
                max_tokens: 150,
                temperature: 0.7,
            },
            {
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${apiKey}`,
                },
            }
        );

        console.log('ChatGPT Response:', response.data.choices[0].message.content);
    } catch (error) {
        console.error('Error communicating with ChatGPT:', error.response?.data || error.message);
    }
}

// Example usage
sendMessageToChatGPT('Hello, how are you?');
