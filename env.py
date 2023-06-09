import os
import re

# SERVER_IP 为 命令行传递的参数
SERVER_IP = os.sys.argv[1]
print("MODE:", SERVER_IP)

# 搜索当前文件夹里的所有.vue、.ts、.js和.go文件
# 除了"node_modules"文件夹里的文件

for root, dirs, files in os.walk("."):
    if "node_modules" in dirs:
        dirs.remove("node_modules")
    for file in files:
        if file.endswith(".vue") or file.endswith(".ts") or file.endswith(".js") or file.endswith(".go"):
            with open(os.path.join(root, file), "r+", encoding="utf-8") as f:
                print("正在处理文件:", os.path.join(root, file))
                # 将"{{server_ip}}"替换为"SERVER_IP"
                content = re.sub(r"{{server_ip}}", SERVER_IP, f.read())
                f.seek(0)
                f.write(content)
                f.truncate()

# Path: env.py
