//CLEAR THE DB
MATCH (n) DETACH DELETE n;
CREATE CONSTRAINT systemCodeUnique IF NOT EXISTS FOR (s:System) REQUIRE s.code IS UNIQUE;

//create systems
CREATE (L1:System {name: 'Laser 1', code: 'L1' })
CREATE (CH1:System {name: 'Chamber 1', code: 'L1CH1' })
CREATE (MIR1:System {name: 'Mirror 1', code: 'L1MI1' })
CREATE (MIR2:System {name: 'Mirror 2', code: 'L1MI2' })
CREATE (CR1:System {name: 'Crystal 1', code: 'L1CR1' })
CREATE (CS1:System {name: 'Control system 1', code: 'L1CS1' })
CREATE (CD1:System {name: 'Control device 1', code: 'L1CS1CDV1' })
CREATE (MOT1:System {name: 'Motor 1', code: 'L1CS1MOT1' })
CREATE (MOT2:System {name: 'Motor 2', code: 'L1CS1MOT2' })
CREATE (PS1:System {name: 'Pressure sensor 1', code: 'L1CS1PS1' })
CREATE (TS1:System {name: 'Temperature sensor 1', code: 'L1CS1TS1' })
CREATE (CAM1:System {name: 'Camera 1', code: 'L1CS1CAM1' })
CREATE (CAM2:System {name: 'Camera 2', code: 'L1CS1CAM2' })
CREATE (CAM3:System {name: 'Camera 3', code: 'L1CS1CAM3' })


//create hierarchical realationships between systems
CREATE (L1)-[:HAS_SUBSYSTEM]->(CH1)
CREATE (CH1)-[:HAS_SUBSYSTEM]->(MIR1)
CREATE (CH1)-[:HAS_SUBSYSTEM]->(MIR2)
CREATE (CH1)-[:HAS_SUBSYSTEM]->(CR1)
CREATE (L1)-[:HAS_SUBSYSTEM]->(CS1)
CREATE (CS1)-[:HAS_SUBSYSTEM]->(CD1)
CREATE (CD1)-[:HAS_SUBSYSTEM]->(MOT1)
CREATE (CD1)-[:HAS_SUBSYSTEM]->(MOT2)
CREATE (CD1)-[:HAS_SUBSYSTEM]->(PS1)
CREATE (CD1)-[:HAS_SUBSYSTEM]->(TS1)
CREATE (CD1)-[:HAS_SUBSYSTEM]->(CAM1)
CREATE (CD1)-[:HAS_SUBSYSTEM]->(CAM2)
CREATE (CD1)-[:HAS_SUBSYSTEM]->(CAM3)

//create users
CREATE (U1:User {username: 'Marie' })
CREATE (U2:User {username: 'Albert' })

//create some maintenance
CREATE (CH1)-[:WAS_MAINTAINED_BY{date:datetime('2022-01-05T15:22')}]->(U1)
CREATE (CH1)-[:WAS_MAINTAINED_BY{date:datetime('2022-02-08T18:19')}]->(U2)
CREATE (CAM1)-[:WAS_MAINTAINED_BY{date:datetime('2022-09-26T10:54')}]->(U1)
CREATE (CAM3)-[:WAS_MAINTAINED_BY{date:datetime('2022-06-11T10:22')}]->(U2)
CREATE (TS1)-[:WAS_MAINTAINED_BY{date:datetime('2022-01-05T15:22')}]->(U1)
CREATE (TS1)-[:WAS_MAINTAINED_BY{date:datetime('2022-08-05T07:22')}]->(U2)
CREATE (TS1)-[:WAS_MAINTAINED_BY{date:datetime('2022-10-02T09:08')}]->(U1)

//create some configuration for the system
CREATE (CFG1:Config {key: 'ExposureMode', value: 'timed' })
CREATE (CFG2:Config {key: 'ExposureTime-us', value: '5000' })
CREATE (CFG3:Config {key: 'TriggerMode', value: 'off' })
CREATE (CFG4:Config {key: 'IP', value: '192.168.1.50' })
CREATE (CAM1)-[:HAS]->(CFG1)
CREATE (CAM1)-[:HAS]->(CFG2)
CREATE (CAM1)-[:HAS]->(CFG3)
CREATE (CAM1)-[:HAS]->(CFG4)

CREATE (CFG5:Config {key: 'ExposureMode', value: 'timed' })
CREATE (CFG6:Config {key: 'ExposureTime-us', value: '1800' })
CREATE (CFG7:Config {key: 'TriggerMode', value: 'on' })
CREATE (CFG8:Config {key: 'IP', value: '192.168.1.51' })
CREATE (CAM2)-[:HAS]->(CFG5)
CREATE (CAM2)-[:HAS]->(CFG6)
CREATE (CAM2)-[:HAS]->(CFG7)
CREATE (CAM2)-[:HAS]->(CFG8)

//generate random sensor data for the temperature sensor
FOREACH ( i IN range(1,1000) |
CREATE(l:TimeValue{ unit: '˚C', time: datetime() + duration({seconds: -i}), value: rand()*50})
CREATE(TS1)-[:LOG]->(l))


//generate random sensor data for the pressure sensor
FOREACH ( i IN range(1,1000) |
CREATE(l:TimeValue{unit: 'mbar', time: datetime() + duration({seconds: -i}), value: rand()*0.001})
CREATE(PS1)-[:LOG]->(l))
