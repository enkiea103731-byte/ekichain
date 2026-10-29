export default function handler(req,res){
 res.setHeader('Access-Control-Allow-Origin','*');
 res.json({name:"EKIChain",chainId:701,status:"ONLINE"});
}
