# kenlm_go
###kenlm go语言封装
编译安装kenLM后使用

```bash
git clone https://github.com/goylord/kenlm.git --depth=1
cd kenlm
mkdir -p build
cd build
cmake ..
make -j 4
```

### 一些运行依赖（未筛选）
```bash
apt install build-essential libboost-system-dev libboost-thread-dev libboost-program-options-dev libboost-test-dev
apt install libbz2-dev zlib1g zlib1g-dev  liblzma-dev
apt install cmake
```