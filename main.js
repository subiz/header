var pb = require('protobufjs')

var root = new pb.Root()
root.resolvePath = (org, tar) =>
	pb.util.path.resolve(tar.startsWith('github.com') ? '../../../' : org, tar)

let FILE = './common.proto'
let TYPE = 'common.Permission'
let OBJ = { fullname: 'Thanh', email: 'abc@gmail.com' }
let BUFFER = new Buffer(
	'10e00118f601200e306438f00150ff0160f00168f001706e78f0018001ff018801f001a001f001a801f001b001f001b801e001c001e001c80140d001f001d80140e00140f001ff01f801f0018002fe018802f00190020f9802c001a00240a802f001',
	'hex'
)

FILE = './payment/payment.proto'
// TYPE = 'payment.InvoiceItem'
 TYPE = 'payment.Limit'
                     //2a104e657720737562736372697074696f6e3801450a57dd424a252a2310021a087374616e64617264280c320466726565408080fc94acfc97b2154dcdcc4c3f
//BUFFER = new Buffer('2a104e657720737562736372697074696f6e380145403566434a252a2310021a087374616e64617264280c32046672656540d4afb5a5eaf29db2154dcdcc4c3f', 'hex')
//BUFFER = new Buffer('2a1252656e657720737562736372697074696f6e3801454035e6424a1a12181a087374616e64617264200c280130eadef9c6e9a698b115', 'hex') // invoice
BUFFER = new Buffer('20904e38004002500160a09c0168904e800103', 'hex') // perm


OBJ = {"maxAutomations":10000,"maxSegments":0,"maxAgents":1,"canBuyAgent":true,"automationWebhookQuota":20000,"automationEmailQuota":10000,"maxRules":3}

root.load(FILE, err => {
	if (err) throw err
	const Type = root.lookupType(TYPE)

	console.log(Type.encode(OBJ).finish().toString('hex'))
	//console.log(JSON.stringify(Type.decode(BUFFER)))
})
