#!/bin/bash

yarn build

tar zcvf dist.tar.gz dist/

scp dist.tar.gz root@34.96.239.6:/home/cnsoftbei/

ssh root@34.96.239.6 "cd /home/cnsoftbei/; tar zxvf dist.tar.gz; "
