Example form
```json
{
  "type": "system",
  "name": "system_appoint_schedule",
  "appoint_schedule": {
    "components": [
      {
        "type": "div",
        "contents": [
          {
            "type": "text",
            "text": "Đặt lịch khám"
          }
        ]
      },
      {
        "type": "text",
        "text": "Hello"
      },
      {
        "type": "input",
        "key": "fullname",
        "default_value": "234",
        "link_attribute": "email",
        "placeholder": "123098123",
        "label": "Họ và tên",
        "input_type": "text",
        "is_required": true
      },
      {
        "type": "input",
        "key": "phone",
        "validation": "phone",
        "label": "SĐT",
        "input_type": "text",
        "is_required": true
      },
      {
        "type": "input",
        "label": "Chọn ca",
        "is_required": true,
        "input_type": "dropdown",
        "key": "shift",
        "input_options": [
          {
            "value": "Buổi sáng",
            "label": "Buổi sáng"
          },
          {
            "value": "Buổi chiều",
            "label": "Buổi chiều"
          }
        ]
      },
      {
        "type": "input",
        "key": "schedule",
        "label": "Chọn ngày",
        "is_required": true,
        "input_type": "date"
      }
    ]
  }
}
```
