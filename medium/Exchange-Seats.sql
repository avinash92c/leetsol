SELECT id, CASE WHEN id%2=0 THEN ISNULL(LAG(student) OVER (order by id),student) ELSE ISNULL(LEAD(student) OVER (order by id),student) END as student from seat 
