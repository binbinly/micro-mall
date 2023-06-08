"use strict";

const CryptoJS = require("crypto-js");

module.exports = {
  /**
   * 执行睡眠
   * @param {*} ms
   * @returns
   */
  sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  },

  /**
   * AES-256-ECB对称加密
   * @param text {string} 要加密的明文
   * @param secretKey {string} 密钥，43位随机大小写与数字
   * @returns {string} 加密后的密文，Base64格式
   */
  aesEcbEncrypt(text) {
    var keyHex = CryptoJS.enc.Base64.parse("ejQL1Ip26nFDG4hWTkfocyBSuMbZiVAr");
    var messageHex = CryptoJS.enc.Utf8.parse(text);
    var encrypted = CryptoJS.AES.encrypt(messageHex, keyHex, {
      mode: CryptoJS.mode.ECB,
      padding: CryptoJS.pad.Pkcs7,
    });
    return encrypted.toString();
  },

  /**
   * AES-256-ECB对称解密
   * @param textBase64 {string} 要解密的密文，Base64格式
   * @param secretKey {string} 密钥，43位随机大小写与数字
   * @returns {string} 解密后的明文
   */
  aesEcbDecrypt(textBase64) {
    var keyHex = CryptoJS.enc.Base64.parse("ejQL1Ip26nFDG4hWTkfocyBSuMbZiVAr");
    var decrypt = CryptoJS.AES.decrypt(textBase64, keyHex, {
      mode: CryptoJS.mode.ECB,
      padding: CryptoJS.pad.Pkcs7,
    });
    return CryptoJS.enc.Utf8.stringify(decrypt);
  },
};
