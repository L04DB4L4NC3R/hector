pragma solidity ^0.5.5;

contract Store {

  event dataSetter(string _class, string _from, string _to, string _timestamp, uint8 _ageSuggestive);
    
  string public from; // encrypted aadhar info
  string public to; // encrypted aadhar info
  string public timestamp;
  string public class;
  uint8 public ageSuggestive;

  // constructor(string memory _class, string memory _from, string memory _to, string memory _timestamp, uint8 _ageSuggestive) public {
  //   from = _from;
  //   to = _to;
  //   class = _class;
  //   timestamp = _timestamp;
  //   ageSuggestive = _ageSuggestive;
  // }

  function setData(string calldata _class, string calldata _from, string calldata _to, string calldata _timestamp, uint8 _ageSuggestive) external {
    from = _from;
    to = _to;
    class = _class;
    timestamp = _timestamp;
    ageSuggestive = _ageSuggestive;
  }
}