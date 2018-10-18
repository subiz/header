var pb = require('protobufjs')

var root = new pb.Root()
root.resolvePath = (org, tar) =>
	pb.util.path.resolve(tar.startsWith('git.subiz.net') ? '../../' : org, tar)

const FILE = './auth/auth.proto'
const TYPE = 'auth.Permission'
const OBJ = { fullname: 'Thanh', email: 'abc@gmail.com' }
const BUFFER = new Buffer(
	'10e00118f601200e306438f00150ff0160f00168f001706e78f0018001ff018801f001a001f001a801f001b001f001b801e001c001e001c80140d001f001d80140e00140f001ff01f801f0018002fe018802f00190020f9802c001a00240a802f001',
	'hex'
)

root.load(FILE, err => {
	if (err) throw err
	const Type = root.lookupType(TYPE)

	console.log(
		Type.encode(OBJ)
			.finish()
			.toString('hex')
	)
	console.log(Type.decode(BUFFER))
})
