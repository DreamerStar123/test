from openai import OpenAI

# Set your OpenAI API key

client = OpenAI()

# Example: Generate a completion using GPT-3.5 Turbo
response = client.completions.create(
    model="gpt-4o-mini",
    prompt="You are a helpful assistant.",
)

print(response.choices[0].text)