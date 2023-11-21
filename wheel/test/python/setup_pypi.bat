@REM Config winget env 
winget source remove winget
winget source add winget https://mirrors.ustc.edu.cn/winget-source

@REM Install pip from winget 
winget install pip

@REM already installed
@REM C:\Users\chisato\anaconda3\envs\pyspark_env\Scripts

@REM Config tingshua env 
pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple

@REM Install pip
@REM python -m pip install --upgrade pip