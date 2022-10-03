// MATCH (root:System)
// WHERE NOT exists( ()-[:HAS_SUBSYSTEM]->(root))
// RETURN root
package services

import (
	"panda/apigateway/models"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type SystemsService struct {
	neo4jDriver neo4j.Driver
}

type ISystemsService interface {
	CreateNewSystem(system models.System) (*models.ResponseMessage, error)
	DeleteSystemByCode(systemCode string) (*models.ResponseMessage, error)
	GetSystemByCode(systemCode string) (*models.System, error)
	GetSystemsByNameOrCode(searchText string, limit int32) ([]models.System, error)
	GetSystemMaintenances(systemCode string) ([]models.Maintenance, error)
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
		_, err := tx.Run(`CREATE (s:System { 
			name: $name, 
			code: $code		
			}) 
		RETURN id(s)`, map[string]interface{}{
			"name": system.Code,
			"code": system.Name,
		})
		if err != nil {
			return nil, err
		} else {
			return nil, nil
		}
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SystemsService) DeleteSystemByCode(systemCode string) (*models.ResponseMessage, error) {
	return nil, nil

}

func (svc *SystemsService) GetSystemByCode(systemCode string) (*models.System, error) {
	return nil, nil

}

func (svc *SystemsService) GetSystemsByNameOrCode(searchText string, limit int32) ([]models.System, error) {

	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	records, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		reader, err := tx.Run(`match(s:System) where ($searchText = '' or ((toLower(s.name) CONTAINS $searchText) or (toLower(s.code) CONTAINS $searchText))) and (not (s)<-[:HAS_SUBSYSTEM]-()) return s.name as name , s.code as code, '' as parent
		union
		match(s:System)<-[r:HAS_SUBSYSTEM]-(parent:System) where $searchText = '' or ((toLower(s.name) CONTAINS $searchText) or (toLower(s.code) CONTAINS $searchText)) return s.name as name, s.code as code, parent.name as parent 
		order by id(s)
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

func (svc *SystemsService) GetSystemMaintenances(systemCode string) ([]models.Maintenance, error) {
	return nil, nil
}

func (svc *SystemsService) DeleteConfigurationByKeyAndSystemCode(systemCode string, key string) (*models.ResponseMessage, error) {
	return nil, nil
}

func (svc *SystemsService) GetSystemConfigurationBySystemCode(systemCode string) ([]models.Configuration, error) {
	return nil, nil
}

func (svc *SystemsService) GetSystemTimeValueLogs(systemCode string) ([]models.TimeValueLog, error) {
	return nil, nil
}

// //create new System as a standalone instance/node without any realtionship
// //return new System id
// func (svc *SystemsService) CreateNewSystem(system models.System) (int64, error) {

// 	result := int64(0)

// 	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
// 	defer session.Close()
// 	res, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
// 		newUid, err := uuid.NewRandom()
// 		if err != nil {
// 			return nil, err
// 		}
// 		records, err := tx.Run(`CREATE (s:System {
// 			name: $name,
// 			uid: $uid,
// 			description: $description,
// 			systemCode: $systemCode,
// 			systemAlias: $systemAlias,
// 			facilityZone: $facilityZone,
// 			location: $location,
// 			owner: $owner,
// 			responsible: $responsible,
// 			maintainedBy: $maintainedBy
// 			})
// 		RETURN id(s)`, map[string]interface{}{
// 			"uid":          newUid.String(),
// 			"name":         system.Name,
// 			"description":  system.Description,
// 			"systemCode":   system.SystemCode,
// 			"systemAlias":  system.SystemAlias,
// 			"facilityZone": system.FacilityZone,
// 			"location":     system.Location,
// 			"owner":        system.OwnerPerson,
// 			"responsible":  system.ResponsiblePerson,
// 			"maintainedBy": system.MaintainedByPerson,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		record, err := records.Single()
// 		if err != nil {
// 			return nil, err
// 		}

// 		id := record.Values[0].(int64)

// 		return id, nil
// 	})

// 	if err != nil {
// 		return result, err
// 	}

// 	result = res.(int64)

// 	return result, nil
// }

// //create new System as a subSystem of and existing System
// //you can pass existing System id, uid or name(be careful because the name could not be unique)
// //return new subSystem id
// func (svc *SystemsService) CreateNewSubsystem(subSystem models.System, parentID int64, parentUID string, parentName string) (int64, error) {

// 	result := int64(0)

// 	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
// 	defer session.Close()
// 	res, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
// 		newUid, err := uuid.NewRandom()
// 		if err != nil {
// 			return nil, err
// 		}
// 		records, err := tx.Run(`
// 		MATCH(parent:System) WHERE ($parentName <> "" AND parent.name = $parentName) OR ($parentId <> -1 AND id(parent) = $parentId) OR ($parentUid <> "" AND parent.uid = $parentUid)
// 			CREATE (s:System {
// 			name: $name,
// 			uid: $uid,
// 			description: $description,
// 			systemCode: $systemCode,
// 			systemAlias: $systemAlias,
// 			facilityZone: $facilityZone,
// 			location: $location,
// 			owner: $owner,
// 			responsible: $responsible,
// 			maintainedBy: $maintainedBy
// 			})
// 		CREATE (parent)-[r:HAS_SUBSYSTEM]->(s)
// 		RETURN id(s)`, map[string]interface{}{
// 			"parentId":     parentID,
// 			"parentUid":    parentUID,
// 			"parentName":   parentName,
// 			"uid":          newUid.String(),
// 			"name":         subSystem.Name,
// 			"description":  subSystem.Description,
// 			"systemCode":   subSystem.SystemCode,
// 			"systemAlias":  subSystem.SystemAlias,
// 			"facilityZone": subSystem.FacilityZone,
// 			"location":     subSystem.Location,
// 			"owner":        subSystem.OwnerPerson,
// 			"responsible":  subSystem.ResponsiblePerson,
// 			"maintainedBy": subSystem.MaintainedByPerson,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		record, err := records.Single()
// 		if err != nil {
// 			return nil, err
// 		}

// 		id := record.Values[0].(int64)

// 		return id, nil
// 	})

// 	if err != nil {
// 		return result, err
// 	}

// 	result = res.(int64)

// 	return result, nil
// }

// //update existing System by id
// //return success message
// func (svc *SystemsService) UpdateSystem(system models.System) (string, error) {

// 	result := "Successfully updated"

// 	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
// 	defer session.Close()
// 	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
// 		records, err := tx.Run(`
// 		MATCH(s:System) WHERE id(s) = $id
// 			SET
// 			s.name = case when $name is not null then $name else s.name end,
// 			s.description= case when $description is not null then $description else s.description end,
// 			s.systemCode= case when $systemCode is not null then $systemCode else s.systemCode end,
// 			s.systemAlias= case when $systemAlias is not null then $systemAlias else s.systemAlias end,
// 			s.facilityZone= case when $facilityZone is not null then $facilityZone else s.facilityZone end,
// 			s.location= case when $location is not null then $location else s.location end,
// 			s.owner= case when $owner is not null then $owner else s.owner end,
// 			s.responsible= case when $responsible is not null then $responsible else s.responsible end,
// 			s.maintainedBy= case when $maintainedBy is not null then $maintainedBy else s.maintainedBy end

// 		RETURN id(s)`, map[string]interface{}{
// 			"id":           system.Id,
// 			"name":         system.Name,
// 			"description":  system.Description,
// 			"systemCode":   system.SystemCode,
// 			"systemAlias":  system.SystemAlias,
// 			"facilityZone": system.FacilityZone,
// 			"location":     system.Location,
// 			"owner":        system.OwnerPerson,
// 			"responsible":  system.ResponsiblePerson,
// 			"maintainedBy": system.MaintainedByPerson,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		summary, err := records.Consume()
// 		if err != nil {
// 			return nil, err
// 		}
// 		return summary, nil
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	return result, nil
// }

// //create relationship between two existing systems
// //you can pass existing id, uid or name(be careful because the name could not be unique)
// //return new relationship id
// func (svc *SystemsService) CreateParentChildRelationship(parentID int64, parentUID string, parentName string, childID int64, childUID string, childName string) (int64, error) {

// 	result := int64(0)

// 	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
// 	defer session.Close()
// 	res, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

// 		records, err := tx.Run(`
// 		MATCH(parent:System) WHERE ($parentName <> "" AND parent.name = $parentName) OR ($parentId <> -1 AND id(parent) = $parentId) OR ($parentUid <> "" AND parent.uid = $parentUid)
// 		MATCH(child:System) WHERE ($childName <> "" AND child.name = $childName) OR ($childId <> -1 AND id(child) = $childId) OR ($childUid <> "" AND child.uid = $childUid)
// 		CREATE (parent)-[r:HAS_SUBSYSTEM]->(child)
// 		RETURN id(r)`, map[string]interface{}{
// 			"parentId":   parentID,
// 			"parentUid":  parentUID,
// 			"parentName": parentName,
// 			"childId":    childID,
// 			"childUid":   childUID,
// 			"childName":  childName,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		record, err := records.Single()
// 		if err != nil {
// 			return nil, err
// 		}

// 		id := record.Values[0].(int64)

// 		return id, nil
// 	})

// 	if err != nil {
// 		return result, err
// 	}

// 	result = res.(int64)

// 	return result, nil
// }

// // delete one System by id and all its relationships
// func (svc *SystemsService) DeleteSystemAndRelationships(systemId int64) (string, error) {

// 	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
// 	defer session.Close()
// 	summary, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

// 		records, err := tx.Run(`MATCH(s:System) WHERE id(s) = $systemId detach delete s`, map[string]interface{}{
// 			"systemId": systemId,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		resultSummary, err := records.Consume()
// 		if err != nil {
// 			return nil, err
// 		}

// 		return resultSummary, nil
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	resultMessage := ""
// 	if summary != nil {
// 		rs := summary.(neo4j.ResultSummary)

// 		resultMessage += strconv.Itoa(rs.Counters().NodesDeleted()) + " node(s) deleted, "
// 		resultMessage += strconv.Itoa(rs.Counters().RelationshipsDeleted()) + " relationship(s) deleted"
// 	}

// 	return resultMessage, nil
// }

// // delete one relationship by id
// func (svc *SystemsService) DeleteRelationshipByID(id int64) (string, error) {

// 	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
// 	defer session.Close()
// 	summary, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

// 		records, err := tx.Run(`MATCH(parent:System)-[r:HAS_SUBSYSTEM]->(child:System) WHERE id(r)=$id DELETE r`, map[string]interface{}{
// 			"id": id,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		resultSummary, err := records.Consume()
// 		if err != nil {
// 			return nil, err
// 		}

// 		return resultSummary, nil
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	resultMessage := ""
// 	if summary != nil {
// 		rs := summary.(neo4j.ResultSummary)

// 		resultMessage += strconv.Itoa(rs.Counters().RelationshipsDeleted()) + " relationship(s) deleted"
// 	}

// 	return resultMessage, nil
// }

// // delete one relationship by parent and child ids
// func (svc *SystemsService) DeleteRelationshipByParentChildIds(parentId int64, childId int64) (string, error) {

// 	session := svc.neo4jDriver.NewSession(neo4j.SessionConfig{})
// 	defer session.Close()
// 	summary, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

// 		records, err := tx.Run(`MATCH (parent:System)-[r:HAS_SUBSYSTEM]->(child:System) WHERE id(parent)=$parentId and id(child)=$childId DELETE r`, map[string]interface{}{
// 			"parentId": parentId,
// 			"childId":  childId,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		resultSummary, err := records.Consume()
// 		if err != nil {
// 			return nil, err
// 		}

// 		return resultSummary, nil
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	resultMessage := ""
// 	if summary != nil {
// 		rs := summary.(neo4j.ResultSummary)

// 		resultMessage += strconv.Itoa(rs.Counters().RelationshipsDeleted()) + " relationship(s) deleted"
// 	}

// 	return resultMessage, nil
// }
