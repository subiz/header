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
	def = replaceDefsToComponentsSchema(def);
	return {
		components:{
			schemas: def
		}
	};
}

function replaceDefsToComponentsSchema(def) {
	return JSON.parse(
		JSON.stringify(def)
			.replace(/"\$ref":"#\/definitions/g, `"$ref":"#/components/schemas`)
			.replace(/accountAgent/g, "Agent")
			.replace(/accountAccount/g, "Account")
			.replace(/accountAgentGroup/g, "AgentGroup")

			.replace(/conversationConversation/g, "Conversation")
			.replace(/conversationButton/g, "Button")
			.replace(/conversationMessage/g, "Message")
			.replace(/conversationAttachment/g, "Attachment")
			.replace(/conversationCannedResponse/g, "CannedResponse")
			.replace(/conversationPostback/g, "Postback")
			.replace(/conversationStartRequest/g, "StartRequest")
			.replace(/conversationRoute/g, "Route")
			.replace(/conversationRule/g, "Rule")
			.replace(/conversationGenericElementTemplate/g, "GenericElementTemplate")
			.replace(/eventRawEvent/g, "RawEvent")
			.replace(/eventBy/g, "By")
	);
}

function filterAccountProp(def) {
	//def.Account = def.accountAccount;
	//def.accountAccount = undefined;

	//def.Agent = def.accountAgent;
	//def.accountAgent = undefined;

	//def.AgentGroup = def.accountAgentGroup;
	//def.accountAgentGroup = undefined;

	def.commonContext = undefined;
	def.authCredential = undefined;
}

fs.writeFileSync(outdst, JSON.stringify(readAllDef(), null, 2));
console.log("\033[0;32mDone. (" + path.resolve(outdst) + ")\033[0;30m\n");
