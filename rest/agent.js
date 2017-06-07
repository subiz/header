exports.tags = {
	name: 'agent',
	description: 'Operattion on agent'
};

var tags = ['agents'];
var consumes = ['applucation/json'];
var produces = ['application/json'];

exports.paths = {
	'/agents': {
		post: {
			tags, consumes, produces,
			summary: 'add new agent to the account',
			description: '',
			operationId: 'AddAgent',
			responses: {
				"405": {
					description: 'invalid input'
				}
			},

		}
	}
};
