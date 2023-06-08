"use strict";

module.exports = {
  /**
   * 接口成功响应
   * @param {*} data
   * @param {*} msg
   * @param {*} code
   */
  apiSuccess(data = "", msg = "ok", code = 200) {
    this.body = { code, msg, data };
  },

  /**
   * 接口失败响应
   * @param {*} code
   * @param {*} msg
   */
  apiFail(code = 400, msg = "fail") {
    this.body = { msg, code };
  },

  /**
   * 生成jwt token
   * @param {*} value
   * @returns
   */
  signToken(value) {
    return this.app.jwt.sign(value, this.app.config.jwt.secret);
  },

  /**
   * 验证jwt token
   * @param {*} token
   * @returns
   */
  verifyToken(token) {
    return this.app.jwt.verify(token, this.app.config.jwt.secret);
  },
};
