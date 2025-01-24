from openai import OpenAI

client = OpenAI(api_key='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')

stream = client.chat.completions.create(
    model="gpt-4o-mini",
    messages=[{"role": "user", "content": "Say this is a test"}],
    stream=True,
)
for chunk in stream:
    if chunk.choices[0].delta.content is not None:
        print(chunk.choices[0].delta.content, end="")