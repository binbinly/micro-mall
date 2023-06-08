// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.0.0/contracts/token/ERC20/IERC20.sol
interface IERC20 {
  //返回ERC20代币的名字，例如”My test token”。
  function name() external view returns (string memory);

  //返回代币的简称，例如：MTT，这个也是我们一般在代币交易所看到的名字。
  function symbol() external view returns (string memory);

  //返回token使用的小数点后几位。比如如果设置为3，就是支持0.001表示。
  function decimals() external view returns (uint8);

  //代币总量
  function totalSupply() external view returns (uint256);

  //返回某个地址(账户)的账户余额
  function balanceOf(address account) external view returns (uint256);

  //从代币合约的调用者地址上转移_value的数量token到的地址_to，并且必须触发Transfer事件。
  function transfer(address _to, uint256 _value) external returns (bool);

  //返回spender仍然被允许从owner提取的金额。
  function allowance(address owner, address spender) external view returns (uint256);

  //允许spender多次取回您的帐户，最高达amount金额。 如果再次调用此函数，它将以amount覆盖当前的余量。
  function approve(address spender, uint256 amount) external returns (bool);

  //从地址_from发送数量为_value的token到地址_to,必须触发Transfer事件。transferFrom方法用于允许合同代理某人转移token。条件是from账户必须经过了approve。
  function transferFrom(
    address _from,
    address _to,
    uint256 _value
  ) external returns (bool);

  event Transfer(address indexed from, address indexed to, uint256 value);
  event Approval(address indexed owner, address indexed spender, uint256 value);
}

abstract contract Context {
  function _msgSender() internal view virtual returns (address) {
    return msg.sender;
  }

  function _msgData() internal view virtual returns (bytes calldata) {
    return msg.data;
  }
}

contract ERC20 is Context, IERC20 {
  mapping(address => uint256) private _balances;

  mapping(address => mapping(address => uint256)) private _allowances;

  uint256 private _totalSupply;

  string private _name;

  string private _symbol;

  uint8 private _decimals;

  constructor(
    string memory name_,
    string memory symbol_,
    uint8 decimals_
  ) {
    _name = name_;
    _symbol = symbol_;
    _decimals = decimals_;
  }

  function name() public view virtual override returns (string memory) {
    return _name;
  }

  function symbol() public view virtual override returns (string memory) {
    return _symbol;
  }

  function decimals() public view virtual override returns (uint8) {
    return _decimals;
  }

  function totalSupply() public view virtual override returns (uint256) {
    return _totalSupply;
  }

  function balanceOf(address account) public view virtual override returns (uint256) {
    return _balances[account];
  }

  function transfer(address recipient, uint256 amount) public virtual override returns (bool) {
    _transfer(_msgSender(), recipient, amount);
    return true;
  }

  function allowance(address owner, address spender) public view virtual override returns (uint256) {
    return _allowances[owner][spender];
  }

  function approve(address spender, uint256 amount) public virtual override returns (bool) {
    _approve(_msgSender(), spender, amount);
    return true;
  }

  function transferFrom(
    address sender,
    address recipient,
    uint256 amount
  ) public virtual override returns (bool) {
    _transfer(sender, recipient, amount);

    uint256 currentAllowance = _allowances[sender][_msgSender()];
    require(currentAllowance >= amount, 'ERC20: transfer amount exceeds allowance');
    unchecked {
      _approve(sender, _msgSender(), currentAllowance - amount);
    }

    return true;
  }

  function increaseAllowance(address spender, uint256 addedValue) public virtual returns (bool) {
    _approve(_msgSender(), spender, _allowances[_msgSender()][spender] + addedValue);
    return true;
  }

  function decreaseAllowance(address spender, uint256 subtractedValue) public virtual returns (bool) {
    uint256 currentAllowance = _allowances[_msgSender()][spender];
    require(currentAllowance >= subtractedValue, 'ERC20: decreased allowance below zero');
    unchecked {
      _approve(_msgSender(), spender, currentAllowance - subtractedValue);
    }

    return true;
  }

  function _transfer(
    address sender,
    address recipient,
    uint256 amount
  ) internal virtual {
    require(sender != address(0), 'ERC20: transfer from the zero address');
    require(recipient != address(0), 'ERC20: transfer to the zero address');

    _beforeTokenTransfer(sender, recipient, amount);

    uint256 senderBalance = _balances[sender];
    require(senderBalance >= amount, 'ERC20: transfer amount exceeds balance');
    unchecked {
      _balances[sender] = senderBalance - amount;
    }
    _balances[recipient] += amount;

    emit Transfer(sender, recipient, amount);

    _afterTokenTransfer(sender, recipient, amount);
  }

  function _mint(address account, uint256 amount) internal virtual {
    require(account != address(0), 'ERC20: mint to the zero address');

    _beforeTokenTransfer(address(0), account, amount);

    _totalSupply += amount;
    _balances[account] += amount;
    emit Transfer(address(0), account, amount);

    _afterTokenTransfer(address(0), account, amount);
  }

  function _burn(address account, uint256 amount) internal virtual {
    require(account != address(0), 'ERC20: burn from the zero address');

    _beforeTokenTransfer(account, address(0), amount);

    uint256 accountBalance = _balances[account];
    require(accountBalance >= amount, 'ERC20: burn amount exceeds balance');
    unchecked {
      _balances[account] = accountBalance - amount;
    }
    _totalSupply -= amount;

    emit Transfer(account, address(0), amount);

    _afterTokenTransfer(account, address(0), amount);
  }

  function _approve(
    address owner,
    address spender,
    uint256 amount
  ) internal virtual {
    require(owner != address(0), 'ERC20: approve from the zero address');
    require(spender != address(0), 'ERC20: approve to the zero address');

    _allowances[owner][spender] = amount;
    emit Approval(owner, spender, amount);
  }

  function _beforeTokenTransfer(
    address from,
    address to,
    uint256 amount
  ) internal virtual {}

  function _afterTokenTransfer(
    address from,
    address to,
    uint256 amount
  ) internal virtual {}
}

/**
 * @dev 发布的token
 */
contract StudyToken is ERC20 {
  // 管理者
  address public owner;

  // 存储指定地址的铸币权限
  mapping(address => bool) public minters;

  // 构造函数，设置代币名称、简称、精度；将发布合约的账号设置为治理账号
  constructor() ERC20('StudyToken.finance', 'STD', 18) {
    owner = msg.sender;
  }

  /**
   * 铸币
   * 拥有铸币权限地址向指定地址铸币
   */
  function mint(address account, uint256 amount) public {
    require(minters[msg.sender], '!minter');
    _mint(account, amount);
  }

  /**
   * 设置治理管理员地址
   */
  function setOwner(address _owner) public {
    // 要求调用者必须为当前治理管理员地址
    require(msg.sender == owner, '!owner');
    // 更新owner
    owner = _owner;
  }

  /**
   * 添加铸币权限函数
   */
  function addMinter(address _minter) public {
    // 要求调用者必须为当前治理管理员地址
    require(msg.sender == owner, '!owner');
    // 变更指定地址_minter的铸币权限为true
    minters[_minter] = true;
  }

  /**
   * 移除铸币权限函数
   */
  function removeMinter(address _minter) public {
    // 要求调用者必须为当前治理管理员地址
    require(msg.sender == owner, '!owner');
    // 变更指定地址_minter的铸币权限为false
    minters[_minter] = false;
  }
}
