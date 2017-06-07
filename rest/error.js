exports.responses = {
	BadRequest: {
		description: 'The server cannot or will not process the request due to an apparent client error (e.g., malformed request syntax, too large size, invalid request message framing, or deceptive request routing)',
		schema: {
			$ref: '#/definitions/Error',
		},
		examples: {
			"application/json": {
				Code: '400010',
				Messages: 'The server cannot or will not process the request due to an apparent client error (e.g., malformed request syntax, too large size, invalid request message framing, or deceptive request routing)',
			},
		},
	},
	NotFound: {
		description: 'The requested resource could not be found but may be available in the future. Subsequent requests by the client are permissible.',
		schema: {
			$ref: '#/definitions/Error',
		},
		examples: {
			"application/json": {
				Code: '404324',
				Messages: 'The requested resource could not be found but may be available in the future. Subsequent requests by the client are permissible.',
			},
		},
	},
	Unauthorized: {
		description: 'The request has not been applied because it lacks valid authentication credentials for the target resource.',
		schema: {
			$ref: '#/definitions/Error',
		},
		examples: {
			"application/json": {
				Code: '401344',
				Messages: 'The request has not been applied because it lacks valid authentication credentials for the target resource.',
			},
		},
	},
	Forbidden: {
		description: 'The request was valid, but the server is refusing action. The user might not have the necessary permissions for a resource.',
		schema: {
			$ref: '#/definitions/Error',
		},
		examples: {
			"application/json": {
				Code: '403002',
				Messages: 'The request was valid, but the server is refusing action. The user might not have the necessary permissions for a resource.',
			},
		},
	},
	TooManyRequests: {
		description: 'The user has sent too many requests in a given amount of time ("rate limiting").',
		headers:{
      'X-RateLimit-Limit': {
        type: 'integer',
        description: 'Request limit per hour.',
			},
      'X-RateLimit-Remaining': {
        type: 'integer',
        description: 'The number of requests left for the time window.',
			},
      'X-RateLimit-Reset': {
        type: 'string',
        format: 'date-time',
				description: 'The UTC date/time at which the current rate limit window',
			}
		},
		schema: {
			$ref: '#/definitions/Error',
		},
		examples: {
			"application/json": {
				Code: '429',
				Messages: 'The server is currently unavailable (because it is overloaded or down for maintenance). Generally, this is a temporary state.',
			},
		},
	},
	InternalServerError: {
		description: 'The server is currently unavailable (because it is overloaded or down for maintenance). Generally, this is a temporary state.',
		schema: {
			$ref: '#/definitions/Error',
		},
		examples: {
			"application/json": {
				Code: '500124',
				Messages: 'The server is currently unavailable (because it is overloaded or down for maintenance). Generally, this is a temporary state.',
			},
		},
	},
};

exports.definitions = {
	Error: {
		type: 'object',
		properties: {
			Code: {
				readOnly: true,
				type: 'integer',
			},
			Message: {
				readOnly: true,
				type: 'string',
			},
			RequestId: {
				readOnly: true,
				type: 'string',
			},
		},
		required: ['Code', 'Message'],
	}
};
