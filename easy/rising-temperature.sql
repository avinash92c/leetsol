# Write your MySQL query statement below
SELECT w.Id FROM Weather w JOIN Weather w1 On  w.Temperature>w1.Temperature AND DATEDIFF(w.RecordDate,w1.RecordDate)=1
