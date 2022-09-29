MATCH (n) DETACH DELETE n;

CREATE (L1:System {name: 'Laser 1', code: 'L1' })
CREATE (L2:System {name: 'Laser 2', code: 'L2' })
CREATE (L3:System {name: 'Laser 3', code: 'L3' })

CREATE (L1)-[:HAS_SUBSYSTEM]->(L2)
CREATE (L2)-[:HAS_SUBSYSTEM]->(L3)

MATCH path = (s:System)-[:HAS_SUBSYSTEM]->(System)
WITH collect(path) AS paths
CALL apoc.convert.toTree(paths)
YIELD value
RETURN value;
