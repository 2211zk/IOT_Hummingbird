# 测试用户注册
Write-Host "Testing User Registration..."
$registerData = @{
    user_name = "testuser"
    user_nickname = "Test User"
    department = 1
    mobile = "13800138000"
    email = "test@example.com"
    password = "password123"
    gender = "male"
    role = 1
    comment = "Test user"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/user/register" -Method POST -Body $registerData -ContentType "application/json"
    Write-Host "Registration Response:" -ForegroundColor Green
    $response | ConvertTo-Json
} catch {
    Write-Host "Registration Error: $($_.Exception.Message)" -ForegroundColor Red
}

# 测试用户登录
Write-Host "`nTesting User Login..."
$loginData = @{
    user_name = "testuser"
    password = "password123"
} | ConvertTo-Json

try {
    $loginResponse = Invoke-RestMethod -Uri "http://localhost:8080/user/login" -Method POST -Body $loginData -ContentType "application/json"
    Write-Host "Login Response:" -ForegroundColor Green
    $loginResponse | ConvertTo-Json
    
    # 保存token用于后续测试
    $token = $loginResponse.token
    
    # 测试获取用户信息（需要认证）
    if ($token) {
        Write-Host "`nTesting Get User (with auth)..."
        $headers = @{
            "Authorization" = "Bearer $token"
        }
        
        try {
            $userResponse = Invoke-RestMethod -Uri "http://localhost:8080/user/get?id=1" -Method GET -Headers $headers
            Write-Host "Get User Response:" -ForegroundColor Green
            $userResponse | ConvertTo-Json
        } catch {
            Write-Host "Get User Error: $($_.Exception.Message)" -ForegroundColor Red
        }
        
        # 测试用户列表（需要认证）
        Write-Host "`nTesting User List (with auth)..."
        try {
            $listResponse = Invoke-RestMethod -Uri "http://localhost:8080/user/list?page=1&page_size=10" -Method GET -Headers $headers
            Write-Host "User List Response:" -ForegroundColor Green
            $listResponse | ConvertTo-Json
        } catch {
            Write-Host "User List Error: $($_.Exception.Message)" -ForegroundColor Red
        }
    }
} catch {
    Write-Host "Login Error: $($_.Exception.Message)" -ForegroundColor Red
}