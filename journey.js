

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
        ui_type: "until",
        condition: {
          all: [
            {
              key: "now",
              datetime: {
                op: "after_relative",
                after_realtive_sec: 86400,
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
              key: "now",
              type: "date",
              datetime: {
                op: "after_var",
                after_var: `user.attributes.#(key=="birthday").value`,
                var_format: "2006-01-02T15:04:05Z07:00", // could set to auto
                after_var_sec: N * 86400,
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
    branches: [{
      event_types: ["scheduled", "order_updated", "order_created"],
      ui_type: "wait",
      condition: {
        all: [{
          key: 'now',
          type: "date",
          date: {
            op: "after_var",
            after_var: `order.fields.#(key="checkin").value`,
            var_format: 'unix_sec', // unix_milli, rfc3999
            after_var_sec: 86400,
            max_after_var_sec: 172800,
          },
        }],
      },
    }],
  },
};

// Wait until last email sent is read
var action = {
  name: "wait until the last email is read or clicked",
  type: "wait",
  wait: {
    branches: [
      {
        event_types: ["message_pong"],
        ui_type: "till message read",
        conditions: {
          all: [
            {
              key: "event.data.message.id",
              type: "text",
              text: {
                op: "eq_var",
                eq_variable: "_result.actions.a434.messages.434.event.id",
              },
            },
            {
              one: [
                {
                  key: `_event.data.message.pongs.#(type=="message_read")#.type`,
                  type: "text",
                  text: {
                    op: "eq",
                    eq: "message_read",
                  },
                },
                {
                  key: `_event.data.message.pongs.#(type=="message_clicked")#.type`,
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

// Wait until a random new order is placed
var action = {
  name: "wait until a new order is placed",
  type: "wait",
  wait: {
    branches: [
      {
        event_types: ["order_created"],
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

// DESIGN
// Next Execution Pointer (NEP)
// NEP trỏ đến action mà engine sẽ thực hiện ở lần breath in tiếp theo
// + Ban đầu: NEP sẽ trỏ tới root action của workflow
// + Khi breath in, engine sẽ thực thi action mà NEP đang trỏ
// + Sau khi thực hiện xong. NEP sẽ trỏ tới action tiếp theo
//   Có 4 loại action quyết định việc chọn NEP
//   1. Action thông thường: NEP sẽ được đưa tới action liền kế tiếp. Ví dụ: cập-nhật-thông-tin-user, gắn-tag-hội-thoại, gửi-tin-nhắn
//   2. Action branch: Có nhiều hơn 1 action liền kề, engine sẽ kiểm tra từng điều kiện và chuyển NEP tới nhánh đầu tiên thỏa mãn
//   3. Action do-waiting: Sẽ kiểm tra event vào, nếu không thỏa mãn, NEP sẽ không thay đổi
//   4. Action jump, end, restart: Chuyển NEP tới action quy định trong nội dung action, có thể là nhảy tới một action bất kỳ, quay lại action root hoặc kết thúc


// Cách expand action:
// Một số action là tổ hợp của các action đơn vị. Hiện tại chỉ có action Wait là tổ hợp của 
// ID của action là một chữ số. Ví dụ: "14", "15" hoặc "18"
// Khi expand, ID sẽ được nối vào action parent ID bằng ký tự "/". Như vậy ID có thể có dạng 
// "1", "1/2", "1/2/3" hoặc "1/2/3/4"
// Cách expand:
// 1. Chuyển ID thành mảng "1/2/3/4" -> ["1", "2", "3", "4"]
// 2. kiểm tra kiểu của "1". Nếu là wait thì gọi expand("wait"). Hàm này sẽ trả về một map các node con. Giả sử
//    { "1" => action1, "2" => action2}
// 3. Kiểm tra kiểu của node con "2". Nếu là kafka thì gọi expand("kafka"). Hàm sẽ trả về map node con khác. Giả sử
//    { "1" => actiona1, "2" => actiona2, "3" => actiona3, "4" => actiona4}
// 4. Kiểm tra kiểu của node cháy "3". Nếu là xyz thì gọi expand("xyz"). Hàm sẽ trả về map node con khác. Giả sử
//    { "1" => action1}
// 5. Nếu có node 4 thì thực hiện node 4, nếu ko có node 4 thì là dữ liệu đã hỏng, thoát ngay với exit code là expansion-failed

// Hành động đợi wait
// Đây là hành động phức tạp nhất
// Wait được expand ra làm 2 hành động nhỏ hơn
// 1. wait-canary/1, sẽ đọc nội dung của action, tìm ra các effect-events và đặt các scheduler cần thiết
// 2. do-wating/2, đóng vai trò sensor, nhận event breath in. Nếu là effect events thì gọi hàm giống wait-canary/1 để đặt các scheduler mới
//    nếu là execution-event  thì kiểm tra điều kiện trong branch. Nếu đúng thì thực thi, không thì thôi 

