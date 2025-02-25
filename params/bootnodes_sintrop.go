package params

// SintropBootnodes are the enode URLs of the P2P bootstrap nodes running
// reliably and availably on the Sintrop network.
// They will be the first point of contact for nodes coming online
// to find peers.
var SintropBootnodes = []string{
	"enode://111dc89e4b54aecac9124497b79cffb5e3c96feff216836d2672a3dce96eb6d9922b9a09b9f17ee144d2da5639102aa77c36fe12cbd0d12d0cb30d23c6f0a320@189.5.4.127:30303",
	"enode://f8c4fd131fa56c05d9111503ee6c3a73e8d83bf32e2eac4dc0554e4507bb5c1d90a30aae741561a556fab09f6bd5cebdbdeed9cc6c1ec4003d1dbc618ace8eab@174.138.124.223:30303",
}

// Once Sintrop network has DNS discovery set up,
// this value can be configured.
// var SintropDNSNetwork = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@example.network"
