# ProxySQL


## 部署
使用和MySQL同一个网络环境

docker镜像默认不自带mysql

使用以下命令管理
```
    mysql -h127.0.0.1 -P16032 -uradmin -pLTYlty0123 --prompt "ProxySQL Admin>"
```
在 MGR 集群中创建proxysql使用的账号
```
create user 'proxysql'@'%' identified by 'proxysql';
grant all on *.* to 'proxysql'@'%';
flush privileges;
```
在MGR集群中，执行addition_to_sys.sql  ，创建系统视图
可通过以下代码查看视图
```
select * from sys.gr_member_routing_candidate_status;
```
在proxysql 中增加管理账号
```
insert into mysql_users(username,password,default_hostgroup) values ('proxysql','proxysql',1);
update global_variables set variable_value='proxysql' where variable_name='mysql-monitor_username';
update global_variables set variable_value='proxysql' where variable_name='mysql-monitor_password';
load mysql servers to runtime;
save mysql servers to disk;
```
配置proxysql
```
delete from mysql_servers;

insert into mysql_servers (hostgroup_id, hostname, port, weight) values(1,'172.20.0.101',3306,100);
insert into mysql_servers (hostgroup_id, hostname, port, weight) values(1,'172.20.0.102',3306,100);
insert into mysql_servers (hostgroup_id, hostname, port, weight) values(1,'172.20.0.103',3306,100);
insert into mysql_servers (hostgroup_id, hostname, port, weight) values(2,'172.20.0.101',3306,100);
insert into mysql_servers (hostgroup_id, hostname, port, weight) values(2,'172.20.0.102',3306,100);
insert into mysql_servers (hostgroup_id, hostname, port, weight) values(2,'172.20.0.103',3306,100);

```
然后可以查看mysql列表
```
select * from mysql_servers;
```

删除读写分离相关规则
```
delete from mysql_query_rules;
commit;
```
把所有信息加载到runtime，然后是加载到disk
```
load mysql variables to runtime;
save mysql variables to disk;
load mysql servers to runtime;
save mysql servers to disk;
load mysql users to runtime;
save mysql users to disk;
```

配置scheduler
```
github 地址：
（同一时间多写）
https://raw.githubusercontent.com/ZzzCrazyPig/proxysql_groupreplication_checker/master/proxysql_groupreplication_checker.sh
chmod a+x /var/lib/proxysql/proxysql_groupreplication_checker.sh
insert into scheduler(id, interval_ms, filename, arg1,arg2,arg3,arg4,arg5) values (1,'10000','/var/lib/proxysql/proxysql_groupreplication_checker.sh','1','2','1','1','/var/lib/proxysql/proxysql_groupreplication_checker.log');
load scheduler to runtime;
save scheduler to disk;

```

设置读写分离
```
insert into mysql_query_rules (active, match_pattern, destination_hostgroup, apply) values (1,"^SELECT",2,1);
insert into mysql_query_rules (active, match_pattern, destination_hostgroup, apply) values (1,'^SELECT.*FOR UPDATE$',1,1);
load mysql query rules to runtime;
save mysql query rules to disk;

```


## 命令

查看当前ProxySQL 集群
```
select * from proxysql_servers;
```
查看当前ProxySQL集群状态
```
select * from stats_proxysql_servers_metrics;
```

查询读写分离情况
```
select hostgroup,username,digest_text,count_star from stats_mysql_query_digest;
```



