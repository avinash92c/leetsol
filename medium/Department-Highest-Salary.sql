SELECT d.Name as Department,e.Name as Employee,e.Salary as Salary FROM Employee e JOIN Department d On e.DepartmentId=d.Id WHERE (e.Salary,e.DepartmentId) IN (
SELECT MAX(Salary) ,DepartmentId  FROM Employee Group By DepartmentId  )
