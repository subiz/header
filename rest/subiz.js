var services = [];
services.push(require('./account'));
services.push(require('./agent'));
services.push(require('./chat'));
services.push(require('./error'));

var paths = [], tags = [], defs = [], resp = [];

var commondef = {
	Id: {
		type: 'object',
		required: ["RequestId", "Ok", "Id"],
		properties: {
			RequestId: {
				type: 'string',
			},
			Ok: {
				type: 'boolean',
			},
			Id: {
				type: 'string',
			},
		}
	}
};

for (var i = 0; i < services.length; i++) {
	var s = services[i];
	if (s.paths) paths.push(s.paths);
	if (s.tags) tags.push(s.tags);
	if (s.definitions) defs.push(s.definitions);
	if (s.responses) resp.push(s.responses);
}

var error = {
	'400': { $ref: '#/responses/BadRequest' },
	'401': { $ref: '#/responses/Unauthorized' },
	'403': { $ref: '#/responses/Forbidden' },
	'404': { $ref: '#/responses/NotFound' },
	'429': { $ref: '#/responses/TooManyRequests' },
	'500': { $ref: '#/responses/InternalServerError' },
};

// Add default error
for (var i = 0; i < paths.length; i++) {
	var pathsi = paths[i];
	for (var p in Object.keys(pathsi)) {
		var path = pathsi[Object.keys(pathsi)[p]];
		for (var m in Object.keys(path)) {
			var method = path[Object.keys(path)[m]];
			path[Object.keys(path)[m]].responses = Object.assign({}, method.responses, error);
		}
	}
}

defs.push(commondef);

exports.def = {
	swagger: '2.0',
	info: {
		description: 'This is a sample server Petstore server. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/). For this sample, you can use the api key `special-key` to test the authorization filters.',
		version: '4.0',
		title: 'Account',
		termsOfService: "https://subiz.com/terms/",
		contact: {
			email: "apiteam@subiz.com"
		},
		license: {
			name: "Apache 2.0",
			url: "http://www.apache.org/licenses/LICENSE-2.0.html"
		}
	},
	host: "api.subiz.com",
	basePath: "/v4",
	tags: tags,
	schemes: ["https"],
	paths: Object.assign({}, ...paths),
	// securityDefinitions: {
	// 	subiz_oauth: {
	// 	type: "oauth2",
	// 	authorizationUrl: "https://oauth2.subiz.com/authorize",
	// 		flow: "implicit"
	// 	},
	// 	api_key: {
	// 		type: "apiKey",
	// 		name: "api_key",
	// 		"in": "header"
	// 	}
	// },
	definitions: Object.assign({}, ...defs),
	responses: Object.assign({}, ...resp),
};
