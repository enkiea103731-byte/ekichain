export default function handler(req,res){
  res.setHeader('Access-Control-Allow-Origin','*');
  res.json({name:"EKIChain",chainId:701,symbol:"EKI",rpc:"https://eki-chain.vercel.app/api/rpc.js",status:"ONLINE"});
}
