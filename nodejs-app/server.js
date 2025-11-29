const http = require('http');
const server = http.createServer((req, res) => {
  const path = req.url.replace(/^\/+/, '');
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  if (!path) { res.end('Welcome to root web'); }
  else { res.end(`you access the ${path}`); }
});
// Allow starting manually AND exporting for tests
if (require.main === module) {
  server.listen(3000, () => console.log('Running on 3000'));
}
module.exports = server;