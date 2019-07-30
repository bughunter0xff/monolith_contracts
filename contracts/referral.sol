pragma solidity ^0.5.10;

import "../externals/ERC721.sol";
import "../internals/ownable.sol";

interface TKN {
    function transfer(address _to, uint256 _value) external returns (bool);
}

contract Referral is ERC721, Ownable {

    event MintedReferralTokens(address _from, uint _amount, uint _newSupply);
    event IssuedReferralTokens(address _from, address _to, uint _amount);
    event TrasferedReferalToken(address _from, address _to,uint _tokenID);
    event TransferReferralBonus(address _from, address _to,uint _amount);

    uint constant private _MAX_REF_TOKENS_GIVEAWAY = 5;

    uint public totalSupply;
    uint public referralIndex;
    uint public TKNBonus;
    uint public mintedTokens;

    TKN tkn;

    mapping (uint => bool) public activated;
    mapping (uint => address) public firstOwner;

    constructor(uint _totalSuply, address _TKNAddress) Ownable(msg.sender, false) public  {
        totalSupply = _totalSuply;
        tkn = TKN(_TKNAddress);
        TKNBonus = 10;
    }

    function mintReferralTokens(uint _amount) external onlyOwner {
        uint newMinted = mintedTokens + _amount;
        require(newMinted > mintedTokens, "overflow or 0 input");
        require(newMinted <= totalSupply, "total supply exceeded");
        for(uint i = mintedTokens; i < newMinted; i++) {
            _mint(msg.sender, i);
        }
        mintedTokens = newMinted;
        emit MintedReferralTokens(msg.sender, _amount, mintedTokens);
    }

    function issueReferralTokens(address _to, uint _amount) external onlyOwner {
        //there is no overflow check because the maximum issuance per tx is capped by _MAX_REF_TOKENS_GIVEN
        require(_amount <= _MAX_REF_TOKENS_GIVEAWAY, "too many referral tokens given!");
        uint toBeIssued;
        if(_amount == 0) {
            toBeIssued = 1;
        }
        else{
            toBeIssued = _amount;
        }
        require(referralIndex + toBeIssued <= mintedTokens, "tokens exceed the current suppply!");
        for(uint tokenID = referralIndex; tokenID < referralIndex + toBeIssued; tokenID++) {
            _transferFrom(msg.sender, _to, tokenID);
            firstOwner[tokenID] = _to;
        }
        referralIndex += toBeIssued;
        emit IssuedReferralTokens(msg.sender, _to, toBeIssued);
    }

    function transferReferralToken(address _to, uint _tokenID) external {
        _transferFrom(msg.sender, _to, _tokenID);
        emit TrasferedReferalToken(msg.sender, _to, _tokenID);
    }

    function transferReferralBonus(uint[] calldata _referralTokens) external onlyOwner {
        for (uint i = 0; i < _referralTokens.length; i++) {
            uint referralToken = _referralTokens[i];
            //do NOT  transfer bonus for previously activated cards
            if(!activated[referralToken]){
                activated[referralToken] = true;
                tkn.transfer(firstOwner[referralToken], TKNBonus);
                emit TransferReferralBonus(msg.sender, firstOwner[referralToken], TKNBonus);
            }
        }
    }

    function setReferralBonus(uint newBonus) external onlyOwner{
        TKNBonus = newBonus;
    }
}
