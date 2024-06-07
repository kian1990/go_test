---
date: 2024-04-26
authors:
  - kian
categories:
  - 文档
---

# **MySQL**

## 常用查询语句
<!-- more -->
```sql
# 统计事件大分类
SELECT SUBSTR(GKLX,1,4),COUNT(SUBSTR(GKLX,1,4)) AS item_count FROM `xnsdsjj_12345cbs` GROUP BY SUBSTR(GKLX,1,4) ORDER BY item_count DESC;

# 统计纠纷分类
SELECT GKLX,COUNT(GKLX) AS item_count FROM `xnsdsjj_12345cbs` GROUP BY GKLX HAVING GKLX LIKE '%纠纷%' ORDER BY item_count DESC;

# 统计指定日期内纠纷分类
CREATE TEMPORARY TABLE temp_table AS SELECT BLBM,SLSJ,GKLX,COUNT(DISTINCT GKLX) AS item_count FROM `xnsdsjj_12345cbs` GROUP BY BLBM,SLSJ,GKLX HAVING SLSJ >= '2023-01-01 00:00:00' AND GKLX LIKE '%纠纷%';
SELECT GKLX,COUNT(GKLX) AS GKLX_COUNT FROM temp_table GROUP BY GKLX ORDER BY GKLX_COUNT DESC;
```
