import 'dotenv/config' // see https://github.com/motdotla/dotenv#how-do-i-use-dotenv-with-import
import { Configuration, OpenAIApi } from "openai";

const configuration = new Configuration({
  apiKey: process.env.OPENAI_API_KEY,
});

async function main() {
  const openai = new OpenAIApi(configuration);
  const prompt = process.argv.slice(2).join(" ") ?? "Say this is a test";
  const resp = await openai.createCompletion({
    model: "text-davinci-003",
    prompt: prompt,
    temperature: 0,
    max_tokens: 7,
  }).catch((err) => {
      console.error(err)
      throw err;
    });
  console.log(JSON.stringify(resp));
}

main().catch((err) => {
  console.error(err);
});
