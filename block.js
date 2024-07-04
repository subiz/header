var textbold = {
  type: "text",
  text: "Xin chao",
  bold: true,
};

var boldAndItalic = {
  type: "text",
  text: "Xin chao",
  bold: true,
  italic: true,
};

var p = {
  type: "paragraph",
  content: [
    {
      type: "text",
      text: "Xin chao",
    },
    {
      type: "text",
      text: "Subiz",
      bold: true,
    },
  ],
};

var div = {
  type: "div",
  content: [
    {
      type: "paragraph",
      content: [
        {
          type: "text",
          text: "Xin chao",
        },
        {
          type: "text",
          text: "Subiz",
          bold: true,
        },
      ],
    },
    {
      type: "text",
      text: "Subiz",
      bold: true,
    },
  ],
};

var css = {
  type: "text",
  text: "Xin chao",
  bold: true,
  italic: true,
  style: {
    color: "red",
    margin_top: "10px",
  },
};

// same as prosemirro
let ul = {
  type: "bullet_list",
  content: [
    {
      type: "list_item",
      content: [
        {
          type: "paragraph",
          content: [
            {
              type: "text",
              text: "Test list",
            },
          ],
        },
      ],
    },
  ],
};

// same as prosemirro.
let ol = {
  type: "ordered_list",
  content: [
    {
      type: "list_item",
      order: 1,
      content: [
        {
          type: "paragraph",
          content: [
            {
              type: "text",
              text: "Test number list",
            },
          ],
        },
      ],
    },
  ],
};

var table = {
  type: "table",
  content: [
    {
      type: "table_row",
      content: [
        {
          type: "table_cell",
          colspan: 1,
          rowspan: 1,
          content: [
            {
              type: "paragraph",
              content: [
                {
                  type: "text",
                  text: "Table",
                },
              ],
            },
          ],
        },
        {
          type: "table_cell",
          colspan: 1,
          rowspan: 1,
          content: [
            {
              type: "paragraph",
              content: [
                {
                  type: "text",
                  text: "table ",
                },
              ],
            },
          ],
        },
        {
          type: "table_cell",
          colspan: 1,
          rowspan: 1,
          content: [
            {
              type: "paragraph",
              content: [
                {
                  type: "text",
                  text: "table",
                },
              ],
            },
          ],
        },
      ],
    },
    {
      type: "table_row",
      content: [
        {
          type: "table_cell",
          colspan: 1,
          rowspan: 1,
          content: [
            {
              type: "paragraph",
              content: [
                {
                  type: "text",
                  text: "Table",
                },
              ],
            },
          ],
        },
        {
          type: "table_cell",
          colspan: 1,
          rowspan: 1,
          content: [
            {
              type: "paragraph",
              content: [
                {
                  type: "text",
                  text: "Tble",
                },
              ],
            },
          ],
        },
        {
          type: "table_cell",
          colspan: 1,
          rowspan: 1,
          content: [
            {
              type: "paragraph",
              content: [
                {
                  type: "text",
                  text: "table",
                },
              ],
            },
          ],
        },
      ],
    },
  ],
};

var h1 = {
  type: "heading",
  level: 1,
  content: [
    {
      type: "text",
      text: "Tiêu đề 1",
    },
  ],
};

var h2 = {
  type: "heading",
  level: 2,
  content: [
    {
      type: "text",
      text: "Tiêu đề 1",
    },
  ],
};

let blockquote = {
  type: "blockquote",
  content: [
    {
      type: "paragraph",
      content: [
        {
          type: "text",
          text: "wrap block",
        },
      ],
    },
  ],
};

let link = {
  type: "link",
  href: "https://gemini.google.com/",
  title: "Gemini",
  text: "https://gemini.google.com/",
  target: "_blank",
};

let section = {
  type: "section",
  text: "A message *with some bold text* and _some italicized text_.",
  format: "mrkdwn",
  fields: [
    {
      type: "mrkdwn",
      text: "High",
    },
    {
      attrs: { code: "angry" },
      type: "emoji",
    },
  ],
};

var hr = {
  type: "horizontal_rule",
};

var image = {
  type: "image",
  alt_text: "Xin chao",
  image: {
    url: "https://imageurl.com",
    width: 100,
    height: 100,
  },
};

// dynamic field
var dynamicfield = {
  attrs: { id: "bafwpjphbvtucvgt", value: "Giới tính", key: "user.gender" },
  text: "user.gender",
  type: "dynamic-field",
};

var emoji = {
  attrs: { code: "angry" },
  type: "emoji",
  text: "angry",
};

var mention = {
  type: "mention-agent",
  text: "@all",
};

// ma yeu cau cua ban la {{ticket.number}}

var title = {
  type: "paragraph",
  content: [
    {
      type: "text",
      text: "ma yeu cau cua ban la",
    },
    {
      type: "dynamic-field",
      text: "ticket.number",
    },
  ],
};
