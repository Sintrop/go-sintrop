package params

// SintropBootnodes are the enode URLs of the P2P bootstrap nodes running
// reliably and availably on the Sintrop network.
// They will be the first point of contact for nodes coming online
// to find peers.
var SintropBootnodes = []string{
	"enode://7aa8546bfb9ea34382d05dcbdba4cde940ed83090d5277e062f2e382ebb65ef0e7b00d293ed9b7641c3921692a40f2d89ac8b6998712f6957d21f0428f5142b1@189.5.4.127:25225",
	"enode://d7ea95f26ad680ddaabdb747f5de7d991f597a1a9662ff1cf3e6079bac1da736d83445b09bb5ba060dfbb88443829a1118fd9dd751c0c5939aba5cb5aa5d9e00@188.245.247.149:25225",
}

// Once Sintrop network has DNS discovery set up,
// this value can be configured.
// var SintropDNSNetwork = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@example.network"
