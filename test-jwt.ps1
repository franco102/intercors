# JWT Authentication Test Suite for Windows PowerShell

$GO_API = "http://localhost:8080"
$NODE_API = "http://localhost:3000"

function Write-Header {
    param([string]$text)
    Write-Host "================================" -ForegroundColor Blue
    Write-Host $text -ForegroundColor Blue
    Write-Host "================================" -ForegroundColor Blue
    Write-Host ""
}

function Write-Test {
    param([string]$text)
    Write-Host $text -ForegroundColor Yellow
}

function Write-Success {
    param([string]$text)
    Write-Host "✓ $text" -ForegroundColor Green
}

function Write-Error {
    param([string]$text)
    Write-Host "✗ $text" -ForegroundColor Red
}

Write-Header "JWT Authentication Test Suite"

# Test 1: Health Check (Public)
Write-Test "Test 1: Health Check (Public)"
Write-Host "GET $GO_API/health"
try {
    $response = Invoke-RestMethod -Uri "$GO_API/health" -Method Get
    Write-Host ($response | ConvertTo-Json) -ForegroundColor Cyan
} catch {
    Write-Error "Health check failed: $_"
}
Write-Host ""

# Test 2: Login - Go API
Write-Test "Test 2: Login - Go API"
Write-Host "POST $GO_API/login"
try {
    $loginBody = @{
        username = "admin"
        password = "password"
    } | ConvertTo-Json

    $response = Invoke-RestMethod -Uri "$GO_API/login" -Method Post `
        -ContentType "application/json" `
        -Body $loginBody

    $GO_TOKEN = $response.token
    if ($GO_TOKEN) {
        Write-Success "Token obtained successfully"
        Write-Host "Token: $($GO_TOKEN.Substring(0, [Math]::Min(50, $GO_TOKEN.Length)))..."
    } else {
        Write-Error "Failed to obtain token"
        exit 1
    }
} catch {
    Write-Error "Login failed: $_"
    exit 1
}
Write-Host ""

# Test 3: Login - Node API
Write-Test "Test 3: Login - Node API"
Write-Host "POST $NODE_API/login"
try {
    $loginBody = @{
        username = "admin"
        password = "password"
    } | ConvertTo-Json

    $response = Invoke-RestMethod -Uri "$NODE_API/login" -Method Post `
        -ContentType "application/json" `
        -Body $loginBody

    $NODE_TOKEN = $response.token
    if ($NODE_TOKEN) {
        Write-Success "Token obtained successfully"
        Write-Host "Token: $($NODE_TOKEN.Substring(0, [Math]::Min(50, $NODE_TOKEN.Length)))..."
    } else {
        Write-Error "Failed to obtain token"
        exit 1
    }
} catch {
    Write-Error "Login failed: $_"
    exit 1
}
Write-Host ""

# Test 4: Rotate Matrix - Without Token (Should Fail)
Write-Test "Test 4: Rotate Matrix - Without Token (Should Fail)"
Write-Host "POST $GO_API/api/rotate (no token)"
try {
    $matrixBody = @{
        data = @(
            @(1, 2, 3),
            @(4, 5, 6),
            @(7, 8, 9)
        )
    } | ConvertTo-Json

    $response = Invoke-RestMethod -Uri "$GO_API/api/rotate" -Method Post `
        -ContentType "application/json" `
        -Body $matrixBody

    Write-Error "Should have rejected request without token"
} catch {
    if ($_.Exception.Response.StatusCode -eq 401) {
        Write-Success "Correctly rejected request without token"
        Write-Host "Response: $($_.Exception.Response.StatusCode)" -ForegroundColor Cyan
    } else {
        Write-Error "Unexpected error: $_"
    }
}
Write-Host ""

# Test 5: Rotate Matrix - With Token (Should Succeed)
Write-Test "Test 5: Rotate Matrix - With Token (Should Succeed)"
Write-Host "POST $GO_API/api/rotate (with token)"
try {
    $matrixBody = @{
        data = @(
            @(1, 2, 3),
            @(4, 5, 6),
            @(7, 8, 9)
        )
    } | ConvertTo-Json

    $headers = @{
        "Authorization" = "Bearer $GO_TOKEN"
    }

    $response = Invoke-RestMethod -Uri "$GO_API/api/rotate" -Method Post `
        -ContentType "application/json" `
        -Body $matrixBody `
        -Headers $headers

    Write-Success "Request succeeded"
    Write-Host ($response | ConvertTo-Json) -ForegroundColor Cyan
} catch {
    Write-Error "Request failed: $_"
}
Write-Host ""

# Test 6: Statistics - Without Token (Should Fail)
Write-Test "Test 6: Statistics - Without Token (Should Fail)"
Write-Host "POST $NODE_API/api/statistics (no token)"
try {
    $matrixBody = @{
        data = @(
            @(1, 2, 3),
            @(4, 5, 6),
            @(7, 8, 9)
        )
    } | ConvertTo-Json

    $response = Invoke-RestMethod -Uri "$NODE_API/api/statistics" -Method Post `
        -ContentType "application/json" `
        -Body $matrixBody

    Write-Error "Should have rejected request without token"
} catch {
    if ($_.Exception.Response.StatusCode -eq 401) {
        Write-Success "Correctly rejected request without token"
        Write-Host "Response: $($_.Exception.Response.StatusCode)" -ForegroundColor Cyan
    } else {
        Write-Error "Unexpected error: $_"
    }
}
Write-Host ""

# Test 7: Statistics - With Token (Should Succeed)
Write-Test "Test 7: Statistics - With Token (Should Succeed)"
Write-Host "POST $NODE_API/api/statistics (with token)"
try {
    $matrixBody = @{
        data = @(
            @(1, 2, 3),
            @(4, 5, 6),
            @(7, 8, 9)
        )
    } | ConvertTo-Json

    $headers = @{
        "Authorization" = "Bearer $NODE_TOKEN"
    }

    $response = Invoke-RestMethod -Uri "$NODE_API/api/statistics" -Method Post `
        -ContentType "application/json" `
        -Body $matrixBody `
        -Headers $headers

    Write-Success "Request succeeded"
    Write-Host ($response | ConvertTo-Json) -ForegroundColor Cyan
} catch {
    Write-Error "Request failed: $_"
}
Write-Host ""

# Test 8: Invalid Token (Should Fail)
Write-Test "Test 8: Invalid Token (Should Fail)"
Write-Host "POST $GO_API/api/rotate (invalid token)"
try {
    $matrixBody = @{
        data = @(
            @(1, 2, 3),
            @(4, 5, 6),
            @(7, 8, 9)
        )
    } | ConvertTo-Json

    $headers = @{
        "Authorization" = "Bearer invalid.token.here"
    }

    $response = Invoke-RestMethod -Uri "$GO_API/api/rotate" -Method Post `
        -ContentType "application/json" `
        -Body $matrixBody `
        -Headers $headers

    Write-Error "Should have rejected invalid token"
} catch {
    if ($_.Exception.Response.StatusCode -eq 401) {
        Write-Success "Correctly rejected invalid token"
        Write-Host "Response: $($_.Exception.Response.StatusCode)" -ForegroundColor Cyan
    } else {
        Write-Error "Unexpected error: $_"
    }
}
Write-Host ""

# Test 9: Wrong Credentials (Should Fail)
Write-Test "Test 9: Wrong Credentials (Should Fail)"
Write-Host "POST $GO_API/login (wrong password)"
try {
    $loginBody = @{
        username = "admin"
        password = "wrongpassword"
    } | ConvertTo-Json

    $response = Invoke-RestMethod -Uri "$GO_API/login" -Method Post `
        -ContentType "application/json" `
        -Body $loginBody

    Write-Error "Should have rejected wrong credentials"
} catch {
    if ($_.Exception.Response.StatusCode -eq 401) {
        Write-Success "Correctly rejected wrong credentials"
        Write-Host "Response: $($_.Exception.Response.StatusCode)" -ForegroundColor Cyan
    } else {
        Write-Error "Unexpected error: $_"
    }
}
Write-Host ""

Write-Header "✓ All tests completed!"
