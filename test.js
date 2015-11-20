const fs = require('fs');
const readme = require('./out.js')['README.md'];

fs.readFile('./README.md', 'utf8', (err, text) => {
  if (text !== readme) {
    console.error('not equal');
    console.error(text);
    console.error(readme);
    process.exit(1);
  }
});
