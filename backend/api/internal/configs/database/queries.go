package database

var getAll = `SELECT 
d.*,
STRING_AGG(pr.depends_on, ', ') AS depends_on_list
FROM
courses c
JOIN
courses_disciplines cd ON c.code = cd.course_id
JOIN
disciplines d ON cd.discipline_id = d.code
JOIN 
prerequisites pr ON c.code = pr.course_id
WHERE
c.code = 316130
GROUP BY
d.code`
