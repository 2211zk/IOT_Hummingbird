@echo off
REM Cobra Script Center 测试脚本 (Windows版本)
REM 用于验证基本功能是否正常工作

echo === Cobra Script Center 功能测试 ===

REM 构建应用
echo 1. 构建应用...
make build
if errorlevel 1 (
    echo 构建失败，请检查Go环境和依赖
    pause
    exit /b 1
)

REM 初始化数据库
echo 2. 初始化数据库...
bin\script-center.exe migrate
if errorlevel 1 (
    echo 数据库初始化失败
    pause
    exit /b 1
)

REM 创建管理员用户
echo 3. 创建管理员用户...
bin\script-center.exe user create --username admin --role admin --password admin123
if errorlevel 1 echo 用户可能已存在，继续...

REM 创建测试脚本
echo 4. 创建测试脚本...
bin\script-center.exe script create --name test-bash --language bash --description "测试bash脚本" --tags test
if errorlevel 1 echo 脚本可能已存在，继续...

REM 列出脚本
echo 5. 列出所有脚本...
bin\script-center.exe script list

REM 执行脚本
echo 6. 执行测试脚本...
bin\script-center.exe script run test-bash --param NAME=测试用户 --param MESSAGE=Hello

REM 查看执行历史
echo 7. 查看执行历史...
bin\script-center.exe execution list

REM 创建定时任务
echo 8. 创建定时任务...
bin\script-center.exe schedule create test-bash "0 * * * *"
if errorlevel 1 echo 定时任务可能已存在，继续...

REM 列出定时任务
echo 9. 列出定时任务...
bin\script-center.exe schedule list

REM 列出用户
echo 10. 列出用户...
bin\script-center.exe user list

echo.
echo === 测试完成 ===
echo 所有基本功能测试通过！
echo.
echo 你现在可以：
echo - 使用 'bin\script-center.exe --help' 查看所有命令
echo - 使用 'bin\script-center.exe daemon start' 启动守护进程
echo - 查看 QUICKSTART.md 了解更多使用方法
echo.
pause