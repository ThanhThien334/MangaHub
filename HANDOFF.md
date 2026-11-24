## Handoff Ngày 1 — A→B
**Đã xong:**
- Khung dự án Go (module github.com/ThanhThien334/MangaHub), ba chương trình api/catalog/tui build được.
- internal/config: Load() trả (Config, error), 8 trường (HTTPAddr, GRPCAddr, TCPAddr, UDPAddr, DBPath, CatalogDBPath, JWTSecret, LogLevel); UsingDefaultSecret().
- internal/logging: New(service, level) dùng slog, ghi ra stderr, có service=<tên>.
- api: GET /api/health, graceful shutdown.
- catalog: tạm thời chỉ mở listener TCP ở cổng gRPC (chưa phải gRPC thật).

**Cách chạy/kiểm tra nhanh:**
    go build ./...
    go vet ./...
    go test ./...
    go run ./cmd/api          # terminal 1
    go run ./cmd/catalog      # terminal 2
    curl.exe -i http://localhost:8080/api/health   # kỳ vọng 200 {"service":"api","status":"ok"}

**API/contract đã chốt:**
- GET /api/health -> 200 {"status":"ok","service":"api"}
- Cổng mặc định: HTTP 8080, gRPC 9090, TCP 9000, UDP 9001 (đổi bằng biến môi trường, xem .env.example)
- Định dạng lỗi HTTP thống nhất: {"error":{"code":"...","message":"..."}} (chưa có endpoint lỗi nào)
- Quy ước log: slog dạng text, mỗi dòng có service=<tên>

**Chưa xong / giới hạn đã biết:**
- catalog chưa phải gRPC thật (giai đoạn 3).
- Chưa có database, chưa có xác thực (giai đoạn 2).
- Gin đang ở chế độ debug (log GIN-debug); đặt GIN_MODE=release khi cần.

**Lỗi phát hiện cần B xem:** (chưa có)

**Việc B cần làm đầu tiên:**
- Chạy đúng các lệnh ở trên.
- Làm phần B của 01-foundation.md: TUI cơ bản với tview, helper ui.Do, APIClient.Health().

**.env Note**
- File .env.example chỉ để tham khảo, chương trình ko đọc file .env, nên sửa trong đó ko có tác dụng (bình thường không cần đổi gì vì mọi cấu hình đã có giá trị mặc định)
- Nếu muốn đổi (ví dụ cổng 8080 bị chiếm), đặt biến trong PowerShell trước khi chạy:
      $env:HTTP_ADDR=":8081"
      go run ./cmd/api