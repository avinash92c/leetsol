# Write your MySQL query statement below
SELECT e.Name as Employee FROM EMPLOYEE e JOIN EMPLOYEE m ON e.ManagerID=m.Id WHERE e.Salary>m.Salary
