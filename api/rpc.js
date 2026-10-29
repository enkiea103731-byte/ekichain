export default function handler(req, res) {
  res.setHeader('Access-Control-Allow-Origin','*');
  res.setHeader('Access-Control-Allow-Methods','*');
  res.setHeader('Access-Control-Allow-Headers','*');
  if (req.method === 'OPTIONS') return res.status(200).end();
  const id = req.body?.id || 1;
  const m = req.body?.method;
  if (m === 'net_version') return res.json({jsonrpc:"2.0",id,result:"701"});
  if (m === 'eth_chainId') return res.json({jsonrpc:"2.0",id,result:"0x2BD"});
  if (m === 'eth_blockNumber') return res.json({jsonrpc:"2.0",id,result:"0x1"});
  // quando a Trust só abre a tela, ela faz GET
  return res.json({jsonrpc:"2.0",id,result:"0x2BD"});
}
