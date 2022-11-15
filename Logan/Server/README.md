# 后端应用

## 数据表

1、The database configuration information is read from the resources / db.properties file,
please replace it with your own database information

2、Database table structure：
CREATE TABLE `logan_task` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `platform` tinyint(11) unsigned NOT NULL COMMENT '平台1android2iOS',
  `amount` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '文件大小',
  `app_id` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'app标识',
  `union_id` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '用户标识',
  `device_id` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '设备标识',
  `app_version` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'app版本',
  `build_version` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '构建版本',
  `log_file_name` varchar(512) CHARACTER SET utf8mb4 DEFAULT '' COMMENT '日志文件所在路径',
  `log_date` bigint(20) unsigned DEFAULT NULL COMMENT '日志所属日期',
  `add_time` bigint(20) unsigned NOT NULL COMMENT '业务字段，日志上报时间',
  `status` tinyint(11) unsigned NOT NULL DEFAULT '0' COMMENT '0未分析过,1分析过',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='上报日志任务表';

CREATE TABLE `logan_log_detail` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `task_id` bigint(11) unsigned NOT NULL COMMENT '所属任务id',
  `log_type` int(11) unsigned NOT NULL COMMENT '日志类型',
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原始日志',
  `log_time` bigint(20) unsigned NOT NULL COMMENT '本条日志产生的具体时间戳',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_update_time` (`update_time`),
  KEY `idx_task_id` (`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='日志解析后的数据详情';

CREATE TABLE `web_task` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `device_id` varchar(128) NOT NULL DEFAULT '' COMMENT '设备id',
  `web_source` varchar(128) DEFAULT NULL COMMENT '来源',
  `environment` varchar(2048) DEFAULT NULL COMMENT '客户端自定义环境信息',
  `page_num` int(11) NOT NULL COMMENT '日志页码',
  `content` mediumtext NOT NULL COMMENT '日志内容',
  `add_time` bigint(20) NOT NULL COMMENT '添加时间',
  `log_date` bigint(20) NOT NULL COMMENT '日志所属日期',
  `status` int(11) NOT NULL COMMENT '日志状态0未解析，1已解析',
  `custom_report_info` varchar(2048) DEFAULT NULL COMMENT '自定义上报信息',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `log_date_deviceid` (`log_date`,`device_id`),
  KEY `add_time_deviceid` (`add_time`,`device_id`),
  KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='H5上报任务表';

CREATE TABLE `web_detail` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `task_id` bigint(20) NOT NULL COMMENT '所属任务id',
  `log_type` int(11) NOT NULL COMMENT '日志类型',
  `content` mediumtext NOT NULL COMMENT '日志内容',
  `log_time` bigint(20) NOT NULL COMMENT '日志所属时间',
  `log_level` int(11) DEFAULT NULL COMMENT '日志等级',
  `add_time` bigint(20) NOT NULL COMMENT '添加时间',
  `minute_offset` int(11) NOT NULL COMMENT '距离当天0点的分钟数',
  PRIMARY KEY (`id`),
  KEY `taskid_logtype` (`task_id`,`log_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='H5日志详情表';

3、The native log files are saved on the local disk.
Please modify fileService according to your actual scenario to suit your storage method.

4、You need to implement the ContentHandler interface to format the display according to the log type.
The getSimpleContent() method displays the summary. The getFormatContent () method displays the Formatted data.

## docker构建镜像

应用编译与Docker镜像构建分离
对于静态编译型语言，我们通常需要将应用编译过程与镜像构建过程分离。主要有以下两个考虑：

最终生成的Docker镜像不应该包含源代码 - 信息安全
最终生成的Docker镜像应该最小化，不应该包含编译时工具 - 加速应用交付，减小攻击面
传统做法是通过一个单独的Dockerfile来构建应用，并将编译结果从Docker镜像中拷贝出来。而在Docker 17.05版本之后，Docker引入了多阶段构建方式，我们可以把镜像构建分解为若干个部分，如下所示

第一阶段负责编译 Java 应用
第一阶段利用最小化的 JRE 镜像 openjdk:8-jre-alpine 来构建应用容器镜像

### First stage - Compiling application

FROM registry.cn-hangzhou.aliyuncs.com/acs/maven:3-jdk-8 AS build-env

ENV MY_HOME=/app
RUN mkdir -p $MY_HOME
WORKDIR $MY_HOME
ADD pom.xml $MY_HOME

### get all the downloads out of the way

RUN ["/usr/local/bin/mvn-entrypoint.sh","mvn","verify","clean","--fail-never"]

1. add source
ADD . $MY_HOME
2. run maven verify
RUN ["/usr/local/bin/mvn-entrypoint.sh","mvn","verify"]
3. Second stage - build image

 ```dockerfile
 FROM openjdk:8-jre-alpine
 COPY --from=build-env /app/target/*.jar /app.jar
 ENV JAVA_OPTS=""
 ENV SERVER_PORT 8080
 EXPOSE ${SERVER_PORT}

 ENTRYPOINT [ "sh", "-c", "java $JAVA_OPTS -Djava.security.egd=file:/dev/urandom -jar /app.jar" ]
 ```

通过这个方法可以将，700多兆的Docker镜像，缩小到100兆，这是巨大的节约。
