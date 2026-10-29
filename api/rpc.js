export default function handler(req,res){
 res.setHeader('Access-Control-Allow-Origin','*');
 res.setHeader('Access-Control-Allow-Methods','*');
 res.setHeader('Access-Control-Allow-Headers','*');
 res.setHeader('Content-Type','application/json');
 if(req.method==='OPTIONS') return res.status(200).end();
 const b=req.body||{}; const id=b.id||1; const m=b.method; const p=b.params||[];
 if(m==='net_version') return res.json({jsonrpc:"2.0",id,result:"701"});
 if(m==='eth_chainId') return res.json({jsonrpc:"2.0",id,result:"0x2BD"});
 if(m==='eth_getBalance'){
   const a=(p[0]||'').toLowerCase();
   if(a==='0x03664899007a53f086d545c34063e9cbfb29c9c2'){
     return res.json({jsonrpc:"2.0",id,result:"0x6c9144c1c690d4cb4000000"});
   }
   return res.json({jsonrpc:"2.0",id,result:"0x0"});
 }
 if(m==='eth_blockNumber') return res.json({jsonrpc:"2.0",id,result:"0x1"});
 if(m==='eth_gasPrice') return res.json({jsonrpc:"2.0",id,result:"0x1"});
 return res.json({jsonrpc:"2.0",id,result:"0x1"});
}
