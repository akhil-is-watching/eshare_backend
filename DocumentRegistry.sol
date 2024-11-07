// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract DocumentRegistry {
    
    struct RootDocument {
        address owner;           // Owner of the document
        bytes32 merkleRoot;      // Merkle root of all shared versions
    }

    // Mapping from IPFS hash (CID) to RootDocument
    mapping(bytes32 => RootDocument) public rootDocuments;

    // Events
    event DocumentRegistered(bytes32 indexed documentHash, address owner);
    event DocumentShared(
        bytes32 indexed rootDocumentHash,
        bytes32 indexed sharedDocumentHash,
        address indexed sharedWith
    );
    event MerkleRootUpdated(bytes32 indexed documentHash, bytes32 newMerkleRoot);

    // Register a new root document
    function registerDocument(bytes32 _ipfsHash) external {
        require(rootDocuments[_ipfsHash].owner == address(0), "Document already registered");
        
        rootDocuments[_ipfsHash] = RootDocument({
            owner: msg.sender,
            merkleRoot: bytes32(0)
        });

        emit DocumentRegistered(_ipfsHash, msg.sender);
    }

    // Update merkle root when sharing document
    function updateMerkleRoot(
        bytes32 _rootDocumentHash,
        bytes32 _newMerkleRoot,
        bytes32 _sharedHash,
        address _sharedWith
    ) external {
        RootDocument storage doc = rootDocuments[_rootDocumentHash];
        require(msg.sender == doc.owner, "Only owner can update merkle root");

        doc.merkleRoot = _newMerkleRoot;

        emit DocumentShared(_rootDocumentHash, _sharedHash, _sharedWith);
        emit MerkleRootUpdated(_rootDocumentHash, _newMerkleRoot);
    }

    // Verify if a shared document hash is valid using merkle proof
    function verifySharedDocument(
        bytes32 _rootDocumentHash,
        bytes32 _sharedHash,
        bytes32[] calldata _merkleProof
    ) external view returns (bool) {
        bytes32 merkleRoot = rootDocuments[_rootDocumentHash].merkleRoot;
        bytes32 leaf = _sharedHash;
        
        for (uint256 i = 0; i < _merkleProof.length; i++) {
            if (leaf < _merkleProof[i]) {
                leaf = keccak256(abi.encodePacked(leaf, _merkleProof[i]));
            } else {
                leaf = keccak256(abi.encodePacked(_merkleProof[i], leaf));
            }
        }
        
        return leaf == merkleRoot;
    }

    // Get root document details
    function getRootDocument(bytes32 _documentHash) external view returns (
        address owner,
        bytes32 merkleRoot
    ) {
        RootDocument memory doc = rootDocuments[_documentHash];
        return (doc.owner, doc.merkleRoot);
    }
}