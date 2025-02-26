package params

// SequoiaBootnodes are the enode URLs of the P2P bootstrap nodes running
// reliably and availably on the Sequoia network.
// They will be the first point of contact for nodes coming online
// to find peers.
var SequoiaBootnodes = []string{
	"enode://135236fd6f19d8a66a44743698ca20b5dd389a80edc45d1c12f072af3bfa75a1499821e30250fc6cc1c3e99e5d731dbda1b11973418352e0ba312f3eeccf0bba@209.38.96.144:30303",
}

// Once Sequoia network has DNS discovery set up,
// this value can be configured.
// var SequoiaDNSNetwork = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@example.network"
