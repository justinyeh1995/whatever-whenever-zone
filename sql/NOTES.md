# Notes

## Many-To-Many Relations Modeling

## Data Normalization

1-NF

## Indexing

## Data Partitioning

## Window Function

```sql
--STANDARD FORMULA
Aggregation_function --like ROW_NUMBER()
OVER(PARTITION BY col1
     ORDER BY col2)
```

Common Aggregation Function

- ROW_NUMBER()
- RANK() & DENSE_RANK()
- LAG(column, n) & LEAD(column, n)
- NTILE()
- MAX()& MIN()&AVG()
- FIRST_VALUE() & LAST_VALUE()

Ref: [十分鐘內快速上手與使用 Window function｜SQL 教學](https://medium.com/%E6%95%B8%E6%93%9A%E4%B8%8D%E6%AD%A2-not-only-data/%E5%A6%82%E4%BD%95%E5%8D%81%E5%88%86%E9%90%98%E5%85%A7%E5%BF%AB%E9%80%9F%E4%B8%8A%E6%89%8B%E8%88%87%E4%BD%BF%E7%94%A8-window-function-e24e0a7e75ba)
