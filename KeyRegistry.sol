// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract KeyRegistry {
    mapping(address => bytes) public ecdhPublicKeys;
    
    event PublicKeyRegistered(address indexed owner, bytes publicKey);
    
    /**
     * @notice Register ECDH public key with signature proof
     * @param publicKey The ECDH public key to register
     * @param signature The signature proving ownership
     */
    function registerPublicKey(bytes calldata publicKey, bytes calldata signature) external {
        require(publicKey.length == 65, "Public key must be 65 bytes");
        require(publicKey[0] == 0x04, "Invalid public key format");
        
        // Derive address from public key
        address derivedAddress = address(uint160(uint256(keccak256(publicKey[1:]))));
        
        // Verify the message signature
        bytes32 messageHash = keccak256(
            abi.encodePacked(
                "\x19Ethereum Signed Message:\n32",
                keccak256(abi.encodePacked("Register ECDH key"))
            )
        );
        
        address recoveredSigner = recoverSigner(messageHash, signature);
        
        // Verify both the derived address and signature match the sender
        require(derivedAddress == msg.sender && recoveredSigner == msg.sender, "Invalid ownership proof");
        
        ecdhPublicKeys[msg.sender] = publicKey;
        emit PublicKeyRegistered(msg.sender, publicKey);
    }
    
    function recoverSigner(bytes32 messageHash, bytes memory signature) internal pure returns (address) {
        require(signature.length == 65, "Invalid signature length");
        
        bytes32 r;
        bytes32 s;
        uint8 v;
        
        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))
            v := byte(0, mload(add(signature, 96)))
        }
        
        return ecrecover(messageHash, v, r, s);
    }
}



