# Write your MySQL query statement below
UPDATE  salary SET SEX=(CASE WHEN sex='m' then 'f' WHEN sex='f' THEN 'm' END)
