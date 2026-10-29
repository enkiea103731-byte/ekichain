export default function handler(req, res) {
  res.setHeader('Access-Control-Allow-Origin', '*');
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS');
  res.setHeader('Access-Control-Allow-Headers', 'Content-Type');
  res.setHeader('Content-Type', 'application/json');
  if (req.method === 'OPTIONS') return res.status(200).end();
  
  const body = req.body || {};
  const id = body.id || 1;
  const method = body.method;

  if (method === 'net_version') {
    return res.status(200).json({jsonrpc:"2.0", id, result:"701"});
  }
  if (method === 'eth_chainId') {
    return res.status(200).json({jsonrpc:"2.0", id, result:"0x2BD"});
  }
  if (method === 'eth_blockNumber' || method === 'eth_getBlockByNumber') {
    return res.status(200).json({jsonrpc:"2.0", id, result:"0x1"});
  }
  // GET que a Trust faz pra testar
  return res.status(200).json({jsonrpc:"2.0", id:1, result:"0x2BD", chainId:701});
}
