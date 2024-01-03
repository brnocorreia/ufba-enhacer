/* Script para criação da tabela de disciplinas */

CREATE TABLE IF NOT EXISTS disciplines (
    code VARCHAR(8) NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    workload INT NOT NULL,
    department VARCHAR(60) NOT NULL,
    program VARCHAR(5000) NOT NULL,
    objective VARCHAR(5000) NOT NULL,
    content VARCHAR(5000) NOT NULL,
    bibliography VARCHAR(5000) NOT NULL
)

/* Script para criação da tabela de cursos */

CREATE TABLE IF NOT EXISTS courses (
	code INT NOT NULL PRIMARY KEY,
	shift VARCHAR(15) NOT NULL,
	min_duration DOUBLE PRECISION NOT NULL,
	max_duration DOUBLE PRECISION NOT NULL,
	legal_base VARCHAR(500) NOT NULL,
	description VARCHAR(5000) NOT NULL,
	ob_workload INT NOT NULL,
	op_workload INT NOT NULL,
	ac_workload INT NOT NULL
)

/* Relationship entre cursos e disciplinas */

CREATE TABLE IF NOT EXISTS courses_disciplines (
	course_id INT REFERENCES courses(code) ON DELETE CASCADE,
	discipline_id VARCHAR(8) REFERENCES disciplines(code) ON DELETE CASCADE,
	
	PRIMARY KEY (course_id, discipline_id)
)

/* Relationship entre disciplinas e seus pre requisitos (disciplinas também) */
CREATE TABLE IF NOT EXISTS prerequisites (
	course_id INT REFERENCES courses(code) ON DELETE CASCADE,
	discipline_id VARCHAR(8) REFERENCES disciplines(code) ON DELETE CASCADE,
	depends_on VARCHAR(8) REFERENCES disciplines(code) ON DELETE CASCADE,
	PRIMARY KEY (course_id, discipline_id, depends_on)
)

/* Query para retornar todas as materias de um curso e seus pre requisitos */

SELECT 
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
    d.code
