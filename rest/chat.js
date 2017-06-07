var fs = require('fs');

var protoswagger = fs.readFileSync('./proto/chat/chat.swagger.json').toString();
var protodef = JSON.parse(protoswagger).definitions;

exports.tags = {
	name: 'chat',
	description: 'Operattion on chat'
};

var tags = ['chats'];
var consumes = ['application/json'];
var produces = ['application/json'];

exports.paths = {
	'/chats': {
		post: {
			tags, consumes, produces,
			summary: 'start new chat',
			description: '',
			operationId: 'StartChat',
			parameters: [{
				name: 'haha',
				'in': 'body',
				description: "123",
				schema: {
					$ref: '#/definitions/ChatEvent'
				}
			}],
			responses: {
				200: {
					description: "Chat created",
					schema: {
						$ref: "#/definitions/Id"
					}
				}
			}
		}
	},
	'/chats/{ChatId}/events': {
		post: {
			tags, consumes, produces,
			summary: 'create new event',
			description: '',
			operationId: 'SendChat',
			parameters: [{
				name: 'ChatId',
				'in': 'path',
				description: 'Id of chat session',
				required: true,
				type: 'string',
			},{
				name: 'chatevent',
				'in': 'body',
				description: 'chat event',
				required: true,
				schema: {
					"$ref": "#/definitions/ChatEvent"
				},
			}],
			responses: {
				'200': {
					description: 'a chat event id',
					schema: {
						$ref: '#/definitions/Id'
					}
				}
			}
		}
	}
};

exports.definitions = Object.assign({}, protodef, {
	ChatEvent: {
		allOf: [{
			$ref: '#/definitions/chatChatEvent',
		}, {
			required: ["Type", "SenderId", "SenderType", "Text", "Format", "Attachment"],
			properties: {
				ChatId: {
					readOnly: true
				},
				Id: {
					readOnly: true,
				},
				CreatedTime: {
					readOnly: true,
				}
			}
		}],
	}
});
