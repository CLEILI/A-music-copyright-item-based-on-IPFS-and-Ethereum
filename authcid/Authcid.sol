// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;
contract authcid{
    mapping(string=>address) public auth;

    function store(string calldata cid) public {
        require(auth[cid]==address(0),"the cid has the author");
        auth[cid]=msg.sender;
    }
    function queryauth(string calldata cid)public view returns(address){
        require(auth[cid]!=address(0),"the cid don't have a author");
        return auth[cid];
    }
}