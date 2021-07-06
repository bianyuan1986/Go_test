#!/bin/bash

#lsof -nP -iTCP:80 
#lsof -nP -iTCP:80 -sTCP:LISTEN
#lsof -nP -sTCP:LISTEN|grep keepper|grep IPv4
#ps -eLF|grep PID
#删除文件名称中包含‘-’字符的文件
#rm -- -help.html
#touch -- -help.html
#查看系统启动后运行的时间和平均负载
#uptime
#系统启动时间和系统空闲时间，单位为秒
#cat /proc/uptime
#添加-x选项获取go build执行的过程
#go build -x
#解压缩rpm包中内容
#rpm2cpio t-yundun-waf-scc-1.0.122-20200908214445.x86_64.rpm |cpio -div
#scp -i ~/.ssh/compute_nglist.pem ./t-yundun-waf-scc-1.0.122-20200909103702.x86_64.rpm root@39.108.71.145:/root
#curl http://183.36.0.197/cc -H "Host:test2.wafqa3.com"
#go build -tags netgo -gcflags "-N -l" -o ${target_path}/bin/keepper main.go
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -gcflags "-N -l" -o ${target_path}/bin/keepper main.go
#cgo交叉编译
#CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build go-lua-compare.go
#CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -a -v go-lua-compare.go
#编译gcc，--disable-multilib表明只支持64位程序
#../gcc-4.9.0/configure --prefix=/home/baohu.rbh/gcc_install --disable-multilib --enable-languages=c,c++,fortran,go
#编译gcc时可能提示"cannot compute suffix of object files: cannot compile"，导出如下环境变量重新编译可通过
#export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/path/to/mpc/lib/
#编译hyperscan
#..//cmake-3.5.2-Linux-x86_64/bin/cmake -DCMAKE_C_COMPILER=/home/baohu.rbh/gcc_install/bin/gcc -DCMAKE_CXX_COMPILER=/home/baohu.rbh/gcc_install/bin/g++ -DCMAKE_BUILD_TYPE=Release -DBOOST_ROOT=/home/baohu.rbh/boost_1_74_0 /home/baohu.rbh/hyperscan-4.5.1/
#gcc -march=native -Q --help=target | grep march
#set scheduler-locking on|off
#curl -v http://127.0.0.1:9090/v1/apps/m.yundun.waf.1/lua/script/add -H "Content-Type: application/json" -d '{"Name":"tiger", "Description":"test case", "Enable":"on", "Content":"print(1)"}'
#curl -v http://127.0.0.1:9090/v1/apps/m.yundun.waf.1/lua/script/mod -H "Content-Type: application/json" -d '{"Name":"tiger", "Description":"test case", "Enable":"on", "Content":"print(1)"}'
#curl -v http://127.0.0.1:9090/v1/apps/m.yundun.waf.1/lua/script/del -d '{"Name":"tiger"}'
#curl -v http://127.0.0.1:9090/v1/apps/m.yundun.waf.1/lua/script/get?name=tiger
#curl -v http://127.0.0.1:9090/v1/apps/m.yundun.waf.1/lua/scripts
#./etcdctl --endpoint="http://127.0.0.1:4001" get /luffy/core/apps/m.yundun.waf.1/plugins/lua/config
#./etcd --listen-client-urls http://localhost:4001 --advertise-client-urls http://localhost:4001
#etcdctl member list
#etcdctl cluster-health
#etcdctl get /luffy/core/apps/m.yundun.waf.1/plugins/lua/config
#curl http://127.0.0.1:4001/v2/keys/luffy/core/apps/m.yundun.waf.1/plugins/lua/config
#curl 'localhost:56688/v1/sconfig/domain/www.tiger.test.com?worker_id=9&version=1607055200'
#curl -v http://127.0.0.1:9090/v1/apps/m.yundun.waf.1/lua/script/add -H "Content-Type: application/json" -d '{ "Name":"tiger20", "Description":"test case", "Enable":"on", "Content":"ruleId=waf.GetMatchRuleId(); str1=waf.GetMatchValue(\"queryarg_values\"); print(\"Match Value:[\", str1, \"]|Rule Id:[\", ruleId, \"]\"); if (string.sub(str1,0,7)==\"select\" and string.find(str1,\"where\",1)) then print \"complete sql\" end" }'
#curl -v http://127.0.0.1:9090/v1/apps/m.yundun.waf.1/lua/script/mod -H "Content-Type: application/json" -d '{ "Name":"tiger20", "Description":"test case", "Enable":"on", "Content":"ruleId=waf.GetMatchRuleId(); str1=waf.GetMatchValue(\"queryarg_values\"); print(\"Match Value:[\", str1, \"]|Rule Id:[\", ruleId, \"]\"); if (string.sub(str1,0,7)==\"select\" and string.find(str1,\"where\",1)) then print \"complete sql\" end" }'
#curl -v http://127.0.0.1:9090/v1/apps/m.yundun.waf.1/lua/scripts
#curl -v http://127.0.0.1:9090/v1/apps/m.yundun.waf.1/lua/bypass -H "Content-Type: application/json" -d '{ "Bypass":"off" }'
#查看glibc版本
#ldd --version
#vim 
#[[跳转到函数开头处 [{跳转到当前函数开头处  }]跳转到当前函数结尾处 ]]跳转到下一个函数开头处
#docker exec -it containerID /bin/bash
#curl -v "http://127.0.0.1/" -H "Host: www.tiger.test.com" -F tiger=@PAYLOAD_FILE_NAME
#curl -v "http://127.0.0.1" -H "Host: www.tiger.test.com" -H "Content-Type: application/x-www-form-urlencoded" -d 'x=select aaaa as sdsd select 1 from 2 where 1=cast(1,2,3) where 12'

#git checkout -b my-test  //在当前分支下创建my-test的本地分支分支
#git push origin my-test  //将my-test分支推送到远程
#git branch --set-upstream-to=origin/my-test //将本地分支my-test关联到远程分支my-test上
#git branch -a //查看远程分支

#显示本地主机上的镜像
#docker images
#如果未指定tag则使用最新版本，Command为要执行的命令
#docker run -t -i IMAGE_NAME:TAG COMMAND(/bin/bash)
#docker pull

#显示本地所有容器，包括停止状态的容器
#docker ps -a
#docker rm CONTAINER_ID
#docker rmi IMAGE_NAME

#yum install PACKAGENAME -b test
#curl -v "http://127.0.0.1/gzipTest" -H "Host: cwaf-rf05-attack2.wafqa3.com" -H "Content-Encoding: gzip" --data-binary @payload

#date -d "1970-01-01 UTC `echo "$(date +%s)-$(cat /proc/uptime|cut -f 1 -d' ')+125248258.304112"|bc ` seconds" +'%Y-%m-%d %H:%M:%S'

#ulimit -c unlimited && GOTRACEBACK=crash ./binary
#ulimit -c unlimited && export GOTRACEBACK=crash
