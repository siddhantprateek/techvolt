import { ethers } from "ethers";
const provider = new ethers.providers.Web3Provider(window.ethereum)

await provider.send("eth_requestAccounts", []);

const signer = provider.getSigner()
