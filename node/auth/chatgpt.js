import OpenAI from "openai";

const openai = new OpenAI({ apiKey: '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' });

async function main() {
    const stream = await openai.chat.completions.create({
        model: "gpt-4o-mini",
        messages: [{ role: "user", content: "Say this is a test" }],
        store: true,
        stream: true,
    });
    for await (const chunk of stream) {
        process.stdout.write(chunk.choices[0]?.delta?.content || "");
    }
}

main();