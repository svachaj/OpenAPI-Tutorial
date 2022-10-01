/*Maintenance history*/
MATCH (s:System)-[m:WAS_MAINTAINED_BY]->(u:User) RETURN m.date, u.username, s.name

