CREATE TABLE table_name (
`xm` STRING COMMENT '本人姓名',
`ssxz` STRING COMMENT '所属乡镇/社区',
`name` STRING COMMENT '小区名',
`number` STRING COMMENT '楼栋户号',
`yhzgx` STRING COMMENT '与户主关系',
`xb` STRING COMMENT '性别',
`source` STRING COMMENT '居住类型',
`contactNumber` STRING COMMENT '联系电话',
`gmsfhm` STRING COMMENT '身份证号',
`mz` STRING COMMENT '名族',
`hjdz` STRING COMMENT '户籍所在地',
`whcd` STRING COMMENT '学历',
`rycbzt` STRING COMMENT '社保状况',
`category` STRING COMMENT '身份类型',
`spStatus` STRING COMMENT '特殊人员类型',
`maritalStatus` STRING COMMENT '婚姻状况',
`politicalStatus` STRING COMMENT '政治面貌',
`focusOn` STRING COMMENT '是否常驻',
`floatOutTarget` STRING COMMENT '流动地点',
`religion` STRING COMMENT '宗教信仰',
`cylb` STRING COMMENT '职业属性',
`cbsj` STRING COMMENT '最近参保年份',
`xz` STRING COMMENT '参保类型',
`dwmc` STRING COMMENT '参保单位',
`occupation` STRING COMMENT '就业状态',
`jyxs` STRING COMMENT '就业方式',
`xydm` STRING COMMENT '就业工种',
`jydszsxq` STRING COMMENT '就业地点'
) ROW FORMAT SERDE 'org.apache.hadoop.hive.serde2.lazy.LazySimpleSerDe'
WITH SERDEPROPERTIES (
"field.delim"='\001',
"escape.delim"='\n',
"colelction.delim"=',',
"mapkey.delim"=':'
)
stored as TEXTFILE;