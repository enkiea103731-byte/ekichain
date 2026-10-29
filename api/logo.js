export default function handler(req,res){
 res.setHeader('Access-Control-Allow-Origin','*');
 res.redirect(302,'/public/logo.png');
}
