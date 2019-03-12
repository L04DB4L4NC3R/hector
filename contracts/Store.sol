pragma solidity ^0.5.5;

contract Store {
    
  string private from;
  string private to;
  string public timestamp;
  string public class;
  uint8 public ageSuggestive;

  constructor(string memory _class, string memory _from, string memory _to, string memory _timestamp, uint8 _ageSuggestive) public {
    from = _from;
    to = _to;
    class = _class;
    timestamp = _timestamp;
    ageSuggestive = _ageSuggestive;
  }

}