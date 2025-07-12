import Replicate from "replicate";
import "dotenv/config";

const replicate = new Replicate({
  auth: process.env.REPLICATE_API_TOKEN,
});

console.log("Running the model...");
const [output] = await replicate.run("meta/meta-llama-3-70b-instruct", {
  input: {
    prompt: "hi",
  },
});

// Save the generated image
import { writeFile } from "node:fs/promises";

await writeFile("./output.png", output);
console.log("Image saved as output.png");
