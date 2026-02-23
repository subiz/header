## quick build go only
./build.sh

## Build docker image subiz/protobuild
```sh
docker build --progress=plain -t live360vn/protobuild:3.2 . && ./build.sh
```

### Merge

conversation keep id
but allow to list using new user id

order

note

attribute?

### Lead source

Bot
Agent
Web
WebPlugin
custom-x-y-z


### aiagent.collect_user_information_strategy

1. Open Helper (Rất nhẹ – Value first)
“Giúp trước – xin sau”
Đặc điểm
	•	Trả lời gần như đầy đủ
	•	Xin SĐT chỉ để hỗ trợ sâu hơn
	•	Không giữ thông tin

Tone: Nhẹ nhàng, chuyên gia
Tỉ lệ xin được SĐT: Thấp
Trải nghiệm KH: Rất tốt

Ngành dùng
	•	SaaS / CRM
	•	Tư vấn B2B
	•	Phần mềm, IT service

2. Permission-Based (Nhẹ – xin có quyền chọn
“Anh/chị chọn cách tiện hơn”

Đặc điểm
	•	Không ép
	•	Cho khách quyền từ chối
	•	Gọi là tuỳ chọn

Tone: Lịch sự, tôn trọng
Ngành dùng
	•	Giáo dục
	•	Tư vấn tài chính cao cấp
	•	Healthcare (phi lâm sàng)

3. Soft Gate (Trung bình – phổ biến nhất)

“Trả lời mồi – giữ phần quan trọng”

Đặc điểm
	•	Trả lời 60–70%
	•	Giữ thông tin ảnh hưởng quyết định
	•	Xin SĐT có lý do rõ ràng

Tone: Chuyên nghiệp, thực tế
Ngành dùng
	•	Bất động sản
	•	Ô tô
	•	Khoá học nghề
	•	Nội thất, xây dựng


4. Risk-Based Gate (Trung bình – dựa vào sợ sai)

“Nhầm là mất tiền”

Đặc điểm
	•	Nhấn mạnh rủi ro nếu hiểu sai
	•	Không nói chi tiết qua chat
	•	Xin SĐT để “tránh sai”

Tone: Cẩn trọng, chuyên gia
Ngành dùng
	•	BĐS (pháp lý)
	•	Bảo hiểm
	•	Đầu tư tài chính
	•	Pháp lý, thuế

5️⃣ Speed-to-Call (Trung bình – ưu tiên gọi)

“Chat chậm, gọi nhanh”

Đặc điểm
	•	Chuyển sang gọi rất sớm
	•	Chat chỉ là cửa vào

Tone: Gọn, nhanh
Ngành dùng
	•	BĐS
	•	Bảo hiểm
	•	Tuyển sinh
	•	Telesale


6️⃣ Scarcity Gate (Trung bình → Mạnh)

“Không gọi là hết”

Đặc điểm
	•	Dùng khan hiếm
	•	Tạo áp lực thời gian

Tone: Gấp nhưng không gắt
Ngành dùng
	•	BĐS dự án
	•	Event / vé
	•	Khuyến mãi giới hạn
	•	Tuyển sinh đợt cuối

⸻

7️⃣ Hard Gate (Aggressive)

“Không SĐT = không thông tin”

Đặc điểm
	•	Giữ toàn bộ thông tin lõi
	•	Xin SĐT ngay hoặc sớm
	•	Không tiếp tục chat nếu khách không cho

Tone: Thẳng, kiểm soát
Ngành dùng
	•	BĐS đại trà
	•	Lead giá rẻ
	•	Affiliate / CPA
	•	Sale mass

⸻

8️⃣ Process-Driven Gate (Bắt buộc về mặt quy trình)

“Thủ tục cần SĐT”

Đặc điểm
	•	SĐT là bước kỹ thuật
	•	Không mang tính sale

Tone: Trung lập
Ngành dùng
	•	Ngân hàng
	•	Bảo hiểm
	•	Hỗ trợ kỹ thuật
	•	CSKH

⸻

9️⃣ Trust-First (Chậm – xây niềm tin)

“Quan hệ trước – thông tin sau”

Đặc điểm
	•	Hội thoại dài
	•	Ít xin SĐT
	•	Ưu tiên độ tin cậy

Tone: Cá nhân hoá
Ngành dùng
	•	BĐS cao cấp
	•	Wealth management
	•	Luật sư riêng
	•	Dịch vụ VIP
