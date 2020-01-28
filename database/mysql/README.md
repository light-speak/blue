# mysql group replication 集群

一主二从

cnf文件代码：
host 和 port 根据实际情况更改

```
[mysqld]
log-error=/var/log/mysql/error.log

# Encoding
collation-server = utf8_unicode_ci
init-connect = 'SET NAMES utf8'
character-set-server = utf8

# Peplaction Framework
gtid_mode = ON
enforce_gtid_consistency = ON
master_info_repository = TABLE
relay_log_info_repository = TABLE
binlog_checksum = NONE
log_slave_updates = ON
log_bin = binlog
binlog_format = ROW

# Host specific replication configuration
server_id = 1
bind-address = "172.20.0.101"
report_host = "172.20.0.101"
loose-group_replication_local_address = "172.20.0.101:33061"

# Group Replication
transaction_write_set_extraction = XXHASH64
loose-group_replication_group_name = "0dacf75b-0227-43ca-ad50-d99212ff3681"
loose-group_replication_start_on_boot = on
loose-group_replication_group_seeds = "172.20.0.101:33061,172.20.0.102:33062,172.20.0.103:33063"
loose-group_replication_bootstrap_group = off
loose-group_replication_ip_whitelist = "172.20.0.1,172.20.0.101,172.20.0.102,172.20.0.103,127.0.0.1/8,172.16.0.0/16"

```

 #### 除了主节点
```
reset master;
```
 #### 所有节点
 rpl_user rpl_pass可变更
```
set sql_log_bin =0;
create user rpl_user@'%' identified by 'rpl_pass';
grant replication slave on *.* to rpl_user@'%';
flush privileges ;
set sql_log_bin =1;
change master to master_user ='rpl_user', master_password ='rpl_pass' for channel 'group_replication_recovery';
install plugin group_replication soname 'group_replication.so';
```

#### 主节点
```
set global group_replication_bootstrap_group=on;
start group_replication;
set global group_replication_bootstrap_group=off;
```
### 从节点
```
start group_replication;
```

### 命令
#### 节点列表
```
select * from performance_schema.replication_group_members;
```
#### 读写
```
show global variables like '%read_only';

如果super_read_only是启动的，那么该成员仅提供读服务；
如果super_read_only是关闭的，并且 replication_group_members 
中正常的成员n 满足 2n+1 > 整个GROUP成员个数，并且该成员的 member state是online，
则该成员可提供读写服务。
```



