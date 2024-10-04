# Sử dụng Golang image
FROM golang:alpine

# Thiết lập thư mục làm việc
WORKDIR /app

# Sao chép module và go.sum
COPY go.mod go.sum ./

# Tải về các phụ thuộc
RUN go mod download

# Sao chép mã nguồn và tệp .env vào container
COPY . .

# Biên dịch ứng dụng
RUN go build -o main .

# Mở cổng 8080
EXPOSE 8080

# Chạy ứng dụng
CMD ["/app/main"]