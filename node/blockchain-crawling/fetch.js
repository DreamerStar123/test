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
  const currentBlock = await provider.getBlockNumber();

  // Query last 1000 blocks
  const fromBlock = currentBlock - 1;

  const logs = await provider.getLogs({
    address: contractAddress,
    topics: [transferTopic],
    fromBlock,
    toBlock: currentBlock,
  });

  logs.forEach((log) => {
    const from = "0x" + log.topics[1].slice(26); // last 40 hex chars = address
    const to = "0x" + log.topics[2].slice(26);

    const [value] = ethers.AbiCoder.defaultAbiCoder().decode(
      ["uint256"],
      log.data
    );

    console.log({ from, to, value: value.toString() });
  });
}

main();
