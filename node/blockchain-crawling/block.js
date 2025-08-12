import { ethers } from "ethers";
import "dotenv/config";

const provider = new ethers.JsonRpcProvider(
  `https://eth-mainnet.g.alchemy.com/v2/${process.env.ALCHEMY_KEY}`
);

console.time("getBlock");

// Query last 1000 blocks
const fromBlock = 23121805;
const block = await provider.getBlock(fromBlock);
console.timeEnd("getBlock");
console.log(block.timestamp);
