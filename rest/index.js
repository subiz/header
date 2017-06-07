var subiz = require('./subiz');
var fs = require('fs');
var path = require("path");

var outdst = './rest/subiz.json';

console.log("\033[0;34m===========BUILDING REST SPEC========");
fs.writeFile(outdst, JSON.stringify(subiz.def), function(err) {
	if (err) {
		console.log("\033[0;31m" + err + "\033[0;30m\n");
		return;
	}
	var absolutePath = path.resolve(outdst);
	console.log("\033[0;32mDone. (" + absolutePath + ")\033[0;30m\n");
});
