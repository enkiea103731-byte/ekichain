export default function handler(req,res){
 res.setHeader('Access-Control-Allow-Origin','*');
 res.setHeader('Access-Control-Allow-Methods','*');
 res.setHeader('Access-Control-Allow-Headers','*');
 res.setHeader('Content-Type','application/json');
 if(req.method==='OPTIONS') return res.status(200).end();
 const b=req.body||{}; const id=b.id||1;
 if(b.method==='net_version') return res.json({jsonrpc:"2.0",id,result:"701"});
 return res.json({jsonrpc:"2.0",id,result:"0x2BD"});
}
