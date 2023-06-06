var date = {
  type: "date",

  // all
  date: [
    {
      op: "lte",
      lte: 328248942,
    },
    {
      op: "lte_variable",
      lte_variable: "user.birthday",
    },
    {
      op: "lte_variable",
      lte_variable: "order.checkout",
    },
    {
      op: "gte_variable",
      gte_variable: "order.checkin",
    },
  ],
};

var event = {
  type: "event",
  events: "message_sent", // message_pong

  // all
  conditions: [
    {
      key: "data.message.id",
      op: "eq_variable",
      eq_variable: "action[ifdfgslf].id",
    },
    {
      key: "data.message.text",
      op: "is_phone",
    },
  ],
};

// Wait until N day after
//    |<--------- N days --------->|
// ---x--------------+>------------o+++++++++>
//   last_run       now           event
var action = {
  name: "Wait until N day after",
  type: "wait",
  wait: {
    branches: [
      {
        ui_type: "until N day",
        event_types: ["scheduled"],
        condition: {
          all: [
            {
              key: "action.last_run",
              date: {
                op: "before_now",
                before_now: N,
                unit: "day", // default ms
              },
            },
          ],
        },
      },
    ],
  },
};

// Wait util 1 days before user birthday
//              |<--- 1 days --->|
// -----+>------+++++++++++++++++x------------------->
//     now                    user.birthday
var action = {
  name: "wait until 1 days before user birthday",
  type: "wait",
  wait: {
    branches: [
      {
        ui_type: "",
        node_id: "next_node",
        event_types: ["scheduled", "user_info_updated"],
        condition: {
          all: [
            {
              key: "user.birthday",
              type: "date",
              date: {
                op: "after_now",
                after_now: N,
                unit: "day",
              },
            },
          ],
        },
      },
    ],
    timeout_branch: "node",
    timeout: 323,
    timeout_unit: "second",
  },
};

// Wait 1 day after checkin but do not cross 2 day
//          |<--- 1 day --->|<--- 1 day --->|
//  --------x---------------+++++++++++++++++-------|>------>
//     order.checkin                                     now
var action = {
  name: "wait 1 day after checkin but do not cross 2 day",
  type: "wait",
  wait: {
    ui_type: "wait",
    event_types: ["scheduled", "order_updated", "order_created"],
    condition: {
      all: [
        {
          key: `order.fields.key=43.`,
          type: "date",
          date: {
            op: "before_now",
            before_now: 1,
            before_now_max: 2,
            unit: "day",
          },
        },
      ],
    },
  },
};

// Wait until last email sent is read
var action = {
  name: "wait until the last email is read or clicked",
  type: "wait",
  wait: {
    event_types: ["message_pong"],
    branches: [
      {
        ui_type: "till message read",
        conditions: {
          all: [
            {
              key: "event.data.message.id",
              type: "text",
              text: {
                op: "eq_var",
                eq_variable: "actions.a434.messages.434.result.event.id",
              },
            },
            {
              one: [
                {
                  key: `event.data.message.pongs.type="message_read".type`,
                  type: "text",
                  text: {
                    op: "eq",
                    eq: "message_read",
                  },
                },
                {
                  key: `event.data.message.pongs.type="message_clicked".type`,
                  type: "text",
                  text: {
                    op: "eq",
                    eq: "message_received",
                  },
                },
              ],
            },
          ],
        },
      },
    ],
  },
};

// Wait until an new order is placed
var action = {
  name: "wait until a new order is placed",
  type: "wait",
  wait: {
    event_types: ["order_created"],
    branches: [
      {
        ui_type: "",
        node_id: "next",
        condition: {
          all: [
            {
              key: "order.value",
              type: "number",
              number: {
                op: "gte",
                gte: 500,
              },
            },
          ],
        },
      },
    ],
  },
};
