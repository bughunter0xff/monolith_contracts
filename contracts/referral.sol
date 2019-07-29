pragma solidity ^0.5.10;

import "../externals/ERC721.sol";
import "../internals/ownable.sol";

interface TKN {
    function transfer(address _to, uint256 _value) external returns (bool);
}

contract Referral is ERC721, Ownable {

    uint constant private _MAX_REF_TOKENS_GIVEN = 5;

    uint public totalSupply;
    uint public referralIndex;
    uint public TKNBonus;
    uint private batch = 10;
    uint public tokensMinted;

    TKN tkn;

    mapping (uint => bool) activated;
    mapping (uint => address) firstOwner;

    constructor(uint _totalSuply, address _TKNAddress) Ownable(msg.sender, false) public  {
        totalSupply = _totalSuply;
        tkn = TKN(_TKNAddress);
        TKNBonus = 10;
    }

    function mintReferralTokens() external onlyOwner {
        uint newMinted = tokensMinted + batch;
        require(newMinted <= totalSupply, "total supply exceeded");
        for(uint i = tokensMinted; i < newMinted; i++) {
            _mint(msg.sender, i);
        }
        tokensMinted = newMinted;
    }

    function issueReferralToken(address _to, uint _count) external onlyOwner {
        require(referralIndex + _count < totalSupply, "tokens exceeded the mac suppply!");
        require(_count <= _MAX_REF_TOKENS_GIVEN, "too many referral tokens given!");
        for(uint tokenid = referralIndex; tokenid < referralIndex + _count; tokenid++) {
            _transferFrom(msg.sender, _to, tokenid);
            firstOwner[tokenid] = _to;
        }
        referralIndex += _count;
    }

    function transferReferralToken(address _to, uint _tokenid) external {
        _transferFrom(msg.sender, _to, _tokenid);
    }

    function transferBonus(uint[] calldata _referralTokens) external onlyOwner {
        for (uint i = 0; i < _referralTokens.length; i++) {
            uint referralToken = _referralTokens[i];
            //do NOT  transfer bonus for previously activated cards
            if(!activated[referralToken]){
                activated[referralToken] = true;
                tkn.transfer(firstOwner[referralToken], TKNBonus);
            }
        }
    }

    function setBonus(uint newBonus) external onlyOwner{
        TKNBonus = newBonus;
    }
}
