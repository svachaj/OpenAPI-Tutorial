// MATCH (root:System)
// WHERE NOT exists( ()-[:HAS_SUBSYSTEM]->(root))
// RETURN root
package services

import (
	"panda/apigateway/models"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type SystemsService struct {
	neo4jDriver neo4j.Driver
}

type ISystemsService interface {
	CreateNewSystem(system models.System) (*models.ResponseMessage, error)
	DeleteSystemByCode(systemCode string) (*models.ResponseMessage, error)
	GetSystemByCode(systemCode string) (models.System, error)
	GetSystemsByNameOrCode(searchText string, limit int32) ([]models.System, error)
	GetSystemMaintenance(systemCode string) ([]models.Maintenance, error)
	DeleteConfigurationByKeyAndSystemCode(systemCode string, key string) (*models.ResponseMessage, error)
	GetSystemConfigurationBySystemCode(systemCode string) ([]models.Configuration, error)
	GetSystemTimeValueLogs(systemCode string) ([]models.TimeValueLog, error)
}

func NewSystemsService(driver neo4j.Driver) ISystemsService {
	return &SystemsService{
		neo4jDriver: driver,
	}
}

//Create new System. If parentSystemCode is specified, create also hierrarchical relationship to this parent System.
func (svc *SystemsService) CreateNewSystem(system models.System) (*models.ResponseMessage, error) {

	result := models.ResponseMessage{Message: "System was succesfuly created."}

	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if system.ParentSystemCode == "" {
			_, err := tx.Run(`CREATE (s:System { 
			name: $name, 
			code: $code		
			}) 
		RETURN id(s)`, map[string]interface{}{
				"name": system.Name,
				"code": system.Code,
			})
			if err != nil {
				return nil, err
			}
			return nil, nil
		} else {
			_, err := tx.Run(`MATCH (parent:System{code:$parentCode})
			CREATE (s:System {name: $name, code: $code })
			CREATE (parent)-[:HAS_SUBSYSTEM]->(s)`, map[string]interface{}{
				"name":       system.Name,
				"code":       system.Code,
				"parentCode": system.ParentSystemCode,
			})
			if err != nil {
				return nil, err
			}
			return nil, nil
		}

	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SystemsService) DeleteSystemByCode(systemCode string) (*models.ResponseMessage, error) {
	result := models.ResponseMessage{Message: "System was succesfuly deleted."}

	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(`MATCH (s:System{code: $code})
		DETACH DELETE s`, map[string]interface{}{
			"code": systemCode,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SystemsService) GetSystemByCode(systemCode string) (models.System, error) {

	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	record, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		reader, err := tx.Run(`MATCH (s:System{code: $code}) return s.code, s.name`, map[string]interface{}{
			"code": systemCode,
		})

		if err != nil {
			return nil, err
		}

		item := models.System{}
		rec, err := reader.Single()

		if err != nil {
			return nil, err
		}
		item.Code = rec.Values[0].(string)
		item.Name = rec.Values[1].(string)

		return item, nil
	})

	if err != nil {
		return models.System{}, err
	}

	return record.(models.System), nil
}

func (svc *SystemsService) GetSystemsByNameOrCode(searchText string, limit int32) ([]models.System, error) {

	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		reader, err := tx.Run(`call {
			match(s:System) where ($searchText = '' or ((toLower(s.name) CONTAINS $searchText) or (toLower(s.code) CONTAINS $searchText))) and (not (s)<-[:HAS_SUBSYSTEM]-()) 
			return s.name as name , s.code as code, '' as parent
			union
			match(s:System)<-[r:HAS_SUBSYSTEM]-(parent:System) where $searchText = '' or ((toLower(s.name) CONTAINS $searchText) or (toLower(s.code) CONTAINS $searchText)) 
			return s.name as name, s.code as code, parent.name as parent 
			}
			return name, code, parent
			limit $limit`, map[string]interface{}{
			"searchText": searchText,
			"limit":      limit,
		})

		if err != nil {
			return nil, err
		}

		list := make([]models.System, 0)

		for reader.Next() {
			list = append(list, models.System{Name: reader.Record().Values[0].(string), Code: reader.Record().Values[1].(string), ParentSystemCode: reader.Record().Values[2].(string)})
		}
		if err = reader.Err(); err != nil {
			return nil, err
		}

		return list, nil
	})

	if err != nil {
		return nil, err
	}

	return records.([]models.System), nil
}

func (svc *SystemsService) GetSystemMaintenance(systemCode string) ([]models.Maintenance, error) {
	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		reader, err := tx.Run(`MATCH (s:System)-[m:WAS_MAINTAINED_BY]->(u:User) WHERE $systemCode = '' or s.code = $systemCode RETURN m.date, u.username, s.name`, map[string]interface{}{
			"systemCode": systemCode,
		})

		if err != nil {
			return nil, err
		}

		list := make([]models.Maintenance, 0)

		for reader.Next() {
			list = append(list, models.Maintenance{When: reader.Record().Values[0].(time.Time), Username: reader.Record().Values[1].(string), SystemName: reader.Record().Values[2].(string)})
		}
		if err = reader.Err(); err != nil {
			return nil, err
		}

		return list, nil
	})

	if err != nil {
		return nil, err
	}

	return records.([]models.Maintenance), nil
}

func (svc *SystemsService) DeleteConfigurationByKeyAndSystemCode(systemCode string, key string) (*models.ResponseMessage, error) {
	result := models.ResponseMessage{Message: "Configuration was succesfuly deleted."}

	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(`match(s:System{code:$systemCode})-[]->(c:Config{key: $key}) detach delete c`, map[string]interface{}{
			"systemCode": systemCode,
			"key":        key,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SystemsService) GetSystemConfigurationBySystemCode(systemCode string) ([]models.Configuration, error) {
	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		reader, err := tx.Run(`match(s:System{code: $systemCode})-[]->(c:Config) return c.key, c.value`, map[string]interface{}{
			"systemCode": systemCode,
		})

		if err != nil {
			return nil, err
		}

		list := make([]models.Configuration, 0)

		for reader.Next() {
			list = append(list, models.Configuration{Key: reader.Record().Values[0].(string), Value: reader.Record().Values[1].(string)})
		}
		if err = reader.Err(); err != nil {
			return nil, err
		}

		return list, nil
	})

	if err != nil {
		return nil, err
	}

	return records.([]models.Configuration), nil
}

func (svc *SystemsService) GetSystemTimeValueLogs(systemCode string) ([]models.TimeValueLog, error) {
	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		reader, err := tx.Run(`match(s:System{code: $systemCode})-[]->(log:TimeValue) 
		
		return log.time, log.value, log.unit order by log.time`, map[string]interface{}{
			"systemCode": systemCode,
		})

		if err != nil {
			return nil, err
		}

		list := make([]models.TimeValueLog, 0)

		for reader.Next() {
			list = append(list, models.TimeValueLog{Time: reader.Record().Values[0].(time.Time), Value: reader.Record().Values[1].(float64), Unit: reader.Record().Values[2].(string)})
		}
		if err = reader.Err(); err != nil {
			return nil, err
		}

		return list, nil
	})

	if err != nil {
		return nil, err
	}

	return records.([]models.TimeValueLog), nil
}
