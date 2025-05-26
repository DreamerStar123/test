from openai import OpenAI

client = OpenAI()

response = client.audio.speech.create(
    model="tts-1",  # or "tts-1-hd"
    voice="nova",   # options: nova, shimmer, echo
    input="Hello, this is a demo of OpenAI's TTS in SDK v1.82.0"
)

# Save the output
with open("output.mp3", "wb") as f:
    f.write(response.content)