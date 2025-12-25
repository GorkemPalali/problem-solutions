SELECT round(SUM(total.tiv_2016),2 )AS tiv_2016
FROM (
    SELECT DISTINCT a.pid, a.tiv_2016
    FROM Insurance AS a
    JOIN Insurance AS b
      ON a.tiv_2015 = b.tiv_2015 
    WHERE a.pid <> b.pid

INTERSECT

SELECT a.pid, a.tiv_2016
FROM Insurance a
WHERE NOT EXISTS (
    SELECT 1
    FROM Insurance b
    WHERE a.pid <> b.pid
      AND a.lat = b.lat
      AND a.lon = b.lon

)
) AS total; 