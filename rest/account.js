exports.tags = {
	name: 'account',
	description: 'Operattion on account'
};

var tags = ['accounts'];
var consumes = ['applucation/json'];
var produces = ['application/json'];

exports.paths = {
	'/accounts/': {
		post: {
			tags, consumes, produces,
			summary: 'add new account to the sotre',
			description: '',
			operationId: 'AddAccount',
			responses: {
			},
		},
	},
};
