@echo off
echo 正在删除workspace.xml文件...

if exist "D:\gowork\src\Hummingbird\.idea\workspace.xml" (
    echo 备份原始文件...
    copy "D:\gowork\src\Hummingbird\.idea\workspace.xml" "D:\gowork\src\Hummingbird\.idea\workspace.xml.bak"
    
    echo 删除原始文件...
    del "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    
    echo 删除完成！请重新打开GoLand，它会自动创建新的workspace.xml文件。
) else (
    echo 错误：找不到文件 D:\gowork\src\Hummingbird\.idea\workspace.xml
)

pause