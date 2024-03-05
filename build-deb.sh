#!/bin/bash

echo "start build EzeFormat debian package"
echo ""

go build EzeFormat.go
echo "build EzeFormat success"
echo ""

# 执行 ./EzeFormat -v, 获取一个版本号 VER_CODE
VER_CODE=$("./EzeFormat" -v)
echo "Version Code: $VER_CODE"

# 获取当前时间，格式为 yy/MM/dd_HH_mm_ss
CURRENT_TIME=$(date +"%y%m%d%H%M%S")
echo "Current Time: $CURRENT_TIME"

# 创建文件夹 ./build-target/deb/VER_CODE_{yy/MM/dd_HH_mm_ss}/EzeFormat-deb
TARGET_DIR="./build-target/deb/${VER_CODE}_${CURRENT_TIME}"
mkdir -p "$TARGET_DIR"

echo "target build path: $TARGET_DIR"
echo ""

# 复制 deb-build-tpl 到 $TARGET_DIR/EzeFormat-deb 里面
# 复制 EzeFormat, config.yaml, res-static/ 到 $TARGET_DIR/EzeFormat-deb/opt/EzeFormat/ 里
cp -r "./deb-build-tpl" "$TARGET_DIR/eze-format-$VER_CODE"
mkdir -p "$TARGET_DIR/eze-format-$VER_CODE/opt/EzeFormat"

cp -r "./EzeFormat" "./res-static" "$TARGET_DIR/eze-format-$VER_CODE/opt/EzeFormat"

# 开始 build deb
cd "$TARGET_DIR"

echo "start build debian package in $TARGET_DIR ..."
echo ""

dpkg-deb --build eze-format-$VER_CODE

echo ""
echo "build success, target deb file: $TARGET_DIR/eze-format-$VER_CODE.deb"
