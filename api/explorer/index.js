import fs from 'fs';
export default function handler(req,res){
 res.setHeader('Access-Control-Allow-Origin','*');
 try{
  const bal = fs.existsSync('./balances.json') ? fs.readFileSync('./balances.json','utf8') : '{"total":"2100000"}';
  res.json({chain:"EKICHAIN", blocks:"block1.go", balances: JSON.parse(bal), explorer:"live"});
 }catch(e){res.json({chain:"EKICHAIN", status:"live", pix:"32468719816"})}
}
