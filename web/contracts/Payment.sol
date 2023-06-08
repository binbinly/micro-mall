// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Payment {
  //管理者
  address payable owner;

  struct Trade {
    string id;
    string order_no;
    uint256 amount;
    uint256 date;
  }

  struct Pay {
    address addr;
    uint256 amount;
    bool is_refund;
  }

  //交易记录
  mapping(address => Trade[]) private trades;

  //支付记录
  mapping(string => Pay) private pays;

  //支付成功回调事件
  event Notify(address addr, string trade_id, string order_no, uint256 date);

  constructor() {
    owner = payable(msg.sender);
  }

  //支付
  function pay(
    string memory id,
    string memory order_no,
    uint256 value
  ) public payable {
    // require(msg.value == value, 'The trade does not match the value of the transaction');
    trades[msg.sender].push(Trade({ id: id, order_no: order_no, amount: msg.value, date: block.timestamp }));
    pays[order_no] = Pay({ addr: msg.sender, is_refund: false, amount: msg.value });
    emit Notify(msg.sender, id, order_no, value);
  }

  //提款
  function withdraw(uint256 amount) public payable {
    require(msg.sender == owner, 'Only owner can withdraw funds');
    owner.transfer(amount);
  }

  //buyer 交易笔数
  function tradesOf(address buyer) public view returns (uint256) {
    return trades[buyer].length;
  }

  //buyer 第 index 笔 交易详情
  function tradesOfAt(address buyer, uint256 index)
    public
    view
    returns (
      string memory,
      string memory,
      uint256 amount,
      uint256 date
    )
  {
    Trade[] memory ts = trades[buyer];
    require(ts.length > index, 'Payment does not exist');
    Trade memory trade = ts[index];
    return (trade.id, trade.order_no, trade.amount, trade.date);
  }

  //原路退款
  function refund(string memory order_no) public payable {
    require(msg.sender == owner, 'Only owner can withdraw funds');
    require(pays[order_no].addr != address(0), 'Pay log not found');
    Pay storage py = pays[order_no];
    payable(py.addr).transfer(py.amount);
    py.is_refund = true;
  }

  //检查是否已支付
  function checkPay(string memory order_no) public view returns (bool) {
    return pays[order_no].addr != address(0);
  }
}
