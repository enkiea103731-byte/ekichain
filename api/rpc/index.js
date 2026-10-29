export default function handler(req,res){
 res.setHeader('Access-Control-Allow-Origin','*');
 res.json({
  chainId: "0x2BD",
  chainName: "EKIChain",
  rpcUrl: "https://eki-chain.vercel.app/api/rpc",
  nativeCurrency: {name:"EKI",symbol:"EKI",decimals:2},
  blockHeight: Math.floor(Date.now()/10000),
  balancesFile: "balances.json na EKICHAIN",
  pix: "32468719816",
  status: "Mainnet Live - Android Native"
 });
}
