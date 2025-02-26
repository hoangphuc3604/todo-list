# To-Do List API & Frontend

## 1. Giới thiệu

Dự án này bao gồm một REST API được xây dựng bằng Golang với framework Gin, cùng với một frontend ReactJS để quản lý danh sách công việc (To-Do List). API hỗ trợ thêm, lấy, cập nhật và xóa công việc.

## 2. Công nghệ sử dụng

- **Backend**: Golang, Gin
- **Frontend**: ReactJS, Axios
- **Deployment**: Render

## 3. Cài đặt và chạy dự án

### 3.1 Backend

#### Yêu cầu

- Golang >= 1.18
- Git

#### Cài đặt và chạy

```sh
# Clone repository
git clone https://github.com/hoangphuc3604/todo-app-api

# Cài đặt dependency
go mod tidy

# Chạy server
go run main.go
```

### 3.2 Frontend

#### Yêu cầu

- Node.js >= 16
- Git

#### Cài đặt và chạy

```sh
# Clone repository
git clone https://github.com/hoangphuc3604/todo-list-ui

# Cài đặt dependency
npm install

# Chạy server
npm start
```

## 4. Triển khai

### 4.1 Backend

#### Deploy trên Render

1. Tạo tài khoản trên [Render](https://render.com/).
2. Tạo dịch vụ mới, kết nối với GitHub repository.
3. Cấu hình runtime là **Golang** và deploy backend.

### 4.2 Frontend

#### Deploy trên Render

1. Tạo dịch vụ mới trên Render, kết nối với GitHub repository.
2. Cấu hình build command:

   ```sh
    npm run build
   ```

3. Deploy frontend.

## 5. Ghi chú

- **CORS** đã được cấu hình để frontend có thể gọi API.

## 6. Các đường dẫn

- Repository [Backend](https://github.com/hoangphuc3604/todo-app-api)
- Repository [Frontend](https://github.com/hoangphuc3604/todo-list-ui)
- Link deploy [Backend](https://todo-list-eczu.onrender.com)
- Link deploy [Frontend](https://todo-list-ui-c5e6.onrender.com)
