const controllers = [
  {
    id: "Brook Controller (STANDARD GAMEPAD",
    label: "Brook UFB (Xbox 360 mode)",
    buttons: {
      jab: 2,
      strong: 3,
      fierce: 5,
      short: 0,
      forward: 1,
      roundhouse: 7,
      ppp: 4,
      kkk: 6,
      up: {
        isButton: true,
        index: 12,
      },
      down: {
        isButton: true,
        index: 13,
      },
      left: {
        isButton: true,
        index: 14,
      },
      right: {
        isButton: true,
        index: 15,
      },
    },
  },
  {
    id: "Brook P4 Wired Gamepad",
    label: "Brook UFB (PS4 mode)",
    buttons: {
      jab: 0,
      strong: 3,
      fierce: 5,
      short: 1,
      forward: 2,
      roundhouse: 7,
      ppp: 4,
      kkk: 6,
      up: {
        isButton: false,
        index: 7,
        value: -1,
      },
      down: {
        isButton: false,
        index: 7,
        value: 1,
      },
      left: {
        isButton: false,
        index: 6,
        value: -1,
      },
      right: {
        isButton: false,
        index: 6,
        value: 1,
      },
    },
  },
  {
    id: "Brook P3 Wired Gamepad",
    label: "Brook UFB (PS3 mode)",
    buttons: {
      jab: 0,
      strong: 3,
      fierce: 5,
      short: 1,
      forward: 2,
      roundhouse: 7,
      ppp: 4,
      kkk: 6,
      up: {
        isButton: false,
        index: 5,
        value: -1,
      },
      down: {
        isButton: false,
        index: 5,
        value: 1,
      },
      left: {
        isButton: false,
        index: 4,
        value: -1,
      },
      right: {
        isButton: false,
        index: 4,
        value: 1,
      },
    },
  },
  {
    id: "Brook POKKEN CONTROLLER",
    label: "Brook UFB (Switch mode)",
    buttons: {
      jab: 0,
      strong: 3,
      fierce: 5,
      short: 1,
      forward: 2,
      roundhouse: 7,
      ppp: 4,
      kkk: 6,
      up: {
        isButton: false,
        index: 5,
        value: -1,
      },
      down: {
        isButton: false,
        index: 5,
        value: 1,
      },
      left: {
        isButton: false,
        index: 4,
        value: -1,
      },
      right: {
        isButton: false,
        index: 4,
        value: 1,
      },
    },
  },
];

export default controllers;
