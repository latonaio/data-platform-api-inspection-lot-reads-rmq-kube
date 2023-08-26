package dpfm_api_output_formatter

import (
	"data-platform-api-inspection-lot-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)
func ConvertToComponentComposition(rows *sql.Rows) (*[]ComponentComposition, error) {
	defer rows.Close()
	componentComposition := make([]ComponentComposition, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ComponentComposition{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.ComponentCompositionType,
			&pm.ComponentCompositionUperLimitInPercent,
			&pm.ComponentCompositionLowerLimitInPercent,
			&pm.ComponentCompositionStandardValueInPercent,
			&pm.Creationstring,
			&pm.LastChangestring,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &componentComposition, err
		}

		data := pm
		componentComposition = append(ComponentComposition, ComponentComposition{
			InspectionLot:                              data.InspectionLot,
			ComponentCompositionType:                   data.ComponentCompositionType,
			ComponentCompositionUperLimitInPercent:     data.ComponentCompositionUperLimitInPercent,
			ComponentCompositionLowerLimitInPercent:    data.ComponentCompositionLowerLimitInPercent,
			ComponentCompositionStandardValueInPercent: data.ComponentCompositionStandardValueInPercent,
			Creationstring:                             data.Creationstring,
			LastChangestring:                           data.LastChangestring,
			IsCancelled:                                data.IsCancelled,
			IsMarkedForDeletion:                        data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &componentComposition, nil
	}

	return &componentComposition, nil
}
func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.InspectionPlan,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
			&pm.Product,
			&pm.ProductSpecification,
			&pm.InspectionSpecification,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.InspectionLotHeaderText,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			InspectionLot:                  data.InspectionLot,
			InspectionPlan:                 data.InspectionPlan,
			InspectionPlantBusinessPartner: data.InspectionPlantBusinessPartner,
			InspectionPlant:                data.InspectionPlant,
			Product:                        data.Product,
			ProductSpecification:           data.ProductSpecification,
			InspectionSpecification:        data.InspectionSpecification,
			ProductionOrder:                data.ProductionOrder,
			ProductionOrderItem:            data.ProductionOrderItem,
			InspectionLotHeaderText:        data.InspectionLotHeaderText,
			CreationDate:                   data.CreationDate,
			LastChangeDate:                 data.LastChangeDate,
			IsCancelled:                    data.IsCancelled,
			IsMarkedForDeletion:            data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToHeaderDoc(rows *sql.Rows) (*[]HeaderDoc, error) {
	defer rows.Close()
	headerDoc := make([]HeaderDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.HeaderDoc{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &headerDoc, err
		}

		data := pm
		headerDoc = append(headerDoc, HeaderDoc{
			InspectionLot:            data.InspectionLot,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &headerDoc, nil
	}

	return &headerDoc, nil
}

func ConvertToInspection(rows *sql.Rows) (*[]Inspection, error) {
	defer rows.Close()
	inspection := make([]Inspection, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Inspection{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.Inspection,
			&pm.InspectionType,
			&pm.InspectionTypeValueUnit,
			&pm.InspectionTypePlannedValue,
			&pm.InspectionTypeCertificateType,
			&pm.InspectionTypeCertificateValueInText,
			&pm.InspectionTypeCertificateValueInQuantity,
			&pm.InspectionLotInspectionText,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &inspection, err
		}

		data := pm
		inspection = append(inspection, Inspection{
			InspectionLot:                            data.InspectionLot,
			Inspection:                               data.Inspection,
			InspectionType:                           data.InspectionType,
			InspectionTypeValueUnit:                  data.InspectionTypeValueUnit,
			InspectionTypePlannedValue:               data.InspectionTypePlannedValue,
			InspectionTypeCertificateType:            data.InspectionTypeCertificateType,
			InspectionTypeCertificateValueInText:     data.InspectionTypeCertificateValueInText,
			InspectionTypeCertificateValueInQuantity: data.InspectionTypeCertificateValueInQuantity,
			InspectionLotInspectionText:              data.InspectionLotInspectionText,
			CreationDate:                             data.CreationDate,
			LastChangeDate:                           data.LastChangeDate,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &inspection, nil
	}

	return &inspection, nil
}
func ConvertToOperation(rows *sql.Rows) (*[]Operation, error) {
	defer rows.Close()
	Operation := make([]Operation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Operation{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.Inspection,
			&pm.OperationType,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
			&pm.Sequence,
			&pm.SequenceText,
			&pm.OperationText,
			&pm.OperationStatus,
			&pm.ResponsiblePlannerGroup,
			&pm.OperationUnit,
			&pm.StandardLotSizeQuantity,
			&pm.MinimumLotSizeQuantity,
			&pm.MaximumLotSizeQuantity,
			&pm.PlainLongText,
			&pm.WorkCenter,
			&pm.CapacityCategoryCode,
			&pm.OperationCostingRelevancyType,
			&pm.OperationSetupType,
			&pm.OperationSetupGroupCategory,
			&pm.OperationSetupGroup,
			&pm.OperationReferenceQuantity,
			&pm.MaximumWaitDuration,
			&pm.StandardWaitDuration,
			&pm.MinimumWaitDuration,
			&pm.WaitDurationUnit,
			&pm.MaximumQueueDuration,
			&pm.StandardQueueDuration,
			&pm.MinimumQueueDuration,
			&pm.QueueDurationUnit,
			&pm.MaximumMoveDuration,
			&pm.StandardMoveDuration,
			&pm.MinimumMoveDuration,
			&pm.MoveDurationUnit,
			&pm.StandardDeliveryDuration,
			&pm.StandardDeliveryDurationUnit,
			&pm.StandardOperationScrapPercent,
			&pm.CostElement,
			&pm.ValidityStart,
			&pm.ValidityEnd,
			&pm.Creation,
			&pm.LastChange,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &operation, err
		}

		data := pm
		operation = append(operation, Operation{
			InspectionLot:                            data.InspectionLot,
			Operations:                               data.Operations,
			OperationsItem:                           data.OperationsItem,
			OperationID:                              data.OperationID,
			Inspection:                               data.Inspection,
			OperationType:                            data.OperationType,
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:        data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:   data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Product:                                  data.Product,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			DeliverToParty:                           data.DeliverToParty,
			DeliverToPlant:                           data.DeliverToPlant,
			DeliverFromParty:                         data.DeliverFromParty,
			DeliverFromPlant:                         data.DeliverFromPlant,
			InspectionPlantBusinessPartner:           data.InspectionPlantBusinessPartner,
			InspectionPlant:                          data.InspectionPlant,
			Sequence:                                 data.Sequence,
			SequenceText:                             data.SequenceText,
			OperationText:                            data.OperationText,
			OperationStatus:                          data.OperationStatus,
			ResponsiblePlannerGroup:                  data.ResponsiblePlannerGroup,
			OperationUnit:                            data.OperationUnit,
			StandardLotSizeQuantity:                  data.StandardLotSizeQuantity,
			MinimumLotSizeQuantity:                   data.MinimumLotSizeQuantity,
			MaximumLotSizeQuantity:                   data.MaximumLotSizeQuantity,
			PlainLongText:                            data.PlainLongText,
			WorkCenter:                               data.WorkCenter,
			CapacityCategoryCode:                     data.CapacityCategoryCode,
			OperationCostingRelevancyType:            data.OperationCostingRelevancyType,
			OperationSetupType:                       data.OperationSetupType,
			OperationSetupGroupCategory:              data.OperationSetupGroupCategory,
			OperationSetupGroup:                      data.OperationSetupGroup,
			OperationReferenceQuantity:               data.OperationReferenceQuantity,
			MaximumWaitDuration:                      data.MaximumWaitDuration,
			StandardWaitDuration:                     data.StandardWaitDuration,
			MinimumWaitDuration:                      data.MinimumWaitDuration,
			WaitDurationUnit:                         data.WaitDurationUnit,
			MaximumQueueDuration:                     data.MaximumQueueDuration,
			StandardQueueDuration:                    data.StandardQueueDuration,
			MinimumQueueDuration:                     data.MinimumQueueDuration,
			QueueDurationUnit:                        data.QueueDurationUnit,
			MaximumMoveDuration:                      data.MaximumMoveDuration,
			StandardMoveDuration:                     data.StandardMoveDuration,
			MinimumMoveDuration:                      data.MinimumMoveDuration,
			MoveDurationUnit:                         data.MoveDurationUnit,
			StandardDeliveryDuration:                 data.StandardDeliveryDuration,
			StandardDeliveryDurationUnit:             data.StandardDeliveryDurationUnit,
			StandardOperationScrapPercent:            data.StandardOperationScrapPercent,
			CostElement:                              data.CostElement,
			ValidityStart:                            data.ValidityStart,
			ValidityEnd:                              data.ValidityEnd,
			Creation:                                 data.Creation,
			LastChange:                               data.LastChange,
			IsCancelled:                              data.IsCancelled,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &operation, nil
	}

	return &operation, nil
}
func ConvertToSpecDetail(rows *sql.Rows) (*[]SpecDetail, error) {
	defer rows.Close()
	specDetail := make([]SpecDetail, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SpecDetail{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.SpecType,
			&pm.UpperLimitValue,
			&pm.LowerLimitValue,
			&pm.StandardValue,
			&pm.SpecTypeUnit,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &specDetail, err
		}

		data := pm
		specDetail = append(specDetail, SpecDetail{
			InspectionLot:       data.InspectionLot,
			SpecType:            data.SpecType,
			UpperLimitValue:     data.UpperLimitValue,
			LowerLimitValue:     data.LowerLimitValue,
			StandardValue:       data.StandardValue,
			SpecTypeUnit:        data.SpecTypeUnit,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsCancelled:         data.IsCancelled,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &specDetail, nil
	}

	return &SpecDetail, nil
}

func ConvertToSpecGeneral(rows *sql.Rows) (*[]SpecGeneral, error) {
	defer rows.Close()
	specGeneral := make([]SpecGeneral, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SpecGeneral{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.HeatNumber,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &specDetail, err
		}

		data := pm
		specGeneral = append(specGeneral, SpecGeneral{
			InspectionLot:       data.InspectionLot,
			HeatNumber:          data.HeatNumber,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsCancelled:         data.IsCancelled,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &specGeneral, nil
	}

	return &SpecGeneral, nil
}
