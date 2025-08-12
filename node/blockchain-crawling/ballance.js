import { ethers } from "ethers";
import "dotenv/config";

const provider = new ethers.JsonRpcProvider(
  `https://eth-mainnet.g.alchemy.com/v2/${process.env.ALCHEMY_KEY}`
);

// Example: USDC token contract
const contractAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48";
const transferTopic =
  "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef";

async function main() {
  const a = await provider.getBalance(
    "0x580abd34ef4e12e79a04b91e04603b7cd08a9655"
  );

  console.log(Number(a) * 1e-18);
}

main();
