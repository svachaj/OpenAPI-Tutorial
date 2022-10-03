// Maintenance history - all
MATCH (s:System)-[m:WAS_MAINTAINED_BY]->(u:User) RETURN m.date, u.username, s.name

//get system time values by system code and time range
match(s:System{code:'L1CS1PS1'})-[]->(log:TimeValue) 
where (null is null or log.time > datetime('2022-10-03T15:45:44')) 
and log.time < datetime('2022-10-03T15:46:04') 
return log.time, log.value, log.unit order by log.time

//get system configuration by system code 
match(s:System{code:'L1CS1CAM1'})-[]->(c:Config) return c.key, c.value

//get all systems and search by name or code - contains
//get also its parent system
call {
match(s:System) where (null is NULL or ((toLower(s.name) CONTAINS null) or (toLower(s.code) CONTAINS null))) and (not (s)<-[:HAS_SUBSYSTEM]-()) 
return s.name as name , s.code as code, '' as parent
union
match(s:System)<-[r:HAS_SUBSYSTEM]-(parent:System) where null is NULL or ((toLower(s.name) CONTAINS null) or (toLower(s.code) CONTAINS null)) 
return s.name as name, s.code as code, parent.name as parent 
}
return name, code, parent
limit 100

//create new system - no parent system
CREATE (s:System {name: 'testnoparent', code: 'test' })

//create new system - with parent system
MATCH (parent:System{code:'L1'})
CREATE (s:System {name: 'testwithparent', code: 'test' })
CREATE (parent)-[:HAS_SUBSYSTEM]->(s)

//delete system
MATCH (s:System{code:'test'})
DETACH DELETE s

//get system info by code
MATCH (s:System{code:'test'}) return s.code, s.name, id(s)