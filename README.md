# Home Tunnel
A from scratch private VPN project.
## How it Functions
### Security
HomeTunnel uses very similar encryption and key creation to Wiregurad a widley used VPN. HomeTunnel works by first generating random 32 bytes where it is then used on curve25519 with a known base to generate a public key, this gives a private and random public and private key. The public key is then made public to anyone. When a connection is made it sends it uses the other parties public key with curve25519 and their private key, both parities do this giving them a common secrete key. This secrete key is then hashed using BLAKE2s, this is then split and gives both parties equal client and server keys to use in each direction never the same. These keys are used to encrypt and decrypt packets sent over the TUN (Network Tunneling Device).

### How does the Network Function
A IP table allows for the routing of packets across your computer, often these tables show 192.168.0.1/24 (Local Network) route to local network device and default to 192.168.0.1 (Router). We can change these IP tables however to say rather then default to our router but default to our TUN which we discuss shortly, however we also need to say that if traffic is going straight to our servers ip then we let it go to the router to avoid us loosing access to the network.

The TUN or Network Tunneling Device is a simple networking tool that allows for routing of network traffic to another location or for it to be altered. Here we use the TUN to take any packets and encrypt them then send them to a specific address aka the server. From their the sever takes the packet decrypts it