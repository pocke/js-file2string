const execSync = require('child_process').execSync;

execSync('go run main.go README.md > out.js');
const readme = require('./out.js')['README.md'];
execSync('rm out.js');

const expected = "" + execSync('cat README.md');


if (readme !== expected) {
  console.error("Expected:\n" + expected + ",\nbut got:\n" + readme);
}
