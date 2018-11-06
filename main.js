var pb = require('protobufjs')

var root = new pb.Root()
root.resolvePath = (org, tar) =>
	pb.util.path.resolve(tar.startsWith('git.subiz.net') ? '../../' : org, tar)

let FILE = './auth/auth.proto'
let TYPE = 'auth.Permission'
let OBJ = { fullname: 'Thanh', email: 'abc@gmail.com' }
let BUFFER = new Buffer(
	'10e00118f601200e306438f00150ff0160f00168f001706e78f0018001ff018801f001a001f001a801f001b001f001b801e001c001e001c80140d001f001d80140e00140f001ff01f801f0018002fe018802f00190020f9802c001a00240a802f001',
	'hex'
)

FILE = './payment/payment.proto'
TYPE = 'payment.InvoiceItem'

                     //2a104e657720737562736372697074696f6e3801450a57dd424a252a2310021a087374616e64617264280c320466726565408080fc94acfc97b2154dcdcc4c3f
//BUFFER = new Buffer('2a104e657720737562736372697074696f6e380145403566434a252a2310021a087374616e64617264280c32046672656540d4afb5a5eaf29db2154dcdcc4c3f', 'hex')
BUFFER = new Buffer('2a104e657720737562736372697074696f6e3801450a57dd424a252a2310021a087374616e64617264280c320466726565408080fc94acfc97b2154dcdcc4c3f', 'hex')

OBJ = {
	description: 'New subscription',
  quantity: 1,
  price: 110.67,
  data: {
    plan: {
      agentCount: 2,
      newPlan: 'standard',
      billingCycleMonth: 6,
      oldPlan: 'free',
      started: 1541437200000000000,
      savePercentage: 0.800000011920929,
		},
	},
}

root.load(FILE, err => {
	if (err) throw err
	const Type = root.lookupType(TYPE)

	console.log(Type.encode(OBJ).finish().toString('hex'))
	console.log(Type.decode(BUFFER))
})
