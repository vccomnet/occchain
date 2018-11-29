// Copyright 2015 The go-blockchain Authors
// This file is part of the go-blockchain library.
//
// The go-blockchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-blockchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-blockchain library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Blockchain network.
var MainnetBootnodes = []string{
	// Blockchain Foundation Go Bootnodes
	//"enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:40404", // IE
	//"enode://3f1d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@13.93.211.84:40404",  // US-WEST
	//"enode://78de8a0916848093c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@191.235.84.50:40404", // BR
	//"enode://158f8aab45f6d19c6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@13.75.154.138:40404", // AU
	//"enode://1118980bf48b0a3640bdba04e0fe78b1add18e1cd99bf22d53daac1fd9972ad650df52176e7c7d89d1114cfef2bc23a2959aa54998a46afcf7d91809f0855082@52.74.57.123:40404",  // SG

	// Blockchain Foundation C++ Bootnodes
	//"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:40404", // DE

        "enode://c3d0ae842a19ef1b2695b5fad17abde1ab51fc265d8c71d1504d02d9123dfeaff908bb459dca458ecab3ed4a91038e2fe51e8c2fd42872552c496f8866b81d39@47.74.14.180:40404",	//OCC-H
	"enode://61d6b7cdd800ef721510ea57369c22465934fc9c50190ddf5aa27d6684a1e5f33db345d9080a1861c3bb5ce22a310213a36d403047f6b9b21e2d38b04496372e@47.91.16.245:40404",	//OCC-G
	"enode://267c0ba45a900d11ec3ca5ad7d1cfee575458454d4e78380f3197fd6f622fd8353d79b54c59f0e459458dcb2e2948536cab8e6b3f267a4d116226e16ac725ea1@47.88.60.4:40404",	//OCC-F
	"enode://43dfb6aa0bbe1d3e6ebb437272fe2683363db11fafc3919ba761c96128f98caf90ffc7ba2301d68099b8fee8db634abae8d5d53438aa177a7b888e695921aa07@47.75.178.173:40404",	//OCC-B
	"enode://7450c1871c4b9fd62bbcb139e7f0754b1afbbde151d0ee0c327cbf263f88fa9d2874396bdcf2f5d56556867b1024cab56cdecdf6e5374f27de4861e237ce5dd1@107.150.103.94:40404",//OCC-D
	"enode://777d0622ddfa8f9bf66fd452444402272a97e8f1ff0bed6721d13fcb9bad7d021143824621a444d51f17182ffc7aa12821d07b45f5cbfe74c0cd6f7cdac4009c@107.150.99.139:40404",//OCC-C
	"enode://4409eb68e6e2466df41981e95a2e280b938ef1c51a2b4943ef41b05e9dae9fe6c11709525788c15290e9603f3344f62d3d388a59e08ac3c7d59a5198e41b3e48@47.75.133.171:40404",	//OCC-A
	"enode://88292e6d150be4a0a8e64f8eb6998c576342ec9a5a8df8f591f32418c377deb5fde4731be69de0a84d8d8600a956e934b05096d61ab57ecddd1139301a73d74d@47.88.46.12:40404",	//OCC-E
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	//"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:40404",    // US-Azure occ
	//"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:40404",  // US-Azure parity
	//"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:40404", // Parity
	//"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:40404", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	//"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:40404", // IE
	//"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:40404",  // INFURA
	//"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:40404", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	//"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:40404",
	//"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	//"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	//"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
