#!/bin/bash
clear 
export GOPATH="$(go env GOPATH)":"$PWD"
#export GOPATH="go env GOPATH":"$PWD"
echo "Current GOPATH:$GOPATH"
# echo "Para:$1"
# go run $1


# 目前的设置情况: 
# 1. 在 bashrc_profile 中设置了 ~/go/:${pwd}
# 2. 在VSC中设置了 go.InferGoPath = true; 之所以设置这个是因为VSC无法从GOPATH获取多个GoPath

# 另外一种可以实践的设置
# 在bashrc中把GOPATH设置成固定的一个

# 另外一种可以实践的设置
# 1.VSC中设置go.InferGoPath = true 设置成功后;会把当前的rootProject这个目录也作为一个goGOpath;这样每一个工程都有一个独立的goPath;
# 但是这样存在的问题是 在VSC的termial中因为没有添加rootProject而报错找不到包; 并且Debug的时候调试器也无法识别这个rootProject这个路径而报错
# 好在我们可以再Debug的时候传入一些参数来解决这个问题.但是在 termial中出现这个错误该如何解决.

# 2.