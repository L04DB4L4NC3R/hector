pragma solidity ^0.4.24;

contract Store {
    
  string private from;
  string private to;
  string public timestamp;
  string public type;
  uint8 public ageSuggestive;

  constructor(string _type, string _from, string _to, string _timestamp, _ageSuggestive) public {
    version = _from;
    from = _from;
    to = _to;
    type = _type;
    timestamp = _timestamp;
    ageSuggestive = _ageSuggestive;
  }

}