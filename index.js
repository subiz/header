var fs = require('fs');
var path = require("path");
const { lstatSync, readdirSync } = require('fs');
const { join } = require('path');

var outdst = './header.swagger.json';

console.log("\033[0;34m# BUILDING REST SPEC");

const isDirectory = source => lstatSync(source).isDirectory();

const isSwaggerFile = source => !lstatSync(source).isDirectory() && source.endsWith(".swagger.json");

const getDirectories = source =>
			readdirSync(source).map(name => join(source, name)).filter(isDirectory);

const getSwaggerFiles = source =>
			readdirSync(source).map(name => join(source, name)).filter(isSwaggerFile);

function readAllDef() {
	var def = {};
	getDirectories("./").map(name => {
		var dname = name;
		getSwaggerFiles(name).map(name => {
			var protoswagger = fs.readFileSync(name).toString();
			var protodef = JSON.parse(protoswagger).definitions;
			delete(protodef[dname + 'AllType']);
			def = Object.assign({}, def, protodef);

			// delete ctx
			Object.keys(def).filter(key => def[key].properties != null).forEach(key => {
				delete(def[key].properties['ctx']);
			});

			// convert int64 to integer
			Object.keys(def).filter(key => def[key].properties != null).forEach(key => {
				var type = key;
				Object.keys(def[key].properties)
					.filter(key => def[type].properties[key].format === 'int64')
					.forEach(key => {
						def[type].properties[key].type = "integer";
					});
			});
		});
	});

	// rename some types
	filterAccountProp(def);
	return def;
}

function filterAccountProp(def) {
	def.Account = def.accountAccount;
	def.accountAccount = undefined;

	def.Agent = def.accountAgent;
	def.accountAgent = undefined;

	def.AgentGroup = def.accountAgentGroup;
	def.accountAgentGroup = undefined;

	def.Invitation = def.accountInvitation;
	def.accountInvitation = undefined;
}

fs.writeFileSync(outdst, JSON.stringify(readAllDef()));
console.log("\033[0;32mDone. (" + path.resolve(outdst) + ")\033[0;30m\n");
