'use strict';

/** @type Egg.EggPlugin */
module.exports = {
  // had enabled by egg
  // static: {
  //   enable: true,
  // }
  mysql: {
    enable: true,
    package: "egg-dbs",
  },

  cors: {
    enable: true,
    package: "egg-cors",
  },

  redis: {
    enable: true,
    package: "egg-redis",
  },
};
