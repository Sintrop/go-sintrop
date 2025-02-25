package params

// SintropBootnodes are the enode URLs of the P2P bootstrap nodes running
// reliably and availably on the Sintrop network.
// They will be the first point of contact for nodes coming online
// to find peers.
var SintropBootnodes = []string{
	"enode://ce9c7db505ad747672f3f6871d95db3921d601bad5e193516e5093090ebe693977221a643e41acca1bd4fbef9acc3942928cdcfb10129735b7aacc01d7961e4b@212.32.255.80:30303",
}

// Once Sintrop network has DNS discovery set up,
// this value can be configured.
// var SintropDNSNetwork = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@example.network"
