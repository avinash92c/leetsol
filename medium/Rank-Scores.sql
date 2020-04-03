/* Write your T-SQL query statement below */
SELECT score,DENSE_RANK() OVER( Order By score DESC) AS 'Rank' FROM Scores
