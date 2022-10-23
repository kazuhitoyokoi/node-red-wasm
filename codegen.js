var fs = require('fs');
console.log('package main');
console.log('func flows() string { return `');
console.log(fs.readFileSync(process.argv[2]).toString());
console.log('`}');
