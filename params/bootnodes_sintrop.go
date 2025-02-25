package params

// SintropBootnodes are the enode URLs of the P2P bootstrap nodes running
// reliably and availably on the Sintrop network.
// They will be the first point of contact for nodes coming online
// to find peers.
var SintropBootnodes = []string{
	"enode://8d17546991a9032fea64a254d31819661e0213fb14bfde3a37826d8699747c19cbdc42b4183138a82aa557b9f270c97ebfa39179a1fa9eccdb895b25f987455d@174.138.124.223:30303",
}

// Once Sintrop network has DNS discovery set up,
// this value can be configured.
// var SintropDNSNetwork = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@example.network"
