@echo off
echo 正在修复workspace.xml文件...

if exist "D:\gowork\src\Hummingbird\.idea\workspace.xml" (
    echo 备份原始文件...
    copy "D:\gowork\src\Hummingbird\.idea\workspace.xml" "D:\gowork\src\Hummingbird\.idea\workspace.xml.bak"
    
    echo 创建新的workspace.xml文件...
    echo ^<?xml version="1.0" encoding="UTF-8"?^> > "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo ^<project version="4"^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^<component name="AutoImportSettings"^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^<option name="autoReloadType" value="ALL" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^</component^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^<component name="ChangeListManager"^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^<list default="true" id="20a1c5d9-3fce-406b-a3ab-a2c1bcec33f0" name="更改" comment="Initial commit"^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo       ^<change beforePath="$PROJECT_DIR$/.idea/workspace.xml" beforeDir="false" afterPath="$PROJECT_DIR$/.idea/workspace.xml" afterDir="false" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^</list^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^<option name="SHOW_DIALOG" value="false" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^<option name="HIGHLIGHT_CONFLICTS" value="true" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^<option name="HIGHLIGHT_NON_ACTIVE_CHANGELIST" value="false" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^<option name="LAST_RESOLUTION" value="IGNORE" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^</component^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^<component name="ProjectId" id="2dXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^<component name="ProjectViewState"^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^<option name="hideEmptyMiddlePackages" value="true" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^<option name="showLibraryContents" value="true" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^</component^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^<component name="SpellCheckerSettings" RuntimeDictionaries="0" Folders="0" CustomDictionaries="0" DefaultDictionary="应用程序级" UseSingleDictionary="true" transferred="true" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^<component name="TypeScriptGeneratedFilesManager"^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo     ^<option name="version" value="3" /^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo   ^</component^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    echo ^</project^> >> "D:\gowork\src\Hummingbird\.idea\workspace.xml"
    
    echo 修复完成！
) else (
    echo 错误：找不到文件 D:\gowork\src\Hummingbird\.idea\workspace.xml
)

pause